package zaysan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ZaysanCounty interface {
	feud.County
	BIvanovka伊万诺夫卡() feud.Barony
	BNovaya诺瓦亚() feud.Barony
	BUbinka乌宾卡() feud.Barony
	BUlba乌尔巴() feud.Barony
	BVolchanka沃尔昌卡() feud.Barony
	BZaysan斋桑() feud.Barony
	BZubair祖拜尔() feud.Barony
}

type 斋桑ZaysanCounty struct {
	feud.BaseCounty
	伊万诺夫卡Ivanovka feud.Barony
	诺瓦亚Novaya     feud.Barony
	乌宾卡Ubinka     feud.Barony
	乌尔巴Ulba       feud.Barony
	沃尔昌卡Volchanka feud.Barony
	斋桑Zaysan      feud.Barony
	祖拜尔Zubair     feud.Barony
}

func (c *斋桑ZaysanCounty) BIvanovka伊万诺夫卡() feud.Barony {
	return c.伊万诺夫卡Ivanovka
}

func (c *斋桑ZaysanCounty) BNovaya诺瓦亚() feud.Barony {
	return c.诺瓦亚Novaya
}

func (c *斋桑ZaysanCounty) BUbinka乌宾卡() feud.Barony {
	return c.乌宾卡Ubinka
}

func (c *斋桑ZaysanCounty) BUlba乌尔巴() feud.Barony {
	return c.乌尔巴Ulba
}

func (c *斋桑ZaysanCounty) BVolchanka沃尔昌卡() feud.Barony {
	return c.沃尔昌卡Volchanka
}

func (c *斋桑ZaysanCounty) BZaysan斋桑() feud.Barony {
	return c.斋桑Zaysan
}

func (c *斋桑ZaysanCounty) BZubair祖拜尔() feud.Barony {
	return c.祖拜尔Zubair
}

var CZaysan斋桑 ZaysanCounty = &斋桑ZaysanCounty{}

func init() {
	f := CZaysan斋桑.(*斋桑ZaysanCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1872",
		Title:     "zaysan",
		TitleName: "斋桑",
		TitleCode: "c_zaysan",
		Baronies:  map[string]feud.Barony{},
	}

	f.伊万诺夫卡Ivanovka = BIvanovka伊万诺夫卡
	f.伊万诺夫卡Ivanovka.SetParent(f)

	f.诺瓦亚Novaya = BNovaya诺瓦亚
	f.诺瓦亚Novaya.SetParent(f)

	f.乌宾卡Ubinka = BUbinka乌宾卡
	f.乌宾卡Ubinka.SetParent(f)

	f.乌尔巴Ulba = BUlba乌尔巴
	f.乌尔巴Ulba.SetParent(f)

	f.沃尔昌卡Volchanka = BVolchanka沃尔昌卡
	f.沃尔昌卡Volchanka.SetParent(f)

	f.斋桑Zaysan = BZaysan斋桑
	f.斋桑Zaysan.SetParent(f)

	f.祖拜尔Zubair = BZubair祖拜尔
	f.祖拜尔Zubair.SetParent(f)

}
