package config

import (
    "os"
    "strings"
    "github.com/arturoguerra/goimgupload/internal/structs"
)

func getSSL() bool {
    value := os.Getenv("SSL")
    if len(value) == 0 {
        return false
    }

    value = strings.ToLower(value)
    if value == "true" {
        return true
    }

    return false
}

func Getenv(key, alt,fallback string) string {
    v := os.Getenv(key)
    if len(v) == 0 {
        if len(alt) != 0 {
            return Getenv(alt, "", fallback)
        }

        return fallback
    }

    return v
}

func Load() *structs.Config {
    endpoint  := Getenv("ENDPOINT", "", "s3.amazonaws.com")
    bucket    := Getenv("BUCKET", "", "")
    location  := Getenv("LOCATION", "REGION", "us-east-1")
    accesskey := Getenv("ACCESS_KEY_ID", "AWS_ACCESS_KEY_ID", "")
    secretkey := Getenv("SECRET_ACCESS_KEY", "AWS_SECRET_ACCESS_KEY", "")
    proxyurl  := Getenv("PROXY_URL", "", "")

    return &structs.Config{
        Endpoint:        endpoint,
        Bucket:          bucket,
        Location:        location,
        AccessKeyID:     accesskey,
        SecretAccessKey: secretkey,
        ProxyURL:        proxyurl,
        SSL:             getSSL(),
    }
}
