package palencia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type PalenciaCounty interface {
	feud.County
	BPalencia帕伦西亚() feud.Barony
	BParedes帕雷德斯() feud.Barony
	BVaccaei瓦凯伊() feud.Barony
}

type 帕伦西亚PalenciaCounty struct {
	feud.BaseCounty
	帕伦西亚Palencia feud.Barony
	帕雷德斯Paredes  feud.Barony
	瓦凯伊Vaccaei   feud.Barony
}

func (c *帕伦西亚PalenciaCounty) BPalencia帕伦西亚() feud.Barony {
	return c.帕伦西亚Palencia
}

func (c *帕伦西亚PalenciaCounty) BParedes帕雷德斯() feud.Barony {
	return c.帕雷德斯Paredes
}

func (c *帕伦西亚PalenciaCounty) BVaccaei瓦凯伊() feud.Barony {
	return c.瓦凯伊Vaccaei
}

var CPalencia帕伦西亚 PalenciaCounty = &帕伦西亚PalenciaCounty{}

func init() {
	f := CPalencia帕伦西亚.(*帕伦西亚PalenciaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1975",
		Title:     "palencia",
		TitleName: "帕伦西亚",
		TitleCode: "c_palencia",
		Baronies:  map[string]feud.Barony{},
	}

	f.帕伦西亚Palencia = BPalencia帕伦西亚
	f.帕伦西亚Palencia.SetParent(f)

	f.帕雷德斯Paredes = BParedes帕雷德斯
	f.帕雷德斯Paredes.SetParent(f)

	f.瓦凯伊Vaccaei = BVaccaei瓦凯伊
	f.瓦凯伊Vaccaei.SetParent(f)

}
