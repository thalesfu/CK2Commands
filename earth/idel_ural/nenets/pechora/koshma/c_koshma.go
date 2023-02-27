package koshma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KoshmaCounty interface {
    feud.County
    BHurai胡赖() 	feud.Barony
    BIlner伊尔涅尔() 	feud.Barony
    BKeshar克沙尔() 	feud.Barony
    BKoshma科沙姆() 	feud.Barony
    BListma利斯特马() 	feud.Barony
    BOurenter奥乌连捷尔() 	feud.Barony
    BPlurel普卢列尔() 	feud.Barony
}

type 科什马KoshmaCounty struct {
	feud.BaseCounty
	胡赖Hurai 	feud.Barony
	伊尔涅尔Ilner 	feud.Barony
	克沙尔Keshar 	feud.Barony
	科沙姆Koshma 	feud.Barony
	利斯特马Listma 	feud.Barony
	奥乌连捷尔Ourenter 	feud.Barony
	普卢列尔Plurel 	feud.Barony
}

func (c *科什马KoshmaCounty) BHurai胡赖() feud.Barony {
	return c.胡赖Hurai
}
    
func (c *科什马KoshmaCounty) BIlner伊尔涅尔() feud.Barony {
	return c.伊尔涅尔Ilner
}
    
func (c *科什马KoshmaCounty) BKeshar克沙尔() feud.Barony {
	return c.克沙尔Keshar
}
    
func (c *科什马KoshmaCounty) BKoshma科沙姆() feud.Barony {
	return c.科沙姆Koshma
}
    
func (c *科什马KoshmaCounty) BListma利斯特马() feud.Barony {
	return c.利斯特马Listma
}
    
func (c *科什马KoshmaCounty) BOurenter奥乌连捷尔() feud.Barony {
	return c.奥乌连捷尔Ourenter
}
    
func (c *科什马KoshmaCounty) BPlurel普卢列尔() feud.Barony {
	return c.普卢列尔Plurel
}
    
var CKoshma科什马 KoshmaCounty = &科什马KoshmaCounty{}

func init() {
	f := CKoshma科什马.(*科什马KoshmaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1831",
		Title:     "koshma",
		TitleName: "科什马",
		TitleCode: "c_koshma",
		Baronies:  map[string]feud.Barony{},
	}

	f.胡赖Hurai = BHurai胡赖
	f.胡赖Hurai.SetParent(f)
	
	f.伊尔涅尔Ilner = BIlner伊尔涅尔
	f.伊尔涅尔Ilner.SetParent(f)
	
	f.克沙尔Keshar = BKeshar克沙尔
	f.克沙尔Keshar.SetParent(f)
	
	f.科沙姆Koshma = BKoshma科沙姆
	f.科沙姆Koshma.SetParent(f)
	
	f.利斯特马Listma = BListma利斯特马
	f.利斯特马Listma.SetParent(f)
	
	f.奥乌连捷尔Ourenter = BOurenter奥乌连捷尔
	f.奥乌连捷尔Ourenter.SetParent(f)
	
	f.普卢列尔Plurel = BPlurel普卢列尔
	f.普卢列尔Plurel.SetParent(f)
	
}
