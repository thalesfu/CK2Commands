package karachev

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KarachevCounty interface {
    feud.County
    BBolkhov博尔霍夫() 	feud.Barony
    BKarachev卡拉切夫() 	feud.Barony
    BKhotynets霍特涅茨() 	feud.Barony
    BKrasnaya克拉斯纳亚() 	feud.Barony
    BNaryshkino纳雷什基诺() 	feud.Barony
    BOrel奥廖尔() 	feud.Barony
    BRyabinovka里亚比诺夫卡() 	feud.Barony
}

type 卡拉切夫KarachevCounty struct {
	feud.BaseCounty
	博尔霍夫Bolkhov 	feud.Barony
	卡拉切夫Karachev 	feud.Barony
	霍特涅茨Khotynets 	feud.Barony
	克拉斯纳亚Krasnaya 	feud.Barony
	纳雷什基诺Naryshkino 	feud.Barony
	奥廖尔Orel 	feud.Barony
	里亚比诺夫卡Ryabinovka 	feud.Barony
}

func (c *卡拉切夫KarachevCounty) BBolkhov博尔霍夫() feud.Barony {
	return c.博尔霍夫Bolkhov
}
    
func (c *卡拉切夫KarachevCounty) BKarachev卡拉切夫() feud.Barony {
	return c.卡拉切夫Karachev
}
    
func (c *卡拉切夫KarachevCounty) BKhotynets霍特涅茨() feud.Barony {
	return c.霍特涅茨Khotynets
}
    
func (c *卡拉切夫KarachevCounty) BKrasnaya克拉斯纳亚() feud.Barony {
	return c.克拉斯纳亚Krasnaya
}
    
func (c *卡拉切夫KarachevCounty) BNaryshkino纳雷什基诺() feud.Barony {
	return c.纳雷什基诺Naryshkino
}
    
func (c *卡拉切夫KarachevCounty) BOrel奥廖尔() feud.Barony {
	return c.奥廖尔Orel
}
    
func (c *卡拉切夫KarachevCounty) BRyabinovka里亚比诺夫卡() feud.Barony {
	return c.里亚比诺夫卡Ryabinovka
}
    
var CKarachev卡拉切夫 KarachevCounty = &卡拉切夫KarachevCounty{}

func init() {
	f := CKarachev卡拉切夫.(*卡拉切夫KarachevCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1666",
		Title:     "karachev",
		TitleName: "卡拉切夫",
		TitleCode: "c_karachev",
		Baronies:  map[string]feud.Barony{},
	}

	f.博尔霍夫Bolkhov = BBolkhov博尔霍夫
	f.博尔霍夫Bolkhov.SetParent(f)
	
	f.卡拉切夫Karachev = BKarachev卡拉切夫
	f.卡拉切夫Karachev.SetParent(f)
	
	f.霍特涅茨Khotynets = BKhotynets霍特涅茨
	f.霍特涅茨Khotynets.SetParent(f)
	
	f.克拉斯纳亚Krasnaya = BKrasnaya克拉斯纳亚
	f.克拉斯纳亚Krasnaya.SetParent(f)
	
	f.纳雷什基诺Naryshkino = BNaryshkino纳雷什基诺
	f.纳雷什基诺Naryshkino.SetParent(f)
	
	f.奥廖尔Orel = BOrel奥廖尔
	f.奥廖尔Orel.SetParent(f)
	
	f.里亚比诺夫卡Ryabinovka = BRyabinovka里亚比诺夫卡
	f.里亚比诺夫卡Ryabinovka.SetParent(f)
	
}
