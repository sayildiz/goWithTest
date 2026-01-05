package main

import (
	"log"
	"os"
	blogposts "sayildiz/blogposts"
)

func main() {
	posts, err := blogposts.NewPostsFromFS(os.DirFS("posts"))

	if err != nil {
		log.Fatal()
	}
	log.Println(posts)
}
