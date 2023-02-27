package iarmond

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type IarmondCounty interface {
    feud.County
    BArdfert阿德弗特() 	feud.Barony
    BCahersiveen凯尔西温() 	feud.Barony
    BInnisfallen因尼斯福伦() 	feud.Barony
    BKillarney基拉尼() 	feud.Barony
    BRoss罗斯() 	feud.Barony
    BSkellig斯凯利格() 	feud.Barony
    BTralee特拉利() 	feud.Barony
}

type 雅尔蒙德IarmondCounty struct {
	feud.BaseCounty
	阿德弗特Ardfert 	feud.Barony
	凯尔西温Cahersiveen 	feud.Barony
	因尼斯福伦Innisfallen 	feud.Barony
	基拉尼Killarney 	feud.Barony
	罗斯Ross 	feud.Barony
	斯凯利格Skellig 	feud.Barony
	特拉利Tralee 	feud.Barony
}

func (c *雅尔蒙德IarmondCounty) BArdfert阿德弗特() feud.Barony {
	return c.阿德弗特Ardfert
}
    
func (c *雅尔蒙德IarmondCounty) BCahersiveen凯尔西温() feud.Barony {
	return c.凯尔西温Cahersiveen
}
    
func (c *雅尔蒙德IarmondCounty) BInnisfallen因尼斯福伦() feud.Barony {
	return c.因尼斯福伦Innisfallen
}
    
func (c *雅尔蒙德IarmondCounty) BKillarney基拉尼() feud.Barony {
	return c.基拉尼Killarney
}
    
func (c *雅尔蒙德IarmondCounty) BRoss罗斯() feud.Barony {
	return c.罗斯Ross
}
    
func (c *雅尔蒙德IarmondCounty) BSkellig斯凯利格() feud.Barony {
	return c.斯凯利格Skellig
}
    
func (c *雅尔蒙德IarmondCounty) BTralee特拉利() feud.Barony {
	return c.特拉利Tralee
}
    
var CIarmond雅尔蒙德 IarmondCounty = &雅尔蒙德IarmondCounty{}

func init() {
	f := CIarmond雅尔蒙德.(*雅尔蒙德IarmondCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1938",
		Title:     "iarmond",
		TitleName: "雅尔蒙德",
		TitleCode: "c_iarmond",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿德弗特Ardfert = BArdfert阿德弗特
	f.阿德弗特Ardfert.SetParent(f)
	
	f.凯尔西温Cahersiveen = BCahersiveen凯尔西温
	f.凯尔西温Cahersiveen.SetParent(f)
	
	f.因尼斯福伦Innisfallen = BInnisfallen因尼斯福伦
	f.因尼斯福伦Innisfallen.SetParent(f)
	
	f.基拉尼Killarney = BKillarney基拉尼
	f.基拉尼Killarney.SetParent(f)
	
	f.罗斯Ross = BRoss罗斯
	f.罗斯Ross.SetParent(f)
	
	f.斯凯利格Skellig = BSkellig斯凯利格
	f.斯凯利格Skellig.SetParent(f)
	
	f.特拉利Tralee = BTralee特拉利
	f.特拉利Tralee.SetParent(f)
	
}
