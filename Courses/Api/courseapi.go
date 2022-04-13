package main

import (
	"flag"
	"fmt"
	"net/http"

	"onlineEudcation/Courses/Api/internal/config"
	"onlineEudcation/Courses/Api/internal/handler"
	"onlineEudcation/Courses/Api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/courseapi.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf, rest.WithCustomCors(nil, notAllowedFn, "http://localhost:8080"))
	defer server.Stop()

	handler.RegisterHandlers(server, ctx)
	//test

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}

func notAllowedFn(w http.ResponseWriter) {
	w.Header().Add("Access-Control-Allow-Headers", "x-token")
}
