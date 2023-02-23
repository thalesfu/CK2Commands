package ghana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type GhanaCounty interface {
	feud.County
	BBaghena巴盖纳() feud.Barony
	BDiara迪亚拉() feud.Barony
	BGoumbou贡布() feud.Barony
	BKughah库加() feud.Barony
	BNema内马() feud.Barony
}

type 瓦加杜GhanaCounty struct {
	feud.BaseCounty
	巴盖纳Baghena feud.Barony
	迪亚拉Diara   feud.Barony
	贡布Goumbou  feud.Barony
	库加Kughah   feud.Barony
	内马Nema     feud.Barony
}

func (c *瓦加杜GhanaCounty) BBaghena巴盖纳() feud.Barony {
	return c.巴盖纳Baghena
}

func (c *瓦加杜GhanaCounty) BDiara迪亚拉() feud.Barony {
	return c.迪亚拉Diara
}

func (c *瓦加杜GhanaCounty) BGoumbou贡布() feud.Barony {
	return c.贡布Goumbou
}

func (c *瓦加杜GhanaCounty) BKughah库加() feud.Barony {
	return c.库加Kughah
}

func (c *瓦加杜GhanaCounty) BNema内马() feud.Barony {
	return c.内马Nema
}

var CGhana瓦加杜 GhanaCounty = &瓦加杜GhanaCounty{}

func init() {
	f := CGhana瓦加杜.(*瓦加杜GhanaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "913",
		Title:     "ghana",
		TitleName: "瓦加杜",
		TitleCode: "c_ghana",
		Baronies:  map[string]feud.Barony{},
	}

	f.巴盖纳Baghena = BBaghena巴盖纳
	f.巴盖纳Baghena.SetParent(f)

	f.迪亚拉Diara = BDiara迪亚拉
	f.迪亚拉Diara.SetParent(f)

	f.贡布Goumbou = BGoumbou贡布
	f.贡布Goumbou.SetParent(f)

	f.库加Kughah = BKughah库加
	f.库加Kughah.SetParent(f)

	f.内马Nema = BNema内马
	f.内马Nema.SetParent(f)

}
