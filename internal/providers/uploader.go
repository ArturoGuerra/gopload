package providers

import (
    "github.com/arturoguerra/goimgupload/internal/providers/utils"
    minio "github.com/minio/minio-go/v6"
    "github.com/arturoguerra/goimgupload/internal/structs"
    "github.com/arturoguerra/go-logging"
)

var log = logging.New()
const policy = `{"Statement":[{"Action":["s3:GetBucketLocation","s3:ListBucket"],"Effect":"Allow","Principal":{"AWS":["*"]},"Resource":["arn:aws:s3:::arturo"]},{"Action":["s3:GetObject"],"Effect":"Allow","Principal":{"AWS":["*"]},"Resource":["arn:aws:s3:::arturo/*"]}],"Version":"2012-10-17"}`

type (
    uploader struct {
        Client *minio.Client
        Cfg    *structs.Config
    }

    Uploader interface {
        Upload(utils.Object) error
    }
)

func New(cfg *structs.Config) (Uploader, error) {
    client, err := minio.NewWithRegion(cfg.Endpoint, cfg.AccessKeyID, cfg.SecretAccessKey, cfg.SSL, cfg.Location)
    if err != nil {
        return nil, err
    }

    return &uploader{
        Client: client,
        Cfg:    cfg,
    }, nil
}

func (u *uploader) Init() error {
    bucket := u.Cfg.Bucket
    region := u.Cfg.Location

    err := u.Client.MakeBucket(bucket, region)
    if err != nil {
        exists, err := u.Client.BucketExists(bucket)
        if err == nil && exists {
            log.Infof("Bucket %s exists", bucket)
        } else {
            return err
        }
    }

    err = u.Client.SetBucketPolicy(bucket, policy)
    if err != nil {
        return err
    }

    return nil
}

func (u *uploader) Upload(object utils.Object) error {
    file := object.GetFile()

    src, err := file.Open()
    if err != nil {
        return err
    }

    defer src.Close()

    opts := minio.PutObjectOptions{
        ContentType: "image/jpeg",
    }

    _, err = u.Client.PutObject(u.Cfg.Bucket, object.GetFilename(), src, -1, opts)
    if err != nil {
        return err
    }

    return nil
}
