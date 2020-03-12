package structs

type Config struct {
    Path            string
    Endpoint        string
    Bucket          string
    Location        string
    AccessKeyID     string
    SecretAccessKey string
    SSL             bool
    PolicyCheck     bool
    ApiKey          string
}
