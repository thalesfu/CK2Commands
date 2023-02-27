package aksum

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AksumCounty interface {
    feud.County
    BAdigrat阿迪格拉特() 	feud.Barony
    BAdwa阿杜瓦() 	feud.Barony
    BAksum阿克苏姆() 	feud.Barony
    BHawzen豪赞() 	feud.Barony
    BMekele默克莱() 	feud.Barony
    BYeha耶哈() 	feud.Barony
}

type 阿克苏姆AksumCounty struct {
	feud.BaseCounty
	阿迪格拉特Adigrat 	feud.Barony
	阿杜瓦Adwa 	feud.Barony
	阿克苏姆Aksum 	feud.Barony
	豪赞Hawzen 	feud.Barony
	默克莱Mekele 	feud.Barony
	耶哈Yeha 	feud.Barony
}

func (c *阿克苏姆AksumCounty) BAdigrat阿迪格拉特() feud.Barony {
	return c.阿迪格拉特Adigrat
}
    
func (c *阿克苏姆AksumCounty) BAdwa阿杜瓦() feud.Barony {
	return c.阿杜瓦Adwa
}
    
func (c *阿克苏姆AksumCounty) BAksum阿克苏姆() feud.Barony {
	return c.阿克苏姆Aksum
}
    
func (c *阿克苏姆AksumCounty) BHawzen豪赞() feud.Barony {
	return c.豪赞Hawzen
}
    
func (c *阿克苏姆AksumCounty) BMekele默克莱() feud.Barony {
	return c.默克莱Mekele
}
    
func (c *阿克苏姆AksumCounty) BYeha耶哈() feud.Barony {
	return c.耶哈Yeha
}
    
var CAksum阿克苏姆 AksumCounty = &阿克苏姆AksumCounty{}

func init() {
	f := CAksum阿克苏姆.(*阿克苏姆AksumCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "875",
		Title:     "aksum",
		TitleName: "阿克苏姆",
		TitleCode: "c_aksum",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿迪格拉特Adigrat = BAdigrat阿迪格拉特
	f.阿迪格拉特Adigrat.SetParent(f)
	
	f.阿杜瓦Adwa = BAdwa阿杜瓦
	f.阿杜瓦Adwa.SetParent(f)
	
	f.阿克苏姆Aksum = BAksum阿克苏姆
	f.阿克苏姆Aksum.SetParent(f)
	
	f.豪赞Hawzen = BHawzen豪赞
	f.豪赞Hawzen.SetParent(f)
	
	f.默克莱Mekele = BMekele默克莱
	f.默克莱Mekele.SetParent(f)
	
	f.耶哈Yeha = BYeha耶哈
	f.耶哈Yeha.SetParent(f)
	
}
