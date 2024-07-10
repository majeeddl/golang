package sockets

import (
	"encoding/json"
	"fmt"

	"github.com/gofiber/contrib/socketio"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

// MessageObject Basic chat message object
type MessageObject struct {
	Data  string `json:"data"`
	From  string `json:"from"`
	Event string `json:"event"`
	To    string `json:"to"`
}

func InitializeSockets(app *fiber.App) error {

	// The key for the map is message.to
	clients := make(map[string]string)

	// create a new socketio instance

	// server := socketio.New(nil)

	// Multiple event handling supported
	socketio.On(socketio.EventConnect, func(ep *socketio.EventPayload) {
		fmt.Println(fmt.Sprintf("Connection event 1 - User: %s", ep.Kws.GetStringAttribute("user_id")))
	})

	// Custom event handling supported
	socketio.On("CUSTOM_EVENT", func(ep *socketio.EventPayload) {
		fmt.Println(fmt.Sprintf("Custom event - User: %s", ep.Kws.GetStringAttribute("user_id")))
		// --->

		// DO YOUR BUSINESS HERE

		// --->
	})

	// On message event
	socketio.On(socketio.EventMessage, func(ep *socketio.EventPayload) {

		fmt.Println(fmt.Sprintf("Message event - User: %s - Message: %s", ep.Kws.GetStringAttribute("user_id"), string(ep.Data)))

		message := MessageObject{}

		// Unmarshal the json message
		// {
		//  "from": "<user-id>",
		//  "to": "<recipient-user-id>",
		//  "event": "CUSTOM_EVENT",
		//  "data": "hello"
		//}
		err := json.Unmarshal(ep.Data, &message)
		if err != nil {
			fmt.Println(err)
			return
		}

		// Fire custom event based on some
		// business logic
		if message.Event != "" {
			ep.Kws.Fire(message.Event, []byte(message.Data))
		}

		// Emit the message directly to specified user
		err = ep.Kws.EmitTo(clients[message.To], ep.Data, socketio.TextMessage)
		if err != nil {
			fmt.Println(err)
		}
	})

	// On disconnect event
	socketio.On(socketio.EventDisconnect, func(ep *socketio.EventPayload) {
		// Remove the user from the local clients
		delete(clients, ep.Kws.GetStringAttribute("user_id"))
		fmt.Println(fmt.Sprintf("Disconnection event - User: %s", ep.Kws.GetStringAttribute("user_id")))
	})

	// On close event
	// This event is called when the server disconnects the user actively with .Close() method
	socketio.On(socketio.EventClose, func(ep *socketio.EventPayload) {
		// Remove the user from the local clients
		delete(clients, ep.Kws.GetStringAttribute("user_id"))
		fmt.Println(fmt.Sprintf("Close event - User: %s", ep.Kws.GetStringAttribute("user_id")))
	})

	// On error event
	socketio.On(socketio.EventError, func(ep *socketio.EventPayload) {
		fmt.Println(fmt.Sprintf("Error event - User: %s", ep.Kws.GetStringAttribute("user_id")))
	})

	app.Use("/ws", func(c *fiber.Ctx) error {
		// IsWebSocketUpgrade returns true if the client
		// requested upgrade to the WebSocket protocol.
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
		// return c.Next()
	})

	app.Get("/ws/:id", socketio.New(func(kws *socketio.Websocket) {

		// Retrieve the user id from endpoint
		userId := kws.Params("id")

		// Add the connection to the list of the connected clients
		// The UUID is generated randomly and is the key that allow
		// socketio to manage Emit/EmitTo/Broadcast
		clients[userId] = kws.UUID

		// Every websocket connection has an optional session key => value storage
		kws.SetAttribute("user_id", userId)

		//Broadcast to all the connected users the newcomer
		// kws.Broadcast([]byte(fmt.Sprintf("New user connected: %s and UUID: %s", userId, kws.UUID)), true, socketio.TextMessage)
		// //Write welcome message
		// kws.Emit([]byte(fmt.Sprintf("Hello user: %s with UUID: %s", userId, kws.UUID)), socketio.TextMessage)
		//kws.Emit([]byte(fmt.Sprintf("Hello user: %s", userId)), socketio.TextMessage)

		// socketio.EmitTo(kws.UUID, []byte(fmt.Sprintf("Hello user: %s", userId)), socketio.TextMessage)
		socketio.Broadcast([]byte(fmt.Sprintf("New user connected: %s and UUID: %s", userId, kws.UUID)))
	}))

	return nil
}
