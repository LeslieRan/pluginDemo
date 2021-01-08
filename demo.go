package pluginDemo

import (
	"context"
	"fmt"
	"net/http"
)

type Config struct {
	//...
	Info string `json:"info,omitempty"`
}

func CreateConfig() *Config {
	fmt.Println("Call CreateConfig()........")
	return &Config{}
}

type Demo struct {
	next http.Handler
	name string
	//...
}

func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	//...
	fmt.Println("Call New()........")
	fmt.Printf("the name is: %s\n", name)
	fmt.Printf("the content of info in configuration is: %s\n", config.Info)

	return &Demo{
		next: next,
		name: name,
	}, nil
}

func (demo *Demo) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	//...
	fmt.Println("Call ServeHTTP()........")
	fmt.Printf("url is: %s\n", req.URL.String())
	demo.next.ServeHTTP(rw, req)
}
