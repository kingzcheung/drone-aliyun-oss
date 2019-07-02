package drone_aliyun_oss

import (
	"drone-aliyun-oss/utils"
	"fmt"
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
)

func TestPlugin_FileName(t *testing.T) {
	as := assert.New(t)
	plugin := &Plugin{
		LocalFile: "dist.tar.gz",
		OSS: OSS{
			FileFormat: "test_{{date|2006-01-02}}.tar.gz",
		},
	}

	as.Equal(plugin.FileName(), "test_2019-07-02.tar.gz")
}

func TestPlugin_FileName2(t *testing.T) {
	str := "test_{{date|2006-01-02}}.tar.gz"
	re := regexp.MustCompile(utils.FormatRE)
	fmt.Println(re.FindAllString(str, -1))
}

func TestPlugin_FileName3(t *testing.T) {
	as := assert.New(t)
	plugin := &Plugin{
		LocalFile: "dist.tar.gz",
		OSS: OSS{
			FileFormat: "test_2006.tar.gz",
		},
	}

	as.Equal(plugin.FileName(), "test_2006.tar.gz")
}
