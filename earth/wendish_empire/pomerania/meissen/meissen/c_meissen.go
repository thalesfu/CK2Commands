package meissen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MeissenCounty interface {
    feud.County
    BBelgern阿尔滕堡() 	feud.Barony
    BDohna多纳() 	feud.Barony
    BDresden德累斯顿() 	feud.Barony
    BMeissen迈森() 	feud.Barony
    BRabenau拉伯瑙() 	feud.Barony
    BRadeburg拉德堡() 	feud.Barony
    BStrehla施特雷拉() 	feud.Barony
    BWettin韦廷() 	feud.Barony
}

type 迈森MeissenCounty struct {
	feud.BaseCounty
	阿尔滕堡Belgern 	feud.Barony
	多纳Dohna 	feud.Barony
	德累斯顿Dresden 	feud.Barony
	迈森Meissen 	feud.Barony
	拉伯瑙Rabenau 	feud.Barony
	拉德堡Radeburg 	feud.Barony
	施特雷拉Strehla 	feud.Barony
	韦廷Wettin 	feud.Barony
}

func (c *迈森MeissenCounty) BBelgern阿尔滕堡() feud.Barony {
	return c.阿尔滕堡Belgern
}
    
func (c *迈森MeissenCounty) BDohna多纳() feud.Barony {
	return c.多纳Dohna
}
    
func (c *迈森MeissenCounty) BDresden德累斯顿() feud.Barony {
	return c.德累斯顿Dresden
}
    
func (c *迈森MeissenCounty) BMeissen迈森() feud.Barony {
	return c.迈森Meissen
}
    
func (c *迈森MeissenCounty) BRabenau拉伯瑙() feud.Barony {
	return c.拉伯瑙Rabenau
}
    
func (c *迈森MeissenCounty) BRadeburg拉德堡() feud.Barony {
	return c.拉德堡Radeburg
}
    
func (c *迈森MeissenCounty) BStrehla施特雷拉() feud.Barony {
	return c.施特雷拉Strehla
}
    
func (c *迈森MeissenCounty) BWettin韦廷() feud.Barony {
	return c.韦廷Wettin
}
    
var CMeissen迈森 MeissenCounty = &迈森MeissenCounty{}

func init() {
	f := CMeissen迈森.(*迈森MeissenCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "312",
		Title:     "meissen",
		TitleName: "迈森",
		TitleCode: "c_meissen",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿尔滕堡Belgern = BBelgern阿尔滕堡
	f.阿尔滕堡Belgern.SetParent(f)
	
	f.多纳Dohna = BDohna多纳
	f.多纳Dohna.SetParent(f)
	
	f.德累斯顿Dresden = BDresden德累斯顿
	f.德累斯顿Dresden.SetParent(f)
	
	f.迈森Meissen = BMeissen迈森
	f.迈森Meissen.SetParent(f)
	
	f.拉伯瑙Rabenau = BRabenau拉伯瑙
	f.拉伯瑙Rabenau.SetParent(f)
	
	f.拉德堡Radeburg = BRadeburg拉德堡
	f.拉德堡Radeburg.SetParent(f)
	
	f.施特雷拉Strehla = BStrehla施特雷拉
	f.施特雷拉Strehla.SetParent(f)
	
	f.韦廷Wettin = BWettin韦廷
	f.韦廷Wettin.SetParent(f)
	
}
