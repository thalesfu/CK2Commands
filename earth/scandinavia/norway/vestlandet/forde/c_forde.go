package forde

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type FordeCounty interface {
    feud.County
    BAudunborg艾于敦堡() 	feud.Barony
    BBygstad比格斯塔() 	feud.Barony
    BDale达勒() 	feud.Barony
    BFlora弗洛拉() 	feud.Barony
    BForde弗勒() 	feud.Barony
    BKinn欣() 	feud.Barony
    BLuster吕斯特() 	feud.Barony
}

type 弗勒FordeCounty struct {
	feud.BaseCounty
	艾于敦堡Audunborg 	feud.Barony
	比格斯塔Bygstad 	feud.Barony
	达勒Dale 	feud.Barony
	弗洛拉Flora 	feud.Barony
	弗勒Forde 	feud.Barony
	欣Kinn 	feud.Barony
	吕斯特Luster 	feud.Barony
}

func (c *弗勒FordeCounty) BAudunborg艾于敦堡() feud.Barony {
	return c.艾于敦堡Audunborg
}
    
func (c *弗勒FordeCounty) BBygstad比格斯塔() feud.Barony {
	return c.比格斯塔Bygstad
}
    
func (c *弗勒FordeCounty) BDale达勒() feud.Barony {
	return c.达勒Dale
}
    
func (c *弗勒FordeCounty) BFlora弗洛拉() feud.Barony {
	return c.弗洛拉Flora
}
    
func (c *弗勒FordeCounty) BForde弗勒() feud.Barony {
	return c.弗勒Forde
}
    
func (c *弗勒FordeCounty) BKinn欣() feud.Barony {
	return c.欣Kinn
}
    
func (c *弗勒FordeCounty) BLuster吕斯特() feud.Barony {
	return c.吕斯特Luster
}
    
var CForde弗勒 FordeCounty = &弗勒FordeCounty{}

func init() {
	f := CForde弗勒.(*弗勒FordeCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1624",
		Title:     "forde",
		TitleName: "弗勒",
		TitleCode: "c_forde",
		Baronies:  map[string]feud.Barony{},
	}

	f.艾于敦堡Audunborg = BAudunborg艾于敦堡
	f.艾于敦堡Audunborg.SetParent(f)
	
	f.比格斯塔Bygstad = BBygstad比格斯塔
	f.比格斯塔Bygstad.SetParent(f)
	
	f.达勒Dale = BDale达勒
	f.达勒Dale.SetParent(f)
	
	f.弗洛拉Flora = BFlora弗洛拉
	f.弗洛拉Flora.SetParent(f)
	
	f.弗勒Forde = BForde弗勒
	f.弗勒Forde.SetParent(f)
	
	f.欣Kinn = BKinn欣
	f.欣Kinn.SetParent(f)
	
	f.吕斯特Luster = BLuster吕斯特
	f.吕斯特Luster.SetParent(f)
	
}
