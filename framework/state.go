package framework

// State represents application state with reactive updates
type State struct {
	data map[string]interface{}
	app  *App
}

// NewState creates a new state instance
func NewState(app *App) *State {
	return &State{
		data: make(map[string]interface{}),
		app:  app,
	}
}

// Set updates a state value and triggers a re-render
func (s *State) Set(key string, value interface{}) {
	s.data[key] = value
	if s.app != nil {
		s.app.Update()
	}
}

// Get retrieves a state value
func (s *State) Get(key string) interface{} {
	return s.data[key]
}

// GetString retrieves a state value as string
func (s *State) GetString(key string) string {
	if val, ok := s.data[key].(string); ok {
		return val
	}
	return ""
}

// GetInt retrieves a state value as int
func (s *State) GetInt(key string) int {
	if val, ok := s.data[key].(int); ok {
		return val
	}
	return 0
}

// GetBool retrieves a state value as bool
func (s *State) GetBool(key string) bool {
	if val, ok := s.data[key].(bool); ok {
		return val
	}
	return false
}

// Has checks if a key exists in state
func (s *State) Has(key string) bool {
	_, exists := s.data[key]
	return exists
}
