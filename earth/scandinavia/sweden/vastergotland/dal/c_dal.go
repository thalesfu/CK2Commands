package dal

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type DalCounty interface {
	feud.County
	BAmordh阿玛德() feud.Barony
	BBolstad博尔斯塔德() feud.Barony
	BDalaborg达拉波雷() feud.Barony
	BMellerud梅勒鲁() feud.Barony
	BNordal诺德尔() feud.Barony
	BSundal松达尔() feud.Barony
	BTossbo塔斯波() feud.Barony
	BVedbo韦德布() feud.Barony
}

type 达尔DalCounty struct {
	feud.BaseCounty
	阿玛德Amordh    feud.Barony
	博尔斯塔德Bolstad feud.Barony
	达拉波雷Dalaborg feud.Barony
	梅勒鲁Mellerud  feud.Barony
	诺德尔Nordal    feud.Barony
	松达尔Sundal    feud.Barony
	塔斯波Tossbo    feud.Barony
	韦德布Vedbo     feud.Barony
}

func (c *达尔DalCounty) BAmordh阿玛德() feud.Barony {
	return c.阿玛德Amordh
}

func (c *达尔DalCounty) BBolstad博尔斯塔德() feud.Barony {
	return c.博尔斯塔德Bolstad
}

func (c *达尔DalCounty) BDalaborg达拉波雷() feud.Barony {
	return c.达拉波雷Dalaborg
}

func (c *达尔DalCounty) BMellerud梅勒鲁() feud.Barony {
	return c.梅勒鲁Mellerud
}

func (c *达尔DalCounty) BNordal诺德尔() feud.Barony {
	return c.诺德尔Nordal
}

func (c *达尔DalCounty) BSundal松达尔() feud.Barony {
	return c.松达尔Sundal
}

func (c *达尔DalCounty) BTossbo塔斯波() feud.Barony {
	return c.塔斯波Tossbo
}

func (c *达尔DalCounty) BVedbo韦德布() feud.Barony {
	return c.韦德布Vedbo
}

var CDal达尔 DalCounty = &达尔DalCounty{}

func init() {
	f := CDal达尔.(*达尔DalCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "295",
		Title:     "dal",
		TitleName: "达尔",
		TitleCode: "c_dal",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿玛德Amordh = BAmordh阿玛德
	f.阿玛德Amordh.SetParent(f)

	f.博尔斯塔德Bolstad = BBolstad博尔斯塔德
	f.博尔斯塔德Bolstad.SetParent(f)

	f.达拉波雷Dalaborg = BDalaborg达拉波雷
	f.达拉波雷Dalaborg.SetParent(f)

	f.梅勒鲁Mellerud = BMellerud梅勒鲁
	f.梅勒鲁Mellerud.SetParent(f)

	f.诺德尔Nordal = BNordal诺德尔
	f.诺德尔Nordal.SetParent(f)

	f.松达尔Sundal = BSundal松达尔
	f.松达尔Sundal.SetParent(f)

	f.塔斯波Tossbo = BTossbo塔斯波
	f.塔斯波Tossbo.SetParent(f)

	f.韦德布Vedbo = BVedbo韦德布
	f.韦德布Vedbo.SetParent(f)

}
