package main

import (
	"github.com/ant0ine/go-urlrouter"
	"github.com/wurkhappy/WH-Payments/handlers"
)

//order matters so most general should go towards the bottom
var router urlrouter.Router = urlrouter.Router{
	Routes: []urlrouter.Route{
		urlrouter.Route{
			PathExp: "/agreements/v/:id/payments",
			Dest: map[string]func(map[string]interface{}, []byte) ([]byte, error, int){
				"POST": handlers.CreatePaymentsByVersionID,
				"GET":  handlers.GetPaymentsByVersionID,
			},
		},
		urlrouter.Route{
			PathExp: "/payments/:id",
			Dest: map[string]func(map[string]interface{}, []byte) ([]byte, error, int){
				"PUT": handlers.UpdatePayment,
			},
		},
		urlrouter.Route{
			PathExp: "/payments/:id/action",
			Dest: map[string]func(map[string]interface{}, []byte) ([]byte, error, int){
				"POST": handlers.UpdateAction,
			},
		},
	},
}
