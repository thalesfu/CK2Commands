package karmanta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KarmantaCounty interface {
    feud.County
    BKarmanta羯曼多() 	feud.Barony
    BLatembarcem卢艺婆旃() 	feud.Barony
    BLavangi卢槃吉() 	feud.Barony
    BMahalpur摩诃罗补罗() 	feud.Barony
    BMainamati梅那摩提() 	feud.Barony
    BUdaipur优陀耶补罗() 	feud.Barony
    BUnakoti优那拘胝() 	feud.Barony
}

type 羯曼多KarmantaCounty struct {
	feud.BaseCounty
	羯曼多Karmanta 	feud.Barony
	卢艺婆旃Latembarcem 	feud.Barony
	卢槃吉Lavangi 	feud.Barony
	摩诃罗补罗Mahalpur 	feud.Barony
	梅那摩提Mainamati 	feud.Barony
	优陀耶补罗Udaipur 	feud.Barony
	优那拘胝Unakoti 	feud.Barony
}

func (c *羯曼多KarmantaCounty) BKarmanta羯曼多() feud.Barony {
	return c.羯曼多Karmanta
}
    
func (c *羯曼多KarmantaCounty) BLatembarcem卢艺婆旃() feud.Barony {
	return c.卢艺婆旃Latembarcem
}
    
func (c *羯曼多KarmantaCounty) BLavangi卢槃吉() feud.Barony {
	return c.卢槃吉Lavangi
}
    
func (c *羯曼多KarmantaCounty) BMahalpur摩诃罗补罗() feud.Barony {
	return c.摩诃罗补罗Mahalpur
}
    
func (c *羯曼多KarmantaCounty) BMainamati梅那摩提() feud.Barony {
	return c.梅那摩提Mainamati
}
    
func (c *羯曼多KarmantaCounty) BUdaipur优陀耶补罗() feud.Barony {
	return c.优陀耶补罗Udaipur
}
    
func (c *羯曼多KarmantaCounty) BUnakoti优那拘胝() feud.Barony {
	return c.优那拘胝Unakoti
}
    
var CKarmanta羯曼多 KarmantaCounty = &羯曼多KarmantaCounty{}

func init() {
	f := CKarmanta羯曼多.(*羯曼多KarmantaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1131",
		Title:     "karmanta",
		TitleName: "羯曼多",
		TitleCode: "c_karmanta",
		Baronies:  map[string]feud.Barony{},
	}

	f.羯曼多Karmanta = BKarmanta羯曼多
	f.羯曼多Karmanta.SetParent(f)
	
	f.卢艺婆旃Latembarcem = BLatembarcem卢艺婆旃
	f.卢艺婆旃Latembarcem.SetParent(f)
	
	f.卢槃吉Lavangi = BLavangi卢槃吉
	f.卢槃吉Lavangi.SetParent(f)
	
	f.摩诃罗补罗Mahalpur = BMahalpur摩诃罗补罗
	f.摩诃罗补罗Mahalpur.SetParent(f)
	
	f.梅那摩提Mainamati = BMainamati梅那摩提
	f.梅那摩提Mainamati.SetParent(f)
	
	f.优陀耶补罗Udaipur = BUdaipur优陀耶补罗
	f.优陀耶补罗Udaipur.SetParent(f)
	
	f.优那拘胝Unakoti = BUnakoti优那拘胝
	f.优那拘胝Unakoti.SetParent(f)
	
}
