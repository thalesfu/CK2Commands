package lower_volga

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type Lower_volgaCounty interface {
    feud.County
    BDubovka杜博夫卡() 	feud.Barony
    BKamyshin卡梅申() 	feud.Barony
    BKotovo科托沃() 	feud.Barony
    BPetrovval彼得罗夫瓦尔() 	feud.Barony
    BPrishib普里希布() 	feud.Barony
    BSarysu萨雷苏() 	feud.Barony
    BSerebrjakovo谢列布里亚科沃() 	feud.Barony
    BTsaritsyn察里津() 	feud.Barony
}

type 下伏尔加Lower_volgaCounty struct {
	feud.BaseCounty
	杜博夫卡Dubovka 	feud.Barony
	卡梅申Kamyshin 	feud.Barony
	科托沃Kotovo 	feud.Barony
	彼得罗夫瓦尔Petrovval 	feud.Barony
	普里希布Prishib 	feud.Barony
	萨雷苏Sarysu 	feud.Barony
	谢列布里亚科沃Serebrjakovo 	feud.Barony
	察里津Tsaritsyn 	feud.Barony
}

func (c *下伏尔加Lower_volgaCounty) BDubovka杜博夫卡() feud.Barony {
	return c.杜博夫卡Dubovka
}
    
func (c *下伏尔加Lower_volgaCounty) BKamyshin卡梅申() feud.Barony {
	return c.卡梅申Kamyshin
}
    
func (c *下伏尔加Lower_volgaCounty) BKotovo科托沃() feud.Barony {
	return c.科托沃Kotovo
}
    
func (c *下伏尔加Lower_volgaCounty) BPetrovval彼得罗夫瓦尔() feud.Barony {
	return c.彼得罗夫瓦尔Petrovval
}
    
func (c *下伏尔加Lower_volgaCounty) BPrishib普里希布() feud.Barony {
	return c.普里希布Prishib
}
    
func (c *下伏尔加Lower_volgaCounty) BSarysu萨雷苏() feud.Barony {
	return c.萨雷苏Sarysu
}
    
func (c *下伏尔加Lower_volgaCounty) BSerebrjakovo谢列布里亚科沃() feud.Barony {
	return c.谢列布里亚科沃Serebrjakovo
}
    
func (c *下伏尔加Lower_volgaCounty) BTsaritsyn察里津() feud.Barony {
	return c.察里津Tsaritsyn
}
    
var CLower_volga下伏尔加 Lower_volgaCounty = &下伏尔加Lower_volgaCounty{}

func init() {
	f := CLower_volga下伏尔加.(*下伏尔加Lower_volgaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "608",
		Title:     "lower_volga",
		TitleName: "下伏尔加",
		TitleCode: "c_lower_volga",
		Baronies:  map[string]feud.Barony{},
	}

	f.杜博夫卡Dubovka = BDubovka杜博夫卡
	f.杜博夫卡Dubovka.SetParent(f)
	
	f.卡梅申Kamyshin = BKamyshin卡梅申
	f.卡梅申Kamyshin.SetParent(f)
	
	f.科托沃Kotovo = BKotovo科托沃
	f.科托沃Kotovo.SetParent(f)
	
	f.彼得罗夫瓦尔Petrovval = BPetrovval彼得罗夫瓦尔
	f.彼得罗夫瓦尔Petrovval.SetParent(f)
	
	f.普里希布Prishib = BPrishib普里希布
	f.普里希布Prishib.SetParent(f)
	
	f.萨雷苏Sarysu = BSarysu萨雷苏
	f.萨雷苏Sarysu.SetParent(f)
	
	f.谢列布里亚科沃Serebrjakovo = BSerebrjakovo谢列布里亚科沃
	f.谢列布里亚科沃Serebrjakovo.SetParent(f)
	
	f.察里津Tsaritsyn = BTsaritsyn察里津
	f.察里津Tsaritsyn.SetParent(f)
	
}
