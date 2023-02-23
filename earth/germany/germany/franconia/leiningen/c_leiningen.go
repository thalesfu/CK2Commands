package leiningen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type LeiningenCounty interface {
	feud.County
	BAschaffenburg阿沙芬堡() feud.Barony
	BBattenberg巴滕贝格() feud.Barony
	BDurkheim迪尔凯姆() feud.Barony
	BHardenburg哈登堡() feud.Barony
	BLeiningen莱宁根() feud.Barony
	BLorsch洛尔施() feud.Barony
	BPfeffingen普费芬根() feud.Barony
	BUngstein翁施泰因() feud.Barony
}

type 阿沙芬堡LeiningenCounty struct {
	feud.BaseCounty
	阿沙芬堡Aschaffenburg feud.Barony
	巴滕贝格Battenberg    feud.Barony
	迪尔凯姆Durkheim      feud.Barony
	哈登堡Hardenburg     feud.Barony
	莱宁根Leiningen      feud.Barony
	洛尔施Lorsch         feud.Barony
	普费芬根Pfeffingen    feud.Barony
	翁施泰因Ungstein      feud.Barony
}

func (c *阿沙芬堡LeiningenCounty) BAschaffenburg阿沙芬堡() feud.Barony {
	return c.阿沙芬堡Aschaffenburg
}

func (c *阿沙芬堡LeiningenCounty) BBattenberg巴滕贝格() feud.Barony {
	return c.巴滕贝格Battenberg
}

func (c *阿沙芬堡LeiningenCounty) BDurkheim迪尔凯姆() feud.Barony {
	return c.迪尔凯姆Durkheim
}

func (c *阿沙芬堡LeiningenCounty) BHardenburg哈登堡() feud.Barony {
	return c.哈登堡Hardenburg
}

func (c *阿沙芬堡LeiningenCounty) BLeiningen莱宁根() feud.Barony {
	return c.莱宁根Leiningen
}

func (c *阿沙芬堡LeiningenCounty) BLorsch洛尔施() feud.Barony {
	return c.洛尔施Lorsch
}

func (c *阿沙芬堡LeiningenCounty) BPfeffingen普费芬根() feud.Barony {
	return c.普费芬根Pfeffingen
}

func (c *阿沙芬堡LeiningenCounty) BUngstein翁施泰因() feud.Barony {
	return c.翁施泰因Ungstein
}

var CLeiningen阿沙芬堡 LeiningenCounty = &阿沙芬堡LeiningenCounty{}

func init() {
	f := CLeiningen阿沙芬堡.(*阿沙芬堡LeiningenCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "122",
		Title:     "leiningen",
		TitleName: "阿沙芬堡",
		TitleCode: "c_leiningen",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿沙芬堡Aschaffenburg = BAschaffenburg阿沙芬堡
	f.阿沙芬堡Aschaffenburg.SetParent(f)

	f.巴滕贝格Battenberg = BBattenberg巴滕贝格
	f.巴滕贝格Battenberg.SetParent(f)

	f.迪尔凯姆Durkheim = BDurkheim迪尔凯姆
	f.迪尔凯姆Durkheim.SetParent(f)

	f.哈登堡Hardenburg = BHardenburg哈登堡
	f.哈登堡Hardenburg.SetParent(f)

	f.莱宁根Leiningen = BLeiningen莱宁根
	f.莱宁根Leiningen.SetParent(f)

	f.洛尔施Lorsch = BLorsch洛尔施
	f.洛尔施Lorsch.SetParent(f)

	f.普费芬根Pfeffingen = BPfeffingen普费芬根
	f.普费芬根Pfeffingen.SetParent(f)

	f.翁施泰因Ungstein = BUngstein翁施泰因
	f.翁施泰因Ungstein.SetParent(f)

}
