package monkh_khairkhan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type Monkh_khairkhanCounty interface {
    feud.County
    BBayan_olgil巴彦乌列盖() 	feud.Barony
    BFuyun富蕴() 	feud.Barony
    BJargalant扎尔嘎朗特() 	feud.Barony
    BMonkh_khairkhan蒙赫海尔汗() 	feud.Barony
    BMunkh蒙赫() 	feud.Barony
    BRashaant拉善特() 	feud.Barony
    BTsenkher臣赫尔() 	feud.Barony
}

type 蒙赫海尔汗Monkh_khairkhanCounty struct {
	feud.BaseCounty
	巴彦乌列盖Bayan_olgil 	feud.Barony
	富蕴Fuyun 	feud.Barony
	扎尔嘎朗特Jargalant 	feud.Barony
	蒙赫海尔汗Monkh_khairkhan 	feud.Barony
	蒙赫Munkh 	feud.Barony
	拉善特Rashaant 	feud.Barony
	臣赫尔Tsenkher 	feud.Barony
}

func (c *蒙赫海尔汗Monkh_khairkhanCounty) BBayan_olgil巴彦乌列盖() feud.Barony {
	return c.巴彦乌列盖Bayan_olgil
}
    
func (c *蒙赫海尔汗Monkh_khairkhanCounty) BFuyun富蕴() feud.Barony {
	return c.富蕴Fuyun
}
    
func (c *蒙赫海尔汗Monkh_khairkhanCounty) BJargalant扎尔嘎朗特() feud.Barony {
	return c.扎尔嘎朗特Jargalant
}
    
func (c *蒙赫海尔汗Monkh_khairkhanCounty) BMonkh_khairkhan蒙赫海尔汗() feud.Barony {
	return c.蒙赫海尔汗Monkh_khairkhan
}
    
func (c *蒙赫海尔汗Monkh_khairkhanCounty) BMunkh蒙赫() feud.Barony {
	return c.蒙赫Munkh
}
    
func (c *蒙赫海尔汗Monkh_khairkhanCounty) BRashaant拉善特() feud.Barony {
	return c.拉善特Rashaant
}
    
func (c *蒙赫海尔汗Monkh_khairkhanCounty) BTsenkher臣赫尔() feud.Barony {
	return c.臣赫尔Tsenkher
}
    
var CMonkh_khairkhan蒙赫海尔汗 Monkh_khairkhanCounty = &蒙赫海尔汗Monkh_khairkhanCounty{}

func init() {
	f := CMonkh_khairkhan蒙赫海尔汗.(*蒙赫海尔汗Monkh_khairkhanCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1897",
		Title:     "monkh_khairkhan",
		TitleName: "蒙赫海尔汗",
		TitleCode: "c_monkh_khairkhan",
		Baronies:  map[string]feud.Barony{},
	}

	f.巴彦乌列盖Bayan_olgil = BBayan_olgil巴彦乌列盖
	f.巴彦乌列盖Bayan_olgil.SetParent(f)
	
	f.富蕴Fuyun = BFuyun富蕴
	f.富蕴Fuyun.SetParent(f)
	
	f.扎尔嘎朗特Jargalant = BJargalant扎尔嘎朗特
	f.扎尔嘎朗特Jargalant.SetParent(f)
	
	f.蒙赫海尔汗Monkh_khairkhan = BMonkh_khairkhan蒙赫海尔汗
	f.蒙赫海尔汗Monkh_khairkhan.SetParent(f)
	
	f.蒙赫Munkh = BMunkh蒙赫
	f.蒙赫Munkh.SetParent(f)
	
	f.拉善特Rashaant = BRashaant拉善特
	f.拉善特Rashaant.SetParent(f)
	
	f.臣赫尔Tsenkher = BTsenkher臣赫尔
	f.臣赫尔Tsenkher.SetParent(f)
	
}
