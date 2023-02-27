package tobol

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TobolCounty interface {
    feud.County
    BCherlak切尔拉克() 	feud.Barony
    BIsilkul伊西利库尔() 	feud.Barony
    BKalachinsk卡拉钦斯克() 	feud.Barony
    BKrasnyyar克拉斯内_亚尔() 	feud.Barony
    BSargatka萨尔加特卡() 	feud.Barony
    BTobol托博尔() 	feud.Barony
    BTyukalinsk秋卡林斯克() 	feud.Barony
}

type 托博尔TobolCounty struct {
	feud.BaseCounty
	切尔拉克Cherlak 	feud.Barony
	伊西利库尔Isilkul 	feud.Barony
	卡拉钦斯克Kalachinsk 	feud.Barony
	克拉斯内_亚尔Krasnyyar 	feud.Barony
	萨尔加特卡Sargatka 	feud.Barony
	托博尔Tobol 	feud.Barony
	秋卡林斯克Tyukalinsk 	feud.Barony
}

func (c *托博尔TobolCounty) BCherlak切尔拉克() feud.Barony {
	return c.切尔拉克Cherlak
}
    
func (c *托博尔TobolCounty) BIsilkul伊西利库尔() feud.Barony {
	return c.伊西利库尔Isilkul
}
    
func (c *托博尔TobolCounty) BKalachinsk卡拉钦斯克() feud.Barony {
	return c.卡拉钦斯克Kalachinsk
}
    
func (c *托博尔TobolCounty) BKrasnyyar克拉斯内_亚尔() feud.Barony {
	return c.克拉斯内_亚尔Krasnyyar
}
    
func (c *托博尔TobolCounty) BSargatka萨尔加特卡() feud.Barony {
	return c.萨尔加特卡Sargatka
}
    
func (c *托博尔TobolCounty) BTobol托博尔() feud.Barony {
	return c.托博尔Tobol
}
    
func (c *托博尔TobolCounty) BTyukalinsk秋卡林斯克() feud.Barony {
	return c.秋卡林斯克Tyukalinsk
}
    
var CTobol托博尔 TobolCounty = &托博尔TobolCounty{}

func init() {
	f := CTobol托博尔.(*托博尔TobolCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "898",
		Title:     "tobol",
		TitleName: "托博尔",
		TitleCode: "c_tobol",
		Baronies:  map[string]feud.Barony{},
	}

	f.切尔拉克Cherlak = BCherlak切尔拉克
	f.切尔拉克Cherlak.SetParent(f)
	
	f.伊西利库尔Isilkul = BIsilkul伊西利库尔
	f.伊西利库尔Isilkul.SetParent(f)
	
	f.卡拉钦斯克Kalachinsk = BKalachinsk卡拉钦斯克
	f.卡拉钦斯克Kalachinsk.SetParent(f)
	
	f.克拉斯内_亚尔Krasnyyar = BKrasnyyar克拉斯内_亚尔
	f.克拉斯内_亚尔Krasnyyar.SetParent(f)
	
	f.萨尔加特卡Sargatka = BSargatka萨尔加特卡
	f.萨尔加特卡Sargatka.SetParent(f)
	
	f.托博尔Tobol = BTobol托博尔
	f.托博尔Tobol.SetParent(f)
	
	f.秋卡林斯克Tyukalinsk = BTyukalinsk秋卡林斯克
	f.秋卡林斯克Tyukalinsk.SetParent(f)
	
}
