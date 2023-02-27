package frisia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type FrisiaCounty interface {
    feud.County
    BAssen阿森() 	feud.Barony
    BBolsward博尔斯瓦德() 	feud.Barony
    BDokkum多克姆() 	feud.Barony
    BFraneker弗拉讷克() 	feud.Barony
    BGroningen格罗宁根() 	feud.Barony
    BHarlingen哈灵根() 	feud.Barony
    BLeeuwarden吕伐登() 	feud.Barony
    BStavoren斯塔福伦() 	feud.Barony
}

type 弗里西亚FrisiaCounty struct {
	feud.BaseCounty
	阿森Assen 	feud.Barony
	博尔斯瓦德Bolsward 	feud.Barony
	多克姆Dokkum 	feud.Barony
	弗拉讷克Franeker 	feud.Barony
	格罗宁根Groningen 	feud.Barony
	哈灵根Harlingen 	feud.Barony
	吕伐登Leeuwarden 	feud.Barony
	斯塔福伦Stavoren 	feud.Barony
}

func (c *弗里西亚FrisiaCounty) BAssen阿森() feud.Barony {
	return c.阿森Assen
}
    
func (c *弗里西亚FrisiaCounty) BBolsward博尔斯瓦德() feud.Barony {
	return c.博尔斯瓦德Bolsward
}
    
func (c *弗里西亚FrisiaCounty) BDokkum多克姆() feud.Barony {
	return c.多克姆Dokkum
}
    
func (c *弗里西亚FrisiaCounty) BFraneker弗拉讷克() feud.Barony {
	return c.弗拉讷克Franeker
}
    
func (c *弗里西亚FrisiaCounty) BGroningen格罗宁根() feud.Barony {
	return c.格罗宁根Groningen
}
    
func (c *弗里西亚FrisiaCounty) BHarlingen哈灵根() feud.Barony {
	return c.哈灵根Harlingen
}
    
func (c *弗里西亚FrisiaCounty) BLeeuwarden吕伐登() feud.Barony {
	return c.吕伐登Leeuwarden
}
    
func (c *弗里西亚FrisiaCounty) BStavoren斯塔福伦() feud.Barony {
	return c.斯塔福伦Stavoren
}
    
var CFrisia弗里西亚 FrisiaCounty = &弗里西亚FrisiaCounty{}

func init() {
	f := CFrisia弗里西亚.(*弗里西亚FrisiaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "84",
		Title:     "frisia",
		TitleName: "弗里西亚",
		TitleCode: "c_frisia",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿森Assen = BAssen阿森
	f.阿森Assen.SetParent(f)
	
	f.博尔斯瓦德Bolsward = BBolsward博尔斯瓦德
	f.博尔斯瓦德Bolsward.SetParent(f)
	
	f.多克姆Dokkum = BDokkum多克姆
	f.多克姆Dokkum.SetParent(f)
	
	f.弗拉讷克Franeker = BFraneker弗拉讷克
	f.弗拉讷克Franeker.SetParent(f)
	
	f.格罗宁根Groningen = BGroningen格罗宁根
	f.格罗宁根Groningen.SetParent(f)
	
	f.哈灵根Harlingen = BHarlingen哈灵根
	f.哈灵根Harlingen.SetParent(f)
	
	f.吕伐登Leeuwarden = BLeeuwarden吕伐登
	f.吕伐登Leeuwarden.SetParent(f)
	
	f.斯塔福伦Stavoren = BStavoren斯塔福伦
	f.斯塔福伦Stavoren.SetParent(f)
	
}
