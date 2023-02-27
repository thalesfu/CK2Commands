package hull

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type HullCounty interface {
    feud.County
    BBeverley贝弗利() 	feud.Barony
    BBridlington布里德灵顿() 	feud.Barony
    BGuildhall吉尔德霍尔() 	feud.Barony
    BHornsea霍恩西() 	feud.Barony
    BHull赫尔() 	feud.Barony
    BStamford_bridge斯坦福桥() 	feud.Barony
    BWyke怀克() 	feud.Barony
}

type 赫尔HullCounty struct {
	feud.BaseCounty
	贝弗利Beverley 	feud.Barony
	布里德灵顿Bridlington 	feud.Barony
	吉尔德霍尔Guildhall 	feud.Barony
	霍恩西Hornsea 	feud.Barony
	赫尔Hull 	feud.Barony
	斯坦福桥Stamford_bridge 	feud.Barony
	怀克Wyke 	feud.Barony
}

func (c *赫尔HullCounty) BBeverley贝弗利() feud.Barony {
	return c.贝弗利Beverley
}
    
func (c *赫尔HullCounty) BBridlington布里德灵顿() feud.Barony {
	return c.布里德灵顿Bridlington
}
    
func (c *赫尔HullCounty) BGuildhall吉尔德霍尔() feud.Barony {
	return c.吉尔德霍尔Guildhall
}
    
func (c *赫尔HullCounty) BHornsea霍恩西() feud.Barony {
	return c.霍恩西Hornsea
}
    
func (c *赫尔HullCounty) BHull赫尔() feud.Barony {
	return c.赫尔Hull
}
    
func (c *赫尔HullCounty) BStamford_bridge斯坦福桥() feud.Barony {
	return c.斯坦福桥Stamford_bridge
}
    
func (c *赫尔HullCounty) BWyke怀克() feud.Barony {
	return c.怀克Wyke
}
    
var CHull赫尔 HullCounty = &赫尔HullCounty{}

func init() {
	f := CHull赫尔.(*赫尔HullCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1941",
		Title:     "hull",
		TitleName: "赫尔",
		TitleCode: "c_hull",
		Baronies:  map[string]feud.Barony{},
	}

	f.贝弗利Beverley = BBeverley贝弗利
	f.贝弗利Beverley.SetParent(f)
	
	f.布里德灵顿Bridlington = BBridlington布里德灵顿
	f.布里德灵顿Bridlington.SetParent(f)
	
	f.吉尔德霍尔Guildhall = BGuildhall吉尔德霍尔
	f.吉尔德霍尔Guildhall.SetParent(f)
	
	f.霍恩西Hornsea = BHornsea霍恩西
	f.霍恩西Hornsea.SetParent(f)
	
	f.赫尔Hull = BHull赫尔
	f.赫尔Hull.SetParent(f)
	
	f.斯坦福桥Stamford_bridge = BStamford_bridge斯坦福桥
	f.斯坦福桥Stamford_bridge.SetParent(f)
	
	f.怀克Wyke = BWyke怀克
	f.怀克Wyke.SetParent(f)
	
}
