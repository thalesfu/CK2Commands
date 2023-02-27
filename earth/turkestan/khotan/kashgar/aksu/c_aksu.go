package aksu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AksuCounty interface {
    feud.County
    BAksu_mongolia阿克苏() 	feud.Barony
    BAral阿拉尔() 	feud.Barony
    BAwat阿瓦提() 	feud.Barony
    BKelpin柯坪() 	feud.Barony
    BLaofutian老付田() 	feud.Barony
    BTumshuk图木舒克() 	feud.Barony
    BWensu温宿() 	feud.Barony
}

type 阿克苏AksuCounty struct {
	feud.BaseCounty
	阿克苏Aksu_mongolia 	feud.Barony
	阿拉尔Aral 	feud.Barony
	阿瓦提Awat 	feud.Barony
	柯坪Kelpin 	feud.Barony
	老付田Laofutian 	feud.Barony
	图木舒克Tumshuk 	feud.Barony
	温宿Wensu 	feud.Barony
}

func (c *阿克苏AksuCounty) BAksu_mongolia阿克苏() feud.Barony {
	return c.阿克苏Aksu_mongolia
}
    
func (c *阿克苏AksuCounty) BAral阿拉尔() feud.Barony {
	return c.阿拉尔Aral
}
    
func (c *阿克苏AksuCounty) BAwat阿瓦提() feud.Barony {
	return c.阿瓦提Awat
}
    
func (c *阿克苏AksuCounty) BKelpin柯坪() feud.Barony {
	return c.柯坪Kelpin
}
    
func (c *阿克苏AksuCounty) BLaofutian老付田() feud.Barony {
	return c.老付田Laofutian
}
    
func (c *阿克苏AksuCounty) BTumshuk图木舒克() feud.Barony {
	return c.图木舒克Tumshuk
}
    
func (c *阿克苏AksuCounty) BWensu温宿() feud.Barony {
	return c.温宿Wensu
}
    
var CAksu阿克苏 AksuCounty = &阿克苏AksuCounty{}

func init() {
	f := CAksu阿克苏.(*阿克苏AksuCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1445",
		Title:     "aksu",
		TitleName: "阿克苏",
		TitleCode: "c_aksu",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿克苏Aksu_mongolia = BAksu_mongolia阿克苏
	f.阿克苏Aksu_mongolia.SetParent(f)
	
	f.阿拉尔Aral = BAral阿拉尔
	f.阿拉尔Aral.SetParent(f)
	
	f.阿瓦提Awat = BAwat阿瓦提
	f.阿瓦提Awat.SetParent(f)
	
	f.柯坪Kelpin = BKelpin柯坪
	f.柯坪Kelpin.SetParent(f)
	
	f.老付田Laofutian = BLaofutian老付田
	f.老付田Laofutian.SetParent(f)
	
	f.图木舒克Tumshuk = BTumshuk图木舒克
	f.图木舒克Tumshuk.SetParent(f)
	
	f.温宿Wensu = BWensu温宿
	f.温宿Wensu.SetParent(f)
	
}
