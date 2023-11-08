// Package upload 上传文件到云存储器
package upload

import (
	"context"
	"fmt"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"go_chat/global"
	"mime/multipart"
)

// UploadImg 上传图片到七牛云
func UploadImg(file multipart.File, fileSize int64) (string, int) {
	var qiniuConfig = global.ServiceConfig.QiNiuConfig
	putPolicy := storage.PutPolicy{
		Scope: qiniuConfig.Bucket,
	}
	// 过期时间
	putPolicy.Expires = 7200
	mac := qbox.NewMac(qiniuConfig.AccessKey, qiniuConfig.SecretKey)
	upToken := putPolicy.UploadToken(mac)
	// 上传配置
	config := &storage.Config{
		Zone: selectZone(qiniuConfig.Zone),
		// 是否用cdn加速
		UseCdnDomains: false,
		// 是否用https
		UseHTTPS: false,
	}
	// 构建上传对象
	upLoader := storage.NewFormUploader(config)
	ret := storage.PutRet{}
	//
	err := upLoader.PutWithoutKey(context.Background(), &ret, upToken, file, fileSize, &storage.PutExtra{})
	if err != nil {
		return "", 500
	}
	fmt.Println(qiniuConfig.QiniuServer)
	return qiniuConfig.QiniuServer + "/" + ret.Key, 0
}

// Zone设置
func selectZone(id int) *storage.Zone {
	switch id {
	case 1:
		return &storage.ZoneHuadong
	case 2:
		return &storage.ZoneHuabei
	case 3:
		return &storage.ZoneHuanan
	default:
		return &storage.ZoneHuadong
	}
}
