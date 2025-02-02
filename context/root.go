package context

type StateKey string

type Context map[StateKey]any

const (
    RenderQueue StateKey = "render-queue"
)

func InitContext() Context {
    return make(Context)
}

func (context *Context) GetValue(key StateKey) (any, bool) {
    value, ok := (*context)[key]
    return value, ok
}

func (context *Context) SetValue(key StateKey, value any) {
    (*context)[key] = value
}
