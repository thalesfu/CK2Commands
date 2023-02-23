package kota

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KotaCounty interface {
	feud.County
	BBundi本迪() feud.Barony
	BKoshvardhan俱舍伐弹那() feud.Barony
	BKota拘吒() feud.Barony
	BRamgarh罗摩姞利呬() feud.Barony
	BUlui乌卢伊() feud.Barony
	BVairampatti毗蓝钵胝() feud.Barony
	BVana婆那() feud.Barony
}

type 拘吒KotaCounty struct {
	feud.BaseCounty
	本迪Bundi          feud.Barony
	俱舍伐弹那Koshvardhan feud.Barony
	拘吒Kota           feud.Barony
	罗摩姞利呬Ramgarh     feud.Barony
	乌卢伊Ului          feud.Barony
	毗蓝钵胝Vairampatti  feud.Barony
	婆那Vana           feud.Barony
}

func (c *拘吒KotaCounty) BBundi本迪() feud.Barony {
	return c.本迪Bundi
}

func (c *拘吒KotaCounty) BKoshvardhan俱舍伐弹那() feud.Barony {
	return c.俱舍伐弹那Koshvardhan
}

func (c *拘吒KotaCounty) BKota拘吒() feud.Barony {
	return c.拘吒Kota
}

func (c *拘吒KotaCounty) BRamgarh罗摩姞利呬() feud.Barony {
	return c.罗摩姞利呬Ramgarh
}

func (c *拘吒KotaCounty) BUlui乌卢伊() feud.Barony {
	return c.乌卢伊Ului
}

func (c *拘吒KotaCounty) BVairampatti毗蓝钵胝() feud.Barony {
	return c.毗蓝钵胝Vairampatti
}

func (c *拘吒KotaCounty) BVana婆那() feud.Barony {
	return c.婆那Vana
}

var CKota拘吒 KotaCounty = &拘吒KotaCounty{}

func init() {
	f := CKota拘吒.(*拘吒KotaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1300",
		Title:     "kota",
		TitleName: "拘吒",
		TitleCode: "c_kota",
		Baronies:  map[string]feud.Barony{},
	}

	f.本迪Bundi = BBundi本迪
	f.本迪Bundi.SetParent(f)

	f.俱舍伐弹那Koshvardhan = BKoshvardhan俱舍伐弹那
	f.俱舍伐弹那Koshvardhan.SetParent(f)

	f.拘吒Kota = BKota拘吒
	f.拘吒Kota.SetParent(f)

	f.罗摩姞利呬Ramgarh = BRamgarh罗摩姞利呬
	f.罗摩姞利呬Ramgarh.SetParent(f)

	f.乌卢伊Ului = BUlui乌卢伊
	f.乌卢伊Ului.SetParent(f)

	f.毗蓝钵胝Vairampatti = BVairampatti毗蓝钵胝
	f.毗蓝钵胝Vairampatti.SetParent(f)

	f.婆那Vana = BVana婆那
	f.婆那Vana.SetParent(f)

}
