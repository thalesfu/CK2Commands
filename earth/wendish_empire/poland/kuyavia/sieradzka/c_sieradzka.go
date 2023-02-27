package sieradzka

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SieradzkaCounty interface {
    feud.County
    BPiotrkow彼得库夫() 	feud.Barony
    BPrzedborz普热德布日() 	feud.Barony
    BRadomsko拉多姆斯科() 	feud.Barony
    BRozprza罗兹普沙() 	feud.Barony
    BSieradz谢拉兹() 	feud.Barony
    BSzadek沙代克() 	feud.Barony
    BTurek图雷克() 	feud.Barony
}

type 谢拉兹SieradzkaCounty struct {
	feud.BaseCounty
	彼得库夫Piotrkow 	feud.Barony
	普热德布日Przedborz 	feud.Barony
	拉多姆斯科Radomsko 	feud.Barony
	罗兹普沙Rozprza 	feud.Barony
	谢拉兹Sieradz 	feud.Barony
	沙代克Szadek 	feud.Barony
	图雷克Turek 	feud.Barony
}

func (c *谢拉兹SieradzkaCounty) BPiotrkow彼得库夫() feud.Barony {
	return c.彼得库夫Piotrkow
}
    
func (c *谢拉兹SieradzkaCounty) BPrzedborz普热德布日() feud.Barony {
	return c.普热德布日Przedborz
}
    
func (c *谢拉兹SieradzkaCounty) BRadomsko拉多姆斯科() feud.Barony {
	return c.拉多姆斯科Radomsko
}
    
func (c *谢拉兹SieradzkaCounty) BRozprza罗兹普沙() feud.Barony {
	return c.罗兹普沙Rozprza
}
    
func (c *谢拉兹SieradzkaCounty) BSieradz谢拉兹() feud.Barony {
	return c.谢拉兹Sieradz
}
    
func (c *谢拉兹SieradzkaCounty) BSzadek沙代克() feud.Barony {
	return c.沙代克Szadek
}
    
func (c *谢拉兹SieradzkaCounty) BTurek图雷克() feud.Barony {
	return c.图雷克Turek
}
    
var CSieradzka谢拉兹 SieradzkaCounty = &谢拉兹SieradzkaCounty{}

func init() {
	f := CSieradzka谢拉兹.(*谢拉兹SieradzkaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1582",
		Title:     "sieradzka",
		TitleName: "谢拉兹",
		TitleCode: "c_sieradzka",
		Baronies:  map[string]feud.Barony{},
	}

	f.彼得库夫Piotrkow = BPiotrkow彼得库夫
	f.彼得库夫Piotrkow.SetParent(f)
	
	f.普热德布日Przedborz = BPrzedborz普热德布日
	f.普热德布日Przedborz.SetParent(f)
	
	f.拉多姆斯科Radomsko = BRadomsko拉多姆斯科
	f.拉多姆斯科Radomsko.SetParent(f)
	
	f.罗兹普沙Rozprza = BRozprza罗兹普沙
	f.罗兹普沙Rozprza.SetParent(f)
	
	f.谢拉兹Sieradz = BSieradz谢拉兹
	f.谢拉兹Sieradz.SetParent(f)
	
	f.沙代克Szadek = BSzadek沙代克
	f.沙代克Szadek.SetParent(f)
	
	f.图雷克Turek = BTurek图雷克
	f.图雷克Turek.SetParent(f)
	
}
