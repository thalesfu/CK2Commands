package sanaa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SanaaCounty interface {
    feud.County
    BDhamar扎马尔() 	feud.Barony
    BHarib哈里卜() 	feud.Barony
    BHodeida荷台达() 	feud.Barony
    BJabalannabishuayb先知舒艾卜峰() 	feud.Barony
    BJabaltiyal贾拜迪亚() 	feud.Barony
    BMarib马里卜() 	feud.Barony
    BQataba加泰拜() 	feud.Barony
    BSanaa萨那() 	feud.Barony
}

type 萨那SanaaCounty struct {
	feud.BaseCounty
	扎马尔Dhamar 	feud.Barony
	哈里卜Harib 	feud.Barony
	荷台达Hodeida 	feud.Barony
	先知舒艾卜峰Jabalannabishuayb 	feud.Barony
	贾拜迪亚Jabaltiyal 	feud.Barony
	马里卜Marib 	feud.Barony
	加泰拜Qataba 	feud.Barony
	萨那Sanaa 	feud.Barony
}

func (c *萨那SanaaCounty) BDhamar扎马尔() feud.Barony {
	return c.扎马尔Dhamar
}
    
func (c *萨那SanaaCounty) BHarib哈里卜() feud.Barony {
	return c.哈里卜Harib
}
    
func (c *萨那SanaaCounty) BHodeida荷台达() feud.Barony {
	return c.荷台达Hodeida
}
    
func (c *萨那SanaaCounty) BJabalannabishuayb先知舒艾卜峰() feud.Barony {
	return c.先知舒艾卜峰Jabalannabishuayb
}
    
func (c *萨那SanaaCounty) BJabaltiyal贾拜迪亚() feud.Barony {
	return c.贾拜迪亚Jabaltiyal
}
    
func (c *萨那SanaaCounty) BMarib马里卜() feud.Barony {
	return c.马里卜Marib
}
    
func (c *萨那SanaaCounty) BQataba加泰拜() feud.Barony {
	return c.加泰拜Qataba
}
    
func (c *萨那SanaaCounty) BSanaa萨那() feud.Barony {
	return c.萨那Sanaa
}
    
var CSanaa萨那 SanaaCounty = &萨那SanaaCounty{}

func init() {
	f := CSanaa萨那.(*萨那SanaaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "860",
		Title:     "sanaa",
		TitleName: "萨那",
		TitleCode: "c_sanaa",
		Baronies:  map[string]feud.Barony{},
	}

	f.扎马尔Dhamar = BDhamar扎马尔
	f.扎马尔Dhamar.SetParent(f)
	
	f.哈里卜Harib = BHarib哈里卜
	f.哈里卜Harib.SetParent(f)
	
	f.荷台达Hodeida = BHodeida荷台达
	f.荷台达Hodeida.SetParent(f)
	
	f.先知舒艾卜峰Jabalannabishuayb = BJabalannabishuayb先知舒艾卜峰
	f.先知舒艾卜峰Jabalannabishuayb.SetParent(f)
	
	f.贾拜迪亚Jabaltiyal = BJabaltiyal贾拜迪亚
	f.贾拜迪亚Jabaltiyal.SetParent(f)
	
	f.马里卜Marib = BMarib马里卜
	f.马里卜Marib.SetParent(f)
	
	f.加泰拜Qataba = BQataba加泰拜
	f.加泰拜Qataba.SetParent(f)
	
	f.萨那Sanaa = BSanaa萨那
	f.萨那Sanaa.SetParent(f)
	
}
