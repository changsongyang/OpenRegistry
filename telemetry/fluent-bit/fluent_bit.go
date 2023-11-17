package fluentbit

import (
	"bytes"
	"context"
	"net/http"
	"time"

	"github.com/containerish/OpenRegistry/config"
	"github.com/google/uuid"
)

type (
	FluentBit interface {
		Send(logBytes []byte)
	}

	fluentBit struct {
		client        *http.Client
		retryMessages map[string]retryLogMsg
		config        *config.OpenRegistryConfig
	}

	retryLogMsg struct {
		content []byte
		count   int64
		done    bool
	}
)

func New(config *config.OpenRegistryConfig) (FluentBit, error) {
	httpClient := &http.Client{
		Timeout: time.Duration(time.Second * 30),
	}

	fbClient := &fluentBit{
		client:        httpClient,
		config:        config,
		retryMessages: make(map[string]retryLogMsg),
	}

	go fbClient.retry()

	return fbClient, nil
}

func (fb *fluentBit) Send(logBytes []byte) {
	// don't send logs to grafana from local instances of OpenRegistry
	local := fb.config.Environment == config.Local
	ci := fb.config.Environment == config.CI

	if local || ci {
		return
	}

	body := bytes.NewBuffer(logBytes)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, fb.config.LogConfig.Endpoint, body)
	if err != nil {
		fb.queueForRetry(logBytes)
		return
	}

	logConfig := fb.config.LogConfig

	// set basic auth creds if auth is enabled
	if logConfig.AuthMethod != "" {
		req.SetBasicAuth(logConfig.Username, logConfig.Password)
	}

	req.Header.Set("content-type", "application/json")
	resp, err := fb.client.Do(req)
	if err != nil {
		fb.queueForRetry(logBytes)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		fb.queueForRetry(logBytes)
	}
}

func (fb *fluentBit) queueForRetry(logBytes []byte) {
	var msgID string
	id, err := uuid.NewRandom()
	if err != nil {
		msgID = time.Now().Format(time.RFC3339Nano)
	} else {
		msgID = id.String()
	}

	fb.retryMessages[msgID] = retryLogMsg{
		content: logBytes,
		count:   0,
		done:    false,
	}

}

func (fb *fluentBit) retry() {
	ticker := time.NewTicker(time.Second * 5) // sort of retry every 5 seconds

	// lets not do more than 5 req/second just to not flood our free instance of grafana cloud
	for range ticker.C {
		for id, logMsg := range fb.retryMessages {
			fb.retrier(logMsg.content, id)
		}
	}
}

func (fb *fluentBit) retrier(logBytes []byte, id string) {
	// TODO - (@jay-dee7) what to do then? maybe have a different way to ship these logs? like via promtail?
	if msg, ok := fb.retryMessages[id]; ok && msg.count > 3 {
		delete(fb.retryMessages, id)
		return
	}

	body := bytes.NewBuffer(logBytes)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, fb.config.LogConfig.Endpoint, body)
	if err != nil {
		fb.queueForRetry(logBytes)
		return
	}

	req.Header.Set("content-type", "application/json")
	resp, err := fb.client.Do(req)
	if err != nil {
		fb.queueForRetry(logBytes)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusCreated {
		delete(fb.retryMessages, id)
		return
	}

	item := fb.retryMessages[id]
	item.count++
	fb.retryMessages[id] = item
}
