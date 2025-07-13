//go:build js && wasm

package http

import "time"

// Instance globale du client HTTP
var globalClient *Client

// Init initialise le client HTTP global
// À appeler UNE SEULE FOIS dans main.go
func Init(baseURL string) *Client {
	globalClient = NewClient(baseURL)
	return globalClient
}

// GetClient retourne l'instance globale du client HTTP
// Utilisé partout dans l'application
func GetClient() *Client {
	if globalClient == nil {
		// Client par défaut si pas initialisé
		globalClient = NewClient("https://jsonplaceholder.typicode.com")
		globalClient.SetContentType("application/json")
	}
	return globalClient
}

// Configure permet de configurer le client global
func Configure(configureFn func(*Client)) {
	client := GetClient()
	configureFn(client)
}

// ============ MÉTHODES GLOBALES DE CONVENANCE ============

// GET effectue un GET avec le client global
func GET(endpoint string, queryParams ...map[string]string) *Response {
	return GetClient().GET(endpoint, queryParams...)
}

// POST effectue un POST avec le client global
func POST(endpoint string, body interface{}, queryParams ...map[string]string) *Response {
	return GetClient().POST(endpoint, body, queryParams...)
}

// PUT effectue un PUT avec le client global
func PUT(endpoint string, body interface{}, queryParams ...map[string]string) *Response {
	return GetClient().PUT(endpoint, body, queryParams...)
}

// PATCH effectue un PATCH avec le client global
func PATCH(endpoint string, body interface{}, queryParams ...map[string]string) *Response {
	return GetClient().PATCH(endpoint, body, queryParams...)
}

// DELETE effectue un DELETE avec le client global
func DELETE(endpoint string, queryParams ...map[string]string) *Response {
	return GetClient().DELETE(endpoint, queryParams...)
}

// ============ CONFIGURATIONS PRÉ-DÉFINIES ============

// InitJSONPlaceholder configure le client global pour JSONPlaceholder
func InitJSONPlaceholder() {
	globalClient = NewClient("https://jsonplaceholder.typicode.com")
	globalClient.SetContentType("application/json")
	globalClient.SetHeader("User-Agent", "Stencil-Framework/1.0")
}

// InitGitHub configure le client global pour GitHub API
func InitGitHub(token string) {
	globalClient = NewClient("https://api.github.com")
	globalClient.SetAuth(token)
	globalClient.SetContentType("application/json")
	globalClient.SetHeader("Accept", "application/vnd.github.v3+json")
	globalClient.SetHeader("User-Agent", "Stencil-Framework/1.0")
}

// InitCustom configure le client global avec une configuration personnalisée
func InitCustom(baseURL string, configureFn func(*Client)) {
	globalClient = NewClient(baseURL)
	if configureFn != nil {
		configureFn(globalClient)
	}
}

// ============ BUILDERS POUR CONFIGURATIONS COURANTES ============

// JSONPlaceholderClient crée un client pré-configuré pour JSONPlaceholder
func JSONPlaceholderClient() *Client {
	client := NewClient("https://jsonplaceholder.typicode.com")
	client.SetContentType("application/json")
	client.SetHeader("User-Agent", "Stencil-Framework/1.0")
	return client
}

// GitHubClient crée un client pré-configuré pour GitHub
func GitHubClient(token string) *Client {
	client := NewClient("https://api.github.com")
	client.SetAuth(token)
	client.SetContentType("application/json")
	client.SetHeader("Accept", "application/vnd.github.v3+json")
	client.SetHeader("User-Agent", "Stencil-Framework/1.0")
	return client
}

// CustomClient crée un client avec configuration personnalisée
func CustomClient(baseURL string, configureFn func(*Client)) *Client {
	client := NewClient(baseURL)
	if configureFn != nil {
		configureFn(client)
	}
	return client
}

// ============ UTILITAIRES ============

// SetGlobalTimeout configure le timeout du client global
func SetGlobalTimeout(timeout time.Duration) {
	GetClient().SetTimeout(timeout)
}

// SetGlobalHeader ajoute un header au client global
func SetGlobalHeader(key, value string) {
	GetClient().SetHeader(key, value)
}

// SetGlobalAuth configure l'authentification du client global
func SetGlobalAuth(token string) {
	GetClient().SetAuth(token)
}
