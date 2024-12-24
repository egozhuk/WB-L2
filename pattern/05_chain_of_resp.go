package main

import "fmt"

// Интерфейс обработчика
type Handler interface {
	SetNext(handler Handler)
	Handle(request string)
}

// Базовый обработчик
type BaseHandler struct {
	next Handler
}

func (h *BaseHandler) SetNext(handler Handler) {
	h.next = handler
}

func (h *BaseHandler) Handle(request string) {
	if h.next != nil {
		h.next.Handle(request)
	}
}

// Конкретные обработчики
type AuthHandler struct {
	BaseHandler
}

func (h *AuthHandler) Handle(request string) {
	if request == "auth" {
		fmt.Println("AuthHandler handled the request")
		return
	}
	h.BaseHandler.Handle(request)
}

type LogHandler struct {
	BaseHandler
}

func (h *LogHandler) Handle(request string) {
	fmt.Println("LogHandler logged the request")
	h.BaseHandler.Handle(request)
}

func main() {
	auth := &AuthHandler{}
	log := &LogHandler{}

	auth.SetNext(log)

	auth.Handle("auth")
	auth.Handle("log")
}
