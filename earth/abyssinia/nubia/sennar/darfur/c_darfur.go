package darfur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type DarfurCounty interface {
    feud.County
    BAd_dukaydik杜凯迪克() 	feud.Barony
    BAin_farah艾因法拉赫() 	feud.Barony
    BEl_fasher法希尔() 	feud.Barony
    BKalkal卡尔卡尔() 	feud.Barony
    BMasalat马萨拉特() 	feud.Barony
    BUm_marrih乌姆马利() 	feud.Barony
    BUri乌里() 	feud.Barony
}

type 达尔富尔DarfurCounty struct {
	feud.BaseCounty
	杜凯迪克Ad_dukaydik 	feud.Barony
	艾因法拉赫Ain_farah 	feud.Barony
	法希尔El_fasher 	feud.Barony
	卡尔卡尔Kalkal 	feud.Barony
	马萨拉特Masalat 	feud.Barony
	乌姆马利Um_marrih 	feud.Barony
	乌里Uri 	feud.Barony
}

func (c *达尔富尔DarfurCounty) BAd_dukaydik杜凯迪克() feud.Barony {
	return c.杜凯迪克Ad_dukaydik
}
    
func (c *达尔富尔DarfurCounty) BAin_farah艾因法拉赫() feud.Barony {
	return c.艾因法拉赫Ain_farah
}
    
func (c *达尔富尔DarfurCounty) BEl_fasher法希尔() feud.Barony {
	return c.法希尔El_fasher
}
    
func (c *达尔富尔DarfurCounty) BKalkal卡尔卡尔() feud.Barony {
	return c.卡尔卡尔Kalkal
}
    
func (c *达尔富尔DarfurCounty) BMasalat马萨拉特() feud.Barony {
	return c.马萨拉特Masalat
}
    
func (c *达尔富尔DarfurCounty) BUm_marrih乌姆马利() feud.Barony {
	return c.乌姆马利Um_marrih
}
    
func (c *达尔富尔DarfurCounty) BUri乌里() feud.Barony {
	return c.乌里Uri
}
    
var CDarfur达尔富尔 DarfurCounty = &达尔富尔DarfurCounty{}

func init() {
	f := CDarfur达尔富尔.(*达尔富尔DarfurCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1772",
		Title:     "darfur",
		TitleName: "达尔富尔",
		TitleCode: "c_darfur",
		Baronies:  map[string]feud.Barony{},
	}

	f.杜凯迪克Ad_dukaydik = BAd_dukaydik杜凯迪克
	f.杜凯迪克Ad_dukaydik.SetParent(f)
	
	f.艾因法拉赫Ain_farah = BAin_farah艾因法拉赫
	f.艾因法拉赫Ain_farah.SetParent(f)
	
	f.法希尔El_fasher = BEl_fasher法希尔
	f.法希尔El_fasher.SetParent(f)
	
	f.卡尔卡尔Kalkal = BKalkal卡尔卡尔
	f.卡尔卡尔Kalkal.SetParent(f)
	
	f.马萨拉特Masalat = BMasalat马萨拉特
	f.马萨拉特Masalat.SetParent(f)
	
	f.乌姆马利Um_marrih = BUm_marrih乌姆马利
	f.乌姆马利Um_marrih.SetParent(f)
	
	f.乌里Uri = BUri乌里
	f.乌里Uri.SetParent(f)
	
}
