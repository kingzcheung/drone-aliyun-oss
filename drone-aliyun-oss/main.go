package main

import (
	dronealiyunoss "drone-aliyun-oss"
	"github.com/urfave/cli"
	"log"
	"os"
)

var Version = "v0.0.1"

func main() {
	app := cli.NewApp()
	app.Name = "Drone Aliyun OSS Plugin"
	app.Usage = "Push files to Aliyun OSS"
	app.Copyright = "© 2019 kingzcheung"
	app.Authors = []cli.Author{
		{
			Name:  "Kingz Cheung",
			Email: "i@kingzcheung.com",
		},
	}
	app.Version = Version
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "local_file",
			Usage:  "local file",
			EnvVar: "PLUGIN_LOCAL_FILE",
		},
		cli.StringFlag{
			Name:   "oss.endpoint",
			Usage:  "OSS Endpoint",
			EnvVar: "PLUGIN_ENDPOINT",
		},
		cli.StringFlag{
			Name:   "oss.access_key_id",
			Usage:  "OSS AccessKeyId",
			EnvVar: "PLUGIN_ACCESS_KEY_ID",
		},
		cli.StringFlag{
			Name:   "oss.access_key_secret",
			Usage:  "OSS AccessKeySecret",
			EnvVar: "PLUGIN_ACCESS_KEY_SECRET",
		},
		cli.StringFlag{
			Name:   "oss.bucket_name",
			Usage:  "OSS BucketName",
			EnvVar: "PLUGIN_BUCKET_NAME",
		},
		cli.StringFlag{
			Name:   "oss.dir",
			Usage:  "OSS DIR",
			EnvVar: "PLUGIN_DIR",
		},
		cli.StringFlag{
			Name:   "oss.object_name",
			Usage:  "OSS Object Name",
			EnvVar: "PLUGIN_OBJECT_NAME",
		},
	}
	app.Action = run

	if err := app.Run(os.Args); nil != err {
		log.Println(err)
	}
}

func run(c *cli.Context) {
	plugin := &dronealiyunoss.Plugin{
		LocalFile: c.String("local_file"),
		OSS: dronealiyunoss.OSS{
			Endpoint:        c.String("oss.endpoint"),
			AccessKeyId:     c.String("oss.access_key_id"),
			AccessKeySecret: c.String("oss.access_key_secret"),
			BucketName:      c.String("oss.bucket_name"),
			Dir:             c.String("oss.dir"),
			ObjectName:      c.String("oss.object_name"),
		},
	}
	err := plugin.Exec()
	if err != nil {
		log.Fatalf("Run Error: %v", err)
		return
	}

	log.Println("Upload Success")
}
