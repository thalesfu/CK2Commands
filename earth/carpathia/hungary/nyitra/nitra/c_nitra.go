package nitra

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type NitraCounty interface {
    feud.County
    BGalgoc高尔戈茨() 	feud.Barony
    BNagysurany瑙吉舒拉尼() 	feud.Barony
    BNagytapolcsany大陶波尔恰尼() 	feud.Barony
    BNyitra尼特劳() 	feud.Barony
    BPostyen珀什真() 	feud.Barony
    BPreuigan普列维甘() 	feud.Barony
    BStbenedek圣拜奈代克() 	feud.Barony
    BZabokreky扎博克雷基() 	feud.Barony
}

type 尼特拉NitraCounty struct {
	feud.BaseCounty
	高尔戈茨Galgoc 	feud.Barony
	瑙吉舒拉尼Nagysurany 	feud.Barony
	大陶波尔恰尼Nagytapolcsany 	feud.Barony
	尼特劳Nyitra 	feud.Barony
	珀什真Postyen 	feud.Barony
	普列维甘Preuigan 	feud.Barony
	圣拜奈代克Stbenedek 	feud.Barony
	扎博克雷基Zabokreky 	feud.Barony
}

func (c *尼特拉NitraCounty) BGalgoc高尔戈茨() feud.Barony {
	return c.高尔戈茨Galgoc
}
    
func (c *尼特拉NitraCounty) BNagysurany瑙吉舒拉尼() feud.Barony {
	return c.瑙吉舒拉尼Nagysurany
}
    
func (c *尼特拉NitraCounty) BNagytapolcsany大陶波尔恰尼() feud.Barony {
	return c.大陶波尔恰尼Nagytapolcsany
}
    
func (c *尼特拉NitraCounty) BNyitra尼特劳() feud.Barony {
	return c.尼特劳Nyitra
}
    
func (c *尼特拉NitraCounty) BPostyen珀什真() feud.Barony {
	return c.珀什真Postyen
}
    
func (c *尼特拉NitraCounty) BPreuigan普列维甘() feud.Barony {
	return c.普列维甘Preuigan
}
    
func (c *尼特拉NitraCounty) BStbenedek圣拜奈代克() feud.Barony {
	return c.圣拜奈代克Stbenedek
}
    
func (c *尼特拉NitraCounty) BZabokreky扎博克雷基() feud.Barony {
	return c.扎博克雷基Zabokreky
}
    
var CNitra尼特拉 NitraCounty = &尼特拉NitraCounty{}

func init() {
	f := CNitra尼特拉.(*尼特拉NitraCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "443",
		Title:     "nitra",
		TitleName: "尼特拉",
		TitleCode: "c_nitra",
		Baronies:  map[string]feud.Barony{},
	}

	f.高尔戈茨Galgoc = BGalgoc高尔戈茨
	f.高尔戈茨Galgoc.SetParent(f)
	
	f.瑙吉舒拉尼Nagysurany = BNagysurany瑙吉舒拉尼
	f.瑙吉舒拉尼Nagysurany.SetParent(f)
	
	f.大陶波尔恰尼Nagytapolcsany = BNagytapolcsany大陶波尔恰尼
	f.大陶波尔恰尼Nagytapolcsany.SetParent(f)
	
	f.尼特劳Nyitra = BNyitra尼特劳
	f.尼特劳Nyitra.SetParent(f)
	
	f.珀什真Postyen = BPostyen珀什真
	f.珀什真Postyen.SetParent(f)
	
	f.普列维甘Preuigan = BPreuigan普列维甘
	f.普列维甘Preuigan.SetParent(f)
	
	f.圣拜奈代克Stbenedek = BStbenedek圣拜奈代克
	f.圣拜奈代克Stbenedek.SetParent(f)
	
	f.扎博克雷基Zabokreky = BZabokreky扎博克雷基
	f.扎博克雷基Zabokreky.SetParent(f)
	
}
