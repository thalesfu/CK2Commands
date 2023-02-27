package gent

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type GentCounty interface {
    feud.County
    BAalst阿尔斯特() 	feud.Barony
    BDendermonde登德尔蒙德() 	feud.Barony
    BDoornik多尔尼克() 	feud.Barony
    BGent根特() 	feud.Barony
    BGeraardsbergen赫拉尔兹贝亨() 	feud.Barony
    BKortrijk科特赖克() 	feud.Barony
    BOudenaarde奥德纳尔德() 	feud.Barony
    BSt_niklaas圣尼克拉斯() 	feud.Barony
}

type 根特GentCounty struct {
	feud.BaseCounty
	阿尔斯特Aalst 	feud.Barony
	登德尔蒙德Dendermonde 	feud.Barony
	多尔尼克Doornik 	feud.Barony
	根特Gent 	feud.Barony
	赫拉尔兹贝亨Geraardsbergen 	feud.Barony
	科特赖克Kortrijk 	feud.Barony
	奥德纳尔德Oudenaarde 	feud.Barony
	圣尼克拉斯St_niklaas 	feud.Barony
}

func (c *根特GentCounty) BAalst阿尔斯特() feud.Barony {
	return c.阿尔斯特Aalst
}
    
func (c *根特GentCounty) BDendermonde登德尔蒙德() feud.Barony {
	return c.登德尔蒙德Dendermonde
}
    
func (c *根特GentCounty) BDoornik多尔尼克() feud.Barony {
	return c.多尔尼克Doornik
}
    
func (c *根特GentCounty) BGent根特() feud.Barony {
	return c.根特Gent
}
    
func (c *根特GentCounty) BGeraardsbergen赫拉尔兹贝亨() feud.Barony {
	return c.赫拉尔兹贝亨Geraardsbergen
}
    
func (c *根特GentCounty) BKortrijk科特赖克() feud.Barony {
	return c.科特赖克Kortrijk
}
    
func (c *根特GentCounty) BOudenaarde奥德纳尔德() feud.Barony {
	return c.奥德纳尔德Oudenaarde
}
    
func (c *根特GentCounty) BSt_niklaas圣尼克拉斯() feud.Barony {
	return c.圣尼克拉斯St_niklaas
}
    
var CGent根特 GentCounty = &根特GentCounty{}

func init() {
	f := CGent根特.(*根特GentCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "93",
		Title:     "gent",
		TitleName: "根特",
		TitleCode: "c_gent",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿尔斯特Aalst = BAalst阿尔斯特
	f.阿尔斯特Aalst.SetParent(f)
	
	f.登德尔蒙德Dendermonde = BDendermonde登德尔蒙德
	f.登德尔蒙德Dendermonde.SetParent(f)
	
	f.多尔尼克Doornik = BDoornik多尔尼克
	f.多尔尼克Doornik.SetParent(f)
	
	f.根特Gent = BGent根特
	f.根特Gent.SetParent(f)
	
	f.赫拉尔兹贝亨Geraardsbergen = BGeraardsbergen赫拉尔兹贝亨
	f.赫拉尔兹贝亨Geraardsbergen.SetParent(f)
	
	f.科特赖克Kortrijk = BKortrijk科特赖克
	f.科特赖克Kortrijk.SetParent(f)
	
	f.奥德纳尔德Oudenaarde = BOudenaarde奥德纳尔德
	f.奥德纳尔德Oudenaarde.SetParent(f)
	
	f.圣尼克拉斯St_niklaas = BSt_niklaas圣尼克拉斯
	f.圣尼克拉斯St_niklaas.SetParent(f)
	
}
