package ancona

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AnconaCounty interface {
	feud.County
	BAncona安科纳() feud.Barony
	BCamerino卡梅里诺() feud.Barony
	BFermo费尔莫() feud.Barony
	BMatelica马泰利卡() feud.Barony
	BOsimo奥西莫() feud.Barony
	BRecanati雷卡纳蒂() feud.Barony
	BSenigallia塞尼加利亚() feud.Barony
}

type 安科纳AnconaCounty struct {
	feud.BaseCounty
	安科纳Ancona       feud.Barony
	卡梅里诺Camerino    feud.Barony
	费尔莫Fermo        feud.Barony
	马泰利卡Matelica    feud.Barony
	奥西莫Osimo        feud.Barony
	雷卡纳蒂Recanati    feud.Barony
	塞尼加利亚Senigallia feud.Barony
}

func (c *安科纳AnconaCounty) BAncona安科纳() feud.Barony {
	return c.安科纳Ancona
}

func (c *安科纳AnconaCounty) BCamerino卡梅里诺() feud.Barony {
	return c.卡梅里诺Camerino
}

func (c *安科纳AnconaCounty) BFermo费尔莫() feud.Barony {
	return c.费尔莫Fermo
}

func (c *安科纳AnconaCounty) BMatelica马泰利卡() feud.Barony {
	return c.马泰利卡Matelica
}

func (c *安科纳AnconaCounty) BOsimo奥西莫() feud.Barony {
	return c.奥西莫Osimo
}

func (c *安科纳AnconaCounty) BRecanati雷卡纳蒂() feud.Barony {
	return c.雷卡纳蒂Recanati
}

func (c *安科纳AnconaCounty) BSenigallia塞尼加利亚() feud.Barony {
	return c.塞尼加利亚Senigallia
}

var CAncona安科纳 AnconaCounty = &安科纳AnconaCounty{}

func init() {
	f := CAncona安科纳.(*安科纳AnconaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "350",
		Title:     "ancona",
		TitleName: "安科纳",
		TitleCode: "c_ancona",
		Baronies:  map[string]feud.Barony{},
	}

	f.安科纳Ancona = BAncona安科纳
	f.安科纳Ancona.SetParent(f)

	f.卡梅里诺Camerino = BCamerino卡梅里诺
	f.卡梅里诺Camerino.SetParent(f)

	f.费尔莫Fermo = BFermo费尔莫
	f.费尔莫Fermo.SetParent(f)

	f.马泰利卡Matelica = BMatelica马泰利卡
	f.马泰利卡Matelica.SetParent(f)

	f.奥西莫Osimo = BOsimo奥西莫
	f.奥西莫Osimo.SetParent(f)

	f.雷卡纳蒂Recanati = BRecanati雷卡纳蒂
	f.雷卡纳蒂Recanati.SetParent(f)

	f.塞尼加利亚Senigallia = BSenigallia塞尼加利亚
	f.塞尼加利亚Senigallia.SetParent(f)

}
