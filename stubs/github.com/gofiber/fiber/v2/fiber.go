package fiber

import (
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "strings"
)

type Map map[string]interface{}

type Ctx struct {
    Request *http.Request
    Response *httptest.ResponseRecorder
    params map[string]string
    status int
}

func (c *Ctx) Params(key string) string {
    if c.params == nil {
        return ""
    }
    return c.params[key]
}

func (c *Ctx) Query(key string) string {
    return c.Request.URL.Query().Get(key)
}

func (c *Ctx) Status(code int) *Ctx {
    c.status = code
    return c
}

func (c *Ctx) JSON(v interface{}) error {
    if c.status == 0 {
        c.status = StatusOK
    }
    c.Response.WriteHeader(c.status)
    return json.NewEncoder(c.Response).Encode(v)
}

type Handler func(*Ctx) error

type route struct {
    method  string
    path    string
    handler Handler
}

type App struct {
    routes []route
}

func New() *App {
    return &App{}
}

func (a *App) Get(path string, h Handler) {
    a.routes = append(a.routes, route{method: "GET", path: path, handler: h})
}

func (a *App) Use(mw ...Handler) {}

func (a *App) Listen(addr string) error { return nil }

func (a *App) Test(req *http.Request) (*http.Response, error) {
    for _, r := range a.routes {
        if r.method != req.Method {
            continue
        }
        params, ok := match(r.path, req.URL.Path)
        if !ok {
            continue
        }
        ctx := &Ctx{Request: req, Response: httptest.NewRecorder(), params: params}
        if err := r.handler(ctx); err != nil {
            return nil, err
        }
        return ctx.Response.Result(), nil
    }
    resp := httptest.NewRecorder()
    resp.WriteHeader(StatusNotFound)
    return resp.Result(), nil
}

func match(pattern, path string) (map[string]string, bool) {
    pParts := strings.Split(strings.Trim(pattern, "/"), "/")
    pathParts := strings.Split(strings.Trim(path, "/"), "/")
    if len(pParts) != len(pathParts) {
        return nil, false
    }
    params := map[string]string{}
    for i, part := range pParts {
        if strings.HasPrefix(part, ":") {
            params[part[1:]] = pathParts[i]
        } else if part != pathParts[i] {
            return nil, false
        }
    }
    return params, true
}

type Error struct {
    Code    int
    Message string
}

func (e *Error) Error() string { return e.Message }

func NewError(code int, message string) error {
    return &Error{Code: code, Message: message}
}

const (
    StatusOK                  = http.StatusOK
    StatusBadRequest          = http.StatusBadRequest
    StatusInternalServerError = http.StatusInternalServerError
    StatusNotFound            = http.StatusNotFound
)
