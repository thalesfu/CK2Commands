package gelre

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type GelreCounty interface {
    feud.County
    BArnhem安亨() 	feud.Barony
    BBergh贝尔赫() 	feud.Barony
    BBronckhorst布隆克霍斯特() 	feud.Barony
    BDuetinghem杜廷赫姆() 	feud.Barony
    BLohn勒讷() 	feud.Barony
    BNijmegen奈梅亨() 	feud.Barony
    BZutphen聚特芬() 	feud.Barony
}

type 海尔雷GelreCounty struct {
	feud.BaseCounty
	安亨Arnhem 	feud.Barony
	贝尔赫Bergh 	feud.Barony
	布隆克霍斯特Bronckhorst 	feud.Barony
	杜廷赫姆Duetinghem 	feud.Barony
	勒讷Lohn 	feud.Barony
	奈梅亨Nijmegen 	feud.Barony
	聚特芬Zutphen 	feud.Barony
}

func (c *海尔雷GelreCounty) BArnhem安亨() feud.Barony {
	return c.安亨Arnhem
}
    
func (c *海尔雷GelreCounty) BBergh贝尔赫() feud.Barony {
	return c.贝尔赫Bergh
}
    
func (c *海尔雷GelreCounty) BBronckhorst布隆克霍斯特() feud.Barony {
	return c.布隆克霍斯特Bronckhorst
}
    
func (c *海尔雷GelreCounty) BDuetinghem杜廷赫姆() feud.Barony {
	return c.杜廷赫姆Duetinghem
}
    
func (c *海尔雷GelreCounty) BLohn勒讷() feud.Barony {
	return c.勒讷Lohn
}
    
func (c *海尔雷GelreCounty) BNijmegen奈梅亨() feud.Barony {
	return c.奈梅亨Nijmegen
}
    
func (c *海尔雷GelreCounty) BZutphen聚特芬() feud.Barony {
	return c.聚特芬Zutphen
}
    
var CGelre海尔雷 GelreCounty = &海尔雷GelreCounty{}

func init() {
	f := CGelre海尔雷.(*海尔雷GelreCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "83",
		Title:     "gelre",
		TitleName: "海尔雷",
		TitleCode: "c_gelre",
		Baronies:  map[string]feud.Barony{},
	}

	f.安亨Arnhem = BArnhem安亨
	f.安亨Arnhem.SetParent(f)
	
	f.贝尔赫Bergh = BBergh贝尔赫
	f.贝尔赫Bergh.SetParent(f)
	
	f.布隆克霍斯特Bronckhorst = BBronckhorst布隆克霍斯特
	f.布隆克霍斯特Bronckhorst.SetParent(f)
	
	f.杜廷赫姆Duetinghem = BDuetinghem杜廷赫姆
	f.杜廷赫姆Duetinghem.SetParent(f)
	
	f.勒讷Lohn = BLohn勒讷
	f.勒讷Lohn.SetParent(f)
	
	f.奈梅亨Nijmegen = BNijmegen奈梅亨
	f.奈梅亨Nijmegen.SetParent(f)
	
	f.聚特芬Zutphen = BZutphen聚特芬
	f.聚特芬Zutphen.SetParent(f)
	
}
