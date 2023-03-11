package saumur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SaumurCounty interface {
    feud.County
    BBeaupreau博普雷欧() 	feud.Barony
    BCholet绍莱() 	feud.Barony
    BDoue_la_fontaine杜埃拉丰坦() 	feud.Barony
    BFontevraud_l_abbaye丰特夫罗拉拜() 	feud.Barony
    BMaulevrier莫莱夫里耶() 	feud.Barony
    BSaumur索米尔() 	feud.Barony
    BVauchretien沃克雷蒂安() 	feud.Barony
}

type 索米尔SaumurCounty struct {
	feud.BaseCounty
	博普雷欧Beaupreau 	feud.Barony
	绍莱Cholet 	feud.Barony
	杜埃拉丰坦Doue_la_fontaine 	feud.Barony
	丰特夫罗拉拜Fontevraud_l_abbaye 	feud.Barony
	莫莱夫里耶Maulevrier 	feud.Barony
	索米尔Saumur 	feud.Barony
	沃克雷蒂安Vauchretien 	feud.Barony
}

func (c *索米尔SaumurCounty) BBeaupreau博普雷欧() feud.Barony {
	return c.博普雷欧Beaupreau
}
    
func (c *索米尔SaumurCounty) BCholet绍莱() feud.Barony {
	return c.绍莱Cholet
}
    
func (c *索米尔SaumurCounty) BDoue_la_fontaine杜埃拉丰坦() feud.Barony {
	return c.杜埃拉丰坦Doue_la_fontaine
}
    
func (c *索米尔SaumurCounty) BFontevraud_l_abbaye丰特夫罗拉拜() feud.Barony {
	return c.丰特夫罗拉拜Fontevraud_l_abbaye
}
    
func (c *索米尔SaumurCounty) BMaulevrier莫莱夫里耶() feud.Barony {
	return c.莫莱夫里耶Maulevrier
}
    
func (c *索米尔SaumurCounty) BSaumur索米尔() feud.Barony {
	return c.索米尔Saumur
}
    
func (c *索米尔SaumurCounty) BVauchretien沃克雷蒂安() feud.Barony {
	return c.沃克雷蒂安Vauchretien
}
    
var CSaumur索米尔 SaumurCounty = &索米尔SaumurCounty{}

func init() {
	f := CSaumur索米尔.(*索米尔SaumurCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1959",
		Title:     "saumur",
		TitleName: "索米尔",
		TitleCode: "c_saumur",
		Baronies:  map[string]feud.Barony{},
	}

	f.博普雷欧Beaupreau = BBeaupreau博普雷欧
	f.博普雷欧Beaupreau.SetParent(f)
	
	f.绍莱Cholet = BCholet绍莱
	f.绍莱Cholet.SetParent(f)
	
	f.杜埃拉丰坦Doue_la_fontaine = BDoue_la_fontaine杜埃拉丰坦
	f.杜埃拉丰坦Doue_la_fontaine.SetParent(f)
	
	f.丰特夫罗拉拜Fontevraud_l_abbaye = BFontevraud_l_abbaye丰特夫罗拉拜
	f.丰特夫罗拉拜Fontevraud_l_abbaye.SetParent(f)
	
	f.莫莱夫里耶Maulevrier = BMaulevrier莫莱夫里耶
	f.莫莱夫里耶Maulevrier.SetParent(f)
	
	f.索米尔Saumur = BSaumur索米尔
	f.索米尔Saumur.SetParent(f)
	
	f.沃克雷蒂安Vauchretien = BVauchretien沃克雷蒂安
	f.沃克雷蒂安Vauchretien.SetParent(f)
	
}
