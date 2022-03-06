package handlers

import (
	"net/http"
	"strconv"

	"simple-blog/models"
	"simple-blog/utils"

	"github.com/gin-gonic/gin"
)

func ShowIndexPage(c *gin.Context) {
	articles := models.GetAllArticles()

	// Call the render function with the name of the template to render
	utils.Render(c, gin.H{
		"title":   "Home Page",
		"payload": articles}, "index.html")
}

func ShowArticleCreationPage(c *gin.Context) {
	// Call the render function with the name of the template to render
	utils.Render(c, gin.H{
		"title": "Create New Article"}, "create-article.html")
}

func GetArticle(c *gin.Context) {
	// Check if the article ID is valid
	if articleID, err := strconv.Atoi(c.Param("article_id")); err == nil {
		// Check if the article exists
		if article, err := models.GetArticleByID(articleID); err == nil {
			// Call the render function with the title, article and the name of the
			// template
			utils.Render(c, gin.H{
				"title":   article.Title,
				"payload": article}, "article.html")

		} else {
			// If the article is not found, abort with an error
			c.AbortWithError(http.StatusNotFound, err)
		}

	} else {
		// If an invalid article ID is specified in the URL, abort with an error
		c.AbortWithStatus(http.StatusNotFound)
	}
}

func CreateArticle(c *gin.Context) {
	// Obtain the POSTed title and content values
	title := c.PostForm("title")
	content := c.PostForm("content")

	if a, err := models.CreateNewArticle(title, content); err == nil {
		// If the article is created successfully, show success message
		utils.Render(c, gin.H{
			"title":   "Submission Successful",
			"payload": a}, "submission-successful.html")
	} else {
		// if there was an error while creating the article, abort with an error
		c.AbortWithStatus(http.StatusBadRequest)
	}
}
