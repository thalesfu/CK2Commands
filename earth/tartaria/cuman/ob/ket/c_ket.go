package ket

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KetCounty interface {
	feud.County
	BAlchach阿尔恰奇() feud.Barony
	BAsino阿西诺() feud.Barony
	BChachamga恰恰姆加() feud.Barony
	BKet克季() feud.Barony
	BKetino克蒂诺() feud.Barony
	BKetkas憩加斯() feud.Barony
	BSochur索丘尔() feud.Barony
}

type 阿努伊KetCounty struct {
	feud.BaseCounty
	阿尔恰奇Alchach   feud.Barony
	阿西诺Asino      feud.Barony
	恰恰姆加Chachamga feud.Barony
	克季Ket         feud.Barony
	克蒂诺Ketino     feud.Barony
	憩加斯Ketkas     feud.Barony
	索丘尔Sochur     feud.Barony
}

func (c *阿努伊KetCounty) BAlchach阿尔恰奇() feud.Barony {
	return c.阿尔恰奇Alchach
}

func (c *阿努伊KetCounty) BAsino阿西诺() feud.Barony {
	return c.阿西诺Asino
}

func (c *阿努伊KetCounty) BChachamga恰恰姆加() feud.Barony {
	return c.恰恰姆加Chachamga
}

func (c *阿努伊KetCounty) BKet克季() feud.Barony {
	return c.克季Ket
}

func (c *阿努伊KetCounty) BKetino克蒂诺() feud.Barony {
	return c.克蒂诺Ketino
}

func (c *阿努伊KetCounty) BKetkas憩加斯() feud.Barony {
	return c.憩加斯Ketkas
}

func (c *阿努伊KetCounty) BSochur索丘尔() feud.Barony {
	return c.索丘尔Sochur
}

var CKet阿努伊 KetCounty = &阿努伊KetCounty{}

func init() {
	f := CKet阿努伊.(*阿努伊KetCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1329",
		Title:     "ket",
		TitleName: "阿努伊",
		TitleCode: "c_ket",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿尔恰奇Alchach = BAlchach阿尔恰奇
	f.阿尔恰奇Alchach.SetParent(f)

	f.阿西诺Asino = BAsino阿西诺
	f.阿西诺Asino.SetParent(f)

	f.恰恰姆加Chachamga = BChachamga恰恰姆加
	f.恰恰姆加Chachamga.SetParent(f)

	f.克季Ket = BKet克季
	f.克季Ket.SetParent(f)

	f.克蒂诺Ketino = BKetino克蒂诺
	f.克蒂诺Ketino.SetParent(f)

	f.憩加斯Ketkas = BKetkas憩加斯
	f.憩加斯Ketkas.SetParent(f)

	f.索丘尔Sochur = BSochur索丘尔
	f.索丘尔Sochur.SetParent(f)

}
