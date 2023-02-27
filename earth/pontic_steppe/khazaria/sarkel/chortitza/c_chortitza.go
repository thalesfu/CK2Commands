package chortitza

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ChortitzaCounty interface {
    feud.County
    BAlexandrowsk亚历山德罗夫斯克() 	feud.Barony
    BBaszmacka巴什马奇卡() 	feud.Barony
    BChortitza霍尔季察() 	feud.Barony
    BKhorolya霍罗利亚() 	feud.Barony
    BLtava勒塔瓦() 	feud.Barony
    BRasumowka拉祖莫夫卡() 	feud.Barony
    BVosnesjensk沃兹涅先斯克() 	feud.Barony
}

type 霍尔季察ChortitzaCounty struct {
	feud.BaseCounty
	亚历山德罗夫斯克Alexandrowsk 	feud.Barony
	巴什马奇卡Baszmacka 	feud.Barony
	霍尔季察Chortitza 	feud.Barony
	霍罗利亚Khorolya 	feud.Barony
	勒塔瓦Ltava 	feud.Barony
	拉祖莫夫卡Rasumowka 	feud.Barony
	沃兹涅先斯克Vosnesjensk 	feud.Barony
}

func (c *霍尔季察ChortitzaCounty) BAlexandrowsk亚历山德罗夫斯克() feud.Barony {
	return c.亚历山德罗夫斯克Alexandrowsk
}
    
func (c *霍尔季察ChortitzaCounty) BBaszmacka巴什马奇卡() feud.Barony {
	return c.巴什马奇卡Baszmacka
}
    
func (c *霍尔季察ChortitzaCounty) BChortitza霍尔季察() feud.Barony {
	return c.霍尔季察Chortitza
}
    
func (c *霍尔季察ChortitzaCounty) BKhorolya霍罗利亚() feud.Barony {
	return c.霍罗利亚Khorolya
}
    
func (c *霍尔季察ChortitzaCounty) BLtava勒塔瓦() feud.Barony {
	return c.勒塔瓦Ltava
}
    
func (c *霍尔季察ChortitzaCounty) BRasumowka拉祖莫夫卡() feud.Barony {
	return c.拉祖莫夫卡Rasumowka
}
    
func (c *霍尔季察ChortitzaCounty) BVosnesjensk沃兹涅先斯克() feud.Barony {
	return c.沃兹涅先斯克Vosnesjensk
}
    
var CChortitza霍尔季察 ChortitzaCounty = &霍尔季察ChortitzaCounty{}

func init() {
	f := CChortitza霍尔季察.(*霍尔季察ChortitzaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "556",
		Title:     "chortitza",
		TitleName: "霍尔季察",
		TitleCode: "c_chortitza",
		Baronies:  map[string]feud.Barony{},
	}

	f.亚历山德罗夫斯克Alexandrowsk = BAlexandrowsk亚历山德罗夫斯克
	f.亚历山德罗夫斯克Alexandrowsk.SetParent(f)
	
	f.巴什马奇卡Baszmacka = BBaszmacka巴什马奇卡
	f.巴什马奇卡Baszmacka.SetParent(f)
	
	f.霍尔季察Chortitza = BChortitza霍尔季察
	f.霍尔季察Chortitza.SetParent(f)
	
	f.霍罗利亚Khorolya = BKhorolya霍罗利亚
	f.霍罗利亚Khorolya.SetParent(f)
	
	f.勒塔瓦Ltava = BLtava勒塔瓦
	f.勒塔瓦Ltava.SetParent(f)
	
	f.拉祖莫夫卡Rasumowka = BRasumowka拉祖莫夫卡
	f.拉祖莫夫卡Rasumowka.SetParent(f)
	
	f.沃兹涅先斯克Vosnesjensk = BVosnesjensk沃兹涅先斯克
	f.沃兹涅先斯克Vosnesjensk.SetParent(f)
	
}
