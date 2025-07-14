package workwx

import (
	"encoding/json"
	"fmt"
	"github.com/bitly/go-simplejson"
	c "github.com/smartystreets/goconvey/convey"
	"strconv"
	"testing"
	"time"
)

type BusinessApply struct {
	Applicant  string `oa:"creator"`  // 特殊字段：申请人
	TemplateID string `oa:"template"` // 特殊字段：模板ID

	//履约主体
	PerformanceSubject string `oa:"control=Selector;id=Selector-1745736605666;option=option-1745736605666"`
	CompanyName        string `oa:"control=Text;id=Text-1745736790921"`
	//ApplyDate          time.Time `oa:"control=Date;id=Date-123;type=day"`
	//Department string  `oa:"control=Selector;id=Selector-456;option=dept_option1"`
	Reason                 string  `oa:"control=Text;id=Text-1745736339267"`
	Amount                 float64 `oa:"control=Money;id=Money-1745736827101"`
	AttachmentID           string  `oa:"control=File;id=File-1745824439696"`
	IsRegulatoryApproval   string  `oa:"control=Selector;id=Selector-1745827549537;option=option-1745827549538"` // 是否监管部门审批
	SettlementBasisType    string  `oa:"control=Selector;id=Selector-1745824160158"`                             // 结算依据提交情况
	OtherDeductions        float64 `oa:"control=Money;id=Money-1745823950698"`                                   // 其他扣款
	ActualPaymentAmount    float64 `oa:"control=Money;id=Money-1745824081652"`                                   // 实际付款金额
	ActualPaymentAmountStr string  `oa:"control=Textarea;id=Textarea-1750233243207"`                             // 实际付款金额描述
	// 明细控件
	Items []Item `oa:"control=Table;id=Table-1745823552331"`
}

// 明细项结构体
type Item struct {
	ProjectCode string  `oa:"control=Text;id=Text-1745823631415"`
	ProjectName string  `oa:"control=Text;id=Text-1745823655223"`
	Price       float64 `oa:"control=Money;id=Money-1745824059595"`
}

func TestBusinessApply(t *testing.T) {

	// 准备业务数据
	apply := BusinessApply{
		PerformanceSubject: "option-1745736605666",
		CompanyName:        "某个乙方有限公司",
		Applicant:          "ZhouEnXian",
		TemplateID:         "3WLuFTUtLstd52F3Fai9ZbeEZphsdGVWTd7FKAsm",
		Reason:             "采购申请",
		Amount:             1999.99,
		//AttachmentID:       mRes.MediaID,
		//AttachmentID: "",
		IsRegulatoryApproval:   "option-1745827549538",
		SettlementBasisType:    "",
		OtherDeductions:        12,
		ActualPaymentAmount:    12321.32,
		ActualPaymentAmountStr: "(12.32)",
		Items: []Item{
			{
				ProjectCode: "20020001",
				ProjectName: "项目1",
				Price:       8999,
			},
			{
				ProjectCode: "20020002",
				ProjectName: "项目2",
				Price:       1999,
			},
		},
	}

	// 创建转换器
	converter := NewConverter("", "") // 参数会被业务数据覆盖
	converter.UseTemplateApprover(true)

	// 转换业务数据
	oaEvent, err := converter.Parse(&apply)
	if err != nil {
		fmt.Printf("转换失败: %v\n", err)
		return
	}

	// 输出结果检查
	resJSON, _ := json.MarshalIndent(oaEvent, "", "    ")
	fmt.Println(string(resJSON))

	//res, err := app.ApplyOAEvent(*oaEvent)
	//if err != nil {
	//	fmt.Printf("提交审批失败: %v\n", err)
	//	return
	//}

	//resJSON, _ = json.MarshalIndent(res, "", "    ")
	//// 输出结果
	//fmt.Println(string(resJSON))
}

func TestConverter(t *testing.T) {
	// 设置JSON输出
	reporter := NewJsonReporter()
	defer func() {
		// 在测试结束后输出JSON结果
		output, _ := json.Marshal(reporter)
		t.Logf("Test Results:\n%s\n", string(output))
	}()

	c.Convey("企业微信审批转换器测试", t, func() {
		c.Convey("基本字段转换", func() {
			type SimpleApply struct {
				Applicant  string    `oa:"creator"`
				TemplateID string    `oa:"template"`
				ApplyDate  time.Time `oa:"control=Date;id=Date-1;type=day"`
				Reason     string    `oa:"control=Text;id=Text-1"`
			}

			now := time.Now()
			apply := SimpleApply{
				Applicant:  "test_user",
				TemplateID: "TPL_SIMPLE",
				ApplyDate:  now,
				Reason:     "测试原因",
			}

			converter := NewConverter("", "")
			event, err := converter.Parse(&apply)

			c.So(err, c.ShouldBeNil)
			c.So(event, c.ShouldNotBeNil)
			c.So(event.CreatorUserID, c.ShouldEqual, "test_user")
			c.So(event.TemplateID, c.ShouldEqual, "TPL_SIMPLE")
			c.So(len(event.ApplyData.Contents), c.ShouldEqual, 2)

			// 验证日期控件
			dateContent := event.ApplyData.Contents[0]
			c.So(dateContent.Control, c.ShouldEqual, OAControlDate)
			c.So(dateContent.ID, c.ShouldEqual, "Date-1")
			c.So(dateContent.Value.Date.Type, c.ShouldEqual, "day")
			c.So(dateContent.Value.Date.Timestamp, c.ShouldEqual, strconv.FormatInt(now.Unix(), 10))

			// 验证文本控件
			textContent := event.ApplyData.Contents[1]
			c.So(textContent.Control, c.ShouldEqual, OAControlText)
			c.So(textContent.ID, c.ShouldEqual, "Text-1")
			c.So(textContent.Value.Text, c.ShouldEqual, "测试原因")
		})

		c.Convey("数字和金额转换", func() {
			type NumberApply struct {
				Quantity int     `oa:"control=Number;id=Number-1"`
				Amount   float64 `oa:"control=Money;id=Money-1"`
			}

			apply := NumberApply{
				Quantity: 5,
				Amount:   1234.56,
			}

			converter := NewConverter("TPL_NUMBER", "number_user")
			event, err := converter.Parse(&apply)

			c.So(err, c.ShouldBeNil)
			c.So(len(event.ApplyData.Contents), c.ShouldEqual, 2)

			// 验证数字控件
			numContent := event.ApplyData.Contents[0]
			c.So(numContent.Control, c.ShouldEqual, OAControlNumber)
			c.So(numContent.Value.Number, c.ShouldEqual, "5")

			// 验证金额控件
			moneyContent := event.ApplyData.Contents[1]
			c.So(moneyContent.Control, c.ShouldEqual, OAControlMoney)
			c.So(moneyContent.Value.Money, c.ShouldEqual, "1234.56")
		})

		c.Convey("选择器转换", func() {
			c.Convey("单选选择器", func() {
				type SingleSelectApply struct {
					Dept string `oa:"control=Selector;id=Selector-1;option=dept_1"`
				}

				apply := SingleSelectApply{Dept: "dept_1"}
				event, err := NewConverter("", "").Parse(&apply)

				c.So(err, c.ShouldBeNil)
				c.So(len(event.ApplyData.Contents), c.ShouldEqual, 1)

				selector := event.ApplyData.Contents[0].Value.Selector
				c.So(selector.Type, c.ShouldEqual, "single")
				c.So(len(selector.Options), c.ShouldEqual, 1)
				c.So(selector.Options[0].Key, c.ShouldEqual, "dept_1")
			})

			c.Convey("多选选择器", func() {
				type MultiSelectApply struct {
					Tags []string `oa:"control=Selector;id=Selector-2;selector_type=multi"`
				}

				apply := MultiSelectApply{Tags: []string{"tag1", "tag2"}}
				event, err := NewConverter("", "").Parse(&apply)

				c.So(err, c.ShouldBeNil)
				selector := event.ApplyData.Contents[0].Value.Selector
				c.So(selector.Type, c.ShouldEqual, "multi")
				c.So(len(selector.Options), c.ShouldEqual, 2)
				c.So(selector.Options[0].Key, c.ShouldEqual, "tag1")
				c.So(selector.Options[1].Key, c.ShouldEqual, "tag2")
			})
		})

		c.Convey("明细表格转换", func() {
			type TableItem struct {
				Product  string  `oa:"control=Text;id=Text-2"`
				Quantity int     `oa:"control=Number;id=Number-2"`
				Price    float64 `oa:"control=Money;id=Money-2"`
			}

			type TableApply struct {
				Items []TableItem `oa:"control=Table;id=Table-1"`
			}

			apply := TableApply{
				Items: []TableItem{
					{
						Product:  "商品A",
						Quantity: 2,
						Price:    99.99,
					},
					{
						Product:  "商品B",
						Quantity: 3,
						Price:    199.99,
					},
				},
			}

			event, err := NewConverter("", "").Parse(&apply)
			c.So(err, c.ShouldBeNil)

			table := event.ApplyData.Contents[0].Value.Table
			c.So(len(table), c.ShouldEqual, 2)

			// 验证第一行
			firstRow := table[0].List
			c.So(len(firstRow), c.ShouldEqual, 3)
			c.So(firstRow[0].Value.Text, c.ShouldEqual, "商品A")
			c.So(firstRow[1].Value.Number, c.ShouldEqual, "2")
			c.So(firstRow[2].Value.Money, c.ShouldEqual, "99.99")

			// 验证第二行
			secondRow := table[1].List
			c.So(secondRow[0].Value.Text, c.ShouldEqual, "商品B")
			c.So(secondRow[1].Value.Number, c.ShouldEqual, "3")
			c.So(secondRow[2].Value.Money, c.ShouldEqual, "199.99")
		})

		c.Convey("错误处理", func() {
			c.Convey("非结构体输入", func() {
				_, err := NewConverter("", "").Parse("not a struct")
				c.So(err, c.ShouldNotBeNil)
				c.So(err.Error(), c.ShouldContainSubstring, "must be a struct or pointer to struct")
			})

			c.Convey("无效控件类型", func() {
				type InvalidApply struct {
					Field string `oa:"control=InvalidType;id=Invalid-1"`
				}

				_, err := NewConverter("", "").Parse(&InvalidApply{})
				c.So(err, c.ShouldBeNil) // 当前实现会忽略无法处理的控件
			})
		})
	})
}

// JsonReporter 自定义JSON报告器
type JsonReporter struct {
	Results []TestResult
}

type TestResult struct {
	TestName string      `json:"test"`
	Outcome  string      `json:"outcome"`
	Message  string      `json:"message,omitempty"`
	Data     interface{} `json:"data,omitempty"`
}

func NewJsonReporter() *JsonReporter {
	return &JsonReporter{
		Results: make([]TestResult, 0),
	}
}

func (r *JsonReporter) Report(testName, outcome, message string) {
	r.Results = append(r.Results, TestResult{
		TestName: testName,
		Outcome:  outcome,
		Message:  message,
	})
}

func (r *JsonReporter) MarshalJSON() ([]byte, error) {
	js := simplejson.New()
	for _, result := range r.Results {
		js.Set(result.TestName, map[string]interface{}{
			"outcome": result.Outcome,
			"message": result.Message,
			"data":    result.Data,
		})
	}
	return js.MarshalJSON()
}
