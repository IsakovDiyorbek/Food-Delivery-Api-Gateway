package handler

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
)


// @Summary Upload a file to MinIO
// @Description Upload a file to MinIO
// @Tags MinIO
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "File to upload"
// @Param filename formData string true "Filename"
// @Success 200 {object} string "Success"
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /minio/upload [post]
func (h *Handler) UploadFile(c *gin.Context) {
	file, _, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File is required"})
		return
	}

	bucketName := "testbucket"
	objectName := "images/" + time.Now().Format("20060102150405") + "-" + c.Request.FormValue("filename")

	_, err = h.miniIO.PutObject(c, bucketName, objectName, file, -1, minio.PutObjectOptions{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload image"})
		return
	}

	presignedURL, err := h.miniIO.PresignedGetObject(c, bucketName, objectName, 24*time.Hour, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate URL"})
		return
	}
	


	c.JSON(http.StatusOK, gin.H{"url": presignedURL.String()})
}



// @Summary Download an image from MinIO
// @Description Download an image from MinIO
// @Tags MinIO
// @Accept  json
// @Produce  json
// @Param bucket path string true "Bucket name"
// @Param object path string true "Object name"
// @Success 200 {object} string "Success"
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /minio/{bucket}/{object} [get]
func (h *Handler) DownloadImage(c *gin.Context) {
	objectName := c.Param("bucket")
	bucketName := c.Param("obkect")

	object, err := h.miniIO.GetObject(context.Background(), bucketName, objectName, minio.GetObjectOptions{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer object.Close()

	c.DataFromReader(http.StatusOK, -1, "application/octet-stream", object, map[string]string{"Content-Disposition": fmt.Sprintf("attachment; filename=%s", objectName)})
}
