package korosten

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KorostenCounty interface {
    feud.County
    BChernobyl切尔诺贝利() 	feud.Barony
    BKorosten科罗斯坚() 	feud.Barony
    BLouhyny卢吉尼() 	feud.Barony
    BMalyn马林() 	feud.Barony
    BNarodytchi纳罗季奇() 	feud.Barony
    BOlevsk奥列夫斯克() 	feud.Barony
    BOvruch奥夫鲁奇() 	feud.Barony
}

type 科罗斯坚KorostenCounty struct {
	feud.BaseCounty
	切尔诺贝利Chernobyl 	feud.Barony
	科罗斯坚Korosten 	feud.Barony
	卢吉尼Louhyny 	feud.Barony
	马林Malyn 	feud.Barony
	纳罗季奇Narodytchi 	feud.Barony
	奥列夫斯克Olevsk 	feud.Barony
	奥夫鲁奇Ovruch 	feud.Barony
}

func (c *科罗斯坚KorostenCounty) BChernobyl切尔诺贝利() feud.Barony {
	return c.切尔诺贝利Chernobyl
}
    
func (c *科罗斯坚KorostenCounty) BKorosten科罗斯坚() feud.Barony {
	return c.科罗斯坚Korosten
}
    
func (c *科罗斯坚KorostenCounty) BLouhyny卢吉尼() feud.Barony {
	return c.卢吉尼Louhyny
}
    
func (c *科罗斯坚KorostenCounty) BMalyn马林() feud.Barony {
	return c.马林Malyn
}
    
func (c *科罗斯坚KorostenCounty) BNarodytchi纳罗季奇() feud.Barony {
	return c.纳罗季奇Narodytchi
}
    
func (c *科罗斯坚KorostenCounty) BOlevsk奥列夫斯克() feud.Barony {
	return c.奥列夫斯克Olevsk
}
    
func (c *科罗斯坚KorostenCounty) BOvruch奥夫鲁奇() feud.Barony {
	return c.奥夫鲁奇Ovruch
}
    
var CKorosten科罗斯坚 KorostenCounty = &科罗斯坚KorostenCounty{}

func init() {
	f := CKorosten科罗斯坚.(*科罗斯坚KorostenCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1649",
		Title:     "korosten",
		TitleName: "科罗斯坚",
		TitleCode: "c_korosten",
		Baronies:  map[string]feud.Barony{},
	}

	f.切尔诺贝利Chernobyl = BChernobyl切尔诺贝利
	f.切尔诺贝利Chernobyl.SetParent(f)
	
	f.科罗斯坚Korosten = BKorosten科罗斯坚
	f.科罗斯坚Korosten.SetParent(f)
	
	f.卢吉尼Louhyny = BLouhyny卢吉尼
	f.卢吉尼Louhyny.SetParent(f)
	
	f.马林Malyn = BMalyn马林
	f.马林Malyn.SetParent(f)
	
	f.纳罗季奇Narodytchi = BNarodytchi纳罗季奇
	f.纳罗季奇Narodytchi.SetParent(f)
	
	f.奥列夫斯克Olevsk = BOlevsk奥列夫斯克
	f.奥列夫斯克Olevsk.SetParent(f)
	
	f.奥夫鲁奇Ovruch = BOvruch奥夫鲁奇
	f.奥夫鲁奇Ovruch.SetParent(f)
	
}
