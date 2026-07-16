package resume

import (
	"app/db"
	"app/internal/services"
	"app/models/bolejiang"
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"io"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func UploadAction(c *gin.Context) {
	data := gin.H{}
	err := func() error {
		accountId := services.AuthGetAccountID(c)
		if accountId == "" {
			return errors.New("请先登陆")
		}
		formFile, err := c.FormFile("file")
		if err != nil {
			return err
		}
		ossSvc := services.GetOssService()
		client, err := ossSvc.Client()
		if err != nil {
			return err
		}
		bucket, err := client.Bucket(ossSvc.Bucket)
		if err != nil {
			return err
		}
		suffix := ""
		a := strings.Split(formFile.Filename, ".")
		if len(a) > 1 {
			suffix = "." + strings.ToLower(a[len(a)-1])
		}

		fd, err := formFile.Open()
		if err != nil {
			return err
		}
		defer fd.Close()
		bs, err := io.ReadAll(fd)
		if err != nil {
			return err
		}
		m := md5.New()
		m.Write(bs)
		var now = time.Now().Format("20060102/1504/")
		filename := "upload/" + now + hex.EncodeToString(m.Sum(nil)) + suffix
		err = bucket.PutObject(filename, bytes.NewReader(bs))
		if err != nil {
			return err
		}
		s := ossSvc.PublicURL(filename)
		data["url"] = s
		var account bolejiang.Account
		ok, err := db.Default().Where("id = ?", accountId).Get(&account)
		if err != nil {
			return err
		}
		if !ok {
			return errors.New("用户不存在")
		}
		account.ResumeUrl = s
		_, err = db.Default().Where("id = ?", account.Id).Cols("resume_url").Update(account)
		if err != nil {
			return err
		}
		return nil
	}()
	if err != nil {
		c.Status(500)
		services.ResponseError(c, -1, err.Error(), nil)
		return
	}
	services.ResponseSuccess(c, data)
}
