package workwx

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func getTestApp() *WorkwxApp {
	// Initialize with test credentials
	// In a real test, you might use environment variables or test config
	// 方式1: 使用配置文件
	config, err := LoadConfig()
	if err != nil {
		panic("failed to load WeChat Work config: " + err.Error())
	}

	client := New(config.CorpID)
	app := client.WithApp(config.CorpSecret, config.AgentID)
	return app
}

func TestCreateDoc(t *testing.T) {
	app := getTestApp() // Helper function to get a test WorkwxApp instance

	t.Run("\"successful document creation", func(t *testing.T) {
		req := CreateDocumentRequest{
			SpaceID:    "s.wx5ce26a78df2d2c7d.7521262158JX",
			FatherID:   "s.wx5ce26a78df2d2c7d.7521262158JX_d.7521266546zmV",
			Type:       DocTypeSpreadsheet, // Invalid type
			Name:       "testDoc" + time.Now().Format("20060102150405"),
			AdminUsers: []string{"ZhouEnXian"},
		}

		resp, err := app.CreateDocument(req)
		fmt.Println(err)
		assert.NoError(t, err)
		assert.NotEmpty(t, resp)
	})
}
func TestStructToSpreadsheet(t *testing.T) {
	type Person struct {
		Name    string    `json:"name" wxdoc:"姓名"`
		Age     int       `json:"age" wxdoc:"年龄"`
		City    string    `json:"city" wxdoc:"城市"`
		Salary  float64   `json:"salary" wxdoc:"薪资"`
		Active  bool      `json:"active" wxdoc:"active"`
		Created time.Time `json:"created_at" wxdoc:"-"`
	}

	now := time.Now()
	testData := []Person{
		{"Alice", 30, "New York", 75000.50, true, now},
		{"Bob", 25, "San Francisco", 80000.00, false, now},
		{"Charlie", 35, "Chicago", 90000.75, true, now},
	}

	t.Run("with headers", func(t *testing.T) {
		updateReq, err := StructToSpreadsheet(testData, "Sheet1", true)
		assert.NoError(t, err)
		assert.NotNil(t, updateReq)
		assert.Equal(t, "Sheet1", updateReq.SheetID)
		assert.Len(t, updateReq.GridData.Rows, 4) // Header + 3 rows

		// Verify headers
		headers := updateReq.GridData.Rows[0].Values
		assert.Len(t, headers, 5)
		assert.Equal(t, "姓名", headers[0].CellValue.Text)
		assert.Equal(t, "年龄", headers[1].CellValue.Text)
		assert.Equal(t, "城市", headers[2].CellValue.Text)
		assert.Equal(t, "薪资", headers[3].CellValue.Text)
		assert.Equal(t, "active", headers[4].CellValue.Text)

		// Verify first data row
		firstRow := updateReq.GridData.Rows[1].Values
		assert.Equal(t, "Alice", firstRow[0].CellValue.Text)
		assert.Equal(t, "30", firstRow[1].CellValue.Text)
		assert.Equal(t, "New York", firstRow[2].CellValue.Text)
		assert.Equal(t, "75000.5", firstRow[3].CellValue.Text)
		assert.Equal(t, "true", firstRow[4].CellValue.Text)
	})

	t.Run("without headers", func(t *testing.T) {
		updateReq, err := StructToSpreadsheet(testData, "Sheet1", false)
		assert.NoError(t, err)
		assert.NotNil(t, updateReq)
		assert.Len(t, updateReq.GridData.Rows, 3) // Just the 3 data rows
	})

	t.Run("invalid input", func(t *testing.T) {
		// Not a slice
		_, err := StructToSpreadsheet(Person{}, "Sheet1", true)
		assert.Error(t, err)

		// Empty slice
		_, err = StructToSpreadsheet([]Person{}, "Sheet1", true)
		assert.Error(t, err)

		// Slice of non-structs
		_, err = StructToSpreadsheet([]string{"a", "b"}, "Sheet1", true)
		assert.Error(t, err)
	})
}

func TestBatchUpdateSpreadsheet(t *testing.T) {
	app := getTestApp()
	docID := createTestDocument(t, app) // Helper to create a test doc
	//docID := "dcL65IIxT37rkUcbB3YHXOejEmICsqsZv8Fv5r2uRfmKkNJq5Vkq-oazTxNuucVYUIzCnqNcJZVUgsSzh4wulKTg" // Helper to create a test doc

	t.Run("add sheet and update cells", func(t *testing.T) {
		porperties, err := app.GetSheet(docID)
		assert.NoError(t, err)
		sheetId := porperties.Properties[0].SheetID

		// Then prepare some data to insert
		type Product struct {
			ID    string  `json:"id" wxdoc:"id"`
			Name  string  `json:"name" wxdoc:"名称"`
			Price float64 `json:"price" wxdoc:"价格"`
			Stock int     `json:"stock" wxdoc:"库存"`
		}

		products := []Product{
			{"P1", "Laptop", 999.99, 15},
			{"P2", "Phone", 699.99, 100},
			{"P3", "Tablet", 399.99, 75},
		}

		updateResult, err := app.AddData(docID, sheetId, products, true)
		assert.NoError(t, err)
		assert.NotNil(t, updateResult)
		assert.NotNil(t, updateResult.Responses[0].UpdateRangeResponse)
		assert.Equal(t, uint32(16), updateResult.Responses[0].UpdateRangeResponse.UpdatedCells) // 4 cols * 3 rows
	})

	t.Run("invalid doc id", func(t *testing.T) {
		batchReq := reqWedocBatchUpdate{
			DocID: "invalid-doc-id",
			Requests: []UpdateRequest{
				{
					AddSheetRequest: &AddSheetRequest{
						Title:       "ShouldFail",
						RowCount:    5,
						ColumnCount: 5,
					},
				},
			},
		}

		_, err := app.batchUpdateSpreadsheet(batchReq)
		assert.Error(t, err)
	})
}

func TestGetSheetRangeData(t *testing.T) {
	app := getTestApp()
	docID := createTestDocument(t, app)
	sheetID := addTestSheet(t, app, docID) // Helper to add a test sheet

	t.Run("get valid range", func(t *testing.T) {
		req := reqWedocGetSheetRangeData{
			DocID:   docID,
			SheetID: sheetID,
			Range:   "A1:D5",
		}

		resp, err := app.GetSheetRangeData(req)
		fmt.Println(resp)
		assert.NoError(t, err)
	})

	t.Run("invalid range", func(t *testing.T) {
		req := reqWedocGetSheetRangeData{
			DocID:   docID,
			SheetID: sheetID,
			Range:   "InvalidRange",
		}

		_, err := app.GetSheetRangeData(req)
		assert.Error(t, err)
	})
}
func createTestDocument(t *testing.T, app *WorkwxApp) string {
	//req := reqWedocCreateDoc{
	//	FatherID:   "s.wx5ce26a78df2d2c7d.7521262158JX_d.7521266546zmV",
	//	DocType:    uint32(DocTypeSpreadsheet), // Invalid type
	//	DocName:    "testDoc" + time.Now().Format("20060102150405"),
	//	AdminUsers: []string{"ZhouEnXian"},
	//}
	req2 := CreateDocumentRequest{
		SpaceID:    "s.wx5ce26a78df2d2c7d.7521262158JX",
		FatherID:   "s.wx5ce26a78df2d2c7d.7521262158JX_d.7521266546zmV",
		Type:       DocTypeSpreadsheet, // Invalid type
		Name:       "testDoc" + time.Now().Format("20060102150405"),
		AdminUsers: []string{"ZhouEnXian"},
	}

	resp, err := app.CreateDocument(req2)
	if err != nil {
		t.Fatalf("Failed to create test document: %v", err)
	}

	return resp.DocID
}

func addTestSheet(t *testing.T, app *WorkwxApp, docID string) string {
	addSheetReq := reqWedocBatchUpdate{
		DocID: docID,
		Requests: []UpdateRequest{
			{
				AddSheetRequest: &AddSheetRequest{
					Title:       "TestSheet",
					RowCount:    10,
					ColumnCount: 5,
				},
			},
		},
	}

	resp, err := app.batchUpdateSpreadsheet(addSheetReq)
	if err != nil {
		t.Fatalf("Failed to add test sheet: %v", err)
	}
	assert.Less(t, resp.Responses, 1)
	return resp.Responses[0].AddSheetResponse.Properties.SheetID

	//assert.Less(t, resp.Data.Responses, 1)
	//return resp.Data.Responses[0].AddSheetResponse.Properties.SheetID
}
