package lombardia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type LombardiaCounty interface {
	feud.County
	BChiavenna基亚文纳() feud.Barony
	BLegnano莱尼亚诺() feud.Barony
	BLodi洛迪() feud.Barony
	BMaggiore马焦雷() feud.Barony
	BMendrisio门德里西奥() feud.Barony
	BMilano米兰() feud.Barony
	BMonza蒙扎() feud.Barony
	BVigevano维杰瓦诺() feud.Barony
}

type 米兰LombardiaCounty struct {
	feud.BaseCounty
	基亚文纳Chiavenna  feud.Barony
	莱尼亚诺Legnano    feud.Barony
	洛迪Lodi         feud.Barony
	马焦雷Maggiore    feud.Barony
	门德里西奥Mendrisio feud.Barony
	米兰Milano       feud.Barony
	蒙扎Monza        feud.Barony
	维杰瓦诺Vigevano   feud.Barony
}

func (c *米兰LombardiaCounty) BChiavenna基亚文纳() feud.Barony {
	return c.基亚文纳Chiavenna
}

func (c *米兰LombardiaCounty) BLegnano莱尼亚诺() feud.Barony {
	return c.莱尼亚诺Legnano
}

func (c *米兰LombardiaCounty) BLodi洛迪() feud.Barony {
	return c.洛迪Lodi
}

func (c *米兰LombardiaCounty) BMaggiore马焦雷() feud.Barony {
	return c.马焦雷Maggiore
}

func (c *米兰LombardiaCounty) BMendrisio门德里西奥() feud.Barony {
	return c.门德里西奥Mendrisio
}

func (c *米兰LombardiaCounty) BMilano米兰() feud.Barony {
	return c.米兰Milano
}

func (c *米兰LombardiaCounty) BMonza蒙扎() feud.Barony {
	return c.蒙扎Monza
}

func (c *米兰LombardiaCounty) BVigevano维杰瓦诺() feud.Barony {
	return c.维杰瓦诺Vigevano
}

var CLombardia米兰 LombardiaCounty = &米兰LombardiaCounty{}

func init() {
	f := CLombardia米兰.(*米兰LombardiaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "235",
		Title:     "lombardia",
		TitleName: "米兰",
		TitleCode: "c_lombardia",
		Baronies:  map[string]feud.Barony{},
	}

	f.基亚文纳Chiavenna = BChiavenna基亚文纳
	f.基亚文纳Chiavenna.SetParent(f)

	f.莱尼亚诺Legnano = BLegnano莱尼亚诺
	f.莱尼亚诺Legnano.SetParent(f)

	f.洛迪Lodi = BLodi洛迪
	f.洛迪Lodi.SetParent(f)

	f.马焦雷Maggiore = BMaggiore马焦雷
	f.马焦雷Maggiore.SetParent(f)

	f.门德里西奥Mendrisio = BMendrisio门德里西奥
	f.门德里西奥Mendrisio.SetParent(f)

	f.米兰Milano = BMilano米兰
	f.米兰Milano.SetParent(f)

	f.蒙扎Monza = BMonza蒙扎
	f.蒙扎Monza.SetParent(f)

	f.维杰瓦诺Vigevano = BVigevano维杰瓦诺
	f.维杰瓦诺Vigevano.SetParent(f)

}
