//go:build js && wasm

package http

import (
	"encoding/json"
	"fmt"
	"syscall/js"
	"time"
)

// Client représente un client HTTP pour WASM
type Client struct {
	BaseURL string
	Headers map[string]string
	Timeout time.Duration
}

// Response représente une réponse HTTP
type Response struct {
	StatusCode int
	Headers    map[string]string
	Body       []byte
	Error      error
}

// NewClient crée un nouveau client HTTP
func NewClient(baseURL string) *Client {
	return &Client{
		BaseURL: baseURL,
		Headers: make(map[string]string),
		Timeout: 30 * time.Second,
	}
}

// SetHeader ajoute un header
func (c *Client) SetHeader(key, value string) *Client {
	c.Headers[key] = value
	return c
}

// SetAuth configure l'authentification Bearer
func (c *Client) SetAuth(token string) *Client {
	c.SetHeader("Authorization", "Bearer "+token)
	return c
}

// SetContentType configure le Content-Type
func (c *Client) SetContentType(contentType string) *Client {
	c.SetHeader("Content-Type", contentType)
	return c
}

// SetTimeout configure le timeout
func (c *Client) SetTimeout(timeout time.Duration) *Client {
	c.Timeout = timeout
	return c
}

// buildURL construit l'URL complète
func (c *Client) buildURL(endpoint string) string {
	if endpoint[0] != '/' {
		endpoint = "/" + endpoint
	}
	return c.BaseURL + endpoint
}

// GET effectue une requête GET
func (c *Client) GET(endpoint string, queryParams ...map[string]string) *Response {
	url := c.buildURL(endpoint)

	// Ajouter les paramètres de requête
	if len(queryParams) > 0 && queryParams[0] != nil {
		params := "?"
		first := true
		for key, value := range queryParams[0] {
			if !first {
				params += "&"
			}
			params += key + "=" + value
			first = false
		}
		url += params
	}

	return c.makeRequest("GET", url, nil)
}

// POST effectue une requête POST
func (c *Client) POST(endpoint string, body interface{}, queryParams ...map[string]string) *Response {
	url := c.buildURL(endpoint)

	// Ajouter les paramètres de requête si nécessaire
	if len(queryParams) > 0 && queryParams[0] != nil {
		params := "?"
		first := true
		for key, value := range queryParams[0] {
			if !first {
				params += "&"
			}
			params += key + "=" + value
			first = false
		}
		url += params
	}

	return c.makeRequest("POST", url, body)
}

// PUT effectue une requête PUT
func (c *Client) PUT(endpoint string, body interface{}, queryParams ...map[string]string) *Response {
	url := c.buildURL(endpoint)

	// Ajouter les paramètres de requête si nécessaire
	if len(queryParams) > 0 && queryParams[0] != nil {
		params := "?"
		first := true
		for key, value := range queryParams[0] {
			if !first {
				params += "&"
			}
			params += key + "=" + value
			first = false
		}
		url += params
	}

	return c.makeRequest("PUT", url, body)
}

// PATCH effectue une requête PATCH
func (c *Client) PATCH(endpoint string, body interface{}, queryParams ...map[string]string) *Response {
	url := c.buildURL(endpoint)

	// Ajouter les paramètres de requête si nécessaire
	if len(queryParams) > 0 && queryParams[0] != nil {
		params := "?"
		first := true
		for key, value := range queryParams[0] {
			if !first {
				params += "&"
			}
			params += key + "=" + value
			first = false
		}
		url += params
	}

	return c.makeRequest("PATCH", url, body)
}

// DELETE effectue une requête DELETE
func (c *Client) DELETE(endpoint string, queryParams ...map[string]string) *Response {
	url := c.buildURL(endpoint)

	// Ajouter les paramètres de requête
	if len(queryParams) > 0 && queryParams[0] != nil {
		params := "?"
		first := true
		for key, value := range queryParams[0] {
			if !first {
				params += "&"
			}
			params += key + "=" + value
			first = false
		}
		url += params
	}

	return c.makeRequest("DELETE", url, nil)
}

// makeRequest effectue la requête HTTP via JavaScript
func (c *Client) makeRequest(method, url string, body interface{}) *Response {
	// Créer les options de la requête
	options := js.Global().Get("Object").New()
	options.Set("method", method)

	// Ajouter les headers
	headers := js.Global().Get("Object").New()
	for key, value := range c.Headers {
		headers.Set(key, value)
	}
	options.Set("headers", headers)

	// Ajouter le corps si nécessaire
	if body != nil {
		switch v := body.(type) {
		case string:
			options.Set("body", v)
		case []byte:
			options.Set("body", string(v))
		default:
			jsonBody, err := json.Marshal(body)
			if err != nil {
				return &Response{Error: fmt.Errorf("erreur de sérialisation JSON: %w", err)}
			}
			options.Set("body", string(jsonBody))
			if c.Headers["Content-Type"] == "" {
				headers.Set("Content-Type", "application/json")
			}
		}
	}

	// Canal pour recevoir la réponse
	responseChan := make(chan *Response, 1)

	// Faire l'appel fetch
	promise := js.Global().Call("fetch", url, options)

	// Gérer la promesse
	promise.Call("then", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		response := args[0]

		// Créer la réponse
		resp := &Response{
			StatusCode: response.Get("status").Int(),
			Headers:    make(map[string]string),
		}

		// Lire le texte de la réponse
		textPromise := response.Call("text")
		textPromise.Call("then", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			text := args[0].String()
			resp.Body = []byte(text)
			responseChan <- resp
			return nil
		}))

		return nil
	})).Call("catch", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		errMsg := "Erreur de requête"
		if len(args) > 0 {
			errMsg = args[0].String()
		}
		responseChan <- &Response{Error: fmt.Errorf("erreur de requête: %s", errMsg)}
		return nil
	}))

	// Attendre la réponse avec timeout
	select {
	case resp := <-responseChan:
		return resp
	case <-time.After(c.Timeout):
		return &Response{Error: fmt.Errorf("timeout de la requête")}
	}
}

// IsSuccess vérifie si la réponse indique un succès
func (r *Response) IsSuccess() bool {
	return r.StatusCode >= 200 && r.StatusCode < 300
}

// JSON désérialise le corps de la réponse en JSON
func (r *Response) JSON(target interface{}) error {
	if r.Error != nil {
		return r.Error
	}

	if len(r.Body) == 0 {
		return fmt.Errorf("corps de réponse vide")
	}

	return json.Unmarshal(r.Body, target)
}

// String retourne le corps de la réponse sous forme de string
func (r *Response) String() string {
	if r.Error != nil {
		return ""
	}
	return string(r.Body)
}
