package bezhetsky_verh

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type Bezhetsky_verhCounty interface {
    feud.County
    BKirovsk基罗夫斯克() 	feud.Barony
    BLadoga拉多加() 	feud.Barony
    BLyuban柳班() 	feud.Barony
    BNaziya纳济亚() 	feud.Barony
    BNovinka诺温卡() 	feud.Barony
    BTosno托斯诺() 	feud.Barony
    BVyritsa维里察() 	feud.Barony
}

type 拉多加Bezhetsky_verhCounty struct {
	feud.BaseCounty
	基罗夫斯克Kirovsk 	feud.Barony
	拉多加Ladoga 	feud.Barony
	柳班Lyuban 	feud.Barony
	纳济亚Naziya 	feud.Barony
	诺温卡Novinka 	feud.Barony
	托斯诺Tosno 	feud.Barony
	维里察Vyritsa 	feud.Barony
}

func (c *拉多加Bezhetsky_verhCounty) BKirovsk基罗夫斯克() feud.Barony {
	return c.基罗夫斯克Kirovsk
}
    
func (c *拉多加Bezhetsky_verhCounty) BLadoga拉多加() feud.Barony {
	return c.拉多加Ladoga
}
    
func (c *拉多加Bezhetsky_verhCounty) BLyuban柳班() feud.Barony {
	return c.柳班Lyuban
}
    
func (c *拉多加Bezhetsky_verhCounty) BNaziya纳济亚() feud.Barony {
	return c.纳济亚Naziya
}
    
func (c *拉多加Bezhetsky_verhCounty) BNovinka诺温卡() feud.Barony {
	return c.诺温卡Novinka
}
    
func (c *拉多加Bezhetsky_verhCounty) BTosno托斯诺() feud.Barony {
	return c.托斯诺Tosno
}
    
func (c *拉多加Bezhetsky_verhCounty) BVyritsa维里察() feud.Barony {
	return c.维里察Vyritsa
}
    
var CBezhetsky_verh拉多加 Bezhetsky_verhCounty = &拉多加Bezhetsky_verhCounty{}

func init() {
	f := CBezhetsky_verh拉多加.(*拉多加Bezhetsky_verhCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "409",
		Title:     "bezhetsky_verh",
		TitleName: "拉多加",
		TitleCode: "c_bezhetsky_verh",
		Baronies:  map[string]feud.Barony{},
	}

	f.基罗夫斯克Kirovsk = BKirovsk基罗夫斯克
	f.基罗夫斯克Kirovsk.SetParent(f)
	
	f.拉多加Ladoga = BLadoga拉多加
	f.拉多加Ladoga.SetParent(f)
	
	f.柳班Lyuban = BLyuban柳班
	f.柳班Lyuban.SetParent(f)
	
	f.纳济亚Naziya = BNaziya纳济亚
	f.纳济亚Naziya.SetParent(f)
	
	f.诺温卡Novinka = BNovinka诺温卡
	f.诺温卡Novinka.SetParent(f)
	
	f.托斯诺Tosno = BTosno托斯诺
	f.托斯诺Tosno.SetParent(f)
	
	f.维里察Vyritsa = BVyritsa维里察
	f.维里察Vyritsa.SetParent(f)
	
}
