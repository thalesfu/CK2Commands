package powys

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type PowysCounty interface {
    feud.County
    BCaersws凯尔苏斯() 	feud.Barony
    BGlascwm格拉斯库姆() 	feud.Barony
    BLlandinam兰迪南() 	feud.Barony
    BLlangollen兰戈伦() 	feud.Barony
    BMathrafal马斯拉法尔() 	feud.Barony
    BMontgomery蒙哥马利() 	feud.Barony
    BRadnor拉德诺() 	feud.Barony
    BRhayader里达尔特() 	feud.Barony
}

type 波伊斯PowysCounty struct {
	feud.BaseCounty
	凯尔苏斯Caersws 	feud.Barony
	格拉斯库姆Glascwm 	feud.Barony
	兰迪南Llandinam 	feud.Barony
	兰戈伦Llangollen 	feud.Barony
	马斯拉法尔Mathrafal 	feud.Barony
	蒙哥马利Montgomery 	feud.Barony
	拉德诺Radnor 	feud.Barony
	里达尔特Rhayader 	feud.Barony
}

func (c *波伊斯PowysCounty) BCaersws凯尔苏斯() feud.Barony {
	return c.凯尔苏斯Caersws
}
    
func (c *波伊斯PowysCounty) BGlascwm格拉斯库姆() feud.Barony {
	return c.格拉斯库姆Glascwm
}
    
func (c *波伊斯PowysCounty) BLlandinam兰迪南() feud.Barony {
	return c.兰迪南Llandinam
}
    
func (c *波伊斯PowysCounty) BLlangollen兰戈伦() feud.Barony {
	return c.兰戈伦Llangollen
}
    
func (c *波伊斯PowysCounty) BMathrafal马斯拉法尔() feud.Barony {
	return c.马斯拉法尔Mathrafal
}
    
func (c *波伊斯PowysCounty) BMontgomery蒙哥马利() feud.Barony {
	return c.蒙哥马利Montgomery
}
    
func (c *波伊斯PowysCounty) BRadnor拉德诺() feud.Barony {
	return c.拉德诺Radnor
}
    
func (c *波伊斯PowysCounty) BRhayader里达尔特() feud.Barony {
	return c.里达尔特Rhayader
}
    
var CPowys波伊斯 PowysCounty = &波伊斯PowysCounty{}

func init() {
	f := CPowys波伊斯.(*波伊斯PowysCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "65",
		Title:     "powys",
		TitleName: "波伊斯",
		TitleCode: "c_powys",
		Baronies:  map[string]feud.Barony{},
	}

	f.凯尔苏斯Caersws = BCaersws凯尔苏斯
	f.凯尔苏斯Caersws.SetParent(f)
	
	f.格拉斯库姆Glascwm = BGlascwm格拉斯库姆
	f.格拉斯库姆Glascwm.SetParent(f)
	
	f.兰迪南Llandinam = BLlandinam兰迪南
	f.兰迪南Llandinam.SetParent(f)
	
	f.兰戈伦Llangollen = BLlangollen兰戈伦
	f.兰戈伦Llangollen.SetParent(f)
	
	f.马斯拉法尔Mathrafal = BMathrafal马斯拉法尔
	f.马斯拉法尔Mathrafal.SetParent(f)
	
	f.蒙哥马利Montgomery = BMontgomery蒙哥马利
	f.蒙哥马利Montgomery.SetParent(f)
	
	f.拉德诺Radnor = BRadnor拉德诺
	f.拉德诺Radnor.SetParent(f)
	
	f.里达尔特Rhayader = BRhayader里达尔特
	f.里达尔特Rhayader.SetParent(f)
	
}
