package charsianon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type CharsianonCounty interface {
	feud.County
	BCharsianon哈耳西亚农() feud.Barony
	BDazimon德治文() feud.Barony
	BMazaca马扎卡() feud.Barony
	BSebastea塞巴斯蒂亚() feud.Barony
	BSebastopolis塞巴斯托波利斯() feud.Barony
	BZela泽拉() feud.Barony
}

type 哈耳西亚农CharsianonCounty struct {
	feud.BaseCounty
	哈耳西亚农Charsianon     feud.Barony
	德治文Dazimon          feud.Barony
	马扎卡Mazaca           feud.Barony
	塞巴斯蒂亚Sebastea       feud.Barony
	塞巴斯托波利斯Sebastopolis feud.Barony
	泽拉Zela              feud.Barony
}

func (c *哈耳西亚农CharsianonCounty) BCharsianon哈耳西亚农() feud.Barony {
	return c.哈耳西亚农Charsianon
}

func (c *哈耳西亚农CharsianonCounty) BDazimon德治文() feud.Barony {
	return c.德治文Dazimon
}

func (c *哈耳西亚农CharsianonCounty) BMazaca马扎卡() feud.Barony {
	return c.马扎卡Mazaca
}

func (c *哈耳西亚农CharsianonCounty) BSebastea塞巴斯蒂亚() feud.Barony {
	return c.塞巴斯蒂亚Sebastea
}

func (c *哈耳西亚农CharsianonCounty) BSebastopolis塞巴斯托波利斯() feud.Barony {
	return c.塞巴斯托波利斯Sebastopolis
}

func (c *哈耳西亚农CharsianonCounty) BZela泽拉() feud.Barony {
	return c.泽拉Zela
}

var CCharsianon哈耳西亚农 CharsianonCounty = &哈耳西亚农CharsianonCounty{}

func init() {
	f := CCharsianon哈耳西亚农.(*哈耳西亚农CharsianonCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1934",
		Title:     "charsianon",
		TitleName: "哈耳西亚农",
		TitleCode: "c_charsianon",
		Baronies:  map[string]feud.Barony{},
	}

	f.哈耳西亚农Charsianon = BCharsianon哈耳西亚农
	f.哈耳西亚农Charsianon.SetParent(f)

	f.德治文Dazimon = BDazimon德治文
	f.德治文Dazimon.SetParent(f)

	f.马扎卡Mazaca = BMazaca马扎卡
	f.马扎卡Mazaca.SetParent(f)

	f.塞巴斯蒂亚Sebastea = BSebastea塞巴斯蒂亚
	f.塞巴斯蒂亚Sebastea.SetParent(f)

	f.塞巴斯托波利斯Sebastopolis = BSebastopolis塞巴斯托波利斯
	f.塞巴斯托波利斯Sebastopolis.SetParent(f)

	f.泽拉Zela = BZela泽拉
	f.泽拉Zela.SetParent(f)

}
