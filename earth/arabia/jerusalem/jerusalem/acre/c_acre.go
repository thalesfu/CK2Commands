package acre

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AcreCounty interface {
    feud.County
    BAcre阿克() 	feud.Barony
    BAdelon阿德朗() 	feud.Barony
    BAthlith阿特利特() 	feud.Barony
    BHaifa海法() 	feud.Barony
    BManawat马纳瓦图() 	feud.Barony
    BMerle梅尔() 	feud.Barony
    BRecordana雷科丹那() 	feud.Barony
    BSyrcaesarea凯撒利亚() 	feud.Barony
}

type 阿克AcreCounty struct {
	feud.BaseCounty
	阿克Acre 	feud.Barony
	阿德朗Adelon 	feud.Barony
	阿特利特Athlith 	feud.Barony
	海法Haifa 	feud.Barony
	马纳瓦图Manawat 	feud.Barony
	梅尔Merle 	feud.Barony
	雷科丹那Recordana 	feud.Barony
	凯撒利亚Syrcaesarea 	feud.Barony
}

func (c *阿克AcreCounty) BAcre阿克() feud.Barony {
	return c.阿克Acre
}
    
func (c *阿克AcreCounty) BAdelon阿德朗() feud.Barony {
	return c.阿德朗Adelon
}
    
func (c *阿克AcreCounty) BAthlith阿特利特() feud.Barony {
	return c.阿特利特Athlith
}
    
func (c *阿克AcreCounty) BHaifa海法() feud.Barony {
	return c.海法Haifa
}
    
func (c *阿克AcreCounty) BManawat马纳瓦图() feud.Barony {
	return c.马纳瓦图Manawat
}
    
func (c *阿克AcreCounty) BMerle梅尔() feud.Barony {
	return c.梅尔Merle
}
    
func (c *阿克AcreCounty) BRecordana雷科丹那() feud.Barony {
	return c.雷科丹那Recordana
}
    
func (c *阿克AcreCounty) BSyrcaesarea凯撒利亚() feud.Barony {
	return c.凯撒利亚Syrcaesarea
}
    
var CAcre阿克 AcreCounty = &阿克AcreCounty{}

func init() {
	f := CAcre阿克.(*阿克AcreCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "772",
		Title:     "acre",
		TitleName: "阿克",
		TitleCode: "c_acre",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿克Acre = BAcre阿克
	f.阿克Acre.SetParent(f)
	
	f.阿德朗Adelon = BAdelon阿德朗
	f.阿德朗Adelon.SetParent(f)
	
	f.阿特利特Athlith = BAthlith阿特利特
	f.阿特利特Athlith.SetParent(f)
	
	f.海法Haifa = BHaifa海法
	f.海法Haifa.SetParent(f)
	
	f.马纳瓦图Manawat = BManawat马纳瓦图
	f.马纳瓦图Manawat.SetParent(f)
	
	f.梅尔Merle = BMerle梅尔
	f.梅尔Merle.SetParent(f)
	
	f.雷科丹那Recordana = BRecordana雷科丹那
	f.雷科丹那Recordana.SetParent(f)
	
	f.凯撒利亚Syrcaesarea = BSyrcaesarea凯撒利亚
	f.凯撒利亚Syrcaesarea.SetParent(f)
	
}
