package google

import (
	"context"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
	"log"
)

type GDrive interface {
	AddPermission(fileID string, permission *drive.Permission) (err error)
}

type gDrive struct {
	*drive.Service
}

func NewGDrive(ctx context.Context) GDrive {
	svc, err := drive.NewService(ctx, option.WithCredentialsFile("../sheet_credential.json"))
	if err != nil {
		log.Fatalf("Fail to create drive service: %v", err)
	}

	return &gDrive{
		svc,
	}
}

func (d *gDrive) AddPermission(fileID string, permission *drive.Permission) (err error) {
	resp, err := d.Permissions.Create(fileID, permission).Do()
	if err != nil {
		return
	}

	log.Printf("Success AddPermission %+v", resp)

	return
}
