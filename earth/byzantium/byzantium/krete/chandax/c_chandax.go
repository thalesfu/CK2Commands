package chandax

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ChandaxCounty interface {
    feud.County
    BAgiosnikolaos圣尼古劳斯() 	feud.Barony
    BIerapetra耶拉佩特拉() 	feud.Barony
    BIraklio伊拉克利翁() 	feud.Barony
    BKastelli卡斯泰利() 	feud.Barony
    BKnossos克诺索斯() 	feud.Barony
    BLassithi拉西锡() 	feud.Barony
    BMalia玛利亚() 	feud.Barony
    BSitia锡蒂亚() 	feud.Barony
}

type 汉达克斯ChandaxCounty struct {
	feud.BaseCounty
	圣尼古劳斯Agiosnikolaos 	feud.Barony
	耶拉佩特拉Ierapetra 	feud.Barony
	伊拉克利翁Iraklio 	feud.Barony
	卡斯泰利Kastelli 	feud.Barony
	克诺索斯Knossos 	feud.Barony
	拉西锡Lassithi 	feud.Barony
	玛利亚Malia 	feud.Barony
	锡蒂亚Sitia 	feud.Barony
}

func (c *汉达克斯ChandaxCounty) BAgiosnikolaos圣尼古劳斯() feud.Barony {
	return c.圣尼古劳斯Agiosnikolaos
}
    
func (c *汉达克斯ChandaxCounty) BIerapetra耶拉佩特拉() feud.Barony {
	return c.耶拉佩特拉Ierapetra
}
    
func (c *汉达克斯ChandaxCounty) BIraklio伊拉克利翁() feud.Barony {
	return c.伊拉克利翁Iraklio
}
    
func (c *汉达克斯ChandaxCounty) BKastelli卡斯泰利() feud.Barony {
	return c.卡斯泰利Kastelli
}
    
func (c *汉达克斯ChandaxCounty) BKnossos克诺索斯() feud.Barony {
	return c.克诺索斯Knossos
}
    
func (c *汉达克斯ChandaxCounty) BLassithi拉西锡() feud.Barony {
	return c.拉西锡Lassithi
}
    
func (c *汉达克斯ChandaxCounty) BMalia玛利亚() feud.Barony {
	return c.玛利亚Malia
}
    
func (c *汉达克斯ChandaxCounty) BSitia锡蒂亚() feud.Barony {
	return c.锡蒂亚Sitia
}
    
var CChandax汉达克斯 ChandaxCounty = &汉达克斯ChandaxCounty{}

func init() {
	f := CChandax汉达克斯.(*汉达克斯ChandaxCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "480",
		Title:     "chandax",
		TitleName: "汉达克斯",
		TitleCode: "c_chandax",
		Baronies:  map[string]feud.Barony{},
	}

	f.圣尼古劳斯Agiosnikolaos = BAgiosnikolaos圣尼古劳斯
	f.圣尼古劳斯Agiosnikolaos.SetParent(f)
	
	f.耶拉佩特拉Ierapetra = BIerapetra耶拉佩特拉
	f.耶拉佩特拉Ierapetra.SetParent(f)
	
	f.伊拉克利翁Iraklio = BIraklio伊拉克利翁
	f.伊拉克利翁Iraklio.SetParent(f)
	
	f.卡斯泰利Kastelli = BKastelli卡斯泰利
	f.卡斯泰利Kastelli.SetParent(f)
	
	f.克诺索斯Knossos = BKnossos克诺索斯
	f.克诺索斯Knossos.SetParent(f)
	
	f.拉西锡Lassithi = BLassithi拉西锡
	f.拉西锡Lassithi.SetParent(f)
	
	f.玛利亚Malia = BMalia玛利亚
	f.玛利亚Malia.SetParent(f)
	
	f.锡蒂亚Sitia = BSitia锡蒂亚
	f.锡蒂亚Sitia.SetParent(f)
	
}
