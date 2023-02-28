package main

import (
	"py_gomall/v2/user_web/user_router"
	"py_gomall/v2/user_web/user_run"
)

func main() {
	r := user_router.NewRouter()
	user_run.R = r
	user_run.Run(r)
}
