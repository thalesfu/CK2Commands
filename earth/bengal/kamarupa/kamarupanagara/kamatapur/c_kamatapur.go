package kamatapur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KamatapurCounty interface {
    feud.County
    BAvadiyarkoil阿婆地耶罗寺() 	feud.Barony
    BBabnergarh乏婆涅姞利呬() 	feud.Barony
    BBhitagarh毗多姞利呬() 	feud.Barony
    BJapyesvara阇比曳湿伐罗() 	feud.Barony
    BKamatapur迦摩多城() 	feud.Barony
    BMainaguri梅那古利() 	feud.Barony
    BNalrajar_garh那罗耶罗堡() 	feud.Barony
}

type 迦摩多城KamatapurCounty struct {
	feud.BaseCounty
	阿婆地耶罗寺Avadiyarkoil 	feud.Barony
	乏婆涅姞利呬Babnergarh 	feud.Barony
	毗多姞利呬Bhitagarh 	feud.Barony
	阇比曳湿伐罗Japyesvara 	feud.Barony
	迦摩多城Kamatapur 	feud.Barony
	梅那古利Mainaguri 	feud.Barony
	那罗耶罗堡Nalrajar_garh 	feud.Barony
}

func (c *迦摩多城KamatapurCounty) BAvadiyarkoil阿婆地耶罗寺() feud.Barony {
	return c.阿婆地耶罗寺Avadiyarkoil
}
    
func (c *迦摩多城KamatapurCounty) BBabnergarh乏婆涅姞利呬() feud.Barony {
	return c.乏婆涅姞利呬Babnergarh
}
    
func (c *迦摩多城KamatapurCounty) BBhitagarh毗多姞利呬() feud.Barony {
	return c.毗多姞利呬Bhitagarh
}
    
func (c *迦摩多城KamatapurCounty) BJapyesvara阇比曳湿伐罗() feud.Barony {
	return c.阇比曳湿伐罗Japyesvara
}
    
func (c *迦摩多城KamatapurCounty) BKamatapur迦摩多城() feud.Barony {
	return c.迦摩多城Kamatapur
}
    
func (c *迦摩多城KamatapurCounty) BMainaguri梅那古利() feud.Barony {
	return c.梅那古利Mainaguri
}
    
func (c *迦摩多城KamatapurCounty) BNalrajar_garh那罗耶罗堡() feud.Barony {
	return c.那罗耶罗堡Nalrajar_garh
}
    
var CKamatapur迦摩多城 KamatapurCounty = &迦摩多城KamatapurCounty{}

func init() {
	f := CKamatapur迦摩多城.(*迦摩多城KamatapurCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1244",
		Title:     "kamatapur",
		TitleName: "迦摩多城",
		TitleCode: "c_kamatapur",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿婆地耶罗寺Avadiyarkoil = BAvadiyarkoil阿婆地耶罗寺
	f.阿婆地耶罗寺Avadiyarkoil.SetParent(f)
	
	f.乏婆涅姞利呬Babnergarh = BBabnergarh乏婆涅姞利呬
	f.乏婆涅姞利呬Babnergarh.SetParent(f)
	
	f.毗多姞利呬Bhitagarh = BBhitagarh毗多姞利呬
	f.毗多姞利呬Bhitagarh.SetParent(f)
	
	f.阇比曳湿伐罗Japyesvara = BJapyesvara阇比曳湿伐罗
	f.阇比曳湿伐罗Japyesvara.SetParent(f)
	
	f.迦摩多城Kamatapur = BKamatapur迦摩多城
	f.迦摩多城Kamatapur.SetParent(f)
	
	f.梅那古利Mainaguri = BMainaguri梅那古利
	f.梅那古利Mainaguri.SetParent(f)
	
	f.那罗耶罗堡Nalrajar_garh = BNalrajar_garh那罗耶罗堡
	f.那罗耶罗堡Nalrajar_garh.SetParent(f)
	
}
