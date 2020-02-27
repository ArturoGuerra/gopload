package main

import (
    "fmt"
    "github.com/minio/minio-go/v6"
)

var (
    endpoint = "10.50.1.210:9000"
    accessKeyID = "AKIAIOSFODNN7EXAMPLE"
    secretAccessKey = "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY"
)

func main() {
    client, err := minio.New(endpoint, accessKeyID, secretAccessKey, false)
    if err != nil {
        fmt.Println(err)
        return
    }

    //bucketName := "killme"
    //region := "us-east-1"

    buckets, err := client.ListBuckets()
    if err != nil {
        fmt.Println(err)
        return
    }

    for _, bucket := range buckets {
        fmt.Println(bucket.Name)
        lifecycle, _ := client.GetBucketLifecycle(bucket.Name)
        fmt.Println(lifecycle)
    }
}
