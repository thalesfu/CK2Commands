package jaunpur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type JaunpurCounty interface {
	feud.County
	BAmethi阿咩底() feud.Barony
	BDohrighat豆利揭() feud.Barony
	BGadhipuri迦提城() feud.Barony
	BHaldi诃罗提() feud.Barony
	BJaunpur善补罗() feud.Barony
	BKerarkot计剌罗拘吒() feud.Barony
	BTanda檀陀() feud.Barony
}

type 善补罗JaunpurCounty struct {
	feud.BaseCounty
	阿咩底Amethi     feud.Barony
	豆利揭Dohrighat  feud.Barony
	迦提城Gadhipuri  feud.Barony
	诃罗提Haldi      feud.Barony
	善补罗Jaunpur    feud.Barony
	计剌罗拘吒Kerarkot feud.Barony
	檀陀Tanda       feud.Barony
}

func (c *善补罗JaunpurCounty) BAmethi阿咩底() feud.Barony {
	return c.阿咩底Amethi
}

func (c *善补罗JaunpurCounty) BDohrighat豆利揭() feud.Barony {
	return c.豆利揭Dohrighat
}

func (c *善补罗JaunpurCounty) BGadhipuri迦提城() feud.Barony {
	return c.迦提城Gadhipuri
}

func (c *善补罗JaunpurCounty) BHaldi诃罗提() feud.Barony {
	return c.诃罗提Haldi
}

func (c *善补罗JaunpurCounty) BJaunpur善补罗() feud.Barony {
	return c.善补罗Jaunpur
}

func (c *善补罗JaunpurCounty) BKerarkot计剌罗拘吒() feud.Barony {
	return c.计剌罗拘吒Kerarkot
}

func (c *善补罗JaunpurCounty) BTanda檀陀() feud.Barony {
	return c.檀陀Tanda
}

var CJaunpur善补罗 JaunpurCounty = &善补罗JaunpurCounty{}

func init() {
	f := CJaunpur善补罗.(*善补罗JaunpurCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1166",
		Title:     "jaunpur",
		TitleName: "善补罗",
		TitleCode: "c_jaunpur",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿咩底Amethi = BAmethi阿咩底
	f.阿咩底Amethi.SetParent(f)

	f.豆利揭Dohrighat = BDohrighat豆利揭
	f.豆利揭Dohrighat.SetParent(f)

	f.迦提城Gadhipuri = BGadhipuri迦提城
	f.迦提城Gadhipuri.SetParent(f)

	f.诃罗提Haldi = BHaldi诃罗提
	f.诃罗提Haldi.SetParent(f)

	f.善补罗Jaunpur = BJaunpur善补罗
	f.善补罗Jaunpur.SetParent(f)

	f.计剌罗拘吒Kerarkot = BKerarkot计剌罗拘吒
	f.计剌罗拘吒Kerarkot.SetParent(f)

	f.檀陀Tanda = BTanda檀陀
	f.檀陀Tanda.SetParent(f)

}
