package handlers

// Handler is a interface for handlers to describe otty
type Handler interface {
	GetName() string
	GetData() []byte
	SetName([]byte)
	SetData([]byte)
}

// InitHandlers initialize basic handlers
func InitHandlers() map[string]Handler {
	routeHandler := &Route{Name: []byte("Route")}
	dataHandler := &Data{Name: []byte("Data")}

	var handlersToReturn map[string]Handler

	handlersToReturn[routeHandler.GetName()] = routeHandler
	handlersToReturn[dataHandler.GetName()] = dataHandler

	return handlersToReturn
}

// FindHandlerByName returns handler which name is equal
func FindHandlerByName(handlersIn map[string]Handler, name string) Handler {
	if handlersIn[name] != nil {
		return handlersIn[name]
	}

	return nil
}
