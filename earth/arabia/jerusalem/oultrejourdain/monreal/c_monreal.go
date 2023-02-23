package monreal

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MonrealCounty interface {
	feud.County
	BBildad比勒达() feud.Barony
	BHurmniz霍姆尼() feud.Barony
	BIdan伊丹() feud.Barony
	BMonreal蒙雷阿尔() feud.Barony
	BParan巴兰() feud.Barony
	BSela西拉() feud.Barony
	BTafila塔菲拉() feud.Barony
}

type 蒙雷阿尔MonrealCounty struct {
	feud.BaseCounty
	比勒达Bildad   feud.Barony
	霍姆尼Hurmniz  feud.Barony
	伊丹Idan      feud.Barony
	蒙雷阿尔Monreal feud.Barony
	巴兰Paran     feud.Barony
	西拉Sela      feud.Barony
	塔菲拉Tafila   feud.Barony
}

func (c *蒙雷阿尔MonrealCounty) BBildad比勒达() feud.Barony {
	return c.比勒达Bildad
}

func (c *蒙雷阿尔MonrealCounty) BHurmniz霍姆尼() feud.Barony {
	return c.霍姆尼Hurmniz
}

func (c *蒙雷阿尔MonrealCounty) BIdan伊丹() feud.Barony {
	return c.伊丹Idan
}

func (c *蒙雷阿尔MonrealCounty) BMonreal蒙雷阿尔() feud.Barony {
	return c.蒙雷阿尔Monreal
}

func (c *蒙雷阿尔MonrealCounty) BParan巴兰() feud.Barony {
	return c.巴兰Paran
}

func (c *蒙雷阿尔MonrealCounty) BSela西拉() feud.Barony {
	return c.西拉Sela
}

func (c *蒙雷阿尔MonrealCounty) BTafila塔菲拉() feud.Barony {
	return c.塔菲拉Tafila
}

var CMonreal蒙雷阿尔 MonrealCounty = &蒙雷阿尔MonrealCounty{}

func init() {
	f := CMonreal蒙雷阿尔.(*蒙雷阿尔MonrealCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "778",
		Title:     "monreal",
		TitleName: "蒙雷阿尔",
		TitleCode: "c_monreal",
		Baronies:  map[string]feud.Barony{},
	}

	f.比勒达Bildad = BBildad比勒达
	f.比勒达Bildad.SetParent(f)

	f.霍姆尼Hurmniz = BHurmniz霍姆尼
	f.霍姆尼Hurmniz.SetParent(f)

	f.伊丹Idan = BIdan伊丹
	f.伊丹Idan.SetParent(f)

	f.蒙雷阿尔Monreal = BMonreal蒙雷阿尔
	f.蒙雷阿尔Monreal.SetParent(f)

	f.巴兰Paran = BParan巴兰
	f.巴兰Paran.SetParent(f)

	f.西拉Sela = BSela西拉
	f.西拉Sela.SetParent(f)

	f.塔菲拉Tafila = BTafila塔菲拉
	f.塔菲拉Tafila.SetParent(f)

}
