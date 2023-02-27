package dakhina_desa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type Dakhina_desaCounty interface {
    feud.County
    BHatthigiripura诃悉帝耆厘补罗() 	feud.Barony
    BJambudoni旬波多尼() 	feud.Barony
    BKotte屈胝() 	feud.Barony
    BMahiyangana摩喜扬伽那() 	feud.Barony
    BRayigama罗依伽摩() 	feud.Barony
    BSamantakuta三曼多矩吒() 	feud.Barony
    BVapinagara婆比那揭罗() 	feud.Barony
}

type 南域Dakhina_desaCounty struct {
	feud.BaseCounty
	诃悉帝耆厘补罗Hatthigiripura 	feud.Barony
	旬波多尼Jambudoni 	feud.Barony
	屈胝Kotte 	feud.Barony
	摩喜扬伽那Mahiyangana 	feud.Barony
	罗依伽摩Rayigama 	feud.Barony
	三曼多矩吒Samantakuta 	feud.Barony
	婆比那揭罗Vapinagara 	feud.Barony
}

func (c *南域Dakhina_desaCounty) BHatthigiripura诃悉帝耆厘补罗() feud.Barony {
	return c.诃悉帝耆厘补罗Hatthigiripura
}
    
func (c *南域Dakhina_desaCounty) BJambudoni旬波多尼() feud.Barony {
	return c.旬波多尼Jambudoni
}
    
func (c *南域Dakhina_desaCounty) BKotte屈胝() feud.Barony {
	return c.屈胝Kotte
}
    
func (c *南域Dakhina_desaCounty) BMahiyangana摩喜扬伽那() feud.Barony {
	return c.摩喜扬伽那Mahiyangana
}
    
func (c *南域Dakhina_desaCounty) BRayigama罗依伽摩() feud.Barony {
	return c.罗依伽摩Rayigama
}
    
func (c *南域Dakhina_desaCounty) BSamantakuta三曼多矩吒() feud.Barony {
	return c.三曼多矩吒Samantakuta
}
    
func (c *南域Dakhina_desaCounty) BVapinagara婆比那揭罗() feud.Barony {
	return c.婆比那揭罗Vapinagara
}
    
var CDakhina_desa南域 Dakhina_desaCounty = &南域Dakhina_desaCounty{}

func init() {
	f := CDakhina_desa南域.(*南域Dakhina_desaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1215",
		Title:     "dakhina_desa",
		TitleName: "南域",
		TitleCode: "c_dakhina_desa",
		Baronies:  map[string]feud.Barony{},
	}

	f.诃悉帝耆厘补罗Hatthigiripura = BHatthigiripura诃悉帝耆厘补罗
	f.诃悉帝耆厘补罗Hatthigiripura.SetParent(f)
	
	f.旬波多尼Jambudoni = BJambudoni旬波多尼
	f.旬波多尼Jambudoni.SetParent(f)
	
	f.屈胝Kotte = BKotte屈胝
	f.屈胝Kotte.SetParent(f)
	
	f.摩喜扬伽那Mahiyangana = BMahiyangana摩喜扬伽那
	f.摩喜扬伽那Mahiyangana.SetParent(f)
	
	f.罗依伽摩Rayigama = BRayigama罗依伽摩
	f.罗依伽摩Rayigama.SetParent(f)
	
	f.三曼多矩吒Samantakuta = BSamantakuta三曼多矩吒
	f.三曼多矩吒Samantakuta.SetParent(f)
	
	f.婆比那揭罗Vapinagara = BVapinagara婆比那揭罗
	f.婆比那揭罗Vapinagara.SetParent(f)
	
}
