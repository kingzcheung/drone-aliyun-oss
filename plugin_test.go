package drone_aliyun_oss

import (
	"os"
	"testing"
	"time"
)

func Test_renderName(t *testing.T) {
	var testDate = "2021-11-17"
	type args struct {
		name    string
		repoTag string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "format date",
			args: args{
				name:    "foo_{{ .date.Format \"2006-01-02\" }}.tar.gz",
				repoTag: "refs/heads/master",
			},
			want: "foo_2021-11-17.tar.gz",
		},
		{
			name: "format branch master",
			args: args{
				name:    "foo_{{ .tag }}.tar.gz",
				repoTag: "refs/heads/master",
			},
			want: "foo_latest.tar.gz",
		},
		{
			name: "format tag",
			args: args{
				name:    "foo_{{ .tag }}.tar.gz",
				repoTag: "refs/tags/v1.0.0",
			},
			want: "foo_1.0.0.tar.gz",
		},
	}
	for _, tt := range tests {
		os.Setenv("DRONE_COMMIT_REF", tt.args.repoTag)
		t.Run(tt.name, func(t *testing.T) {
			if got := renderName(tt.args.name, func() time.Time {
				date, _ := time.Parse("2006-01-02", testDate)
				return date
			}); got != tt.want {
				t.Errorf("renderName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isTemplateName(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"", args{name: "foo_{{.tag}}.tar.gz"}, true},
		{"", args{name: "foo_{{.tag.tar.gz"}, false},
		{"", args{name: "foo_bar.tag.tar.gz"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isTemplateName(tt.args.name); got != tt.want {
				t.Errorf("isTemplateName() = %v, want %v", got, tt.want)
			}
		})
	}
}
