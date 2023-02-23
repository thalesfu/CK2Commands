package gabiyaha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type GabiyahaCounty interface {
	feud.County
	BButo布托() feud.Barony
	BDisuq迪苏格() feud.Barony
	BEdkhou埃德胡() feud.Barony
	BFuwa弗瓦() feud.Barony
	BHermopolis赫尔莫波利斯() feud.Barony
	BMutubis穆图比斯() feud.Barony
	BRosetta拉希德() feud.Barony
}

type 拉希德GabiyahaCounty struct {
	feud.BaseCounty
	布托Buto           feud.Barony
	迪苏格Disuq         feud.Barony
	埃德胡Edkhou        feud.Barony
	弗瓦Fuwa           feud.Barony
	赫尔莫波利斯Hermopolis feud.Barony
	穆图比斯Mutubis      feud.Barony
	拉希德Rosetta       feud.Barony
}

func (c *拉希德GabiyahaCounty) BButo布托() feud.Barony {
	return c.布托Buto
}

func (c *拉希德GabiyahaCounty) BDisuq迪苏格() feud.Barony {
	return c.迪苏格Disuq
}

func (c *拉希德GabiyahaCounty) BEdkhou埃德胡() feud.Barony {
	return c.埃德胡Edkhou
}

func (c *拉希德GabiyahaCounty) BFuwa弗瓦() feud.Barony {
	return c.弗瓦Fuwa
}

func (c *拉希德GabiyahaCounty) BHermopolis赫尔莫波利斯() feud.Barony {
	return c.赫尔莫波利斯Hermopolis
}

func (c *拉希德GabiyahaCounty) BMutubis穆图比斯() feud.Barony {
	return c.穆图比斯Mutubis
}

func (c *拉希德GabiyahaCounty) BRosetta拉希德() feud.Barony {
	return c.拉希德Rosetta
}

var CGabiyaha拉希德 GabiyahaCounty = &拉希德GabiyahaCounty{}

func init() {
	f := CGabiyaha拉希德.(*拉希德GabiyahaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "799",
		Title:     "gabiyaha",
		TitleName: "拉希德",
		TitleCode: "c_gabiyaha",
		Baronies:  map[string]feud.Barony{},
	}

	f.布托Buto = BButo布托
	f.布托Buto.SetParent(f)

	f.迪苏格Disuq = BDisuq迪苏格
	f.迪苏格Disuq.SetParent(f)

	f.埃德胡Edkhou = BEdkhou埃德胡
	f.埃德胡Edkhou.SetParent(f)

	f.弗瓦Fuwa = BFuwa弗瓦
	f.弗瓦Fuwa.SetParent(f)

	f.赫尔莫波利斯Hermopolis = BHermopolis赫尔莫波利斯
	f.赫尔莫波利斯Hermopolis.SetParent(f)

	f.穆图比斯Mutubis = BMutubis穆图比斯
	f.穆图比斯Mutubis.SetParent(f)

	f.拉希德Rosetta = BRosetta拉希德
	f.拉希德Rosetta.SetParent(f)

}
