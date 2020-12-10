package main

import (
	"fmt"

	"koa.go"
)

func main() {
	app := koa.New()

	app.Use("/", func(err error, ctx *koa.Context, next koa.NextCb) {
		fmt.Println("test1")
		next(err)
		fmt.Println("test1")
	})
	app.Use(func(err error, ctx *koa.Context, next koa.NextCb) {
		fmt.Println("test1")
		next(err)
		fmt.Println("test1")
	})

	app.Use(func(err error, ctx *koa.Context, next koa.NextCb) {
		fmt.Println("test2")
		next(nil)
		fmt.Println("test2")
	})
	err := app.Run(8080)
	if err != nil {
		fmt.Println(err)
	}
}
