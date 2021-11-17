# Drone CI Aliyun OSS Plugin

### 上传示例
`.drone.yml` 配置:
```yaml
    ...
  - name: publish oss
    image: bbking/drone-aliyun-oss
    settings:
      local_file: dist.tar.gz
      endpoint: http://oss-cn-*.aliyuncs.com
      access_key_id: *
      access_key_secret: *
      bucket_name: test_name
      dir: test
      object_name: foo.tar.gz
```

### Plugin Parameter Reference

* `local_file`(string): 上传的文件名称,必填
* `endpoint`(string): 地域节点,必填
* `access_key_id`(string): AccessKeyId,必填
* `access_key_secret`(string): AccessKeySecret,必填
* `bucket_name`(string): Bucket，必填
* `dir`(string): 存放的二级目录，可选
* `object_name`(string): 上传后的文件名,可以是模板格式，内置 `date` 和 `tag` 变量：
  * `foo_{{ .tag }}.tar.gz`
  * `foo_{{ .date.Format "2006-01-02" }}.tar.gz`