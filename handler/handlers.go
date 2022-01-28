package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/techagent/ceeyuapp/models"
	"log"
	"strconv"
	"sync"
)

var count = 0

var article = []models.Article{
	{
		ID:      "1",
		Title:   "Cybersecurity",
		Author:  "Ken",
		Content: "Cyber Security",
	},
	{
		ID:      "2",
		Title:   "Software Development",
		Author:  "Kenneth",
		Content: "Computer Science",
	},
}

func Index() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.HTML(200, "index.gohtml", gin.H{"articles": article})
	}
}

func View() gin.HandlerFunc {
	return func(c *gin.Context) {
		articleID, _ := strconv.Atoi(c.Param("id"))
		article := article[articleID-1]
		c.HTML(200, "view.gohtml", gin.H{"article": article})
	}
}

func Admin() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(200, "admin.gohtml", gin.H{"count": count})
	}
}
func Counter() gin.HandlerFunc {
	return func(c *gin.Context) {
		var wg sync.Mutex
		counter := &struct {
			Count int	`json:"count"`
		}{}
		err := c.BindJSON(counter)
		wg.Lock()

		count++

		wg.Unlock()
		if err != nil {
			log.Println(err)
		}

		c.JSON(200, gin.H{"count": count})
	}
}

//func Comment() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		var comment models.Comment
//		articleComment := c.PostForm("comment")
//
//	}
//}
