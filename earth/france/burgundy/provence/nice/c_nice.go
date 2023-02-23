package nice

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type NiceCounty interface {
	feud.County
	BAntibes昂蒂布() feud.Barony
	BCampogrosso坎波格罗索() feud.Barony
	BContes孔特() feud.Barony
	BLantosque朗托斯克() feud.Barony
	BMentone门托内() feud.Barony
	BMonaco摩纳哥() feud.Barony
	BNizza尼扎() feud.Barony
}

type 尼斯NiceCounty struct {
	feud.BaseCounty
	昂蒂布Antibes       feud.Barony
	坎波格罗索Campogrosso feud.Barony
	孔特Contes         feud.Barony
	朗托斯克Lantosque    feud.Barony
	门托内Mentone       feud.Barony
	摩纳哥Monaco        feud.Barony
	尼扎Nizza          feud.Barony
}

func (c *尼斯NiceCounty) BAntibes昂蒂布() feud.Barony {
	return c.昂蒂布Antibes
}

func (c *尼斯NiceCounty) BCampogrosso坎波格罗索() feud.Barony {
	return c.坎波格罗索Campogrosso
}

func (c *尼斯NiceCounty) BContes孔特() feud.Barony {
	return c.孔特Contes
}

func (c *尼斯NiceCounty) BLantosque朗托斯克() feud.Barony {
	return c.朗托斯克Lantosque
}

func (c *尼斯NiceCounty) BMentone门托内() feud.Barony {
	return c.门托内Mentone
}

func (c *尼斯NiceCounty) BMonaco摩纳哥() feud.Barony {
	return c.摩纳哥Monaco
}

func (c *尼斯NiceCounty) BNizza尼扎() feud.Barony {
	return c.尼扎Nizza
}

var CNice尼斯 NiceCounty = &尼斯NiceCounty{}

func init() {
	f := CNice尼斯.(*尼斯NiceCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "230",
		Title:     "nice",
		TitleName: "尼斯",
		TitleCode: "c_nice",
		Baronies:  map[string]feud.Barony{},
	}

	f.昂蒂布Antibes = BAntibes昂蒂布
	f.昂蒂布Antibes.SetParent(f)

	f.坎波格罗索Campogrosso = BCampogrosso坎波格罗索
	f.坎波格罗索Campogrosso.SetParent(f)

	f.孔特Contes = BContes孔特
	f.孔特Contes.SetParent(f)

	f.朗托斯克Lantosque = BLantosque朗托斯克
	f.朗托斯克Lantosque.SetParent(f)

	f.门托内Mentone = BMentone门托内
	f.门托内Mentone.SetParent(f)

	f.摩纳哥Monaco = BMonaco摩纳哥
	f.摩纳哥Monaco.SetParent(f)

	f.尼扎Nizza = BNizza尼扎
	f.尼扎Nizza.SetParent(f)

}
