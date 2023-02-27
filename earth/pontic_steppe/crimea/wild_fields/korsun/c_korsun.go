package korsun

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KorsunCounty interface {
    feud.County
    BCherkassy切尔卡瑟() 	feud.Barony
    BChyhyryn奇吉林() 	feud.Barony
    BKirovohrad基洛夫格勒() 	feud.Barony
    BKorsun科尔孙() 	feud.Barony
    BOleksandriia亚历山德里亚() 	feud.Barony
    BSmila斯米拉() 	feud.Barony
    BZnamianka兹纳米安卡() 	feud.Barony
}

type 科尔孙KorsunCounty struct {
	feud.BaseCounty
	切尔卡瑟Cherkassy 	feud.Barony
	奇吉林Chyhyryn 	feud.Barony
	基洛夫格勒Kirovohrad 	feud.Barony
	科尔孙Korsun 	feud.Barony
	亚历山德里亚Oleksandriia 	feud.Barony
	斯米拉Smila 	feud.Barony
	兹纳米安卡Znamianka 	feud.Barony
}

func (c *科尔孙KorsunCounty) BCherkassy切尔卡瑟() feud.Barony {
	return c.切尔卡瑟Cherkassy
}
    
func (c *科尔孙KorsunCounty) BChyhyryn奇吉林() feud.Barony {
	return c.奇吉林Chyhyryn
}
    
func (c *科尔孙KorsunCounty) BKirovohrad基洛夫格勒() feud.Barony {
	return c.基洛夫格勒Kirovohrad
}
    
func (c *科尔孙KorsunCounty) BKorsun科尔孙() feud.Barony {
	return c.科尔孙Korsun
}
    
func (c *科尔孙KorsunCounty) BOleksandriia亚历山德里亚() feud.Barony {
	return c.亚历山德里亚Oleksandriia
}
    
func (c *科尔孙KorsunCounty) BSmila斯米拉() feud.Barony {
	return c.斯米拉Smila
}
    
func (c *科尔孙KorsunCounty) BZnamianka兹纳米安卡() feud.Barony {
	return c.兹纳米安卡Znamianka
}
    
var CKorsun科尔孙 KorsunCounty = &科尔孙KorsunCounty{}

func init() {
	f := CKorsun科尔孙.(*科尔孙KorsunCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "544",
		Title:     "korsun",
		TitleName: "科尔孙",
		TitleCode: "c_korsun",
		Baronies:  map[string]feud.Barony{},
	}

	f.切尔卡瑟Cherkassy = BCherkassy切尔卡瑟
	f.切尔卡瑟Cherkassy.SetParent(f)
	
	f.奇吉林Chyhyryn = BChyhyryn奇吉林
	f.奇吉林Chyhyryn.SetParent(f)
	
	f.基洛夫格勒Kirovohrad = BKirovohrad基洛夫格勒
	f.基洛夫格勒Kirovohrad.SetParent(f)
	
	f.科尔孙Korsun = BKorsun科尔孙
	f.科尔孙Korsun.SetParent(f)
	
	f.亚历山德里亚Oleksandriia = BOleksandriia亚历山德里亚
	f.亚历山德里亚Oleksandriia.SetParent(f)
	
	f.斯米拉Smila = BSmila斯米拉
	f.斯米拉Smila.SetParent(f)
	
	f.兹纳米安卡Znamianka = BZnamianka兹纳米安卡
	f.兹纳米安卡Znamianka.SetParent(f)
	
}
