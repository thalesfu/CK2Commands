package delingha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type DelinghaCounty interface {
	feud.County
	BDelingha德令哈() feud.Barony
	BGahai尕海() feud.Barony
	BHaixi海西() feud.Barony
	BHedong河东() feud.Barony
	BHexi河西() feud.Barony
	BKeluke柯鲁柯() feud.Barony
	BXuji蓄集() feud.Barony
}

type 德令哈DelinghaCounty struct {
	feud.BaseCounty
	德令哈Delingha feud.Barony
	尕海Gahai     feud.Barony
	海西Haixi     feud.Barony
	河东Hedong    feud.Barony
	河西Hexi      feud.Barony
	柯鲁柯Keluke   feud.Barony
	蓄集Xuji      feud.Barony
}

func (c *德令哈DelinghaCounty) BDelingha德令哈() feud.Barony {
	return c.德令哈Delingha
}

func (c *德令哈DelinghaCounty) BGahai尕海() feud.Barony {
	return c.尕海Gahai
}

func (c *德令哈DelinghaCounty) BHaixi海西() feud.Barony {
	return c.海西Haixi
}

func (c *德令哈DelinghaCounty) BHedong河东() feud.Barony {
	return c.河东Hedong
}

func (c *德令哈DelinghaCounty) BHexi河西() feud.Barony {
	return c.河西Hexi
}

func (c *德令哈DelinghaCounty) BKeluke柯鲁柯() feud.Barony {
	return c.柯鲁柯Keluke
}

func (c *德令哈DelinghaCounty) BXuji蓄集() feud.Barony {
	return c.蓄集Xuji
}

var CDelingha德令哈 DelinghaCounty = &德令哈DelinghaCounty{}

func init() {
	f := CDelingha德令哈.(*德令哈DelinghaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1569",
		Title:     "delingha",
		TitleName: "德令哈",
		TitleCode: "c_delingha",
		Baronies:  map[string]feud.Barony{},
	}

	f.德令哈Delingha = BDelingha德令哈
	f.德令哈Delingha.SetParent(f)

	f.尕海Gahai = BGahai尕海
	f.尕海Gahai.SetParent(f)

	f.海西Haixi = BHaixi海西
	f.海西Haixi.SetParent(f)

	f.河东Hedong = BHedong河东
	f.河东Hedong.SetParent(f)

	f.河西Hexi = BHexi河西
	f.河西Hexi.SetParent(f)

	f.柯鲁柯Keluke = BKeluke柯鲁柯
	f.柯鲁柯Keluke.SetParent(f)

	f.蓄集Xuji = BXuji蓄集
	f.蓄集Xuji.SetParent(f)

}
