package meath

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MeathCounty interface {
	feud.County
	BCiannachta奇纳赫塔() feud.Barony
	BForradh福拉() feud.Barony
	BKnowth瑙斯() feud.Barony
}

type 米斯MeathCounty struct {
	feud.BaseCounty
	奇纳赫塔Ciannachta feud.Barony
	福拉Forradh      feud.Barony
	瑙斯Knowth       feud.Barony
}

func (c *米斯MeathCounty) BCiannachta奇纳赫塔() feud.Barony {
	return c.奇纳赫塔Ciannachta
}

func (c *米斯MeathCounty) BForradh福拉() feud.Barony {
	return c.福拉Forradh
}

func (c *米斯MeathCounty) BKnowth瑙斯() feud.Barony {
	return c.瑙斯Knowth
}

var CMeath米斯 MeathCounty = &米斯MeathCounty{}

func init() {
	f := CMeath米斯.(*米斯MeathCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1951",
		Title:     "meath",
		TitleName: "米斯",
		TitleCode: "c_meath",
		Baronies:  map[string]feud.Barony{},
	}

	f.奇纳赫塔Ciannachta = BCiannachta奇纳赫塔
	f.奇纳赫塔Ciannachta.SetParent(f)

	f.福拉Forradh = BForradh福拉
	f.福拉Forradh.SetParent(f)

	f.瑙斯Knowth = BKnowth瑙斯
	f.瑙斯Knowth.SetParent(f)

}
