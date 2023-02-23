package ephesos

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type EphesosCounty interface {
	feud.County
	BEphesos以弗所() feud.Barony
	BIassos伊阿索斯() feud.Barony
	BLebedos雷贝多斯() feud.Barony
	BMiletos米利都() feud.Barony
	BPalation帕拉提翁() feud.Barony
	BPetron帕特仑() feud.Barony
	BTralles特拉莱斯() feud.Barony
}

type 以弗所EphesosCounty struct {
	feud.BaseCounty
	以弗所Ephesos   feud.Barony
	伊阿索斯Iassos   feud.Barony
	雷贝多斯Lebedos  feud.Barony
	米利都Miletos   feud.Barony
	帕拉提翁Palation feud.Barony
	帕特仑Petron    feud.Barony
	特拉莱斯Tralles  feud.Barony
}

func (c *以弗所EphesosCounty) BEphesos以弗所() feud.Barony {
	return c.以弗所Ephesos
}

func (c *以弗所EphesosCounty) BIassos伊阿索斯() feud.Barony {
	return c.伊阿索斯Iassos
}

func (c *以弗所EphesosCounty) BLebedos雷贝多斯() feud.Barony {
	return c.雷贝多斯Lebedos
}

func (c *以弗所EphesosCounty) BMiletos米利都() feud.Barony {
	return c.米利都Miletos
}

func (c *以弗所EphesosCounty) BPalation帕拉提翁() feud.Barony {
	return c.帕拉提翁Palation
}

func (c *以弗所EphesosCounty) BPetron帕特仑() feud.Barony {
	return c.帕特仑Petron
}

func (c *以弗所EphesosCounty) BTralles特拉莱斯() feud.Barony {
	return c.特拉莱斯Tralles
}

var CEphesos以弗所 EphesosCounty = &以弗所EphesosCounty{}

func init() {
	f := CEphesos以弗所.(*以弗所EphesosCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "746",
		Title:     "ephesos",
		TitleName: "以弗所",
		TitleCode: "c_ephesos",
		Baronies:  map[string]feud.Barony{},
	}

	f.以弗所Ephesos = BEphesos以弗所
	f.以弗所Ephesos.SetParent(f)

	f.伊阿索斯Iassos = BIassos伊阿索斯
	f.伊阿索斯Iassos.SetParent(f)

	f.雷贝多斯Lebedos = BLebedos雷贝多斯
	f.雷贝多斯Lebedos.SetParent(f)

	f.米利都Miletos = BMiletos米利都
	f.米利都Miletos.SetParent(f)

	f.帕拉提翁Palation = BPalation帕拉提翁
	f.帕拉提翁Palation.SetParent(f)

	f.帕特仑Petron = BPetron帕特仑
	f.帕特仑Petron.SetParent(f)

	f.特拉莱斯Tralles = BTralles特拉莱斯
	f.特拉莱斯Tralles.SetParent(f)

}
