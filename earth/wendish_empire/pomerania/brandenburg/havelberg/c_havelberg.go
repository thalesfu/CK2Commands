package havelberg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type HavelbergCounty interface {
    feud.County
    BHavelberg哈弗尔贝格() 	feud.Barony
    BJerichow耶里肖() 	feud.Barony
    BLenzen伦岑() 	feud.Barony
    BPlattenburg普拉滕堡() 	feud.Barony
    BSchonhausen申豪森() 	feud.Barony
    BWilsnack维尔斯纳克() 	feud.Barony
    BWittstock维特施托克() 	feud.Barony
}

type 哈弗尔贝格HavelbergCounty struct {
	feud.BaseCounty
	哈弗尔贝格Havelberg 	feud.Barony
	耶里肖Jerichow 	feud.Barony
	伦岑Lenzen 	feud.Barony
	普拉滕堡Plattenburg 	feud.Barony
	申豪森Schonhausen 	feud.Barony
	维尔斯纳克Wilsnack 	feud.Barony
	维特施托克Wittstock 	feud.Barony
}

func (c *哈弗尔贝格HavelbergCounty) BHavelberg哈弗尔贝格() feud.Barony {
	return c.哈弗尔贝格Havelberg
}
    
func (c *哈弗尔贝格HavelbergCounty) BJerichow耶里肖() feud.Barony {
	return c.耶里肖Jerichow
}
    
func (c *哈弗尔贝格HavelbergCounty) BLenzen伦岑() feud.Barony {
	return c.伦岑Lenzen
}
    
func (c *哈弗尔贝格HavelbergCounty) BPlattenburg普拉滕堡() feud.Barony {
	return c.普拉滕堡Plattenburg
}
    
func (c *哈弗尔贝格HavelbergCounty) BSchonhausen申豪森() feud.Barony {
	return c.申豪森Schonhausen
}
    
func (c *哈弗尔贝格HavelbergCounty) BWilsnack维尔斯纳克() feud.Barony {
	return c.维尔斯纳克Wilsnack
}
    
func (c *哈弗尔贝格HavelbergCounty) BWittstock维特施托克() feud.Barony {
	return c.维特施托克Wittstock
}
    
var CHavelberg哈弗尔贝格 HavelbergCounty = &哈弗尔贝格HavelbergCounty{}

func init() {
	f := CHavelberg哈弗尔贝格.(*哈弗尔贝格HavelbergCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1988",
		Title:     "havelberg",
		TitleName: "哈弗尔贝格",
		TitleCode: "c_havelberg",
		Baronies:  map[string]feud.Barony{},
	}

	f.哈弗尔贝格Havelberg = BHavelberg哈弗尔贝格
	f.哈弗尔贝格Havelberg.SetParent(f)
	
	f.耶里肖Jerichow = BJerichow耶里肖
	f.耶里肖Jerichow.SetParent(f)
	
	f.伦岑Lenzen = BLenzen伦岑
	f.伦岑Lenzen.SetParent(f)
	
	f.普拉滕堡Plattenburg = BPlattenburg普拉滕堡
	f.普拉滕堡Plattenburg.SetParent(f)
	
	f.申豪森Schonhausen = BSchonhausen申豪森
	f.申豪森Schonhausen.SetParent(f)
	
	f.维尔斯纳克Wilsnack = BWilsnack维尔斯纳克
	f.维尔斯纳克Wilsnack.SetParent(f)
	
	f.维特施托克Wittstock = BWittstock维特施托克
	f.维特施托克Wittstock.SetParent(f)
	
}
