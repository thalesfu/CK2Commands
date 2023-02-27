package koln

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KolnCounty interface {
    feud.County
    BBonn波恩() 	feud.Barony
    BBrauweiler布劳韦勒() 	feud.Barony
    BDietz迪茨() 	feud.Barony
    BHochstaden霍赫施塔登() 	feud.Barony
    BKoln科隆() 	feud.Barony
    BMark马克() 	feud.Barony
    BSaffenburg萨芬堡() 	feud.Barony
}

type 科隆KolnCounty struct {
	feud.BaseCounty
	波恩Bonn 	feud.Barony
	布劳韦勒Brauweiler 	feud.Barony
	迪茨Dietz 	feud.Barony
	霍赫施塔登Hochstaden 	feud.Barony
	科隆Koln 	feud.Barony
	马克Mark 	feud.Barony
	萨芬堡Saffenburg 	feud.Barony
}

func (c *科隆KolnCounty) BBonn波恩() feud.Barony {
	return c.波恩Bonn
}
    
func (c *科隆KolnCounty) BBrauweiler布劳韦勒() feud.Barony {
	return c.布劳韦勒Brauweiler
}
    
func (c *科隆KolnCounty) BDietz迪茨() feud.Barony {
	return c.迪茨Dietz
}
    
func (c *科隆KolnCounty) BHochstaden霍赫施塔登() feud.Barony {
	return c.霍赫施塔登Hochstaden
}
    
func (c *科隆KolnCounty) BKoln科隆() feud.Barony {
	return c.科隆Koln
}
    
func (c *科隆KolnCounty) BMark马克() feud.Barony {
	return c.马克Mark
}
    
func (c *科隆KolnCounty) BSaffenburg萨芬堡() feud.Barony {
	return c.萨芬堡Saffenburg
}
    
var CKoln科隆 KolnCounty = &科隆KolnCounty{}

func init() {
	f := CKoln科隆.(*科隆KolnCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "119",
		Title:     "koln",
		TitleName: "科隆",
		TitleCode: "c_koln",
		Baronies:  map[string]feud.Barony{},
	}

	f.波恩Bonn = BBonn波恩
	f.波恩Bonn.SetParent(f)
	
	f.布劳韦勒Brauweiler = BBrauweiler布劳韦勒
	f.布劳韦勒Brauweiler.SetParent(f)
	
	f.迪茨Dietz = BDietz迪茨
	f.迪茨Dietz.SetParent(f)
	
	f.霍赫施塔登Hochstaden = BHochstaden霍赫施塔登
	f.霍赫施塔登Hochstaden.SetParent(f)
	
	f.科隆Koln = BKoln科隆
	f.科隆Koln.SetParent(f)
	
	f.马克Mark = BMark马克
	f.马克Mark.SetParent(f)
	
	f.萨芬堡Saffenburg = BSaffenburg萨芬堡
	f.萨芬堡Saffenburg.SetParent(f)
	
}
