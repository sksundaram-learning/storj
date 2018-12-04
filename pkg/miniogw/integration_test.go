// Copyright (C) 2018 Storj Labs, Inc.
// See LICENSE for copying information.

package miniogw

import (
	"context"
	"flag"
	"os"
	"testing"
	"time"

	"github.com/minio/cli"
	minio "github.com/minio/minio/cmd"
	"github.com/stretchr/testify/assert"

	"storj.io/storj/internal/identity"
	"storj.io/storj/internal/s3client"
	"storj.io/storj/internal/testcontext"
	"storj.io/storj/internal/testplanet"
	"storj.io/storj/pkg/cfgstruct"
	"storj.io/storj/pkg/miniogw/logging"
	"storj.io/storj/pkg/provider"
)

func TestUploadDownload(t *testing.T) {
	var (
		apiKey     = "apiKey"
		encKey     = "encKey"
		bucket     = "bucket"
		objectName = "testdata"
		ctx = testcontext.New(t)
	)

	planet, err := testplanet.New(t, 1, 40, 0)
	assert.NoError(t, err)

	defer func() {
		err = planet.Shutdown()
		if err != nil {
			t.Fatal(err)
		}
	}()

	// create temporary directory for minio
	tmpDir := ctx.Dir()
	
	defer ctx.Cleanup()

	err = flag.Set("pointer-db.auth.api-key", apiKey)
	assert.NoError(t, err)

	// bind default values to config
	var gwCfg Config
	cfgstruct.Bind(&flag.FlagSet{}, &gwCfg)

	// minio config directory
	gwCfg.MinioDir = tmpDir

	// addresses
	gwCfg.Address = "127.0.0.1:7777"
	gwCfg.OverlayAddr = planet.Satellites[0].Addr()
	gwCfg.PointerDBAddr = planet.Satellites[0].Addr()

	// keys
	gwCfg.APIKey = apiKey
	gwCfg.EncKey = encKey

	// redundancy
	gwCfg.MinThreshold = 20
	gwCfg.RepairThreshold = 25
	gwCfg.SuccessThreshold = 30
	gwCfg.MaxThreshold = 40

	planet.Start(ctx)

	time.Sleep(2 * time.Second)

	ca, err := testidentity.NewTestCA(ctx)
	assert.NoError(t, err)
	identity, err := ca.NewIdentity()
	assert.NoError(t, err)

	errch := make(chan error)

	// setup and start gateway
	go func() {
		errch <-runGateway(ctx, gwCfg, identity)
	}()

	time.Sleep(100 * time.Millisecond)

	client, err := s3client.NewMinio(s3client.Config{
		S3Gateway:     gwCfg.Address,
		Satellite:     planet.Satellites[0].Addr(),
		AccessKey:     gwCfg.AccessKey,
		SecretKey:     gwCfg.SecretKey,
		APIKey:        apiKey,
		EncryptionKey: encKey,
		NoSSL:         true,
	})
	assert.NoError(t, err)

	err = client.MakeBucket(bucket, "")
	assert.NoError(t, err)

	// generate enough data for a remote segment
	data := []byte{}
	for i := 0; i < 5000; i++ {
		data = append(data, 'a')
	}

	err = client.Upload(bucket, objectName, data)
	assert.NoError(t, err)

	// buffer := make([]byte, len(data))

	// bytes, err := client.Download(bucket, tobjectName, buffer)
	// assert.NoError(t, err)

	// assert.Equal(t, string(data), string(bytes))

	select {
	// case <-ctx.Done():
	case err = <-errch:
		t.Fatal(err)
	default:
	}
}

// runGateway registers and starts a gateway
func runGateway(ctx context.Context, c Config, identity *provider.FullIdentity) (err error) {
	err = minio.RegisterGatewayCommand(cli.Command{
		Name:  "storj",
		Usage: "Storj",
		Action: func(cliCtx *cli.Context) error {
			gw, err := c.NewGateway(ctx, identity)
			if err != nil {
				return err
			}

			minio.StartGateway(cliCtx, logging.Gateway(gw))
			return Error.New("unexpected minio exit")
		},
		HideHelpCommand: true,
	})
	if err != nil {
		return err
	}

	err = os.Setenv("MINIO_ACCESS_KEY", c.AccessKey)
	if err != nil {
		return err
	}

	err = os.Setenv("MINIO_SECRET_KEY", c.SecretKey)
	if err != nil {
		return err
	}

	minio.Main([]string{"storj", "gateway", "storj", "--address", c.Address, "--config-dir", c.MinioDir, "--quiet"})
	return Error.New("unexpected minio exit")
}
