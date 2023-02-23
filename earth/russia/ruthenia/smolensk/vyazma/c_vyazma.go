package vyazma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type VyazmaCounty interface {
	feud.County
	BDorogobyzh多罗戈布日() feud.Barony
	BIgnatkovo伊格纳特科沃() feud.Barony
	BMokhnatka莫赫纳特卡() feud.Barony
	BSafonovo萨福诺沃() feud.Barony
	BTesnoye捷斯诺耶() feud.Barony
	BVyazma维亚济马() feud.Barony
	BYelnya叶利尼亚() feud.Barony
}

type 维亚济马VyazmaCounty struct {
	feud.BaseCounty
	多罗戈布日Dorogobyzh feud.Barony
	伊格纳特科沃Ignatkovo feud.Barony
	莫赫纳特卡Mokhnatka  feud.Barony
	萨福诺沃Safonovo    feud.Barony
	捷斯诺耶Tesnoye     feud.Barony
	维亚济马Vyazma      feud.Barony
	叶利尼亚Yelnya      feud.Barony
}

func (c *维亚济马VyazmaCounty) BDorogobyzh多罗戈布日() feud.Barony {
	return c.多罗戈布日Dorogobyzh
}

func (c *维亚济马VyazmaCounty) BIgnatkovo伊格纳特科沃() feud.Barony {
	return c.伊格纳特科沃Ignatkovo
}

func (c *维亚济马VyazmaCounty) BMokhnatka莫赫纳特卡() feud.Barony {
	return c.莫赫纳特卡Mokhnatka
}

func (c *维亚济马VyazmaCounty) BSafonovo萨福诺沃() feud.Barony {
	return c.萨福诺沃Safonovo
}

func (c *维亚济马VyazmaCounty) BTesnoye捷斯诺耶() feud.Barony {
	return c.捷斯诺耶Tesnoye
}

func (c *维亚济马VyazmaCounty) BVyazma维亚济马() feud.Barony {
	return c.维亚济马Vyazma
}

func (c *维亚济马VyazmaCounty) BYelnya叶利尼亚() feud.Barony {
	return c.叶利尼亚Yelnya
}

var CVyazma维亚济马 VyazmaCounty = &维亚济马VyazmaCounty{}

func init() {
	f := CVyazma维亚济马.(*维亚济马VyazmaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "569",
		Title:     "vyazma",
		TitleName: "维亚济马",
		TitleCode: "c_vyazma",
		Baronies:  map[string]feud.Barony{},
	}

	f.多罗戈布日Dorogobyzh = BDorogobyzh多罗戈布日
	f.多罗戈布日Dorogobyzh.SetParent(f)

	f.伊格纳特科沃Ignatkovo = BIgnatkovo伊格纳特科沃
	f.伊格纳特科沃Ignatkovo.SetParent(f)

	f.莫赫纳特卡Mokhnatka = BMokhnatka莫赫纳特卡
	f.莫赫纳特卡Mokhnatka.SetParent(f)

	f.萨福诺沃Safonovo = BSafonovo萨福诺沃
	f.萨福诺沃Safonovo.SetParent(f)

	f.捷斯诺耶Tesnoye = BTesnoye捷斯诺耶
	f.捷斯诺耶Tesnoye.SetParent(f)

	f.维亚济马Vyazma = BVyazma维亚济马
	f.维亚济马Vyazma.SetParent(f)

	f.叶利尼亚Yelnya = BYelnya叶利尼亚
	f.叶利尼亚Yelnya.SetParent(f)

}
