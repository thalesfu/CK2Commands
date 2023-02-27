package sandomierskie

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SandomierskieCounty interface {
    feud.County
    BKielce凯尔采() 	feud.Barony
    BOlescnica伊乌扎() 	feud.Barony
    BOpatow奥帕图夫() 	feud.Barony
    BOpoczno奥波奇诺() 	feud.Barony
    BRadom拉多姆() 	feud.Barony
    BSandomierz桑多梅日() 	feud.Barony
    BWislica维希利察() 	feud.Barony
}

type 桑多梅日SandomierskieCounty struct {
	feud.BaseCounty
	凯尔采Kielce 	feud.Barony
	伊乌扎Olescnica 	feud.Barony
	奥帕图夫Opatow 	feud.Barony
	奥波奇诺Opoczno 	feud.Barony
	拉多姆Radom 	feud.Barony
	桑多梅日Sandomierz 	feud.Barony
	维希利察Wislica 	feud.Barony
}

func (c *桑多梅日SandomierskieCounty) BKielce凯尔采() feud.Barony {
	return c.凯尔采Kielce
}
    
func (c *桑多梅日SandomierskieCounty) BOlescnica伊乌扎() feud.Barony {
	return c.伊乌扎Olescnica
}
    
func (c *桑多梅日SandomierskieCounty) BOpatow奥帕图夫() feud.Barony {
	return c.奥帕图夫Opatow
}
    
func (c *桑多梅日SandomierskieCounty) BOpoczno奥波奇诺() feud.Barony {
	return c.奥波奇诺Opoczno
}
    
func (c *桑多梅日SandomierskieCounty) BRadom拉多姆() feud.Barony {
	return c.拉多姆Radom
}
    
func (c *桑多梅日SandomierskieCounty) BSandomierz桑多梅日() feud.Barony {
	return c.桑多梅日Sandomierz
}
    
func (c *桑多梅日SandomierskieCounty) BWislica维希利察() feud.Barony {
	return c.维希利察Wislica
}
    
var CSandomierskie桑多梅日 SandomierskieCounty = &桑多梅日SandomierskieCounty{}

func init() {
	f := CSandomierskie桑多梅日.(*桑多梅日SandomierskieCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "531",
		Title:     "sandomierskie",
		TitleName: "桑多梅日",
		TitleCode: "c_sandomierskie",
		Baronies:  map[string]feud.Barony{},
	}

	f.凯尔采Kielce = BKielce凯尔采
	f.凯尔采Kielce.SetParent(f)
	
	f.伊乌扎Olescnica = BOlescnica伊乌扎
	f.伊乌扎Olescnica.SetParent(f)
	
	f.奥帕图夫Opatow = BOpatow奥帕图夫
	f.奥帕图夫Opatow.SetParent(f)
	
	f.奥波奇诺Opoczno = BOpoczno奥波奇诺
	f.奥波奇诺Opoczno.SetParent(f)
	
	f.拉多姆Radom = BRadom拉多姆
	f.拉多姆Radom.SetParent(f)
	
	f.桑多梅日Sandomierz = BSandomierz桑多梅日
	f.桑多梅日Sandomierz.SetParent(f)
	
	f.维希利察Wislica = BWislica维希利察
	f.维希利察Wislica.SetParent(f)
	
}
