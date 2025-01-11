package context

type StateKey string

type Context map[string]any

const (
    Buffer StateKey = "buffer"
)

func NewContext() Context {
    return make(Context)
}

func (context *Context) 
