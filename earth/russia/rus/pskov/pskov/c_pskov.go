package pskov

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type PskovCounty interface {
    feud.County
    BDedovichi杰多维奇() 	feud.Barony
    BGolubovo戈卢博沃() 	feud.Barony
    BGoristo戈里斯托() 	feud.Barony
    BKhilovo希洛沃() 	feud.Barony
    BOstrov奥斯特罗夫() 	feud.Barony
    BPorkhov波尔霍夫() 	feud.Barony
    BPskov普斯科夫() 	feud.Barony
}

type 普斯科夫PskovCounty struct {
	feud.BaseCounty
	杰多维奇Dedovichi 	feud.Barony
	戈卢博沃Golubovo 	feud.Barony
	戈里斯托Goristo 	feud.Barony
	希洛沃Khilovo 	feud.Barony
	奥斯特罗夫Ostrov 	feud.Barony
	波尔霍夫Porkhov 	feud.Barony
	普斯科夫Pskov 	feud.Barony
}

func (c *普斯科夫PskovCounty) BDedovichi杰多维奇() feud.Barony {
	return c.杰多维奇Dedovichi
}
    
func (c *普斯科夫PskovCounty) BGolubovo戈卢博沃() feud.Barony {
	return c.戈卢博沃Golubovo
}
    
func (c *普斯科夫PskovCounty) BGoristo戈里斯托() feud.Barony {
	return c.戈里斯托Goristo
}
    
func (c *普斯科夫PskovCounty) BKhilovo希洛沃() feud.Barony {
	return c.希洛沃Khilovo
}
    
func (c *普斯科夫PskovCounty) BOstrov奥斯特罗夫() feud.Barony {
	return c.奥斯特罗夫Ostrov
}
    
func (c *普斯科夫PskovCounty) BPorkhov波尔霍夫() feud.Barony {
	return c.波尔霍夫Porkhov
}
    
func (c *普斯科夫PskovCounty) BPskov普斯科夫() feud.Barony {
	return c.普斯科夫Pskov
}
    
var CPskov普斯科夫 PskovCounty = &普斯科夫PskovCounty{}

func init() {
	f := CPskov普斯科夫.(*普斯科夫PskovCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "413",
		Title:     "pskov",
		TitleName: "普斯科夫",
		TitleCode: "c_pskov",
		Baronies:  map[string]feud.Barony{},
	}

	f.杰多维奇Dedovichi = BDedovichi杰多维奇
	f.杰多维奇Dedovichi.SetParent(f)
	
	f.戈卢博沃Golubovo = BGolubovo戈卢博沃
	f.戈卢博沃Golubovo.SetParent(f)
	
	f.戈里斯托Goristo = BGoristo戈里斯托
	f.戈里斯托Goristo.SetParent(f)
	
	f.希洛沃Khilovo = BKhilovo希洛沃
	f.希洛沃Khilovo.SetParent(f)
	
	f.奥斯特罗夫Ostrov = BOstrov奥斯特罗夫
	f.奥斯特罗夫Ostrov.SetParent(f)
	
	f.波尔霍夫Porkhov = BPorkhov波尔霍夫
	f.波尔霍夫Porkhov.SetParent(f)
	
	f.普斯科夫Pskov = BPskov普斯科夫
	f.普斯科夫Pskov.SetParent(f)
	
}
