package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"os"
	"path/filepath"
)

type ObjectInfo struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Size    int64  `json:"size"`
	ModTime int64  `json:"mod_time"`
}

type ObjectApi struct {
	dbPath string
}

func NewObjectApi() *ObjectApi {
	const EnvDataPath = "LB_DATA_PATH"
	const DefaultDataPath = "./data"
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
	zap.S().Infow("success open data path",
		"data_path", path)
	return &ObjectApi{
		dbPath: path,
	}
}

func (a *ObjectApi) List(c *gin.Context) {
	zap.S().Infow("looking up files in data_path",
		"data_path", a.dbPath,
	)
	files, _ := os.ReadDir(a.dbPath)
	data := make([]ObjectInfo, len(files))
	for index, entry := range files {
		if entry.IsDir() {
			continue
		}
		fileDetail, _ := entry.Info()
		data[index] = ObjectInfo{
			Id:      entry.Name(),
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
}

func (a *ObjectApi) Add(c *gin.Context) {
	object, _ := c.FormFile("object")
	objectName := c.Query("objectName")
	if objectName == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    "400",
			"message": "objectName is required and should not be empty.",
		})
		return
	}
	object.Filename = objectName
	savePath := filepath.Join(a.dbPath, object.Filename)
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
		"data": ObjectInfo{
			Id:      objectName,
			Name:    objectName,
			Size:    fileDetail.Size(),
			ModTime: fileDetail.ModTime().Unix(),
		},
	})
}

func (a *ObjectApi) Get(c *gin.Context) {
	objId := c.Query("name")
	if objId == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    "400",
			"message": "objId为空",
		})
		return
	}
	objPath := filepath.Join(a.dbPath, objId)
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
}
