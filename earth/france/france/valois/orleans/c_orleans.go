package orleans

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type OrleansCounty interface {
    feud.County
    BFleury弗勒里() 	feud.Barony
    BJanville让维尔() 	feud.Barony
    BJargeau雅尔若() 	feud.Barony
    BLepuiset勒皮伊塞() 	feud.Barony
    BMeung默恩() 	feud.Barony
    BOrleans奥尔良() 	feud.Barony
    BSully叙利() 	feud.Barony
}

type 奥尔良OrleansCounty struct {
	feud.BaseCounty
	弗勒里Fleury 	feud.Barony
	让维尔Janville 	feud.Barony
	雅尔若Jargeau 	feud.Barony
	勒皮伊塞Lepuiset 	feud.Barony
	默恩Meung 	feud.Barony
	奥尔良Orleans 	feud.Barony
	叙利Sully 	feud.Barony
}

func (c *奥尔良OrleansCounty) BFleury弗勒里() feud.Barony {
	return c.弗勒里Fleury
}
    
func (c *奥尔良OrleansCounty) BJanville让维尔() feud.Barony {
	return c.让维尔Janville
}
    
func (c *奥尔良OrleansCounty) BJargeau雅尔若() feud.Barony {
	return c.雅尔若Jargeau
}
    
func (c *奥尔良OrleansCounty) BLepuiset勒皮伊塞() feud.Barony {
	return c.勒皮伊塞Lepuiset
}
    
func (c *奥尔良OrleansCounty) BMeung默恩() feud.Barony {
	return c.默恩Meung
}
    
func (c *奥尔良OrleansCounty) BOrleans奥尔良() feud.Barony {
	return c.奥尔良Orleans
}
    
func (c *奥尔良OrleansCounty) BSully叙利() feud.Barony {
	return c.叙利Sully
}
    
var COrleans奥尔良 OrleansCounty = &奥尔良OrleansCounty{}

func init() {
	f := COrleans奥尔良.(*奥尔良OrleansCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "138",
		Title:     "orleans",
		TitleName: "奥尔良",
		TitleCode: "c_orleans",
		Baronies:  map[string]feud.Barony{},
	}

	f.弗勒里Fleury = BFleury弗勒里
	f.弗勒里Fleury.SetParent(f)
	
	f.让维尔Janville = BJanville让维尔
	f.让维尔Janville.SetParent(f)
	
	f.雅尔若Jargeau = BJargeau雅尔若
	f.雅尔若Jargeau.SetParent(f)
	
	f.勒皮伊塞Lepuiset = BLepuiset勒皮伊塞
	f.勒皮伊塞Lepuiset.SetParent(f)
	
	f.默恩Meung = BMeung默恩
	f.默恩Meung.SetParent(f)
	
	f.奥尔良Orleans = BOrleans奥尔良
	f.奥尔良Orleans.SetParent(f)
	
	f.叙利Sully = BSully叙利
	f.叙利Sully.SetParent(f)
	
}
