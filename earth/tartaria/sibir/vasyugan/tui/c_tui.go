package tui

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TuiCounty interface {
    feud.County
    BAbaul阿巴乌尔() 	feud.Barony
    BAksurka阿克苏尔卡() 	feud.Barony
    BAzy阿济() 	feud.Barony
    BSaurgachi绍尔加奇() 	feud.Barony
    BShabry沙布雷() 	feud.Barony
    BSupra苏普拉() 	feud.Barony
    BTui图伊() 	feud.Barony
}

type 图伊TuiCounty struct {
	feud.BaseCounty
	阿巴乌尔Abaul 	feud.Barony
	阿克苏尔卡Aksurka 	feud.Barony
	阿济Azy 	feud.Barony
	绍尔加奇Saurgachi 	feud.Barony
	沙布雷Shabry 	feud.Barony
	苏普拉Supra 	feud.Barony
	图伊Tui 	feud.Barony
}

func (c *图伊TuiCounty) BAbaul阿巴乌尔() feud.Barony {
	return c.阿巴乌尔Abaul
}
    
func (c *图伊TuiCounty) BAksurka阿克苏尔卡() feud.Barony {
	return c.阿克苏尔卡Aksurka
}
    
func (c *图伊TuiCounty) BAzy阿济() feud.Barony {
	return c.阿济Azy
}
    
func (c *图伊TuiCounty) BSaurgachi绍尔加奇() feud.Barony {
	return c.绍尔加奇Saurgachi
}
    
func (c *图伊TuiCounty) BShabry沙布雷() feud.Barony {
	return c.沙布雷Shabry
}
    
func (c *图伊TuiCounty) BSupra苏普拉() feud.Barony {
	return c.苏普拉Supra
}
    
func (c *图伊TuiCounty) BTui图伊() feud.Barony {
	return c.图伊Tui
}
    
var CTui图伊 TuiCounty = &图伊TuiCounty{}

func init() {
	f := CTui图伊.(*图伊TuiCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1865",
		Title:     "tui",
		TitleName: "图伊",
		TitleCode: "c_tui",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿巴乌尔Abaul = BAbaul阿巴乌尔
	f.阿巴乌尔Abaul.SetParent(f)
	
	f.阿克苏尔卡Aksurka = BAksurka阿克苏尔卡
	f.阿克苏尔卡Aksurka.SetParent(f)
	
	f.阿济Azy = BAzy阿济
	f.阿济Azy.SetParent(f)
	
	f.绍尔加奇Saurgachi = BSaurgachi绍尔加奇
	f.绍尔加奇Saurgachi.SetParent(f)
	
	f.沙布雷Shabry = BShabry沙布雷
	f.沙布雷Shabry.SetParent(f)
	
	f.苏普拉Supra = BSupra苏普拉
	f.苏普拉Supra.SetParent(f)
	
	f.图伊Tui = BTui图伊
	f.图伊Tui.SetParent(f)
	
}
