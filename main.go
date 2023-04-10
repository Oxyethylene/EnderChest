package main

import (
	"fmt"
	"github.com/Oxyethylene/littlebox/logging"
	"github.com/Oxyethylene/littlebox/middleware"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"os"
	"path/filepath"
)

type fileInfo struct {
	Name    string `json:"name"`
	Size    int64  `json:"size"`
	ModTime int64  `json:"mod_time"`
}

var DataPath string

func init() {
	logging.InitLogger()
	const EnvDataPath = "LB_DATA_PATH"
	const DefaultDataPath = "./data"
	zap.S().Info("init started")
	zap.S().Infow(" try load data path from env",
		"env", EnvDataPath,
	)
	value, exist := os.LookupEnv(EnvDataPath)
	if !exist {
		value = DefaultDataPath
		zap.S().Infow("can't read data path from env, fallback to default",
			"env", EnvDataPath,
			"default", DefaultDataPath,
		)
	}
	path, err := filepath.Abs(value)
	if err != nil {
		zap.S().Fatal("can't located data path",
			zap.Error(err),
			"data_path", path,
		)
	}
	DataPath = path
	zap.S().Infow("data path inited",
		"data_path", path,
	)
}

func main() {
	defer zap.S().Sync()
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(
		middleware.GinLogger(),
		middleware.GinRecovery(true),
		middleware.GinCors(),
	)
	r.GET("/files", func(c *gin.Context) {
		zap.S().Infow("looking up files in data_path",
			"data_path", DataPath,
		)
		files, _ := os.ReadDir(DataPath)
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
		response := gin.H{
			"code":    200,
			"message": "query success",
			"data":    data,
		}
		zap.S().Infow("success response",
			"response", response,
		)
		c.JSON(http.StatusOK, response)
	})

	r.POST("file", func(c *gin.Context) {
		object, _ := c.FormFile("object")
		objectName := c.Query("objectName")
		object.Filename = objectName
		savePath := filepath.Join(DataPath, object.Filename)
		zap.S().Infow("attempted saving obj",
			"target_path", savePath,
		)
		err := c.SaveUploadedFile(object, savePath)
		if err != nil {
			zap.S().Error(err)
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
		objPath := filepath.Join(DataPath, objId)
		zap.S().Info(fmt.Sprintf("searching file with path %s", objPath))
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
