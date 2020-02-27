package utils

import (
    "fmt"
    "strings"
    "mime/multipart"
    "time"
    "math/rand"
)

const charset  = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
var seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))

type (
    object struct {
        File      *multipart.FileHeader
        Filename  string
        Extension string
    }

    Object interface {
        GetFile() *multipart.FileHeader
        GetFilename() string
        GetExt() string
    }
)

func randString(length int) string {
  b := make([]byte, length)
  for i := range b {
    b[i] = charset[seededRand.Intn(len(charset))]
  }
  return string(b)
}

func NewFile(file *multipart.FileHeader) Object {
    raw := file.Filename
    name := randString(10)
    ext := strings.Split(raw, ".")[1]
    fname  := fmt.Sprintf("%s.%s", name, ext)

    return &object{
        File:      file,
        Filename:  fname,
        Extension: ext,
    }
}

func (o *object) GetFile() *multipart.FileHeader {
    return o.File
}

func (o *object) GetExt() string {
    return o.Extension
}

func (o *object) GetFilename() string {
    return o.Filename
}
