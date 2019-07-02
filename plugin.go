package drone_aliyun_oss

import (
	"drone-aliyun-oss/utils"
	"errors"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"strings"
)

type Plugin struct {
	OSS       OSS
	LocalFile string
}
type OSS struct {
	Endpoint        string
	AccessKeyId     string
	AccessKeySecret string
	BucketName      string
	Dir             string
	FileFormat      string
}

func (p *Plugin) check() error {

	if p.LocalFile == "" {
		return errors.New("LocalFile can not nil")
	}

	if p.OSS.Endpoint == "" {
		return errors.New("endpoint can not nil")
	}

	if p.OSS.AccessKeyId == "" {
		return errors.New("AccessKeyId can not nil")
	}

	if p.OSS.AccessKeySecret == "" {
		return errors.New("AccessKeySecret can not nil")
	}

	if p.OSS.BucketName == "" {
		return errors.New("BucketName can not nil")
	}

	return nil
}

func (p *Plugin) Exec() error {

	if err := p.check(); err != nil {
		return err
	}

	client, err := oss.New(p.OSS.Endpoint, p.OSS.AccessKeyId, p.OSS.AccessKeySecret)
	if err != nil {
		return err
	}

	bucket, err := client.Bucket(p.OSS.BucketName)

	if err != nil {
		return err
	}
	name := p.FileName()

	objectKey := fmt.Sprintf("%s/%s", p.OSS.Dir, name)
	if p.OSS.Dir == "" {
		objectKey = strings.TrimLeft(objectKey, "/")
	}
	return bucket.PutObjectFromFile(objectKey, p.LocalFile)
	//return err
}

func (p *Plugin) FileName() string {
	if p.OSS.FileFormat == "" {
		f := strings.Split(p.LocalFile, "/")
		return f[len(f)-1]
	}

	return utils.Replace(p.OSS.FileFormat)
}
