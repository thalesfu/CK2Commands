package osma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type OsmaCounty interface {
	feud.County
	BCalatanazor卡拉塔尼亚索尔() feud.Barony
	BDuero杜罗() feud.Barony
	BOsma奥斯马() feud.Barony
	BUcero乌塞罗() feud.Barony
	BUxama乌克萨马() feud.Barony
}

type 奥斯马OsmaCounty struct {
	feud.BaseCounty
	卡拉塔尼亚索尔Calatanazor feud.Barony
	杜罗Duero            feud.Barony
	奥斯马Osma            feud.Barony
	乌塞罗Ucero           feud.Barony
	乌克萨马Uxama          feud.Barony
}

func (c *奥斯马OsmaCounty) BCalatanazor卡拉塔尼亚索尔() feud.Barony {
	return c.卡拉塔尼亚索尔Calatanazor
}

func (c *奥斯马OsmaCounty) BDuero杜罗() feud.Barony {
	return c.杜罗Duero
}

func (c *奥斯马OsmaCounty) BOsma奥斯马() feud.Barony {
	return c.奥斯马Osma
}

func (c *奥斯马OsmaCounty) BUcero乌塞罗() feud.Barony {
	return c.乌塞罗Ucero
}

func (c *奥斯马OsmaCounty) BUxama乌克萨马() feud.Barony {
	return c.乌克萨马Uxama
}

var COsma奥斯马 OsmaCounty = &奥斯马OsmaCounty{}

func init() {
	f := COsma奥斯马.(*奥斯马OsmaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1977",
		Title:     "osma",
		TitleName: "奥斯马",
		TitleCode: "c_osma",
		Baronies:  map[string]feud.Barony{},
	}

	f.卡拉塔尼亚索尔Calatanazor = BCalatanazor卡拉塔尼亚索尔
	f.卡拉塔尼亚索尔Calatanazor.SetParent(f)

	f.杜罗Duero = BDuero杜罗
	f.杜罗Duero.SetParent(f)

	f.奥斯马Osma = BOsma奥斯马
	f.奥斯马Osma.SetParent(f)

	f.乌塞罗Ucero = BUcero乌塞罗
	f.乌塞罗Ucero.SetParent(f)

	f.乌克萨马Uxama = BUxama乌克萨马
	f.乌克萨马Uxama.SetParent(f)

}
