package tabaristan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TabaristanCounty interface {
	feud.County
	BAmul阿莫勒() feud.Barony
	BChamnoo察姆诺() feud.Barony
	BFarim法里曼() feud.Barony
	BFiruzkuh菲鲁兹库赫() feud.Barony
	BLavij拉瓦基() feud.Barony
	BMamatir玛马提尔() feud.Barony
	BRarem拉雷姆() feud.Barony
	BSari萨里() feud.Barony
}

type 陀拔斯单TabaristanCounty struct {
	feud.BaseCounty
	阿莫勒Amul       feud.Barony
	察姆诺Chamnoo    feud.Barony
	法里曼Farim      feud.Barony
	菲鲁兹库赫Firuzkuh feud.Barony
	拉瓦基Lavij      feud.Barony
	玛马提尔Mamatir   feud.Barony
	拉雷姆Rarem      feud.Barony
	萨里Sari        feud.Barony
}

func (c *陀拔斯单TabaristanCounty) BAmul阿莫勒() feud.Barony {
	return c.阿莫勒Amul
}

func (c *陀拔斯单TabaristanCounty) BChamnoo察姆诺() feud.Barony {
	return c.察姆诺Chamnoo
}

func (c *陀拔斯单TabaristanCounty) BFarim法里曼() feud.Barony {
	return c.法里曼Farim
}

func (c *陀拔斯单TabaristanCounty) BFiruzkuh菲鲁兹库赫() feud.Barony {
	return c.菲鲁兹库赫Firuzkuh
}

func (c *陀拔斯单TabaristanCounty) BLavij拉瓦基() feud.Barony {
	return c.拉瓦基Lavij
}

func (c *陀拔斯单TabaristanCounty) BMamatir玛马提尔() feud.Barony {
	return c.玛马提尔Mamatir
}

func (c *陀拔斯单TabaristanCounty) BRarem拉雷姆() feud.Barony {
	return c.拉雷姆Rarem
}

func (c *陀拔斯单TabaristanCounty) BSari萨里() feud.Barony {
	return c.萨里Sari
}

var CTabaristan陀拔斯单 TabaristanCounty = &陀拔斯单TabaristanCounty{}

func init() {
	f := CTabaristan陀拔斯单.(*陀拔斯单TabaristanCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "661",
		Title:     "tabaristan",
		TitleName: "陀拔斯单",
		TitleCode: "c_tabaristan",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿莫勒Amul = BAmul阿莫勒
	f.阿莫勒Amul.SetParent(f)

	f.察姆诺Chamnoo = BChamnoo察姆诺
	f.察姆诺Chamnoo.SetParent(f)

	f.法里曼Farim = BFarim法里曼
	f.法里曼Farim.SetParent(f)

	f.菲鲁兹库赫Firuzkuh = BFiruzkuh菲鲁兹库赫
	f.菲鲁兹库赫Firuzkuh.SetParent(f)

	f.拉瓦基Lavij = BLavij拉瓦基
	f.拉瓦基Lavij.SetParent(f)

	f.玛马提尔Mamatir = BMamatir玛马提尔
	f.玛马提尔Mamatir.SetParent(f)

	f.拉雷姆Rarem = BRarem拉雷姆
	f.拉雷姆Rarem.SetParent(f)

	f.萨里Sari = BSari萨里
	f.萨里Sari.SetParent(f)

}
