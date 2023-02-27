package nyland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type NylandCounty interface {
    feud.County
    BEspoo埃斯波() 	feud.Barony
    BHamina哈米纳() 	feud.Barony
    BHanko汉科() 	feud.Barony
    BKotka科特卡() 	feud.Barony
    BLoviisa洛维萨() 	feud.Barony
    BPorvoo波尔沃() 	feud.Barony
    BSvartholm斯瓦特霍尔姆() 	feud.Barony
}

type 新地NylandCounty struct {
	feud.BaseCounty
	埃斯波Espoo 	feud.Barony
	哈米纳Hamina 	feud.Barony
	汉科Hanko 	feud.Barony
	科特卡Kotka 	feud.Barony
	洛维萨Loviisa 	feud.Barony
	波尔沃Porvoo 	feud.Barony
	斯瓦特霍尔姆Svartholm 	feud.Barony
}

func (c *新地NylandCounty) BEspoo埃斯波() feud.Barony {
	return c.埃斯波Espoo
}
    
func (c *新地NylandCounty) BHamina哈米纳() feud.Barony {
	return c.哈米纳Hamina
}
    
func (c *新地NylandCounty) BHanko汉科() feud.Barony {
	return c.汉科Hanko
}
    
func (c *新地NylandCounty) BKotka科特卡() feud.Barony {
	return c.科特卡Kotka
}
    
func (c *新地NylandCounty) BLoviisa洛维萨() feud.Barony {
	return c.洛维萨Loviisa
}
    
func (c *新地NylandCounty) BPorvoo波尔沃() feud.Barony {
	return c.波尔沃Porvoo
}
    
func (c *新地NylandCounty) BSvartholm斯瓦特霍尔姆() feud.Barony {
	return c.斯瓦特霍尔姆Svartholm
}
    
var CNyland新地 NylandCounty = &新地NylandCounty{}

func init() {
	f := CNyland新地.(*新地NylandCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "381",
		Title:     "nyland",
		TitleName: "新地",
		TitleCode: "c_nyland",
		Baronies:  map[string]feud.Barony{},
	}

	f.埃斯波Espoo = BEspoo埃斯波
	f.埃斯波Espoo.SetParent(f)
	
	f.哈米纳Hamina = BHamina哈米纳
	f.哈米纳Hamina.SetParent(f)
	
	f.汉科Hanko = BHanko汉科
	f.汉科Hanko.SetParent(f)
	
	f.科特卡Kotka = BKotka科特卡
	f.科特卡Kotka.SetParent(f)
	
	f.洛维萨Loviisa = BLoviisa洛维萨
	f.洛维萨Loviisa.SetParent(f)
	
	f.波尔沃Porvoo = BPorvoo波尔沃
	f.波尔沃Porvoo.SetParent(f)
	
	f.斯瓦特霍尔姆Svartholm = BSvartholm斯瓦特霍尔姆
	f.斯瓦特霍尔姆Svartholm.SetParent(f)
	
}
