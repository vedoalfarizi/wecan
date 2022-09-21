package google

import (
	"context"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
	"log"
)

type GSheet interface {
	GetSpreadsheet(ID string) (resp *sheets.Spreadsheet, err error)
	CreateSpreadsheet(ctx context.Context, sheet *sheets.Spreadsheet) (sheetID string, err error)
	UpdateSpreadsheet(ctx context.Context, ID string, values *sheets.ValueRange) (err error)
	AppendValue(ctx context.Context, ID string, values *sheets.ValueRange) (err error)
	BatchUpdateSpreadsheet(ctx context.Context, ID string, req *sheets.BatchUpdateSpreadsheetRequest) (err error)
}

type gSheet struct {
	*sheets.Service
}

func NewGSheet(ctx context.Context) GSheet {
	svc, err := sheets.NewService(ctx, option.WithCredentialsFile("../sheet_credential.json"))
	if err != nil {
		log.Fatalf("Fail to create sheets service: %v", err)
	}

	return &gSheet{
		svc,
	}
}

func (g *gSheet) GetSpreadsheet(ID string) (resp *sheets.Spreadsheet, err error) {
	resp, err = g.Spreadsheets.Get(ID).Do()
	if err != nil {
		return
	}

	log.Println("Success GetSpreadsheet")

	return
}

func (g *gSheet) CreateSpreadsheet(ctx context.Context, sheet *sheets.Spreadsheet) (sheetID string, err error) {
	resp, err := g.Spreadsheets.Create(sheet).Context(ctx).Do()
	if err != nil {
		return
	}

	log.Println("Success CreateSpreadsheet")

	sheetID = resp.SpreadsheetId
	return
}

func (g *gSheet) UpdateSpreadsheet(ctx context.Context, ID string, values *sheets.ValueRange) (err error) {
	_, err = g.Spreadsheets.Values.Update(ID, values.Range, values).ValueInputOption("USER_ENTERED").Context(ctx).Do()
	if err != nil {
		return
	}

	log.Println("Success UpdateSpreadsheet")

	return
}

func (g *gSheet) AppendValue(ctx context.Context, ID string, values *sheets.ValueRange) (err error) {
	_, err = g.Spreadsheets.Values.Append(ID, values.Range, values).ValueInputOption("USER_ENTERED").Context(ctx).Do()
	if err != nil {
		return
	}

	log.Println("Success AppendValue")

	return
}

func (g *gSheet) BatchUpdateSpreadsheet(ctx context.Context, ID string, req *sheets.BatchUpdateSpreadsheetRequest) (err error) {
	_, err = g.Spreadsheets.BatchUpdate(ID, req).Context(ctx).Do()
	if err != nil {
		return
	}

	log.Println("Success BatchUpdateSpreadsheet")

	return
}
