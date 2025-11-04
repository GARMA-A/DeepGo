package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Post struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func DeletePost(c *gin.Context) {
	id := c.Param("id")
	intid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ID",
		})
		return
	}

	_, err = db.Exec("DELETE FROM posts WHERE id = ?", intid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Good practice: check if you actually deleted something
	// res.RowsAffected()

	c.JSON(http.StatusOK, gin.H{
		"message": "Delete Post",
	})
}

func UpdatePost(c *gin.Context) {
	idStr := c.Param("id")
	intid, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ID",
		})
		return
	}

	var post Post
	if err := c.BindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	_, err = db.Exec("UPDATE posts SET title = ?, content = ? WHERE id = ?", post.Title, post.Content, intid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Post updated successfully",
	})
}

func CreatePost(c *gin.Context) {
	var post Post
	if err := c.BindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid input",
		})
		return
	}

	_, err := db.Exec("INSERT INTO posts (id, title, content) VALUES (?, ?, ?)", post.ID, post.Title, post.Content)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Create Post",
		"post":    post,
	})
}

func GetPosts(c *gin.Context) {
	rows, err := db.Query("SELECT id, title, content FROM posts")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	defer rows.Close()

	posts := []Post{}

	for rows.Next() {
		var post Post

		err = rows.Scan(&post.ID, &post.Title, &post.Content)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to scan row: " + err.Error(),
			})
			return
		}

		posts = append(posts, post)
	}

	if err = rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error during rows iteration: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, posts)
}
