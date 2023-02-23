package consenza

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ConsenzaCounty interface {
	feud.County
	BArgentano阿尔真塔诺() feud.Barony
	BCerenzia切伦齐亚() feud.Barony
	BCosenza科森扎() feud.Barony
	BRossano罗萨诺() feud.Barony
	BScalea斯卡莱阿() feud.Barony
	BStrongoli斯特龙戈利() feud.Barony
	BUmbriatico温布里亚蒂科() feud.Barony
}

type 科森扎ConsenzaCounty struct {
	feud.BaseCounty
	阿尔真塔诺Argentano   feud.Barony
	切伦齐亚Cerenzia     feud.Barony
	科森扎Cosenza       feud.Barony
	罗萨诺Rossano       feud.Barony
	斯卡莱阿Scalea       feud.Barony
	斯特龙戈利Strongoli   feud.Barony
	温布里亚蒂科Umbriatico feud.Barony
}

func (c *科森扎ConsenzaCounty) BArgentano阿尔真塔诺() feud.Barony {
	return c.阿尔真塔诺Argentano
}

func (c *科森扎ConsenzaCounty) BCerenzia切伦齐亚() feud.Barony {
	return c.切伦齐亚Cerenzia
}

func (c *科森扎ConsenzaCounty) BCosenza科森扎() feud.Barony {
	return c.科森扎Cosenza
}

func (c *科森扎ConsenzaCounty) BRossano罗萨诺() feud.Barony {
	return c.罗萨诺Rossano
}

func (c *科森扎ConsenzaCounty) BScalea斯卡莱阿() feud.Barony {
	return c.斯卡莱阿Scalea
}

func (c *科森扎ConsenzaCounty) BStrongoli斯特龙戈利() feud.Barony {
	return c.斯特龙戈利Strongoli
}

func (c *科森扎ConsenzaCounty) BUmbriatico温布里亚蒂科() feud.Barony {
	return c.温布里亚蒂科Umbriatico
}

var CConsenza科森扎 ConsenzaCounty = &科森扎ConsenzaCounty{}

func init() {
	f := CConsenza科森扎.(*科森扎ConsenzaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "337",
		Title:     "consenza",
		TitleName: "科森扎",
		TitleCode: "c_consenza",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿尔真塔诺Argentano = BArgentano阿尔真塔诺
	f.阿尔真塔诺Argentano.SetParent(f)

	f.切伦齐亚Cerenzia = BCerenzia切伦齐亚
	f.切伦齐亚Cerenzia.SetParent(f)

	f.科森扎Cosenza = BCosenza科森扎
	f.科森扎Cosenza.SetParent(f)

	f.罗萨诺Rossano = BRossano罗萨诺
	f.罗萨诺Rossano.SetParent(f)

	f.斯卡莱阿Scalea = BScalea斯卡莱阿
	f.斯卡莱阿Scalea.SetParent(f)

	f.斯特龙戈利Strongoli = BStrongoli斯特龙戈利
	f.斯特龙戈利Strongoli.SetParent(f)

	f.温布里亚蒂科Umbriatico = BUmbriatico温布里亚蒂科
	f.温布里亚蒂科Umbriatico.SetParent(f)

}
