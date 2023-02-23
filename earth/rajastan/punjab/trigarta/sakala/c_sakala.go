package sakala

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SakalaCounty interface {
	feud.County
	BBahu巴胡() feud.Barony
	BBhuttar普多尔() feud.Barony
	BJammu查谟() feud.Barony
	BPasrur帕萨鲁尔() feud.Barony
	BSakala奢羯罗() feud.Barony
	BVallapura筏罗补罗() feud.Barony
}

type 奢羯罗SakalaCounty struct {
	feud.BaseCounty
	巴胡Bahu        feud.Barony
	普多尔Bhuttar    feud.Barony
	查谟Jammu       feud.Barony
	帕萨鲁尔Pasrur    feud.Barony
	奢羯罗Sakala     feud.Barony
	筏罗补罗Vallapura feud.Barony
}

func (c *奢羯罗SakalaCounty) BBahu巴胡() feud.Barony {
	return c.巴胡Bahu
}

func (c *奢羯罗SakalaCounty) BBhuttar普多尔() feud.Barony {
	return c.普多尔Bhuttar
}

func (c *奢羯罗SakalaCounty) BJammu查谟() feud.Barony {
	return c.查谟Jammu
}

func (c *奢羯罗SakalaCounty) BPasrur帕萨鲁尔() feud.Barony {
	return c.帕萨鲁尔Pasrur
}

func (c *奢羯罗SakalaCounty) BSakala奢羯罗() feud.Barony {
	return c.奢羯罗Sakala
}

func (c *奢羯罗SakalaCounty) BVallapura筏罗补罗() feud.Barony {
	return c.筏罗补罗Vallapura
}

var CSakala奢羯罗 SakalaCounty = &奢羯罗SakalaCounty{}

func init() {
	f := CSakala奢羯罗.(*奢羯罗SakalaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1179",
		Title:     "sakala",
		TitleName: "奢羯罗",
		TitleCode: "c_sakala",
		Baronies:  map[string]feud.Barony{},
	}

	f.巴胡Bahu = BBahu巴胡
	f.巴胡Bahu.SetParent(f)

	f.普多尔Bhuttar = BBhuttar普多尔
	f.普多尔Bhuttar.SetParent(f)

	f.查谟Jammu = BJammu查谟
	f.查谟Jammu.SetParent(f)

	f.帕萨鲁尔Pasrur = BPasrur帕萨鲁尔
	f.帕萨鲁尔Pasrur.SetParent(f)

	f.奢羯罗Sakala = BSakala奢羯罗
	f.奢羯罗Sakala.SetParent(f)

	f.筏罗补罗Vallapura = BVallapura筏罗补罗
	f.筏罗补罗Vallapura.SetParent(f)

}
