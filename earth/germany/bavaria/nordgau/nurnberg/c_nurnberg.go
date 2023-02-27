package nurnberg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type NurnbergCounty interface {
    feud.County
    BEllwangen埃尔旺根() 	feud.Barony
    BErlangen埃尔朗根() 	feud.Barony
    BFurth菲尔特() 	feud.Barony
    BHohenburg霍恩堡() 	feud.Barony
    BKulmbach库尔姆巴赫() 	feud.Barony
    BNordlingen讷德林根() 	feud.Barony
    BNurnberg纽伦堡() 	feud.Barony
}

type 纽伦堡NurnbergCounty struct {
	feud.BaseCounty
	埃尔旺根Ellwangen 	feud.Barony
	埃尔朗根Erlangen 	feud.Barony
	菲尔特Furth 	feud.Barony
	霍恩堡Hohenburg 	feud.Barony
	库尔姆巴赫Kulmbach 	feud.Barony
	讷德林根Nordlingen 	feud.Barony
	纽伦堡Nurnberg 	feud.Barony
}

func (c *纽伦堡NurnbergCounty) BEllwangen埃尔旺根() feud.Barony {
	return c.埃尔旺根Ellwangen
}
    
func (c *纽伦堡NurnbergCounty) BErlangen埃尔朗根() feud.Barony {
	return c.埃尔朗根Erlangen
}
    
func (c *纽伦堡NurnbergCounty) BFurth菲尔特() feud.Barony {
	return c.菲尔特Furth
}
    
func (c *纽伦堡NurnbergCounty) BHohenburg霍恩堡() feud.Barony {
	return c.霍恩堡Hohenburg
}
    
func (c *纽伦堡NurnbergCounty) BKulmbach库尔姆巴赫() feud.Barony {
	return c.库尔姆巴赫Kulmbach
}
    
func (c *纽伦堡NurnbergCounty) BNordlingen讷德林根() feud.Barony {
	return c.讷德林根Nordlingen
}
    
func (c *纽伦堡NurnbergCounty) BNurnberg纽伦堡() feud.Barony {
	return c.纽伦堡Nurnberg
}
    
var CNurnberg纽伦堡 NurnbergCounty = &纽伦堡NurnbergCounty{}

func init() {
	f := CNurnberg纽伦堡.(*纽伦堡NurnbergCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "314",
		Title:     "nurnberg",
		TitleName: "纽伦堡",
		TitleCode: "c_nurnberg",
		Baronies:  map[string]feud.Barony{},
	}

	f.埃尔旺根Ellwangen = BEllwangen埃尔旺根
	f.埃尔旺根Ellwangen.SetParent(f)
	
	f.埃尔朗根Erlangen = BErlangen埃尔朗根
	f.埃尔朗根Erlangen.SetParent(f)
	
	f.菲尔特Furth = BFurth菲尔特
	f.菲尔特Furth.SetParent(f)
	
	f.霍恩堡Hohenburg = BHohenburg霍恩堡
	f.霍恩堡Hohenburg.SetParent(f)
	
	f.库尔姆巴赫Kulmbach = BKulmbach库尔姆巴赫
	f.库尔姆巴赫Kulmbach.SetParent(f)
	
	f.讷德林根Nordlingen = BNordlingen讷德林根
	f.讷德林根Nordlingen.SetParent(f)
	
	f.纽伦堡Nurnberg = BNurnberg纽伦堡
	f.纽伦堡Nurnberg.SetParent(f)
	
}
