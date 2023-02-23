package chartres

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ChartresCounty interface {
	feud.County
	BBretigny布雷蒂尼() feud.Barony
	BChartres沙特尔() feud.Barony
	BDreux德勒() feud.Barony
	BEpernon埃佩尔农() feud.Barony
	BGallardon加拉尔东() feud.Barony
	BNogent诺让() feud.Barony
}

type 沙特尔ChartresCounty struct {
	feud.BaseCounty
	布雷蒂尼Bretigny  feud.Barony
	沙特尔Chartres   feud.Barony
	德勒Dreux       feud.Barony
	埃佩尔农Epernon   feud.Barony
	加拉尔东Gallardon feud.Barony
	诺让Nogent      feud.Barony
}

func (c *沙特尔ChartresCounty) BBretigny布雷蒂尼() feud.Barony {
	return c.布雷蒂尼Bretigny
}

func (c *沙特尔ChartresCounty) BChartres沙特尔() feud.Barony {
	return c.沙特尔Chartres
}

func (c *沙特尔ChartresCounty) BDreux德勒() feud.Barony {
	return c.德勒Dreux
}

func (c *沙特尔ChartresCounty) BEpernon埃佩尔农() feud.Barony {
	return c.埃佩尔农Epernon
}

func (c *沙特尔ChartresCounty) BGallardon加拉尔东() feud.Barony {
	return c.加拉尔东Gallardon
}

func (c *沙特尔ChartresCounty) BNogent诺让() feud.Barony {
	return c.诺让Nogent
}

var CChartres沙特尔 ChartresCounty = &沙特尔ChartresCounty{}

func init() {
	f := CChartres沙特尔.(*沙特尔ChartresCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "111",
		Title:     "chartres",
		TitleName: "沙特尔",
		TitleCode: "c_chartres",
		Baronies:  map[string]feud.Barony{},
	}

	f.布雷蒂尼Bretigny = BBretigny布雷蒂尼
	f.布雷蒂尼Bretigny.SetParent(f)

	f.沙特尔Chartres = BChartres沙特尔
	f.沙特尔Chartres.SetParent(f)

	f.德勒Dreux = BDreux德勒
	f.德勒Dreux.SetParent(f)

	f.埃佩尔农Epernon = BEpernon埃佩尔农
	f.埃佩尔农Epernon.SetParent(f)

	f.加拉尔东Gallardon = BGallardon加拉尔东
	f.加拉尔东Gallardon.SetParent(f)

	f.诺让Nogent = BNogent诺让
	f.诺让Nogent.SetParent(f)

}
