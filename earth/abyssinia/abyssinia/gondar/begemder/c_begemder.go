package begemder

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BegemderCounty interface {
	feud.County
	BAwariya阿瓦里亚() feud.Barony
	BChekerefta杰克雷夫塔() feud.Barony
	BDanya但雅() feud.Barony
	BFilakit菲拉基特() feud.Barony
	BGereger格雷格尔() feud.Barony
	BMaryamu玛丽亚穆() feud.Barony
}

type 贝格姆迪尔BegemderCounty struct {
	feud.BaseCounty
	阿瓦里亚Awariya     feud.Barony
	杰克雷夫塔Chekerefta feud.Barony
	但雅Danya         feud.Barony
	菲拉基特Filakit     feud.Barony
	格雷格尔Gereger     feud.Barony
	玛丽亚穆Maryamu     feud.Barony
}

func (c *贝格姆迪尔BegemderCounty) BAwariya阿瓦里亚() feud.Barony {
	return c.阿瓦里亚Awariya
}

func (c *贝格姆迪尔BegemderCounty) BChekerefta杰克雷夫塔() feud.Barony {
	return c.杰克雷夫塔Chekerefta
}

func (c *贝格姆迪尔BegemderCounty) BDanya但雅() feud.Barony {
	return c.但雅Danya
}

func (c *贝格姆迪尔BegemderCounty) BFilakit菲拉基特() feud.Barony {
	return c.菲拉基特Filakit
}

func (c *贝格姆迪尔BegemderCounty) BGereger格雷格尔() feud.Barony {
	return c.格雷格尔Gereger
}

func (c *贝格姆迪尔BegemderCounty) BMaryamu玛丽亚穆() feud.Barony {
	return c.玛丽亚穆Maryamu
}

var CBegemder贝格姆迪尔 BegemderCounty = &贝格姆迪尔BegemderCounty{}

func init() {
	f := CBegemder贝格姆迪尔.(*贝格姆迪尔BegemderCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1429",
		Title:     "begemder",
		TitleName: "贝格姆迪尔",
		TitleCode: "c_begemder",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿瓦里亚Awariya = BAwariya阿瓦里亚
	f.阿瓦里亚Awariya.SetParent(f)

	f.杰克雷夫塔Chekerefta = BChekerefta杰克雷夫塔
	f.杰克雷夫塔Chekerefta.SetParent(f)

	f.但雅Danya = BDanya但雅
	f.但雅Danya.SetParent(f)

	f.菲拉基特Filakit = BFilakit菲拉基特
	f.菲拉基特Filakit.SetParent(f)

	f.格雷格尔Gereger = BGereger格雷格尔
	f.格雷格尔Gereger.SetParent(f)

	f.玛丽亚穆Maryamu = BMaryamu玛丽亚穆
	f.玛丽亚穆Maryamu.SetParent(f)

}
