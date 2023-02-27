package vodi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type VodiCounty interface {
    feud.County
    BIvanovskaya泰于西内() 	feud.Barony
    BJogopera约格佩雷() 	feud.Barony
    BKhotchino霍奇诺() 	feud.Barony
    BKingisepp亚马() 	feud.Barony
    BLiivtsula利夫特许莱() 	feud.Barony
    BNosok诺索克() 	feud.Barony
    BNoteborg佩赫基内林纳() 	feud.Barony
    BNyen尼恩() 	feud.Barony
}

type 英格里亚VodiCounty struct {
	feud.BaseCounty
	泰于西内Ivanovskaya 	feud.Barony
	约格佩雷Jogopera 	feud.Barony
	霍奇诺Khotchino 	feud.Barony
	亚马Kingisepp 	feud.Barony
	利夫特许莱Liivtsula 	feud.Barony
	诺索克Nosok 	feud.Barony
	佩赫基内林纳Noteborg 	feud.Barony
	尼恩Nyen 	feud.Barony
}

func (c *英格里亚VodiCounty) BIvanovskaya泰于西内() feud.Barony {
	return c.泰于西内Ivanovskaya
}
    
func (c *英格里亚VodiCounty) BJogopera约格佩雷() feud.Barony {
	return c.约格佩雷Jogopera
}
    
func (c *英格里亚VodiCounty) BKhotchino霍奇诺() feud.Barony {
	return c.霍奇诺Khotchino
}
    
func (c *英格里亚VodiCounty) BKingisepp亚马() feud.Barony {
	return c.亚马Kingisepp
}
    
func (c *英格里亚VodiCounty) BLiivtsula利夫特许莱() feud.Barony {
	return c.利夫特许莱Liivtsula
}
    
func (c *英格里亚VodiCounty) BNosok诺索克() feud.Barony {
	return c.诺索克Nosok
}
    
func (c *英格里亚VodiCounty) BNoteborg佩赫基内林纳() feud.Barony {
	return c.佩赫基内林纳Noteborg
}
    
func (c *英格里亚VodiCounty) BNyen尼恩() feud.Barony {
	return c.尼恩Nyen
}
    
var CVodi英格里亚 VodiCounty = &英格里亚VodiCounty{}

func init() {
	f := CVodi英格里亚.(*英格里亚VodiCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "411",
		Title:     "vodi",
		TitleName: "英格里亚",
		TitleCode: "c_vodi",
		Baronies:  map[string]feud.Barony{},
	}

	f.泰于西内Ivanovskaya = BIvanovskaya泰于西内
	f.泰于西内Ivanovskaya.SetParent(f)
	
	f.约格佩雷Jogopera = BJogopera约格佩雷
	f.约格佩雷Jogopera.SetParent(f)
	
	f.霍奇诺Khotchino = BKhotchino霍奇诺
	f.霍奇诺Khotchino.SetParent(f)
	
	f.亚马Kingisepp = BKingisepp亚马
	f.亚马Kingisepp.SetParent(f)
	
	f.利夫特许莱Liivtsula = BLiivtsula利夫特许莱
	f.利夫特许莱Liivtsula.SetParent(f)
	
	f.诺索克Nosok = BNosok诺索克
	f.诺索克Nosok.SetParent(f)
	
	f.佩赫基内林纳Noteborg = BNoteborg佩赫基内林纳
	f.佩赫基内林纳Noteborg.SetParent(f)
	
	f.尼恩Nyen = BNyen尼恩
	f.尼恩Nyen.SetParent(f)
	
}
