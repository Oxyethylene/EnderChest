package api

import (
	"errors"
	"fmt"
	"github.com/Oxyethylene/littlebox/config"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"os"
	"path/filepath"
)

const (
	success = 200 + iota
)

const (
	clientError = 400 + iota
	invalidObjId
)

const (
	serverError = 500 + iota
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
	fullPath, err := filepath.Abs(config.Store.DbPath)
	exist, err := pathExists(fullPath)
	if err != nil {
		zap.S().Fatal("error check data_path exists",
			zap.Error(err),
			"data_path", config.Store.DbPath,
		)
	}
	if !exist {
		zap.S().Fatal("data_path not exists",
			zap.Error(errors.New("data_path not exist")),
			"data_path", fullPath,
		)
	} else {
		if dir := isDir(fullPath); !dir {
			zap.S().Fatal("error check data_path exists",
				zap.Error(errors.New("data_path is file")),
				"data_path", fullPath,
			)
		}
	}
	zap.S().Infow("success loaded data_path",
		"data_path", fullPath,
	)
	return &ObjectApi{
		dbPath: fullPath,
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
		"code":    success,
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
			"code":    clientError,
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
			"code":    clientError,
			"message": fmt.Sprintf("err upload object: %v", err),
		})
		return
	}
	fileDetail, _ := os.Stat(savePath)
	c.JSON(http.StatusOK, gin.H{
		"code":    success,
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
			"code":    clientError,
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
				"code":    invalidObjId,
				"message": fmt.Sprintf("can't find object with id %s", objId),
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":    serverError,
				"message": fmt.Sprintf("server internal error: %v", err),
			})
		}
		return
	}
	c.File(objPath)
}

func (a *ObjectApi) Remove(c *gin.Context) {
	objId := c.Query("name")
	if objId == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    clientError,
			"message": "objId为空",
		})
		return
	}
	objPath := filepath.Join(a.dbPath, objId)
	err := os.Remove(objPath)
	if err != nil {
		zap.S().Errorw("remove object failed", "path", objPath)
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    serverError,
			"message": "can't remove objects",
		})
		return
	}
	zap.S().Infow("remove object success", "objId", objId, "dbPath", a.dbPath)
	c.JSON(http.StatusBadRequest, gin.H{
		"code":    success,
		"message": "remove success",
	})
}

func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	//IsNotExist来判断，是不是不存在的错误
	if os.IsNotExist(err) { //如果返回的错误类型使用os.isNotExist()判断为true，说明文件或者文件夹不存在
		return false, nil
	}
	return false, err //如果有错误了，但是不是不存在的错误，所以把这个错误原封不动的返回
}

func isDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {

		return false
	}
	return s.IsDir()
}

func isFile(path string) bool {
	return !isDir(path)
}
