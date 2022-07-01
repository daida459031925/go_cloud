package goTest

import (
	"fmt"
	err "github.com/daida459031925/common/error"
	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/text/feature/plural"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"golang.org/x/text/message/catalog"
	"reflect"
	"testing"
)

var i18nMap = make(map[string]*message.Printer)

func init() {
	initEnglish()
	initChinese()
	initNewPrinter()
}

func initNewPrinter() {
	i18nMap["zh"] = message.NewPrinter(language.Chinese)
	i18nMap["en"] = message.NewPrinter(language.English)
	logx.Info("国际化翻译器初始化成功")
}

func initEnglish() {
	message.Set(language.English, "我有%d个苹果",
		plural.Selectf(1, "%d",
			"=1", "I have an apple",
			"=2", "I have two apples",
			"other", "I have %[1]d apples",
		))
	message.Set(language.English, "还剩余%d天",
		plural.Selectf(1, "%d",
			"one", "One day left",
			"other", "%[1]d days left",
		))

	message.Set(language.English, "你迟了%d分钟。",
		catalog.Var("m", plural.Selectf(1, "%d",
			"one", "minute",
			"other", "minutes")),
		catalog.String("You are %[1]d ${m} late."))
	logx.Info("初始化英文翻译内容")
}

func initChinese() {
	message.SetString(language.Chinese, "%s went to %s.", "%s去了%s。")
	message.SetString(language.Chinese, "%s has been stolen.", "%s被偷走了。")
	message.SetString(language.Chinese, "How are you?", "你好吗?.")
	logx.Info("初始化中文翻译内容")
}

func GetI18nTrans(trans string) (*message.Printer, error) {
	item := i18nMap[trans]
	if item == nil {
		return nil, err.New("未找到对应翻译器")
	}
	return item, nil
}

func TestI18n(t *testing.T) {
	//itemzh, err := GetI18nTrans("zh1")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	//itemen, err := GetI18nTrans("en2")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}

	//指针

	animal := Animal{}
	animal1 := animal
	animal2 := animal1
	value := reflect.ValueOf(animal2)
	if value.Kind() != reflect.Ptr {
		value = reflect.ValueOf(&animal2)
	}
	f := value.MethodByName("Eat") //通过反射获取它对应的函数，然后通过call来调用
	fmt.Println(f.Call([]reflect.Value{}))

	//p := itemzh
	//fmt.Println(p.Sprintf("%s went to %s.", "彼得", "英格兰"))
	//fmt.Println(p.Sprintf("%s has been stolen.", "宝石"))
	//
	//p = message.NewPrinter(language.AmericanEnglish)
	//
	//fmt.Println(p.Sprintf("%s went to %s.", "Peter", "England"))
	//fmt.Println(p.Sprintf("%s has been stolen.", "The Gem"))
	//
	//p = itemen
	//p.Printf("我有 %d 个苹果", 1)
	//fmt.Println()
	//p.Printf("我有 %d 个苹果", 2)
	//fmt.Println()
	//p.Printf("我有 %d 个苹果", 5)
	//fmt.Println()
	//p.Printf("还剩余 %d 天", 1)
	//fmt.Println()
	//p.Printf("还剩余 %d 天", 10)
	//fmt.Println()
	//
	//p = message.NewPrinter(language.English)
	//p.Printf("你迟了 %d 分钟。", 1)
	//// 打印： You are 1 minute late.
	//fmt.Println()
	//p.Printf("你迟了 %d 分钟。", 10)
	//// 打印： You are 10 minutes late.
	//fmt.Println()
	//
	//p = message.NewPrinter(language.English)
	//p.Printf("%d", currency.Symbol(currency.USD.Amount(0.1)))
	//fmt.Println()
	//p.Printf("%d", currency.NarrowSymbol(currency.JPY.Amount(1.6)))
	//fmt.Println()
	//p.Printf("%d", currency.ISO.Kind(currency.Cash)(currency.EUR.Amount(12.255)))
	//fmt.Println()
	//p.Printf("%d", currency.EUR.Amount(12.255))
	//fmt.Println()
}

type Animal struct {
}

func (m *Animal) Eat() string {
	return "Eat111111111111"
}
