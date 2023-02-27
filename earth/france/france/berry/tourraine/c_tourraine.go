package tourraine

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TourraineCounty interface {
    feud.County
    BAmboise安布瓦兹() 	feud.Barony
    BBeaulieu博略() 	feud.Barony
    BChinon希农() 	feud.Barony
    BFierbois菲尔布瓦() 	feud.Barony
    BLangeais朗热() 	feud.Barony
    BLoches洛什() 	feud.Barony
    BMontbazon蒙巴宗() 	feud.Barony
    BTours图尔() 	feud.Barony
}

type 图尔TourraineCounty struct {
	feud.BaseCounty
	安布瓦兹Amboise 	feud.Barony
	博略Beaulieu 	feud.Barony
	希农Chinon 	feud.Barony
	菲尔布瓦Fierbois 	feud.Barony
	朗热Langeais 	feud.Barony
	洛什Loches 	feud.Barony
	蒙巴宗Montbazon 	feud.Barony
	图尔Tours 	feud.Barony
}

func (c *图尔TourraineCounty) BAmboise安布瓦兹() feud.Barony {
	return c.安布瓦兹Amboise
}
    
func (c *图尔TourraineCounty) BBeaulieu博略() feud.Barony {
	return c.博略Beaulieu
}
    
func (c *图尔TourraineCounty) BChinon希农() feud.Barony {
	return c.希农Chinon
}
    
func (c *图尔TourraineCounty) BFierbois菲尔布瓦() feud.Barony {
	return c.菲尔布瓦Fierbois
}
    
func (c *图尔TourraineCounty) BLangeais朗热() feud.Barony {
	return c.朗热Langeais
}
    
func (c *图尔TourraineCounty) BLoches洛什() feud.Barony {
	return c.洛什Loches
}
    
func (c *图尔TourraineCounty) BMontbazon蒙巴宗() feud.Barony {
	return c.蒙巴宗Montbazon
}
    
func (c *图尔TourraineCounty) BTours图尔() feud.Barony {
	return c.图尔Tours
}
    
var CTourraine图尔 TourraineCounty = &图尔TourraineCounty{}

func init() {
	f := CTourraine图尔.(*图尔TourraineCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "140",
		Title:     "tourraine",
		TitleName: "图尔",
		TitleCode: "c_tourraine",
		Baronies:  map[string]feud.Barony{},
	}

	f.安布瓦兹Amboise = BAmboise安布瓦兹
	f.安布瓦兹Amboise.SetParent(f)
	
	f.博略Beaulieu = BBeaulieu博略
	f.博略Beaulieu.SetParent(f)
	
	f.希农Chinon = BChinon希农
	f.希农Chinon.SetParent(f)
	
	f.菲尔布瓦Fierbois = BFierbois菲尔布瓦
	f.菲尔布瓦Fierbois.SetParent(f)
	
	f.朗热Langeais = BLangeais朗热
	f.朗热Langeais.SetParent(f)
	
	f.洛什Loches = BLoches洛什
	f.洛什Loches.SetParent(f)
	
	f.蒙巴宗Montbazon = BMontbazon蒙巴宗
	f.蒙巴宗Montbazon.SetParent(f)
	
	f.图尔Tours = BTours图尔
	f.图尔Tours.SetParent(f)
	
}
