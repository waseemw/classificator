package post_model

import (
	"fmt"
	"github.com/cdipaolo/goml/base"
	"github.com/cdipaolo/goml/text"
)

func train(postModel *PostModel, posts []Post) *text.NaiveBayes {
	for i, p := range posts {
		if _, ok := postModel.groupsNameToId[p.GroupName]; !ok {
			postModel.groupsNameToId[p.GroupName] = len(postModel.groupsNameToId)
		}
		posts[i].GroupId = postModel.groupsNameToId[p.GroupName]
	}
	stream := make(chan base.TextDatapoint, 100)
	errors := make(chan error)
	model := text.NewNaiveBayes(stream, uint8(len(postModel.groupsNameToId)), base.OnlyWords)
	go model.OnlineLearn(errors)

	for _, p := range posts {
		stream <- base.TextDatapoint{
			X: p.Content,
			Y: uint8(p.GroupId),
		}
	}
	close(stream)
	for {
		err, _ := <-errors
		if err != nil {
			fmt.Printf("Error passed: %v", err)
		} else {
			// training is done!
			break
		}
	}
	return model
}
