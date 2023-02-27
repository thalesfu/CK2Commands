package kurmanchal

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KurmanchalCounty interface {
    feud.County
    BAskot_kurmanchal阿尸拘吒() 	feud.Barony
    BChampawat瞻波婆多() 	feud.Barony
    BDidihat迪迪哈特() 	feud.Barony
    BGovishana瞿毗霜那() 	feud.Barony
    BKartikeyapura迦絺吉夜补罗() 	feud.Barony
    BKatarmal迦吒罗末罗() 	feud.Barony
    BRanikhet罗耳契吒() 	feud.Barony
}

type 俱利曼遮罗KurmanchalCounty struct {
	feud.BaseCounty
	阿尸拘吒Askot_kurmanchal 	feud.Barony
	瞻波婆多Champawat 	feud.Barony
	迪迪哈特Didihat 	feud.Barony
	瞿毗霜那Govishana 	feud.Barony
	迦絺吉夜补罗Kartikeyapura 	feud.Barony
	迦吒罗末罗Katarmal 	feud.Barony
	罗耳契吒Ranikhet 	feud.Barony
}

func (c *俱利曼遮罗KurmanchalCounty) BAskot_kurmanchal阿尸拘吒() feud.Barony {
	return c.阿尸拘吒Askot_kurmanchal
}
    
func (c *俱利曼遮罗KurmanchalCounty) BChampawat瞻波婆多() feud.Barony {
	return c.瞻波婆多Champawat
}
    
func (c *俱利曼遮罗KurmanchalCounty) BDidihat迪迪哈特() feud.Barony {
	return c.迪迪哈特Didihat
}
    
func (c *俱利曼遮罗KurmanchalCounty) BGovishana瞿毗霜那() feud.Barony {
	return c.瞿毗霜那Govishana
}
    
func (c *俱利曼遮罗KurmanchalCounty) BKartikeyapura迦絺吉夜补罗() feud.Barony {
	return c.迦絺吉夜补罗Kartikeyapura
}
    
func (c *俱利曼遮罗KurmanchalCounty) BKatarmal迦吒罗末罗() feud.Barony {
	return c.迦吒罗末罗Katarmal
}
    
func (c *俱利曼遮罗KurmanchalCounty) BRanikhet罗耳契吒() feud.Barony {
	return c.罗耳契吒Ranikhet
}
    
var CKurmanchal俱利曼遮罗 KurmanchalCounty = &俱利曼遮罗KurmanchalCounty{}

func init() {
	f := CKurmanchal俱利曼遮罗.(*俱利曼遮罗KurmanchalCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1471",
		Title:     "kurmanchal",
		TitleName: "俱利曼遮罗",
		TitleCode: "c_kurmanchal",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿尸拘吒Askot_kurmanchal = BAskot_kurmanchal阿尸拘吒
	f.阿尸拘吒Askot_kurmanchal.SetParent(f)
	
	f.瞻波婆多Champawat = BChampawat瞻波婆多
	f.瞻波婆多Champawat.SetParent(f)
	
	f.迪迪哈特Didihat = BDidihat迪迪哈特
	f.迪迪哈特Didihat.SetParent(f)
	
	f.瞿毗霜那Govishana = BGovishana瞿毗霜那
	f.瞿毗霜那Govishana.SetParent(f)
	
	f.迦絺吉夜补罗Kartikeyapura = BKartikeyapura迦絺吉夜补罗
	f.迦絺吉夜补罗Kartikeyapura.SetParent(f)
	
	f.迦吒罗末罗Katarmal = BKatarmal迦吒罗末罗
	f.迦吒罗末罗Katarmal.SetParent(f)
	
	f.罗耳契吒Ranikhet = BRanikhet罗耳契吒
	f.罗耳契吒Ranikhet.SetParent(f)
	
}
