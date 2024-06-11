package gee

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type H map[string]interface{}

type Context struct {
	Writer http.ResponseWriter
	Req    *http.Request

	Method string
	Path   string

	StatusCode int
}

func newContext(writer http.ResponseWriter, req *http.Request) *Context {
	return &Context{writer, req, req.Method, req.URL.Path, 0}
}

// PostForm Post Param
func (c *Context) PostForm(key string) string {
	return c.Req.FormValue(key)
}

// Query Get Param
func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

func (c *Context) SetStatus(status int) {
	c.StatusCode = status
	c.Writer.WriteHeader(status)
}

func (c *Context) SetHeader(key string, value string) {
	c.Writer.Header().Set(key, value)
}

func (c *Context) Data(status int, data []byte) {
	_, err := c.Writer.Write(data)
	if err != nil {
		fmt.Println(err)
	}
}
func (c *Context) String(status int, format string, values ...interface{}) {
	c.SetStatus(status)
	c.SetHeader("Content-Type", "text/plain")
	_, err := fmt.Fprintf(c.Writer, format, values...)
	if err != nil {
		fmt.Println(err)
	}
}

func (c *Context) JSON(status int, data interface{}) {
	c.SetStatus(status)
	c.SetHeader("Content-Type", "application/json")
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(data); err != nil {
		fmt.Println(err)
	}
}
func (c *Context) HTML(status int, html string) {
	c.SetStatus(status)
	c.SetHeader("Content-Type", "text/html")
	_, err := c.Writer.Write([]byte(html))
	if err != nil {
		fmt.Println(err)
	}
}
