package config

type MinIO struct {
	EndPoint        string `mapstructure:"end-point" json:"endPoint" yaml:"end-point"`                        // 服务器地址
	AccessKeyID     string `mapstructure:"access-key-id" json:"accessKeyId" yaml:"access-key-id"`             // 访问密钥
	SecretAccessKey string `mapstructure:"secret-access-key" json:"secretAccessKey" yaml:"secret-access-key"` // 访问密钥
	UseSSL          bool   `mapstructure:"use-ssl" json:"useSSL" yaml:"use-ssl"`                              // 是否使用SSL
	BucketName      string `mapstructure:"bucket-name" json:"bucketName" yaml:"bucket-name"`                  // 存储空间名称
}
