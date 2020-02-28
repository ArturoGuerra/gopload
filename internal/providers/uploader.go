package providers

import (
    "github.com/arturoguerra/goimgupload/internal/providers/utils"
    minio "github.com/minio/minio-go/v6"
    "github.com/arturoguerra/goimgupload/internal/structs"
    "github.com/arturoguerra/go-logging"

    "context"
    "time"
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
        Init() error
    }
)

func New(cfg *structs.Config) (Uploader, error) {
    client, err := minio.New(cfg.Endpoint, cfg.AccessKeyID, cfg.SecretAccessKey, cfg.SSL)
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

    log.Info("Creating bucket")
    err := u.Client.MakeBucket(bucket, region)
    if err != nil {
        log.Error(err)
        return nil
    }

    log.Info("Setting bucket policy")
    err = u.Client.SetBucketPolicy(bucket, policy)
    if err != nil {
        return err
    }

    return nil
}

func (u *uploader) Upload(object utils.Object) error {
    file := object.GetFile()

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
    defer cancel()

    src, err := file.Open()
    if err != nil {
        return err
    }

    defer src.Close()

    opts := minio.PutObjectOptions{
        ContentType: "image/jpeg",
    }

    log.Infof("Uploading: %s of size: %d", object.GetFilename(), file.Size)

    _, err = u.Client.PutObjectWithContext(ctx, u.Cfg.Bucket, object.GetFilename(), src, file.Size, opts)
    if err != nil {
        return err
    }

    return nil
}
