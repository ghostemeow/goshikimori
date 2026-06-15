package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"

	g "github.com/ghostemeow/goshikimori"
	"github.com/ghostemeow/goshikimori/constants"
)

func config() *g.Configuration {
	return g.SetConfiguration(
		"APPLICATION_NAME",
		"PRIVATE_KEY",
	)
}

func readMessage() {
	c := config()
	// method #1. find my id and read inbox.
	fast, status, err := c.FastIdUser("arctica")
	if status != 200 || err != nil {
		fmt.Println(status, err)
		return
	}
	messages, status, err := fast.UserMessages(&g.Options{
		Type: constants.MESSAGE_TYPE_INBOX, Page: 1, Limit: 10,
	})
	if status != 200 || err != nil {
		fmt.Println(err)
		return
	}
	for _, v := range messages {
		message, status, err := c.ReadMessage(v.Id)
		if status != 200 || err != nil {
			fmt.Println(status, err)
			return
		}
		fmt.Println(
			message.Body,
			message.From.Id, message.From.Nickname,
			message.To.Id, message.To.Nickname,
		)
	}
	// method #2. read the dialogs, find the kind and id.
	d, status, err := c.Dialogs()
	if status != 200 || err != nil {
		fmt.Println(status, err)
		return
	}
	for _, v := range d {
		if v.Message.Kind == constants.MESSAGE_TYPE_PRIVATE && v.Message.Body != "" {
			message, status, err := c.ReadMessage(v.Message.Id)
			if status != 200 || err != nil {
				fmt.Println(status, err)
				return
			}
			fmt.Println(
				message.Body,
				message.From.Id, message.From.Nickname,
				message.To.Id, message.To.Nickname,
			)
		}
	}
	// method #3. friend nickname search.
	// In this case all messages from the last to the first are found.
	f, status, err := c.FastIdUser("morr")
	if status != 200 || err != nil {
		fmt.Println(status, err)
		return
	}
	sd, status, err := f.SearchDialogs()
	if status != 200 || err != nil {
		fmt.Println(status, err)
		return
	}
	for _, v := range sd {
		fmt.Println(
			v.Body,
			v.From.Id, v.From.Nickname,
			v.To.Id, v.To.Nickname,
		)
	}
}

func sendMessage() {
	c := config()
	me, status, err := c.WhoAmi()
	if status != 200 || err != nil {
		fmt.Println(status, err)
		return
	}
	to, status, err := c.FastIdUser("morr")
	if status != 200 || err != nil {
		fmt.Println(status, err)
		return
	}
	message := "test message from API"
	send, status, err := c.SendMessage(me.Id, to.Id, message)
	if status != 201 || err != nil {
		fmt.Println(status, err)
		return
	}
	fmt.Println(send.Id, send.Body, send.From.Nickname, send.To.Nickname)
}

func changeMesage() {
	c := config()
	d, status, err := c.Dialogs()
	if status != 200 || err != nil {
		fmt.Println(status, err)
		return
	}
	for _, v := range d {
		if v.Message.Kind == constants.MESSAGE_TYPE_PRIVATE && v.Message.Body != "" {
			new_message := "changed the message from API :)"
			message, status, err := c.ChangeMessage(v.Message.Id, new_message)
			if status != 200 || err != nil {
				fmt.Println(status, err)
				return
			}
			fmt.Println(
				message.Body,
				message.From.Id, message.From.Nickname,
				message.To.Id, message.To.Nickname,
			)
		}
	}
}

func deleteMessage() {
	c := config()
	d, status, err := c.Dialogs()
	if status != 200 || err != nil {
		fmt.Println(status, err)
		return
	}
	for _, v := range d {
		if v.Message.Kind == constants.MESSAGE_TYPE_PRIVATE && v.Message.Body != "" {
			status, err := c.DeleteMessage(v.Message.Id)
			if status != 204 || err != nil {
				fmt.Println(status, err)
				return
			}
			fmt.Println(status) // only status 204 is returned
		}
	}
}

func idsToString(slice []int) string {
	var res bytes.Buffer
	for i := range slice {
		res.WriteString(strconv.Itoa(slice[i]))
		res.WriteString(",")
	}
	return strings.TrimSuffix(res.String(), ",")
}

func markReadUnreadMessages() {
	c := config()

	var count int

	fast, status, err := c.FastIdUser("arctica")
	if status != 200 || err != nil {
		fmt.Println(status, err)
		return
	}

	ids, status, err := fast.UnreadMessagesIds("messages")
	if status != 200 || err != nil {
		fmt.Println(err)
		return
	}

	messages, status, err := fast.UserMessages(&g.Options{
		Type: constants.MESSAGE_TYPE_INBOX, Page: 1, Limit: 10,
	})
	if status != 200 || err != nil {
		fmt.Println(err)
		return
	}
	for _, v := range messages {
		message, status, err := c.ReadMessage(v.Id)
		if status != 200 || err != nil {
			fmt.Println(status, err)
			return
		}
		// You can set all sorts of filters for easy array compilation.
		if message.From.Nickname == "morr" && !message.Read {
			ids[count] = message.Id
			count++
		}
	}

	read, err := c.MarkReadMessages(idsToString(ids), 1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(read)
}

func readAll() {
	c := config()
	read, err := c.ReadAllMessages("notifications")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(read)
}

func deleteAll() {
	c := config()
	del, err := c.DeleteAllMessages("notifications")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(del)
}

func main() {
	// not to use all requests at the same time!
	readMessage()
	sendMessage()
	changeMesage()
	deleteMessage()

	markReadUnreadMessages()
	readAll()
	deleteAll()
}
