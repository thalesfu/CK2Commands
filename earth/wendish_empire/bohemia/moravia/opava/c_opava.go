package opava

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type OpavaCounty interface {
    feud.County
    BBruntal布伦塔尔() 	feud.Barony
    BHlubcice格武布奇采() 	feud.Barony
    BHolasovice霍拉索维采() 	feud.Barony
    BHradec_nad_moravici摩拉维采河畔赫拉德茨() 	feud.Barony
    BKrnov克尔诺夫() 	feud.Barony
    BOldrisov奥尔德日绍夫() 	feud.Barony
    BOpava奥帕瓦() 	feud.Barony
}

type 奥帕瓦OpavaCounty struct {
	feud.BaseCounty
	布伦塔尔Bruntal 	feud.Barony
	格武布奇采Hlubcice 	feud.Barony
	霍拉索维采Holasovice 	feud.Barony
	摩拉维采河畔赫拉德茨Hradec_nad_moravici 	feud.Barony
	克尔诺夫Krnov 	feud.Barony
	奥尔德日绍夫Oldrisov 	feud.Barony
	奥帕瓦Opava 	feud.Barony
}

func (c *奥帕瓦OpavaCounty) BBruntal布伦塔尔() feud.Barony {
	return c.布伦塔尔Bruntal
}
    
func (c *奥帕瓦OpavaCounty) BHlubcice格武布奇采() feud.Barony {
	return c.格武布奇采Hlubcice
}
    
func (c *奥帕瓦OpavaCounty) BHolasovice霍拉索维采() feud.Barony {
	return c.霍拉索维采Holasovice
}
    
func (c *奥帕瓦OpavaCounty) BHradec_nad_moravici摩拉维采河畔赫拉德茨() feud.Barony {
	return c.摩拉维采河畔赫拉德茨Hradec_nad_moravici
}
    
func (c *奥帕瓦OpavaCounty) BKrnov克尔诺夫() feud.Barony {
	return c.克尔诺夫Krnov
}
    
func (c *奥帕瓦OpavaCounty) BOldrisov奥尔德日绍夫() feud.Barony {
	return c.奥尔德日绍夫Oldrisov
}
    
func (c *奥帕瓦OpavaCounty) BOpava奥帕瓦() feud.Barony {
	return c.奥帕瓦Opava
}
    
var COpava奥帕瓦 OpavaCounty = &奥帕瓦OpavaCounty{}

func init() {
	f := COpava奥帕瓦.(*奥帕瓦OpavaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1927",
		Title:     "opava",
		TitleName: "奥帕瓦",
		TitleCode: "c_opava",
		Baronies:  map[string]feud.Barony{},
	}

	f.布伦塔尔Bruntal = BBruntal布伦塔尔
	f.布伦塔尔Bruntal.SetParent(f)
	
	f.格武布奇采Hlubcice = BHlubcice格武布奇采
	f.格武布奇采Hlubcice.SetParent(f)
	
	f.霍拉索维采Holasovice = BHolasovice霍拉索维采
	f.霍拉索维采Holasovice.SetParent(f)
	
	f.摩拉维采河畔赫拉德茨Hradec_nad_moravici = BHradec_nad_moravici摩拉维采河畔赫拉德茨
	f.摩拉维采河畔赫拉德茨Hradec_nad_moravici.SetParent(f)
	
	f.克尔诺夫Krnov = BKrnov克尔诺夫
	f.克尔诺夫Krnov.SetParent(f)
	
	f.奥尔德日绍夫Oldrisov = BOldrisov奥尔德日绍夫
	f.奥尔德日绍夫Oldrisov.SetParent(f)
	
	f.奥帕瓦Opava = BOpava奥帕瓦
	f.奥帕瓦Opava.SetParent(f)
	
}
