package pfalz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type PfalzCounty interface {
    feud.County
    BHagenau阿格诺() 	feud.Barony
    BKaiserslautern凯撒斯劳滕() 	feud.Barony
    BSpeyer施派尔() 	feud.Barony
    BStahleck施塔莱克() 	feud.Barony
    BTrifels特里费尔斯() 	feud.Barony
    BWorms沃尔姆斯() 	feud.Barony
    BZweibrucken茨韦布吕肯() 	feud.Barony
}

type 沃尔姆斯PfalzCounty struct {
	feud.BaseCounty
	阿格诺Hagenau 	feud.Barony
	凯撒斯劳滕Kaiserslautern 	feud.Barony
	施派尔Speyer 	feud.Barony
	施塔莱克Stahleck 	feud.Barony
	特里费尔斯Trifels 	feud.Barony
	沃尔姆斯Worms 	feud.Barony
	茨韦布吕肯Zweibrucken 	feud.Barony
}

func (c *沃尔姆斯PfalzCounty) BHagenau阿格诺() feud.Barony {
	return c.阿格诺Hagenau
}
    
func (c *沃尔姆斯PfalzCounty) BKaiserslautern凯撒斯劳滕() feud.Barony {
	return c.凯撒斯劳滕Kaiserslautern
}
    
func (c *沃尔姆斯PfalzCounty) BSpeyer施派尔() feud.Barony {
	return c.施派尔Speyer
}
    
func (c *沃尔姆斯PfalzCounty) BStahleck施塔莱克() feud.Barony {
	return c.施塔莱克Stahleck
}
    
func (c *沃尔姆斯PfalzCounty) BTrifels特里费尔斯() feud.Barony {
	return c.特里费尔斯Trifels
}
    
func (c *沃尔姆斯PfalzCounty) BWorms沃尔姆斯() feud.Barony {
	return c.沃尔姆斯Worms
}
    
func (c *沃尔姆斯PfalzCounty) BZweibrucken茨韦布吕肯() feud.Barony {
	return c.茨韦布吕肯Zweibrucken
}
    
var CPfalz沃尔姆斯 PfalzCounty = &沃尔姆斯PfalzCounty{}

func init() {
	f := CPfalz沃尔姆斯.(*沃尔姆斯PfalzCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "124",
		Title:     "pfalz",
		TitleName: "沃尔姆斯",
		TitleCode: "c_pfalz",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿格诺Hagenau = BHagenau阿格诺
	f.阿格诺Hagenau.SetParent(f)
	
	f.凯撒斯劳滕Kaiserslautern = BKaiserslautern凯撒斯劳滕
	f.凯撒斯劳滕Kaiserslautern.SetParent(f)
	
	f.施派尔Speyer = BSpeyer施派尔
	f.施派尔Speyer.SetParent(f)
	
	f.施塔莱克Stahleck = BStahleck施塔莱克
	f.施塔莱克Stahleck.SetParent(f)
	
	f.特里费尔斯Trifels = BTrifels特里费尔斯
	f.特里费尔斯Trifels.SetParent(f)
	
	f.沃尔姆斯Worms = BWorms沃尔姆斯
	f.沃尔姆斯Worms.SetParent(f)
	
	f.茨韦布吕肯Zweibrucken = BZweibrucken茨韦布吕肯
	f.茨韦布吕肯Zweibrucken.SetParent(f)
	
}
