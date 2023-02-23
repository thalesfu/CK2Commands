package valais

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ValaisCounty interface {
	feud.County
	BAigle艾格勒() feud.Barony
	BBrig布里格() feud.Barony
	BGreyerz格雷耶茨() feud.Barony
	BMartigny马蒂尼() feud.Barony
	BMonthey蒙泰() feud.Barony
	BSiders西德尔斯() feud.Barony
	BSitten锡滕() feud.Barony
}

type 瓦莱ValaisCounty struct {
	feud.BaseCounty
	艾格勒Aigle    feud.Barony
	布里格Brig     feud.Barony
	格雷耶茨Greyerz feud.Barony
	马蒂尼Martigny feud.Barony
	蒙泰Monthey   feud.Barony
	西德尔斯Siders  feud.Barony
	锡滕Sitten    feud.Barony
}

func (c *瓦莱ValaisCounty) BAigle艾格勒() feud.Barony {
	return c.艾格勒Aigle
}

func (c *瓦莱ValaisCounty) BBrig布里格() feud.Barony {
	return c.布里格Brig
}

func (c *瓦莱ValaisCounty) BGreyerz格雷耶茨() feud.Barony {
	return c.格雷耶茨Greyerz
}

func (c *瓦莱ValaisCounty) BMartigny马蒂尼() feud.Barony {
	return c.马蒂尼Martigny
}

func (c *瓦莱ValaisCounty) BMonthey蒙泰() feud.Barony {
	return c.蒙泰Monthey
}

func (c *瓦莱ValaisCounty) BSiders西德尔斯() feud.Barony {
	return c.西德尔斯Siders
}

func (c *瓦莱ValaisCounty) BSitten锡滕() feud.Barony {
	return c.锡滕Sitten
}

var CValais瓦莱 ValaisCounty = &瓦莱ValaisCounty{}

func init() {
	f := CValais瓦莱.(*瓦莱ValaisCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "238",
		Title:     "valais",
		TitleName: "瓦莱",
		TitleCode: "c_valais",
		Baronies:  map[string]feud.Barony{},
	}

	f.艾格勒Aigle = BAigle艾格勒
	f.艾格勒Aigle.SetParent(f)

	f.布里格Brig = BBrig布里格
	f.布里格Brig.SetParent(f)

	f.格雷耶茨Greyerz = BGreyerz格雷耶茨
	f.格雷耶茨Greyerz.SetParent(f)

	f.马蒂尼Martigny = BMartigny马蒂尼
	f.马蒂尼Martigny.SetParent(f)

	f.蒙泰Monthey = BMonthey蒙泰
	f.蒙泰Monthey.SetParent(f)

	f.西德尔斯Siders = BSiders西德尔斯
	f.西德尔斯Siders.SetParent(f)

	f.锡滕Sitten = BSitten锡滕
	f.锡滕Sitten.SetParent(f)

}
