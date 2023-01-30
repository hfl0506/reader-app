package controller

import (
	"fmt"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/hfl0506/reader-app/internal/utils"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

func (h *Handler) UploadImage(e echo.Context) error {
	sess := e.Get("sess").(*session.Session)

	uploader := s3manager.NewUploader(sess)

	bucket := viper.GetString("AWS_BUCKET")

	fmt.Println(bucket)

	file, err := e.FormFile("image")

	if err != nil {
		return e.JSON(http.StatusBadRequest, utils.NewError(err))
	}

	filename := file.Filename

	fileSrc, err := file.Open()

	if err != nil {
		return e.JSON(http.StatusBadRequest, utils.NewError(err))
	}

	up, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(filename),
		Body:   fileSrc,
	})

	if err != nil {
		return e.JSON(http.StatusInternalServerError, echo.Map{
			"error":    err,
			"uploader": up,
		})
	}

	filepath := "https://" + bucket + ".s3.amazonaws.com/" + filename

	return e.JSON(http.StatusCreated, echo.Map{
		"uri": filepath,
	})

}

func (h *Handler) DownloadFile(e echo.Context) error {
	sess := e.Get("sess").(*session.Session)

	filename := e.QueryParam("file")

	downloader := s3manager.NewDownloader(sess)

	bucket := viper.GetString("AWS_BUCKET")

	buffer := &aws.WriteAtBuffer{}

	dl, err := downloader.Download(buffer, &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(filename),
	})

	if err != nil {
		return e.JSON(http.StatusBadRequest, echo.Map{
			"error":      err,
			"downloader": dl,
		})
	}

	data := buffer.Bytes()

	return e.JSON(http.StatusAccepted, echo.Map{
		"data": data,
	})

}
