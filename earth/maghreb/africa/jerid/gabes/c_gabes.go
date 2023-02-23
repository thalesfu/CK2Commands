package gabes

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type GabesCounty interface {
	feud.County
	BAlqalah卡拉罕() feud.Barony
	BGabes加贝斯() feud.Barony
	BHaddej哈德贾() feud.Barony
	BKebili吉比利() feud.Barony
	BMedenine梅德宁() feud.Barony
	BSteftimi塞夫提米() feud.Barony
}

type 加贝斯GabesCounty struct {
	feud.BaseCounty
	卡拉罕Alqalah   feud.Barony
	加贝斯Gabes     feud.Barony
	哈德贾Haddej    feud.Barony
	吉比利Kebili    feud.Barony
	梅德宁Medenine  feud.Barony
	塞夫提米Steftimi feud.Barony
}

func (c *加贝斯GabesCounty) BAlqalah卡拉罕() feud.Barony {
	return c.卡拉罕Alqalah
}

func (c *加贝斯GabesCounty) BGabes加贝斯() feud.Barony {
	return c.加贝斯Gabes
}

func (c *加贝斯GabesCounty) BHaddej哈德贾() feud.Barony {
	return c.哈德贾Haddej
}

func (c *加贝斯GabesCounty) BKebili吉比利() feud.Barony {
	return c.吉比利Kebili
}

func (c *加贝斯GabesCounty) BMedenine梅德宁() feud.Barony {
	return c.梅德宁Medenine
}

func (c *加贝斯GabesCounty) BSteftimi塞夫提米() feud.Barony {
	return c.塞夫提米Steftimi
}

var CGabes加贝斯 GabesCounty = &加贝斯GabesCounty{}

func init() {
	f := CGabes加贝斯.(*加贝斯GabesCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "814",
		Title:     "gabes",
		TitleName: "加贝斯",
		TitleCode: "c_gabes",
		Baronies:  map[string]feud.Barony{},
	}

	f.卡拉罕Alqalah = BAlqalah卡拉罕
	f.卡拉罕Alqalah.SetParent(f)

	f.加贝斯Gabes = BGabes加贝斯
	f.加贝斯Gabes.SetParent(f)

	f.哈德贾Haddej = BHaddej哈德贾
	f.哈德贾Haddej.SetParent(f)

	f.吉比利Kebili = BKebili吉比利
	f.吉比利Kebili.SetParent(f)

	f.梅德宁Medenine = BMedenine梅德宁
	f.梅德宁Medenine.SetParent(f)

	f.塞夫提米Steftimi = BSteftimi塞夫提米
	f.塞夫提米Steftimi.SetParent(f)

}
