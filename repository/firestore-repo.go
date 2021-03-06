package repository

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	"github.com/msrshahrukh100/Golang-Clean-Architecture-webapi/entity"
	"google.golang.org/api/iterator"
)

const (
	projectID      string = "burgerbuilder-52330"
	collectionName string = "posts"
)

// PostRepository ..
type PostRepository struct{}

// NewFirestoreRepository sdf
func NewFirestoreRepository() Repository {
	return &PostRepository{}
}

func (*PostRepository) Save(post *entity.Post) (*entity.Post, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create Firebase client : %v", err)
		return nil, err
	}

	defer client.Close()

	_, _, err1 := client.Collection(collectionName).Add(ctx, map[string]interface{}{"ID": post.Id, "Title": post.Title, "Text": post.Text})

	if err1 != nil {
		log.Fatalf("Failed adding new post : %v", err)
		return nil, err
	}

	return post, nil
}

func (*PostRepository) FindAll() ([]entity.Post, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create Firebase client : %v", err)
		return nil, err
	}

	defer client.Close()

	var posts []entity.Post

	it := client.Collection(collectionName).Documents(ctx)

	for {
		doc, err := it.Next()
		if err == iterator.Done {
			break
		}

		if err != nil {
			log.Fatalf("Failed iterate the list of posts : %v", err)
			return nil, err
		}
		post := entity.Post{
			Id:    doc.Data()["ID"].(int64),
			Title: doc.Data()["Title"].(string),
			Text:  doc.Data()["Text"].(string),
		}
		posts = append(posts, post)
	}
	return posts, nil
}
