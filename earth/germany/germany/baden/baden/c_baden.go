package baden

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BadenCounty interface {
    feud.County
    BBaden巴登() 	feud.Barony
    BCalw卡尔夫() 	feud.Barony
    BDurlach杜尔拉赫() 	feud.Barony
    BKarlsruhe卡尔斯鲁厄() 	feud.Barony
    BNeuhausen诺伊豪森() 	feud.Barony
    BPforzheim普福尔茨海姆() 	feud.Barony
    BRastatt拉施塔特() 	feud.Barony
    BWimpfen温普芬() 	feud.Barony
}

type 巴登BadenCounty struct {
	feud.BaseCounty
	巴登Baden 	feud.Barony
	卡尔夫Calw 	feud.Barony
	杜尔拉赫Durlach 	feud.Barony
	卡尔斯鲁厄Karlsruhe 	feud.Barony
	诺伊豪森Neuhausen 	feud.Barony
	普福尔茨海姆Pforzheim 	feud.Barony
	拉施塔特Rastatt 	feud.Barony
	温普芬Wimpfen 	feud.Barony
}

func (c *巴登BadenCounty) BBaden巴登() feud.Barony {
	return c.巴登Baden
}
    
func (c *巴登BadenCounty) BCalw卡尔夫() feud.Barony {
	return c.卡尔夫Calw
}
    
func (c *巴登BadenCounty) BDurlach杜尔拉赫() feud.Barony {
	return c.杜尔拉赫Durlach
}
    
func (c *巴登BadenCounty) BKarlsruhe卡尔斯鲁厄() feud.Barony {
	return c.卡尔斯鲁厄Karlsruhe
}
    
func (c *巴登BadenCounty) BNeuhausen诺伊豪森() feud.Barony {
	return c.诺伊豪森Neuhausen
}
    
func (c *巴登BadenCounty) BPforzheim普福尔茨海姆() feud.Barony {
	return c.普福尔茨海姆Pforzheim
}
    
func (c *巴登BadenCounty) BRastatt拉施塔特() feud.Barony {
	return c.拉施塔特Rastatt
}
    
func (c *巴登BadenCounty) BWimpfen温普芬() feud.Barony {
	return c.温普芬Wimpfen
}
    
var CBaden巴登 BadenCounty = &巴登BadenCounty{}

func init() {
	f := CBaden巴登.(*巴登BadenCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "125",
		Title:     "baden",
		TitleName: "巴登",
		TitleCode: "c_baden",
		Baronies:  map[string]feud.Barony{},
	}

	f.巴登Baden = BBaden巴登
	f.巴登Baden.SetParent(f)
	
	f.卡尔夫Calw = BCalw卡尔夫
	f.卡尔夫Calw.SetParent(f)
	
	f.杜尔拉赫Durlach = BDurlach杜尔拉赫
	f.杜尔拉赫Durlach.SetParent(f)
	
	f.卡尔斯鲁厄Karlsruhe = BKarlsruhe卡尔斯鲁厄
	f.卡尔斯鲁厄Karlsruhe.SetParent(f)
	
	f.诺伊豪森Neuhausen = BNeuhausen诺伊豪森
	f.诺伊豪森Neuhausen.SetParent(f)
	
	f.普福尔茨海姆Pforzheim = BPforzheim普福尔茨海姆
	f.普福尔茨海姆Pforzheim.SetParent(f)
	
	f.拉施塔特Rastatt = BRastatt拉施塔特
	f.拉施塔特Rastatt.SetParent(f)
	
	f.温普芬Wimpfen = BWimpfen温普芬
	f.温普芬Wimpfen.SetParent(f)
	
}
