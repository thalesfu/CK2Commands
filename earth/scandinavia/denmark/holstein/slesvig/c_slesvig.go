package slesvig

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SlesvigCounty interface {
    feud.County
    BAabenraa奥本罗() 	feud.Barony
    BFlensborg弗伦斯堡() 	feud.Barony
    BHaderslev哈泽斯莱乌() 	feud.Barony
    BHedeby海泽比() 	feud.Barony
    BKolding科灵() 	feud.Barony
    BRibe里伯() 	feud.Barony
    BSlesvig石勒苏益格() 	feud.Barony
    BSonderborg森讷堡() 	feud.Barony
    BTonder岑讷() 	feud.Barony
}

type 石勒苏益格SlesvigCounty struct {
	feud.BaseCounty
	奥本罗Aabenraa 	feud.Barony
	弗伦斯堡Flensborg 	feud.Barony
	哈泽斯莱乌Haderslev 	feud.Barony
	海泽比Hedeby 	feud.Barony
	科灵Kolding 	feud.Barony
	里伯Ribe 	feud.Barony
	石勒苏益格Slesvig 	feud.Barony
	森讷堡Sonderborg 	feud.Barony
	岑讷Tonder 	feud.Barony
}

func (c *石勒苏益格SlesvigCounty) BAabenraa奥本罗() feud.Barony {
	return c.奥本罗Aabenraa
}
    
func (c *石勒苏益格SlesvigCounty) BFlensborg弗伦斯堡() feud.Barony {
	return c.弗伦斯堡Flensborg
}
    
func (c *石勒苏益格SlesvigCounty) BHaderslev哈泽斯莱乌() feud.Barony {
	return c.哈泽斯莱乌Haderslev
}
    
func (c *石勒苏益格SlesvigCounty) BHedeby海泽比() feud.Barony {
	return c.海泽比Hedeby
}
    
func (c *石勒苏益格SlesvigCounty) BKolding科灵() feud.Barony {
	return c.科灵Kolding
}
    
func (c *石勒苏益格SlesvigCounty) BRibe里伯() feud.Barony {
	return c.里伯Ribe
}
    
func (c *石勒苏益格SlesvigCounty) BSlesvig石勒苏益格() feud.Barony {
	return c.石勒苏益格Slesvig
}
    
func (c *石勒苏益格SlesvigCounty) BSonderborg森讷堡() feud.Barony {
	return c.森讷堡Sonderborg
}
    
func (c *石勒苏益格SlesvigCounty) BTonder岑讷() feud.Barony {
	return c.岑讷Tonder
}
    
var CSlesvig石勒苏益格 SlesvigCounty = &石勒苏益格SlesvigCounty{}

func init() {
	f := CSlesvig石勒苏益格.(*石勒苏益格SlesvigCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "264",
		Title:     "slesvig",
		TitleName: "石勒苏益格",
		TitleCode: "c_slesvig",
		Baronies:  map[string]feud.Barony{},
	}

	f.奥本罗Aabenraa = BAabenraa奥本罗
	f.奥本罗Aabenraa.SetParent(f)
	
	f.弗伦斯堡Flensborg = BFlensborg弗伦斯堡
	f.弗伦斯堡Flensborg.SetParent(f)
	
	f.哈泽斯莱乌Haderslev = BHaderslev哈泽斯莱乌
	f.哈泽斯莱乌Haderslev.SetParent(f)
	
	f.海泽比Hedeby = BHedeby海泽比
	f.海泽比Hedeby.SetParent(f)
	
	f.科灵Kolding = BKolding科灵
	f.科灵Kolding.SetParent(f)
	
	f.里伯Ribe = BRibe里伯
	f.里伯Ribe.SetParent(f)
	
	f.石勒苏益格Slesvig = BSlesvig石勒苏益格
	f.石勒苏益格Slesvig.SetParent(f)
	
	f.森讷堡Sonderborg = BSonderborg森讷堡
	f.森讷堡Sonderborg.SetParent(f)
	
	f.岑讷Tonder = BTonder岑讷
	f.岑讷Tonder.SetParent(f)
	
}
