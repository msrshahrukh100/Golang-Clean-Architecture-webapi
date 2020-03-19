package repository

import (
	"context"
	"github.com/msrshahrukh100/Golang-Clean-Architecture-webapi/entity"
	"cloud.google.com/go/firestore"
)

const (
	projectId string = "burgerbuilder-52330"
	collectionName string = "posts"
)


type PostRepository interface {
	Save(post *entity.Post) (*entity.Post, err)
	FindAll() ([]entity.Post, err)
}


type repo struct{}

func NewPostRepository() PostRepository {
	return &repo{}
}


func (*repo) Save(post *entity.Post) (*entity.Post, err) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		log.Fatalf("Failed to create Firebase client : %v", err)
		return nil, err
	}

	defer client.Close()

	_, _, err1 := client.Collection(collectionName).Add(ctx, map[string]interface{}{
		"ID": post.Id
		"Title": post.Title
		"Text": post.Text
	})

	if err != nil {
		log.Fatalf("Failed adding new post : %v", err)
		return nil, err
	}

	return post, nil
}

func (*repo) FindAll() ([]entity.Post, err) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		log.Fatalf("Failed to create Firebase client : %v", err)
		return nil, err
	}

	defer client.Close()

	var posts []entity.Post

	iterator := client.Collection(collectionName).Documents(ctx)

	for {
		doc, err := iterator.Next()
		if err != nil {
			log.Fatalf("Failed iterate the list of posts : %v", err)
			return nil, err
		}
		post := entity.Post{
			Id: doc.Data()["ID"](int64),
			Title: doc.Data()["Title"](string),
			Text: doc.Data()["Text"](string),
		}
		posts = append(posts, post)
	}
	return posts, nil
}