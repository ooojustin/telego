package telegram

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type WebhookInfo struct {
	URL                          string   `json:"url"`
	HasCustomCertificate         bool     `json:"has_custom_certificate"`
	PendingUpdateCount           int      `json:"pending_update_count"`
	IpAddress                    string   `json:"ip_address"`
	LastErrorDate                int      `json:"last_error_date"`
	LastErrorMessage             string   `json:"last_error_message"`
	LastSynchronizationErrorDate int      `json:"last_synchronization_error_date"`
	MaxConnections               int      `json:"max_connections"`
	AllowedUpdates               []string `json:"allowed_updates"`
}

type SetWebhookResponse struct {
	TelegramResponse
	Result bool `json:"result"`
}

type GetWebhookInfoResponse struct {
	Ok     bool        `json:"ok"`
	Result WebhookInfo `json:"result"`
}

func (tc *TelegramClient) SetWebhook(url string, allowedUpdates []string, secretToken string) error {
	var resp *http.Response
	var err error

	req := IMap{"url": url}

	if len(allowedUpdates) > 0 {
		req["allowed_updates"] = allowedUpdates
	}

	if len(secretToken) > 0 {
		req["secret_token"] = secretToken
	}

	resp, err = tc.SendRequest(POST, "setWebhook", &req)
	if err != nil {
		return fmt.Errorf("SetWebhook failed: %w", err)
	}

	defer resp.Body.Close()

	var data SetWebhookResponse
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return fmt.Errorf("SetWebhook failed: %w", err)
	}

	if data.Ok {
		return nil
	} else if len(data.Description) > 0 {
		err = errors.New(data.Description)
		return fmt.Errorf("SetWebhook failed: %w", err)
	}

	return fmt.Errorf("SetWebhook failed: %w", UnknownError)
}

func (tc *TelegramClient) GetWebhookInfo() (*WebhookInfo, error) {
	var resp *http.Response
	var err error

	resp, err = tc.SendRequest(GET, "getWebhookInfo", nil)
	if err != nil {
		return nil, fmt.Errorf("GetWebhookInfo failed: %w", err)
	}

	defer resp.Body.Close()

	var data GetWebhookInfoResponse
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, fmt.Errorf("GetWebhookInfoResponse failed: %w", err)
	}

	if data.Ok {
		return &data.Result, nil
	}

	return nil, fmt.Errorf("GetWebhookInfo failed: %w", UnknownError)
}
