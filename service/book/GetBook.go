package service

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/oranges0da/entertainment-api/model"
	"github.com/oranges0da/entertainment-api/service/utils"
)

func GetMovieByTitle(client *http.Client, title string) model.Book {
	key := os.Getenv("BOOK_KEY") // get key to book api from env
	var url string = "https://www.googleapis.com/books/v1/volumes?q=" + title + "&key=" + key + "&intitle"
	log.Print(url)

	req, err := http.NewRequest("GET", url, nil) // create request, not send it

	if err != nil {
		log.Fatal(err)
	}

	res, err := client.Do(req) // send request

	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close() // close the body after reading

	data := utils.UnpackBook(res)

	return data
}

func GetBookByAuthor(client *http.Client, author string) model.Book {
	key := os.Getenv("BOOK_KEY") // get key to book api from env
	var url string = "https://www.googleapis.com/books/v1/volumes?q=" + author + "&key=" + key + "&inauthor"
	log.Print(url)

	req, err := http.NewRequest("GET", url, nil) // create request, not send it

	if err != nil {
		log.Fatal(err)
	}

	res, err := client.Do(req) // send request

	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close() // close the body after reading

	data := utils.UnpackBook(res)

	return data
}

func GetBook() gin.HandlerFunc {
	return func(c *gin.Context) {
		titleQuery := c.Query("title")                        // get title from url
		bookTitle := strings.ReplaceAll(titleQuery, " ", "+") // replace space with + for annoying query

		authorQuery := c.Query("author")                    // get author from url
		author := strings.ReplaceAll(authorQuery, " ", "+") // replace space with + for annoying query

		if bookTitle != "" {
			bookData := GetMovieByTitle(http.DefaultClient, bookTitle)

			c.JSON(http.StatusOK, gin.H{
				"data":   bookData,
				"errors": []string{},
			})
		} else if author != "" {
			bookData := GetBookByAuthor(http.DefaultClient, author)

			c.JSON(http.StatusOK, gin.H{
				"data":   bookData,
				"errors": []string{},
			})
		}
	}
}
