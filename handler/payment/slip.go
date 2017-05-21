package payment

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"regexp"
	"time"

	"io"

	"github.com/ekkapob/saucony/handler/mail"
	"github.com/golang/glog"
	"github.com/google/uuid"
)

func UploadTransferSlip(w http.ResponseWriter, r *http.Request) {
	// limit 3.5 megabyte
	r.Body = http.MaxBytesReader(w, r.Body, 3.5*(1<<(10*2)))
	orderId := r.FormValue("orderId")
	file, header, err := r.FormFile("file")
	if err != nil {
		glog.Error(err)
		formatError(w, errors.New("image file might be bigger than 1 MB"))
		return
	}
	defer file.Close()

	matched, err := regexp.MatchString("^image", header.Header.Get("Content-Type"))
	if err != nil {
		glog.Error(err)
		formatError(w, err)
		return
	}
	if !matched {
		errorMessage := "only image file can be uploade"
		glog.Error(errors.New(errorMessage))
		formatError(w, errors.New(errorMessage))
		return
	}

	t := time.Now()

	if orderId == "" {
		orderId = "no-orderid"
	}
	filePath := fmt.Sprint(
		os.Getenv("UPLOAD_TRANSFER_SLIP_PATH"),
		"/",
		orderId,
		"-",
		t.Format("020106-"),
		uuid.New().String(),
		"-",
		header.Filename)
	destFile, err := os.Create(filePath)

	if err != nil {
		glog.Error(err)
		formatError(w, err)
		return
	}
	defer destFile.Close()

	if _, err = io.Copy(destFile, file); err != nil {
		glog.Error(err)
		formatError(w, err)
		return
	}
	go (func() {
		_, err := mail.TransferSlipUploadNotify(orderId,

			fmt.Sprint(
				os.Getenv("HOST"),
				"/",
				filePath))
		if err != nil {
			glog.Error(err)
		}
	})()

	w.WriteHeader(http.StatusOK)
}

func formatError(w http.ResponseWriter, err error) {
	json, _ := json.Marshal(
		struct {
			Error string
		}{
			Error: err.Error(),
		})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnprocessableEntity)
	w.Write(json)
}
