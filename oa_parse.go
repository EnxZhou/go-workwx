package workwx

import (
	"reflect"
	"strconv"
	"time"
)

const (
	tagName          = "oa"
	tagCreator       = "creator"
	tagTemplate      = "template"
	tagControl       = "control"
	tagID            = "id"
	tagType          = "type"
	tagOption        = "option"
	tagDateFormat    = "format"
	tagSelectorType  = "selector_type"
	tagMultiSelector = "multi"
)

// Converter 业务数据到OAApplyEvent的转换器
type Converter struct {
	templateID          string
	creatorUserID       string
	useTemplateApprover uint8
	approvers           []OAApprover
	notifiers           []string
	notifyType          *uint8
	summaryList         []OASummaryList
	handlers            map[OAControl]func(reflect.Value, map[string]string) OAContentValue
}

// NewConverter 创建新的转换器实例
func NewConverter(templateID, creatorID string) *Converter {
	c := &Converter{
		templateID:          templateID,
		creatorUserID:       creatorID,
		useTemplateApprover: 0,
		handlers:            make(map[OAControl]func(reflect.Value, map[string]string) OAContentValue),
	}

	// 注册默认的控件处理器
	c.RegisterHandler(OAControlText, c.handleText)
	c.RegisterHandler(OAControlTextarea, c.handleText)
	c.RegisterHandler(OAControlNumber, c.handleNumber)
	c.RegisterHandler(OAControlMoney, c.handleMoney)
	c.RegisterHandler(OAControlDate, c.handleDate)
	c.RegisterHandler(OAControlSelector, c.handleSelector)
	c.RegisterHandler(OAControlContact, c.handleContact)
	c.RegisterHandler(OAControlFile, c.handleFile)
	c.RegisterHandler(OAControlTable, c.handleTable)

	return c
}

// RegisterHandler 注册自定义控件处理器
func (c *Converter) RegisterHandler(control OAControl, handler func(reflect.Value, map[string]string) OAContentValue) {
	c.handlers[control] = handler
}

// UseTemplateApprover 设置使用模板审批流程
func (c *Converter) UseTemplateApprover(use bool) {
	if use {
		c.useTemplateApprover = 1
	} else {
		c.useTemplateApprover = 0
	}
}

// WithApprovers 设置审批人
func (c *Converter) WithApprovers(approvers []OAApprover) {
	c.approvers = approvers
}

// WithNotifiers 设置抄送人
func (c *Converter) WithNotifiers(notifiers []string, notifyType uint8) {
	c.notifiers = notifiers
	c.notifyType = &notifyType
}

// WithSummary 设置摘要信息
func (c *Converter) WithSummary(summary []OASummaryList) {
	c.summaryList = summary
}

// Parse 解析业务结构体为OAApplyEvent
func (c *Converter) Parse(businessData interface{}) (*OAApplyEvent, error) {
	val := reflect.ValueOf(businessData)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	if val.Kind() != reflect.Struct {
		return nil, NewInvalidBusinessDataError("businessData must be a struct or pointer to struct")
	}

	event := &OAApplyEvent{
		CreatorUserID:       c.creatorUserID,
		TemplateID:          c.templateID,
		UseTemplateApprover: c.useTemplateApprover,
		Approver:            c.approvers,
		Notifier:            c.notifiers,
		NotifyType:          c.notifyType,
		SummaryList:         c.summaryList,
		ApplyData: OAContents{
			Contents: make([]OAContent, 0),
		},
	}

	typ := val.Type()
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := typ.Field(i)
		tag := fieldType.Tag.Get(tagName)

		if tag == "" {
			continue
		}

		// 处理特殊字段
		switch tag {
		case tagCreator:
			event.CreatorUserID = field.String()
			continue
		case tagTemplate:
			event.TemplateID = field.String()
			continue
		}

		// 解析控件参数
		params := parseTagParams(tag)
		controlType := params[tagControl]
		if controlType == "" {
			continue
		}

		content := OAContent{
			Control: OAControl(controlType),
			ID:      params[tagID],
		}

		// 调用注册的处理器处理控件值
		if handler, ok := c.handlers[OAControl(controlType)]; ok {
			content.Value = handler(field, params)
		}

		event.ApplyData.Contents = append(event.ApplyData.Contents, content)
	}

	return event, nil
}

// parseTagParams 解析结构体标签参数
func parseTagParams(tag string) map[string]string {
	params := make(map[string]string)
	pairs := splitTag(tag, ";")

	for _, pair := range pairs {
		kv := splitTag(pair, "=")
		if len(kv) == 2 {
			params[kv[0]] = kv[1]
		} else if len(kv) == 1 {
			params[kv[0]] = ""
		}
	}

	return params
}

// splitTag 分割标签字符串
func splitTag(tag, sep string) []string {
	var parts []string
	start := 0
	inQuote := false

	for i := 0; i < len(tag); i++ {
		if tag[i] == '"' {
			inQuote = !inQuote
		} else if !inQuote && tag[i] == sep[0] {
			if i > start {
				parts = append(parts, tag[start:i])
			}
			start = i + 1
		}
	}

	if start < len(tag) {
		parts = append(parts, tag[start:])
	}

	return parts
}

// 以下是各控件的处理函数

func (c *Converter) handleText(field reflect.Value, params map[string]string) OAContentValue {
	return OAContentValue{
		Text: field.String(),
	}
}

func (c *Converter) handleNumber(field reflect.Value, params map[string]string) OAContentValue {
	var numStr string
	switch field.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		numStr = strconv.FormatInt(field.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		numStr = strconv.FormatUint(field.Uint(), 10)
	case reflect.Float32, reflect.Float64:
		numStr = strconv.FormatFloat(field.Float(), 'f', -1, 64)
	default:
		numStr = "0"
	}

	return OAContentValue{
		Number: numStr,
	}
}

func (c *Converter) handleMoney(field reflect.Value, params map[string]string) OAContentValue {
	var moneyStr string
	switch field.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		moneyStr = strconv.FormatInt(field.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		moneyStr = strconv.FormatUint(field.Uint(), 10)
	case reflect.Float32, reflect.Float64:
		moneyStr = strconv.FormatFloat(field.Float(), 'f', 2, 64)
	default:
		moneyStr = "0.00"
	}

	return OAContentValue{
		Money: moneyStr,
	}
}

func (c *Converter) handleDate(field reflect.Value, params map[string]string) OAContentValue {
	var timestamp int64
	var dateType string

	// 确定日期类型
	if t, ok := params[tagType]; ok {
		dateType = t
	} else {
		dateType = "day" // 默认按天
	}

	// 处理不同类型的时间字段
	switch v := field.Interface().(type) {
	case time.Time:
		timestamp = v.Unix()
	case int, int8, int16, int32, int64:
		timestamp = field.Int()
	case uint, uint8, uint16, uint32, uint64:
		timestamp = int64(field.Uint())
	default:
		timestamp = time.Now().Unix()
	}

	return OAContentValue{
		Date: OAContentDate{
			Type:      dateType,
			Timestamp: strconv.FormatInt(timestamp, 10),
		},
	}
}

func (c *Converter) handleSelector(field reflect.Value, params map[string]string) OAContentValue {
	selectorType := "single" // 默认单选
	if t, ok := params[tagSelectorType]; ok {
		selectorType = t
	}

	// 处理多选情况
	if multi, ok := params[tagMultiSelector]; ok && multi == "true" {
		selectorType = "multi"
	}

	var options []OAContentSelectorOption
	if selectorType == "multi" {
		// 处理多选情况（slice类型）
		if field.Kind() == reflect.Slice || field.Kind() == reflect.Array {
			for i := 0; i < field.Len(); i++ {
				options = append(options, OAContentSelectorOption{
					Key: field.Index(i).String(),
				})
			}
		}
	} else {
		// 处理单选情况
		optionKey := field.String()
		if opt, ok := params[tagOption]; ok {
			optionKey = opt
		}
		options = append(options, OAContentSelectorOption{
			Key: optionKey,
		})
	}

	return OAContentValue{
		Selector: OAContentSelector{
			Type:    selectorType,
			Options: options,
		},
	}
}

func (c *Converter) handleContact(field reflect.Value, params map[string]string) OAContentValue {
	contactType := "single" // 默认单选
	if t, ok := params[tagType]; ok {
		contactType = t
	}

	mode := "user" // 默认成员
	if m, ok := params["mode"]; ok {
		mode = m
	}

	if mode == "user" {
		var members []OAContentMember
		if contactType == "multi" && (field.Kind() == reflect.Slice || field.Kind() == reflect.Array) {
			for i := 0; i < field.Len(); i++ {
				members = append(members, OAContentMember{
					UserID: field.Index(i).String(),
				})
			}
		} else {
			members = append(members, OAContentMember{
				UserID: field.String(),
			})
		}
		return OAContentValue{
			Members: members,
		}
	}

	// 处理部门
	var depts []OAContentDepartment
	if contactType == "multi" && (field.Kind() == reflect.Slice || field.Kind() == reflect.Array) {
		for i := 0; i < field.Len(); i++ {
			depts = append(depts, OAContentDepartment{
				OpenAPIID: field.Index(i).String(),
			})
		}
	} else {
		depts = append(depts, OAContentDepartment{
			OpenAPIID: field.String(),
		})
	}
	return OAContentValue{
		Departments: depts,
	}
}

func (c *Converter) handleFile(field reflect.Value, params map[string]string) OAContentValue {
	var files []OAContentFile
	if field.Kind() == reflect.Slice || field.Kind() == reflect.Array {
		for i := 0; i < field.Len(); i++ {
			files = append(files, OAContentFile{
				FileID: field.Index(i).String(),
			})
		}
	} else {
		files = append(files, OAContentFile{
			FileID: field.String(),
		})
	}
	return OAContentValue{
		Files: files,
	}
}

func (c *Converter) handleTable(field reflect.Value, params map[string]string) OAContentValue {
	var table []OAContentTableList
	if field.Kind() == reflect.Slice || field.Kind() == reflect.Array {
		for i := 0; i < field.Len(); i++ {
			item := field.Index(i)
			if item.Kind() == reflect.Ptr {
				item = item.Elem()
			}

			if item.Kind() != reflect.Struct {
				continue
			}

			var list []OAContent
			itemType := item.Type()
			for j := 0; j < item.NumField(); j++ {
				subField := item.Field(j)
				subFieldType := itemType.Field(j)
				tag := subFieldType.Tag.Get(tagName)

				if tag == "" {
					continue
				}

				subParams := parseTagParams(tag)
				controlType := subParams[tagControl]
				if controlType == "" {
					continue
				}

				content := OAContent{
					Control: OAControl(controlType),
					ID:      subParams[tagID],
				}

				if handler, ok := c.handlers[OAControl(controlType)]; ok {
					content.Value = handler(subField, subParams)
				}

				list = append(list, content)
			}

			table = append(table, OAContentTableList{
				List: list,
			})
		}
	}
	return OAContentValue{
		Table: table,
	}
}

// InvalidBusinessDataError 业务数据结构错误
type InvalidBusinessDataError struct {
	message string
}

func (e *InvalidBusinessDataError) Error() string {
	return e.message
}

func NewInvalidBusinessDataError(message string) *InvalidBusinessDataError {
	return &InvalidBusinessDataError{message: message}
}
