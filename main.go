package main

import (
	"fmt"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/upload", handleUpload)

	if err := r.Run(":8000"); err != nil {
		log.Fatal(err)
	}
}

func handleUpload(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	files := form.File["files"]
	for key, file := range files {
		err := saveFile(file, key)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "Files uploaded successfully"})
}

func saveFile(fileHeader *multipart.FileHeader, key int) error {
	fmt.Println(fileHeader.Filename)
	ext := strings.Split(fileHeader.Filename, ".")
	// if ext[1] != "png" {
	// 	return errors.New("pass a valid png")
	// }

	src, err := fileHeader.Open()
	if err != nil {
		return err
	}

	defer src.Close()

	// Create a new file in the desired destination folder
	dstPath := filepath.Join("./destination", ext[0]+strconv.Itoa(key)+"."+ext[1])
	dst, err := os.Create(dstPath)
	if err != nil {
		return err
	}
	defer dst.Close()

	// fmt.Printf("File saved to: %s\n", dstPath)
	return nil
}
