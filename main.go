package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

type fileInfo struct {
	Name    string `json:"name"`
	Size    int64  `json:"size"`
	ModTime int64  `json:"mod_time"`
}

func main() {
	r := gin.Default()
	r.GET("/files", func(c *gin.Context) {
		files, _ := os.ReadDir("./data")
		data := make([]fileInfo, len(files))
		for index, entry := range files {
			if entry.IsDir() {
				continue
			}
			fileDetail, _ := entry.Info()
			data[index] = fileInfo{
				Name:    entry.Name(),
				Size:    fileDetail.Size(),
				ModTime: fileDetail.ModTime().Unix(),
			}
		}
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "query success",
			"data":    data,
		})
	})

	r.POST("file", func(c *gin.Context) {
		object, _ := c.FormFile("object")
		objectName := c.Query("objectName")
		object.Filename = objectName
		savePath := fmt.Sprintf("data/%s", object.Filename)
		err := c.SaveUploadedFile(object, savePath)
		if err != nil {
			fmt.Println(err.Error())
			c.JSON(http.StatusOK, gin.H{
				"code":    400,
				"message": fmt.Sprintf("err upload object: %v", err),
			})
			return
		}
		fileDetail, _ := os.Stat(savePath)
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": fmt.Sprintf("'%s' uploaded!", object.Filename),
			"data": fileInfo{
				Name:    objectName,
				Size:    fileDetail.Size(),
				ModTime: fileDetail.ModTime().Unix(),
			},
		})
	})

	r.GET("/file", func(c *gin.Context) {
		objId := c.Query("name")
		if objId == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    "400",
				"message": "objId为空",
			})
			return
		}
		objPath := fmt.Sprintf("data/%s", objId)
		_, err := os.Stat(objPath)
		if err != nil {
			if os.IsNotExist(err) {
				c.JSON(http.StatusBadRequest, gin.H{
					"code":    "401",
					"message": fmt.Sprintf("can't find object with id %s", objId),
				})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{
					"code":    "500",
					"message": fmt.Sprintf("server internal error: %v", err),
				})
			}
			return
		}
		c.File(objPath)
	})

	r.Run(":8080")
}
