package gevaudan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type GevaudanCounty interface {
    feud.County
    BApchier阿普谢() 	feud.Barony
    BFlorac弗洛拉克() 	feud.Barony
    BGrezes格雷泽() 	feud.Barony
    BLepuy勒皮() 	feud.Barony
    BMarvejols马尔沃若勒() 	feud.Barony
    BMende芒德() 	feud.Barony
    BStsauveur圣索弗尔() 	feud.Barony
    BTournel图尔内尔() 	feud.Barony
}

type 热沃当GevaudanCounty struct {
	feud.BaseCounty
	阿普谢Apchier 	feud.Barony
	弗洛拉克Florac 	feud.Barony
	格雷泽Grezes 	feud.Barony
	勒皮Lepuy 	feud.Barony
	马尔沃若勒Marvejols 	feud.Barony
	芒德Mende 	feud.Barony
	圣索弗尔Stsauveur 	feud.Barony
	图尔内尔Tournel 	feud.Barony
}

func (c *热沃当GevaudanCounty) BApchier阿普谢() feud.Barony {
	return c.阿普谢Apchier
}
    
func (c *热沃当GevaudanCounty) BFlorac弗洛拉克() feud.Barony {
	return c.弗洛拉克Florac
}
    
func (c *热沃当GevaudanCounty) BGrezes格雷泽() feud.Barony {
	return c.格雷泽Grezes
}
    
func (c *热沃当GevaudanCounty) BLepuy勒皮() feud.Barony {
	return c.勒皮Lepuy
}
    
func (c *热沃当GevaudanCounty) BMarvejols马尔沃若勒() feud.Barony {
	return c.马尔沃若勒Marvejols
}
    
func (c *热沃当GevaudanCounty) BMende芒德() feud.Barony {
	return c.芒德Mende
}
    
func (c *热沃当GevaudanCounty) BStsauveur圣索弗尔() feud.Barony {
	return c.圣索弗尔Stsauveur
}
    
func (c *热沃当GevaudanCounty) BTournel图尔内尔() feud.Barony {
	return c.图尔内尔Tournel
}
    
var CGevaudan热沃当 GevaudanCounty = &热沃当GevaudanCounty{}

func init() {
	f := CGevaudan热沃当.(*热沃当GevaudanCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "219",
		Title:     "gevaudan",
		TitleName: "热沃当",
		TitleCode: "c_gevaudan",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿普谢Apchier = BApchier阿普谢
	f.阿普谢Apchier.SetParent(f)
	
	f.弗洛拉克Florac = BFlorac弗洛拉克
	f.弗洛拉克Florac.SetParent(f)
	
	f.格雷泽Grezes = BGrezes格雷泽
	f.格雷泽Grezes.SetParent(f)
	
	f.勒皮Lepuy = BLepuy勒皮
	f.勒皮Lepuy.SetParent(f)
	
	f.马尔沃若勒Marvejols = BMarvejols马尔沃若勒
	f.马尔沃若勒Marvejols.SetParent(f)
	
	f.芒德Mende = BMende芒德
	f.芒德Mende.SetParent(f)
	
	f.圣索弗尔Stsauveur = BStsauveur圣索弗尔
	f.圣索弗尔Stsauveur.SetParent(f)
	
	f.图尔内尔Tournel = BTournel图尔内尔
	f.图尔内尔Tournel.SetParent(f)
	
}
