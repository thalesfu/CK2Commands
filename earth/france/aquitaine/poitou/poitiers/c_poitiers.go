package poitiers

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type PoitiersCounty interface {
    feud.County
    BChatellerault沙泰勒罗() 	feud.Barony
    BChauvigny绍维尼() 	feud.Barony
    BLoudun卢丹() 	feud.Barony
    BMirebeau米尔博() 	feud.Barony
    BMoncontour蒙孔图尔() 	feud.Barony
    BParthenay帕尔特奈() 	feud.Barony
    BPoitiers普瓦捷() 	feud.Barony
    BStsavin圣萨万() 	feud.Barony
}

type 普瓦捷PoitiersCounty struct {
	feud.BaseCounty
	沙泰勒罗Chatellerault 	feud.Barony
	绍维尼Chauvigny 	feud.Barony
	卢丹Loudun 	feud.Barony
	米尔博Mirebeau 	feud.Barony
	蒙孔图尔Moncontour 	feud.Barony
	帕尔特奈Parthenay 	feud.Barony
	普瓦捷Poitiers 	feud.Barony
	圣萨万Stsavin 	feud.Barony
}

func (c *普瓦捷PoitiersCounty) BChatellerault沙泰勒罗() feud.Barony {
	return c.沙泰勒罗Chatellerault
}
    
func (c *普瓦捷PoitiersCounty) BChauvigny绍维尼() feud.Barony {
	return c.绍维尼Chauvigny
}
    
func (c *普瓦捷PoitiersCounty) BLoudun卢丹() feud.Barony {
	return c.卢丹Loudun
}
    
func (c *普瓦捷PoitiersCounty) BMirebeau米尔博() feud.Barony {
	return c.米尔博Mirebeau
}
    
func (c *普瓦捷PoitiersCounty) BMoncontour蒙孔图尔() feud.Barony {
	return c.蒙孔图尔Moncontour
}
    
func (c *普瓦捷PoitiersCounty) BParthenay帕尔特奈() feud.Barony {
	return c.帕尔特奈Parthenay
}
    
func (c *普瓦捷PoitiersCounty) BPoitiers普瓦捷() feud.Barony {
	return c.普瓦捷Poitiers
}
    
func (c *普瓦捷PoitiersCounty) BStsavin圣萨万() feud.Barony {
	return c.圣萨万Stsavin
}
    
var CPoitiers普瓦捷 PoitiersCounty = &普瓦捷PoitiersCounty{}

func init() {
	f := CPoitiers普瓦捷.(*普瓦捷PoitiersCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "141",
		Title:     "poitiers",
		TitleName: "普瓦捷",
		TitleCode: "c_poitiers",
		Baronies:  map[string]feud.Barony{},
	}

	f.沙泰勒罗Chatellerault = BChatellerault沙泰勒罗
	f.沙泰勒罗Chatellerault.SetParent(f)
	
	f.绍维尼Chauvigny = BChauvigny绍维尼
	f.绍维尼Chauvigny.SetParent(f)
	
	f.卢丹Loudun = BLoudun卢丹
	f.卢丹Loudun.SetParent(f)
	
	f.米尔博Mirebeau = BMirebeau米尔博
	f.米尔博Mirebeau.SetParent(f)
	
	f.蒙孔图尔Moncontour = BMoncontour蒙孔图尔
	f.蒙孔图尔Moncontour.SetParent(f)
	
	f.帕尔特奈Parthenay = BParthenay帕尔特奈
	f.帕尔特奈Parthenay.SetParent(f)
	
	f.普瓦捷Poitiers = BPoitiers普瓦捷
	f.普瓦捷Poitiers.SetParent(f)
	
	f.圣萨万Stsavin = BStsavin圣萨万
	f.圣萨万Stsavin.SetParent(f)
	
}
