package main

import (
	"fmt"

	g "github.com/ghostemeow/goshikimori"
	"github.com/ghostemeow/goshikimori/constants"
)

func config() *g.Configuration {
	return g.SetConfiguration(
		"APPLICATION_NAME",
		"PRIVATE_KEY",
	)
}

func searchComments() {
	c := config()

	// Search comments of a topic.
	o := &g.Options{
		Commentable_id:   368370,
		Commentable_type: constants.COMMENTABLE_TYPE_TOPIC,
		Page:             1,
		Limit:            10,
	}
	comments, status, err := c.SearchComments(o)
	if status != 200 || err != nil {
		fmt.Println(status, err)
		return
	}
	for _, v := range comments {
		comment, status, err := c.ReadComment(v.Id)
		if status != 200 || err != nil {
			fmt.Println(status, err)
			return
		}
		fmt.Println(comment.Id, comment.Body, comment.User.Nickname)
	}
}

func createUpdateDeleteComment() {
	c := config()

	// Create a comment.
	comment, status, err := c.CreateComment(g.CommentParams{
		Body:             "comment body",
		Commentable_id:   368370,
		Commentable_type: constants.COMMENTABLE_TYPE_TOPIC,
	}, false)
	if status != 201 || err != nil {
		fmt.Println(status, err)
		return
	}
	fmt.Println(comment.Id, comment.Body, comment.User.Nickname)

	// Update the created comment.
	update, status, err := c.UpdateComment(comment.Id, g.CommentParams{Body: "new body"})
	if status != 200 || err != nil {
		fmt.Println(status, err)
		return
	}
	fmt.Println(update.Id, update.Body)

	// Delete the created comment.
	delete_comment, status, err := c.DeleteComment(comment.Id)
	if status != 200 || err != nil {
		fmt.Println(status, err)
		return
	}
	fmt.Println(delete_comment.Notice)
}

func main() {
	searchComments()
	createUpdateDeleteComment()
}
