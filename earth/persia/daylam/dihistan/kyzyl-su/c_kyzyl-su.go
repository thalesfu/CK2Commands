package kyzyl-su

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type Kyzyl-suCounty interface {
    feud.County
    BAkdash阿克达什() 	feud.Barony
    BAwaza阿瓦扎() 	feud.Barony
    BBelek别列克() 	feud.Barony
    BDzhanga占加() 	feud.Barony
    BKenar克纳尔() 	feud.Barony
    BKyzyl-su克孜勒苏() 	feud.Barony
    BYangadzha扬加贾() 	feud.Barony
}

type 克孜勒苏Kyzyl-suCounty struct {
	feud.BaseCounty
	阿克达什Akdash 	feud.Barony
	阿瓦扎Awaza 	feud.Barony
	别列克Belek 	feud.Barony
	占加Dzhanga 	feud.Barony
	克纳尔Kenar 	feud.Barony
	克孜勒苏Kyzyl-su 	feud.Barony
	扬加贾Yangadzha 	feud.Barony
}

func (c *克孜勒苏Kyzyl-suCounty) BAkdash阿克达什() feud.Barony {
	return c.阿克达什Akdash
}
    
func (c *克孜勒苏Kyzyl-suCounty) BAwaza阿瓦扎() feud.Barony {
	return c.阿瓦扎Awaza
}
    
func (c *克孜勒苏Kyzyl-suCounty) BBelek别列克() feud.Barony {
	return c.别列克Belek
}
    
func (c *克孜勒苏Kyzyl-suCounty) BDzhanga占加() feud.Barony {
	return c.占加Dzhanga
}
    
func (c *克孜勒苏Kyzyl-suCounty) BKenar克纳尔() feud.Barony {
	return c.克纳尔Kenar
}
    
func (c *克孜勒苏Kyzyl-suCounty) BKyzyl-su克孜勒苏() feud.Barony {
	return c.克孜勒苏Kyzyl-su
}
    
func (c *克孜勒苏Kyzyl-suCounty) BYangadzha扬加贾() feud.Barony {
	return c.扬加贾Yangadzha
}
    
var CKyzyl-su克孜勒苏 Kyzyl-suCounty = &克孜勒苏Kyzyl-suCounty{}

func init() {
	f := CKyzyl-su克孜勒苏.(*克孜勒苏Kyzyl-suCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1789",
		Title:     "kyzyl-su",
		TitleName: "克孜勒苏",
		TitleCode: "c_kyzyl-su",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿克达什Akdash = BAkdash阿克达什
	f.阿克达什Akdash.SetParent(f)
	
	f.阿瓦扎Awaza = BAwaza阿瓦扎
	f.阿瓦扎Awaza.SetParent(f)
	
	f.别列克Belek = BBelek别列克
	f.别列克Belek.SetParent(f)
	
	f.占加Dzhanga = BDzhanga占加
	f.占加Dzhanga.SetParent(f)
	
	f.克纳尔Kenar = BKenar克纳尔
	f.克纳尔Kenar.SetParent(f)
	
	f.克孜勒苏Kyzyl-su = BKyzyl-su克孜勒苏
	f.克孜勒苏Kyzyl-su.SetParent(f)
	
	f.扬加贾Yangadzha = BYangadzha扬加贾
	f.扬加贾Yangadzha.SetParent(f)
	
}
