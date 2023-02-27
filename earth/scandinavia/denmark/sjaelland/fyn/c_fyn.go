package fyn

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type FynCounty interface {
    feud.County
    BAssens阿森斯() 	feud.Barony
    BBogense博恩瑟() 	feud.Barony
    BFaaborg福堡() 	feud.Barony
    BKerteminde凯特明讷() 	feud.Barony
    BMiddelfart米泽尔法特() 	feud.Barony
    BNyborg尼堡() 	feud.Barony
    BOdense欧登塞() 	feud.Barony
    BSvendborg斯文堡() 	feud.Barony
}

type 菲英FynCounty struct {
	feud.BaseCounty
	阿森斯Assens 	feud.Barony
	博恩瑟Bogense 	feud.Barony
	福堡Faaborg 	feud.Barony
	凯特明讷Kerteminde 	feud.Barony
	米泽尔法特Middelfart 	feud.Barony
	尼堡Nyborg 	feud.Barony
	欧登塞Odense 	feud.Barony
	斯文堡Svendborg 	feud.Barony
}

func (c *菲英FynCounty) BAssens阿森斯() feud.Barony {
	return c.阿森斯Assens
}
    
func (c *菲英FynCounty) BBogense博恩瑟() feud.Barony {
	return c.博恩瑟Bogense
}
    
func (c *菲英FynCounty) BFaaborg福堡() feud.Barony {
	return c.福堡Faaborg
}
    
func (c *菲英FynCounty) BKerteminde凯特明讷() feud.Barony {
	return c.凯特明讷Kerteminde
}
    
func (c *菲英FynCounty) BMiddelfart米泽尔法特() feud.Barony {
	return c.米泽尔法特Middelfart
}
    
func (c *菲英FynCounty) BNyborg尼堡() feud.Barony {
	return c.尼堡Nyborg
}
    
func (c *菲英FynCounty) BOdense欧登塞() feud.Barony {
	return c.欧登塞Odense
}
    
func (c *菲英FynCounty) BSvendborg斯文堡() feud.Barony {
	return c.斯文堡Svendborg
}
    
var CFyn菲英 FynCounty = &菲英FynCounty{}

func init() {
	f := CFyn菲英.(*菲英FynCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "265",
		Title:     "fyn",
		TitleName: "菲英",
		TitleCode: "c_fyn",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿森斯Assens = BAssens阿森斯
	f.阿森斯Assens.SetParent(f)
	
	f.博恩瑟Bogense = BBogense博恩瑟
	f.博恩瑟Bogense.SetParent(f)
	
	f.福堡Faaborg = BFaaborg福堡
	f.福堡Faaborg.SetParent(f)
	
	f.凯特明讷Kerteminde = BKerteminde凯特明讷
	f.凯特明讷Kerteminde.SetParent(f)
	
	f.米泽尔法特Middelfart = BMiddelfart米泽尔法特
	f.米泽尔法特Middelfart.SetParent(f)
	
	f.尼堡Nyborg = BNyborg尼堡
	f.尼堡Nyborg.SetParent(f)
	
	f.欧登塞Odense = BOdense欧登塞
	f.欧登塞Odense.SetParent(f)
	
	f.斯文堡Svendborg = BSvendborg斯文堡
	f.斯文堡Svendborg.SetParent(f)
	
}
