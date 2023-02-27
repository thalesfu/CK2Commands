package chach

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ChachCounty interface {
    feud.County
    BChach赭石() 	feud.Barony
    BIsbijab白水城() 	feud.Barony
    BNavekat新城() 	feud.Barony
    BPskent别斯干() 	feud.Barony
    BSayram塞蓝() 	feud.Barony
    BShymkent赤麦干() 	feud.Barony
    BTurbat土尔巴特() 	feud.Barony
}

type 赭石ChachCounty struct {
	feud.BaseCounty
	赭石Chach 	feud.Barony
	白水城Isbijab 	feud.Barony
	新城Navekat 	feud.Barony
	别斯干Pskent 	feud.Barony
	塞蓝Sayram 	feud.Barony
	赤麦干Shymkent 	feud.Barony
	土尔巴特Turbat 	feud.Barony
}

func (c *赭石ChachCounty) BChach赭石() feud.Barony {
	return c.赭石Chach
}
    
func (c *赭石ChachCounty) BIsbijab白水城() feud.Barony {
	return c.白水城Isbijab
}
    
func (c *赭石ChachCounty) BNavekat新城() feud.Barony {
	return c.新城Navekat
}
    
func (c *赭石ChachCounty) BPskent别斯干() feud.Barony {
	return c.别斯干Pskent
}
    
func (c *赭石ChachCounty) BSayram塞蓝() feud.Barony {
	return c.塞蓝Sayram
}
    
func (c *赭石ChachCounty) BShymkent赤麦干() feud.Barony {
	return c.赤麦干Shymkent
}
    
func (c *赭石ChachCounty) BTurbat土尔巴特() feud.Barony {
	return c.土尔巴特Turbat
}
    
var CChach赭石 ChachCounty = &赭石ChachCounty{}

func init() {
	f := CChach赭石.(*赭石ChachCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1186",
		Title:     "chach",
		TitleName: "赭石",
		TitleCode: "c_chach",
		Baronies:  map[string]feud.Barony{},
	}

	f.赭石Chach = BChach赭石
	f.赭石Chach.SetParent(f)
	
	f.白水城Isbijab = BIsbijab白水城
	f.白水城Isbijab.SetParent(f)
	
	f.新城Navekat = BNavekat新城
	f.新城Navekat.SetParent(f)
	
	f.别斯干Pskent = BPskent别斯干
	f.别斯干Pskent.SetParent(f)
	
	f.塞蓝Sayram = BSayram塞蓝
	f.塞蓝Sayram.SetParent(f)
	
	f.赤麦干Shymkent = BShymkent赤麦干
	f.赤麦干Shymkent.SetParent(f)
	
	f.土尔巴特Turbat = BTurbat土尔巴特
	f.土尔巴特Turbat.SetParent(f)
	
}
