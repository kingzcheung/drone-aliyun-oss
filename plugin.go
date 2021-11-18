package drone_aliyun_oss

import (
	"bytes"
	"errors"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"log"
	"os"
	"path"
	"regexp"
	"strings"
	"text/template"
	"time"
)

var commitRef = os.Getenv("DRONE_COMMIT_REF")
var repoBranch = os.Getenv("DRONE_REPO_BRANCH")

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
	ObjectName      string
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
	log.Printf("Info: FileName=%s\n", name)
	log.Printf("Info: commitRef=%s,repoBranch=%s\n", commitRef, repoBranch)
	objectKey := path.Join(p.OSS.Dir, name)

	if p.OSS.Dir == "" {
		objectKey = strings.TrimLeft(objectKey, "/")
	}
	return bucket.PutObjectFromFile(objectKey, p.LocalFile)
	//return err
}

func (p *Plugin) FileName() string {
	if p.OSS.ObjectName == "" {
		f := strings.Split(p.LocalFile, "/")
		return f[len(f)-1]
	}
	if isTemplateName(p.OSS.ObjectName) {
		if UseDefaultTag(commitRef, repoBranch) {
			return renderName(p.OSS.ObjectName, func() time.Time {
				return time.Now()
			})
		}
		return p.OSS.ObjectName
	}

	return p.OSS.ObjectName
}

func isTemplateName(name string) bool {
	reg := regexp.MustCompile(`{{.*}}`)
	all := reg.FindAllString(name, -1)
	return len(all) > 0
}

func renderName(tmplString string, fn func() time.Time) string {
	repoTag, err := DefaultTag(commitRef)
	if err != nil {
		log.Println("Warning: ", err)
		repoTag = ""
	}
	tmpl, err := template.New("file_name").Parse(tmplString)
	if err != nil {
		return tmplString
	}
	var bf bytes.Buffer
	err = tmpl.Execute(&bf, map[string]interface{}{
		"date": fn(),
		"tag":  repoTag,
	})
	if err != nil {
		return tmplString
	}
	return bf.String()
}
