package main

import (
	"internal/reflectlite"
	"regexp"
)

type controllerInfo struct {
	regex          *regexp.Regexp
	params         map[int]string
	controllerType reflectlite.Type
}

type ControllerRegistory struct {
	routers     []*controllerInfo
	Application *App
}

func main() {

}
