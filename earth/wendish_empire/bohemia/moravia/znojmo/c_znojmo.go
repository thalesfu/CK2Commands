package znojmo

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ZnojmoCounty interface {
    feud.County
    BBitov比托夫() 	feud.Barony
    BEibenshitz艾本希茨() 	feud.Barony
    BJihlava伊赫拉瓦() 	feud.Barony
    BLouka洛乌卡() 	feud.Barony
    BMikulov米库洛夫() 	feud.Barony
    BPohansko波汉斯科() 	feud.Barony
    BTelc泰尔奇() 	feud.Barony
    BTrebic特热比奇() 	feud.Barony
    BVranov弗拉诺夫() 	feud.Barony
    BZnojmo兹诺伊莫() 	feud.Barony
}

type 兹诺伊莫ZnojmoCounty struct {
	feud.BaseCounty
	比托夫Bitov 	feud.Barony
	艾本希茨Eibenshitz 	feud.Barony
	伊赫拉瓦Jihlava 	feud.Barony
	洛乌卡Louka 	feud.Barony
	米库洛夫Mikulov 	feud.Barony
	波汉斯科Pohansko 	feud.Barony
	泰尔奇Telc 	feud.Barony
	特热比奇Trebic 	feud.Barony
	弗拉诺夫Vranov 	feud.Barony
	兹诺伊莫Znojmo 	feud.Barony
}

func (c *兹诺伊莫ZnojmoCounty) BBitov比托夫() feud.Barony {
	return c.比托夫Bitov
}
    
func (c *兹诺伊莫ZnojmoCounty) BEibenshitz艾本希茨() feud.Barony {
	return c.艾本希茨Eibenshitz
}
    
func (c *兹诺伊莫ZnojmoCounty) BJihlava伊赫拉瓦() feud.Barony {
	return c.伊赫拉瓦Jihlava
}
    
func (c *兹诺伊莫ZnojmoCounty) BLouka洛乌卡() feud.Barony {
	return c.洛乌卡Louka
}
    
func (c *兹诺伊莫ZnojmoCounty) BMikulov米库洛夫() feud.Barony {
	return c.米库洛夫Mikulov
}
    
func (c *兹诺伊莫ZnojmoCounty) BPohansko波汉斯科() feud.Barony {
	return c.波汉斯科Pohansko
}
    
func (c *兹诺伊莫ZnojmoCounty) BTelc泰尔奇() feud.Barony {
	return c.泰尔奇Telc
}
    
func (c *兹诺伊莫ZnojmoCounty) BTrebic特热比奇() feud.Barony {
	return c.特热比奇Trebic
}
    
func (c *兹诺伊莫ZnojmoCounty) BVranov弗拉诺夫() feud.Barony {
	return c.弗拉诺夫Vranov
}
    
func (c *兹诺伊莫ZnojmoCounty) BZnojmo兹诺伊莫() feud.Barony {
	return c.兹诺伊莫Znojmo
}
    
var CZnojmo兹诺伊莫 ZnojmoCounty = &兹诺伊莫ZnojmoCounty{}

func init() {
	f := CZnojmo兹诺伊莫.(*兹诺伊莫ZnojmoCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "446",
		Title:     "znojmo",
		TitleName: "兹诺伊莫",
		TitleCode: "c_znojmo",
		Baronies:  map[string]feud.Barony{},
	}

	f.比托夫Bitov = BBitov比托夫
	f.比托夫Bitov.SetParent(f)
	
	f.艾本希茨Eibenshitz = BEibenshitz艾本希茨
	f.艾本希茨Eibenshitz.SetParent(f)
	
	f.伊赫拉瓦Jihlava = BJihlava伊赫拉瓦
	f.伊赫拉瓦Jihlava.SetParent(f)
	
	f.洛乌卡Louka = BLouka洛乌卡
	f.洛乌卡Louka.SetParent(f)
	
	f.米库洛夫Mikulov = BMikulov米库洛夫
	f.米库洛夫Mikulov.SetParent(f)
	
	f.波汉斯科Pohansko = BPohansko波汉斯科
	f.波汉斯科Pohansko.SetParent(f)
	
	f.泰尔奇Telc = BTelc泰尔奇
	f.泰尔奇Telc.SetParent(f)
	
	f.特热比奇Trebic = BTrebic特热比奇
	f.特热比奇Trebic.SetParent(f)
	
	f.弗拉诺夫Vranov = BVranov弗拉诺夫
	f.弗拉诺夫Vranov.SetParent(f)
	
	f.兹诺伊莫Znojmo = BZnojmo兹诺伊莫
	f.兹诺伊莫Znojmo.SetParent(f)
	
}
