package workwx

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

// DocType 文档类型
type DocType uint32

// Document types
const (
	DocTypeDocument    DocType = 3
	DocTypeSpreadsheet DocType = 4
	DocTypeSmartSheet  DocType = 10
)

// Dimension types
const (
	DimensionRow    = "ROW"
	DimensionColumn = "COLUMN"
)

// reqWedocCreateDoc represents a request to create a new document
type reqWedocCreateDoc struct {
	SpaceID    string   `json:"spaceid,omitempty"`
	FatherID   string   `json:"fatherid,omitempty"`
	DocType    uint32   `json:"doc_type"`
	DocName    string   `json:"doc_name"`
	AdminUsers []string `json:"admin_users,omitempty"`
}

func (r reqWedocCreateDoc) intoBody() ([]byte, error) {
	return marshalIntoJSONBody(r)
}

// respWedocCreateDoc represents the response from creating a new document
type respWedocCreateDoc struct {
	respCommon
	URL   string `json:"url,omitempty"`
	DocID string `json:"docid,omitempty"`
}

type WedocCreateDocResult struct {
	URL   string `json:"url,omitempty"`
	DocID string `json:"docid,omitempty"`
}

func (x respWedocCreateDoc) intoWedocCreateDocResult() (WedocCreateDocResult, error) {
	if x.ErrCode != 0 {
		return WedocCreateDocResult{}, errors.New(x.ErrMsg)
	}
	return WedocCreateDocResult{
		URL:   x.URL,
		DocID: x.DocID,
	}, nil
}

// reqWedocBatchUpdate represents a batch update request for spreadsheets
type reqWedocBatchUpdate struct {
	DocID    string          `json:"docid"`
	Requests []UpdateRequest `json:"requests"`
}

func (r reqWedocBatchUpdate) intoBody() ([]byte, error) {
	return marshalIntoJSONBody(r)
}

// UpdateRequest represents a single update operation
type UpdateRequest struct {
	AddSheetRequest        *AddSheetRequest        `json:"add_sheet_request,omitempty"`
	DeleteSheetRequest     *DeleteSheetRequest     `json:"delete_sheet_request,omitempty"`
	UpdateRangeRequest     *UpdateRangeRequest     `json:"update_range_request,omitempty"`
	DeleteDimensionRequest *DeleteDimensionRequest `json:"delete_dimension_request,omitempty"`
}

// AddSheetRequest represents a request to add a new sheet
type AddSheetRequest struct {
	Title       string `json:"title"`
	RowCount    uint32 `json:"row_count"`
	ColumnCount uint32 `json:"column_count"`
}

// DeleteSheetRequest represents a request to delete a sheet
type DeleteSheetRequest struct {
	SheetID string `json:"sheet_id"`
}

// UpdateRangeRequest represents a request to update a range of cells
type UpdateRangeRequest struct {
	SheetID  string    `json:"sheet_id"`
	GridData *GridData `json:"grid_data"`
}

// GridData represents the data for a range of cells
type GridData struct {
	StartRow    uint32    `json:"start_row"`
	StartColumn uint32    `json:"start_column"`
	Rows        []RowData `json:"rows"`
}

// RowData represents data for a single row
type RowData struct {
	Values []CellData `json:"values"`
}

// CellData represents data for a single cell
type CellData struct {
	CellValue *CellValue `json:"cell_value"`
	//CellValue *CellValue `json:"cell_value,omitempty"`
	//CellFormat *CellFormat `json:"cell_format,omitempty"`
}

// CellValue represents the value of a cell
type CellValue struct {
	//Text string     `json:"text,omitempty"`
	Text string `json:"text"`
	//Link *WeDocLink `json:"link,omitempty"`
}

// WeDocLink represents a hyperlink in a cell
type WeDocLink struct {
	Text string `json:"text"`
	URL  string `json:"url"`
}

// CellFormat represents the formatting of a cell
type CellFormat struct {
	TextFormat *TextFormat `json:"text_format,omitempty"`
}

// TextFormat represents text formatting
type TextFormat struct {
	Font          string `json:"font,omitempty"`
	FontSize      uint32 `json:"font_size,omitempty"`
	Bold          bool   `json:"bold,omitempty"`
	Italic        bool   `json:"italic,omitempty"`
	Strikethrough bool   `json:"strikethrough,omitempty"`
	Underline     bool   `json:"underline,omitempty"`
	Color         *Color `json:"color,omitempty"`
}

// Color represents a color in RGBA format
type Color struct {
	Red   uint32 `json:"red"`
	Green uint32 `json:"green"`
	Blue  uint32 `json:"blue"`
	Alpha uint32 `json:"alpha"`
}

// DeleteDimensionRequest represents a request to delete rows or columns
type DeleteDimensionRequest struct {
	SheetID    string `json:"sheet_id"`
	Dimension  string `json:"dimension"`
	StartIndex uint32 `json:"start_index"`
	EndIndex   uint32 `json:"end_index"`
}

// respWedocBatchUpdate represents the response from a batch update
type respWedocBatchUpdate struct {
	respCommon
	//ErrCode int    `json:"errcode"`
	//ErrMsg  string `json:"errmsg"`
	//Data    BatchUpdateData `json:"data,omitempty"`
	Responses []UpdateResponse `json:"responses"`
}

type WedocBatchUpdateResult struct {
	//Data BatchUpdateData
	Responses []UpdateResponse `json:"responses"`
}

func (x respWedocBatchUpdate) intoWedocBatchUpdateResult() (WedocBatchUpdateResult, error) {
	if x.ErrCode != 0 {
		return WedocBatchUpdateResult{}, errors.New(x.ErrMsg)
	}
	return WedocBatchUpdateResult{
		//Data: x.Data,
		Responses: x.Responses,
	}, nil
}

// BatchUpdateData contains the responses from batch operations
type BatchUpdateData struct {
	Responses []UpdateResponse `json:"responses"`
}

// UpdateResponse represents a response to an update operation
type UpdateResponse struct {
	AddSheetResponse        *AddSheetResponse        `json:"add_sheet_response,omitempty"`
	DeleteSheetResponse     *DeleteSheetResponse     `json:"delete_sheet_response,omitempty"`
	UpdateRangeResponse     *UpdateRangeResponse     `json:"update_range_response,omitempty"`
	DeleteDimensionResponse *DeleteDimensionResponse `json:"delete_dimension_response,omitempty"`
}

// AddSheetResponse represents a response to adding a sheet
type AddSheetResponse struct {
	Properties *Properties `json:"properties,omitempty"`
}

// Properties represents sheet properties
type Properties struct {
	SheetID string `json:"sheet_id,omitempty"`
	Title   string `json:"title,omitempty"`
	Index   uint32 `json:"index,omitempty"`
}

// DeleteSheetResponse represents a response to deleting a sheet
type DeleteSheetResponse struct {
	SheetID string `json:"sheet_id"`
}

// UpdateRangeResponse represents a response to updating a range
type UpdateRangeResponse struct {
	UpdatedCells uint32 `json:"updated_cells"`
}

// DeleteDimensionResponse represents a response to deleting dimensions
type DeleteDimensionResponse struct {
	Deleted uint32 `json:"deleted"`
}

// reqWedocGetSheetRangeData represents a request to get sheet range data
type reqWedocGetSheetRangeData struct {
	DocID   string `json:"docid"`
	SheetID string `json:"sheet_id"`
	Range   string `json:"range"`
}

func (r reqWedocGetSheetRangeData) intoBody() ([]byte, error) {
	return marshalIntoJSONBody(r)
}

// respWedocGetSheetRangeData represents the response from getting sheet data
type respWedocGetSheetRangeData struct {
	respCommon
	ErrCode int       `json:"errcode"`
	ErrMsg  string    `json:"errmsg"`
	Data    SheetData `json:"data,omitempty"`
}

type WedocGetSheetRangeDataResult struct {
	Data SheetData
}

func (x respWedocGetSheetRangeData) intoWedocGetSheetRangeDataResult() (WedocGetSheetRangeDataResult, error) {
	if x.ErrCode != 0 {
		return WedocGetSheetRangeDataResult{}, errors.New(x.ErrMsg)
	}
	return WedocGetSheetRangeDataResult{
		Data: x.Data,
	}, nil
}

// reqWedocGetSheetProperties represents a request to get sheet properties
type reqWedocGetSheetProperties struct {
	DocID string `json:"docid"`
}

func (r reqWedocGetSheetProperties) intoBody() ([]byte, error) {
	return marshalIntoJSONBody(r)
}

// SheetProperty represents the properties of a sheet
type SheetProperty struct {
	SheetID     string `json:"sheet_id"`
	Title       string `json:"title"`
	RowCount    uint32 `json:"row_count"`
	ColumnCount uint32 `json:"column_count"`
}

// respWedocGetSheetProperties represents the response from getting sheet properties
type respWedocGetSheetProperties struct {
	respCommon
	Properties []SheetProperty `json:"properties,omitempty"`
}

// WedocGetSheetPropertiesResult represents the final result of getting sheet properties
type WedocGetSheetPropertiesResult struct {
	Properties []SheetProperty
}

func (x respWedocGetSheetProperties) intoWedocGetSheetPropertiesResult() (WedocGetSheetPropertiesResult, error) {
	if x.ErrCode != 0 {
		return WedocGetSheetPropertiesResult{}, errors.New(x.ErrMsg)
	}
	return WedocGetSheetPropertiesResult{
		Properties: x.Properties,
	}, nil
}

// SheetData contains the sheet data
type SheetData struct {
	Result GridData `json:"result"`
}

// createDoc creates a new document in WeChat Work
func (c *WorkwxApp) createDoc(req reqWedocCreateDoc) (*WedocCreateDocResult, error) {
	resp, err := c.execWedocCreateDoc(req)
	if err != nil {
		return nil, err
	}
	obj, err := resp.intoWedocCreateDocResult()
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

// batchUpdateSpreadsheet performs batch updates on a spreadsheet
func (c *WorkwxApp) batchUpdateSpreadsheet(req reqWedocBatchUpdate) (*WedocBatchUpdateResult, error) {
	resp, err := c.execWedocBatchUpdate(req)
	if err != nil {
		return nil, err
	}
	obj, err := resp.intoWedocBatchUpdateResult()
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

// GetSheetRangeData retrieves data from a specified range in a sheet
func (c *WorkwxApp) GetSheetRangeData(req reqWedocGetSheetRangeData) (*WedocGetSheetRangeDataResult, error) {
	resp, err := c.execWedocGetSheetRangeData(req)
	if err != nil {
		return nil, err
	}
	obj, err := resp.intoWedocGetSheetRangeDataResult()
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

// getSheetProperties retrieves data from properties in a sheet
func (c *WorkwxApp) getSheetProperties(req reqWedocGetSheetProperties) (*WedocGetSheetPropertiesResult, error) {
	resp, err := c.execWedocGetSheetProperties(req)
	if err != nil {
		return nil, err
	}
	obj, err := resp.intoWedocGetSheetPropertiesResult()
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

// CreateDocumentRequest 创建文档请求
type CreateDocumentRequest struct {
	SpaceID    string   // 空间ID
	FatherID   string   // 父目录ID
	Name       string   // 文档名称
	Type       DocType  // 文档类型
	AdminUsers []string // 管理员用户列表
}

// CreateDocument 创建新文档
func (c *WorkwxApp) CreateDocument(req CreateDocumentRequest) (*WedocCreateDocResult, error) {
	// 转换请求格式
	apiReq := reqWedocCreateDoc{
		SpaceID:    req.SpaceID,
		FatherID:   req.FatherID,
		DocType:    uint32(req.Type),
		DocName:    req.Name,
		AdminUsers: req.AdminUsers,
	}

	// 调用API
	resp, err := c.createDoc(apiReq)
	if err != nil {
		return nil, fmt.Errorf("failed to create document: %w", err)
	}

	return resp, nil
}

// GetDefaultSheet 获取文档的默认Sheet1
func (c *WorkwxApp) GetSheet(docId string) (*WedocGetSheetPropertiesResult, error) {
	// 获取文档的所有sheet
	req := reqWedocGetSheetProperties{
		DocID: docId,
	}

	resp, err := c.getSheetProperties(req)
	if err != nil {
		return nil, fmt.Errorf("failed to get sheet properties: %w", err)
	}

	if len(resp.Properties) == 0 {
		return nil, errors.New("no sheets found in document")
	}

	return resp, nil
}

// AddData 向sheet中添加数据
func (c *WorkwxApp) AddData(docId, sheetId string, data interface{}, includeHeaders bool) (*WedocBatchUpdateResult, error) {
	// 将数据转换为更新请求
	updateReq, err := StructToSpreadsheet(data, sheetId, includeHeaders)
	if err != nil {
		return nil, fmt.Errorf("failed to convert data to spreadsheet format: %w", err)
	}

	// 准备批量更新请求
	batchReq := reqWedocBatchUpdate{
		DocID: docId,
		Requests: []UpdateRequest{
			{
				UpdateRangeRequest: updateReq,
			},
		},
	}

	// 执行更新
	resp, err := c.batchUpdateSpreadsheet(batchReq)
	if err != nil {
		return nil, fmt.Errorf("failed to update spreadsheet: %w", err)
	}

	return resp, nil
}

// StructToSpreadsheet converts a struct slice to spreadsheet data
func StructToSpreadsheet(data interface{}, sheetId string, includeHeaders bool) (*UpdateRangeRequest, error) {
	val := reflect.ValueOf(data)
	if val.Kind() != reflect.Slice {
		return nil, errors.New("slice must be a struct slice")
	}

	if val.Len() == 0 {
		return nil, errors.New("slice must not be empty")
	}

	// Get the type of the slice elements
	elemType := val.Type().Elem()
	if elemType.Kind() == reflect.Ptr {
		elemType = elemType.Elem()
	}

	if elemType.Kind() != reflect.Struct {
		return nil, errors.New("slice elements must be structs")
	}

	// Prepare grid data
	gridData := &GridData{
		StartRow:    0,
		StartColumn: 0,
	}

	// Add headers if requested
	if includeHeaders {
		headerRow := RowData{}
		for i := 0; i < elemType.NumField(); i++ {
			field := elemType.Field(i)
			tag := field.Tag.Get("wxdoc")
			if tag == "" {
				tag = field.Name
			}
			if tag == "-" {
				continue
			}
			headerRow.Values = append(headerRow.Values, CellData{
				CellValue: &CellValue{Text: tag},
			})
		}
		gridData.Rows = append(gridData.Rows, headerRow)
	}

	// Add data rows
	for i := 0; i < val.Len(); i++ {
		elem := val.Index(i)
		if elem.Kind() == reflect.Ptr {
			elem = elem.Elem()
		}

		row := RowData{}
		for j := 0; j < elem.NumField(); j++ {
			field := elem.Field(j)
			var cellValue string

			switch field.Kind() {
			case reflect.String:
				cellValue = field.String()
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				cellValue = strconv.FormatInt(field.Int(), 10)
			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
				cellValue = strconv.FormatUint(field.Uint(), 10)
			case reflect.Float32, reflect.Float64:
				cellValue = strconv.FormatFloat(field.Float(), 'f', -1, 64)
			case reflect.Bool:
				cellValue = strconv.FormatBool(field.Bool())
			default:
				cellValue = "Unsupported type"
			}

			row.Values = append(row.Values, CellData{
				CellValue: &CellValue{Text: cellValue},
			})
		}
		gridData.Rows = append(gridData.Rows, row)
	}

	return &UpdateRangeRequest{
		SheetID:  sheetId,
		GridData: gridData,
	}, nil
}
