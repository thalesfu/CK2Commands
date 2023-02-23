package taranto

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TarantoCounty interface {
	feud.County
	BCassano卡萨诺() feud.Barony
	BCastellaneta卡斯泰拉内塔() feud.Barony
	BGravina格拉维纳() feud.Barony
	BMontepeloso蒙特佩洛索() feud.Barony
	BMottola莫托拉() feud.Barony
	BTaranto塔兰托() feud.Barony
	BTricanico特里卡尼科() feud.Barony
	BTursi图尔西() feud.Barony
}

type 塔兰托TarantoCounty struct {
	feud.BaseCounty
	卡萨诺Cassano         feud.Barony
	卡斯泰拉内塔Castellaneta feud.Barony
	格拉维纳Gravina        feud.Barony
	蒙特佩洛索Montepeloso   feud.Barony
	莫托拉Mottola         feud.Barony
	塔兰托Taranto         feud.Barony
	特里卡尼科Tricanico     feud.Barony
	图尔西Tursi           feud.Barony
}

func (c *塔兰托TarantoCounty) BCassano卡萨诺() feud.Barony {
	return c.卡萨诺Cassano
}

func (c *塔兰托TarantoCounty) BCastellaneta卡斯泰拉内塔() feud.Barony {
	return c.卡斯泰拉内塔Castellaneta
}

func (c *塔兰托TarantoCounty) BGravina格拉维纳() feud.Barony {
	return c.格拉维纳Gravina
}

func (c *塔兰托TarantoCounty) BMontepeloso蒙特佩洛索() feud.Barony {
	return c.蒙特佩洛索Montepeloso
}

func (c *塔兰托TarantoCounty) BMottola莫托拉() feud.Barony {
	return c.莫托拉Mottola
}

func (c *塔兰托TarantoCounty) BTaranto塔兰托() feud.Barony {
	return c.塔兰托Taranto
}

func (c *塔兰托TarantoCounty) BTricanico特里卡尼科() feud.Barony {
	return c.特里卡尼科Tricanico
}

func (c *塔兰托TarantoCounty) BTursi图尔西() feud.Barony {
	return c.图尔西Tursi
}

var CTaranto塔兰托 TarantoCounty = &塔兰托TarantoCounty{}

func init() {
	f := CTaranto塔兰托.(*塔兰托TarantoCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "344",
		Title:     "taranto",
		TitleName: "塔兰托",
		TitleCode: "c_taranto",
		Baronies:  map[string]feud.Barony{},
	}

	f.卡萨诺Cassano = BCassano卡萨诺
	f.卡萨诺Cassano.SetParent(f)

	f.卡斯泰拉内塔Castellaneta = BCastellaneta卡斯泰拉内塔
	f.卡斯泰拉内塔Castellaneta.SetParent(f)

	f.格拉维纳Gravina = BGravina格拉维纳
	f.格拉维纳Gravina.SetParent(f)

	f.蒙特佩洛索Montepeloso = BMontepeloso蒙特佩洛索
	f.蒙特佩洛索Montepeloso.SetParent(f)

	f.莫托拉Mottola = BMottola莫托拉
	f.莫托拉Mottola.SetParent(f)

	f.塔兰托Taranto = BTaranto塔兰托
	f.塔兰托Taranto.SetParent(f)

	f.特里卡尼科Tricanico = BTricanico特里卡尼科
	f.特里卡尼科Tricanico.SetParent(f)

	f.图尔西Tursi = BTursi图尔西
	f.图尔西Tursi.SetParent(f)

}
