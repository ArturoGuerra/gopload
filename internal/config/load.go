package config

import (
    "os"
    "strings"
    "github.com/arturoguerra/goimgupload/internal/structs"
)

func GetBoolEnv(key string) bool {
    value := os.Getenv(key)
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
    path      := Getenv("ENDPOINT_PATH", "", "/upload")
    endpoint  := Getenv("ENDPOINT", "", "s3.amazonaws.com")
    bucket    := Getenv("BUCKET", "", "")
    location  := Getenv("LOCATION", "REGION", "us-east-1")
    accesskey := Getenv("ACCESS_KEY_ID", "AWS_ACCESS_KEY_ID", "")
    secretkey := Getenv("SECRET_ACCESS_KEY", "AWS_SECRET_ACCESS_KEY", "")

    return &structs.Config{
        Path:            path,
        Endpoint:        endpoint,
        Bucket:          bucket,
        Location:        location,
        AccessKeyID:     accesskey,
        SecretAccessKey: secretkey,
        SSL:             GetBoolEnv("SSL"),
        PolicyCheck:     GetBoolEnv("POLICY_CHECK"),
    }
}
