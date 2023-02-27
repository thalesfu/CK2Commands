package zavarot

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ZavarotCounty interface {
    feud.County
    BAndeg安杰格() 	feud.Barony
    BKhongurey洪古列伊() 	feud.Barony
    BKuya库亚() 	feud.Barony
    BMakarovo马卡罗沃() 	feud.Barony
    BOksino奥克西诺() 	feud.Barony
    BPylemets佩列梅茨() 	feud.Barony
    BZavarot扎瓦罗特() 	feud.Barony
}

type 扎瓦罗特ZavarotCounty struct {
	feud.BaseCounty
	安杰格Andeg 	feud.Barony
	洪古列伊Khongurey 	feud.Barony
	库亚Kuya 	feud.Barony
	马卡罗沃Makarovo 	feud.Barony
	奥克西诺Oksino 	feud.Barony
	佩列梅茨Pylemets 	feud.Barony
	扎瓦罗特Zavarot 	feud.Barony
}

func (c *扎瓦罗特ZavarotCounty) BAndeg安杰格() feud.Barony {
	return c.安杰格Andeg
}
    
func (c *扎瓦罗特ZavarotCounty) BKhongurey洪古列伊() feud.Barony {
	return c.洪古列伊Khongurey
}
    
func (c *扎瓦罗特ZavarotCounty) BKuya库亚() feud.Barony {
	return c.库亚Kuya
}
    
func (c *扎瓦罗特ZavarotCounty) BMakarovo马卡罗沃() feud.Barony {
	return c.马卡罗沃Makarovo
}
    
func (c *扎瓦罗特ZavarotCounty) BOksino奥克西诺() feud.Barony {
	return c.奥克西诺Oksino
}
    
func (c *扎瓦罗特ZavarotCounty) BPylemets佩列梅茨() feud.Barony {
	return c.佩列梅茨Pylemets
}
    
func (c *扎瓦罗特ZavarotCounty) BZavarot扎瓦罗特() feud.Barony {
	return c.扎瓦罗特Zavarot
}
    
var CZavarot扎瓦罗特 ZavarotCounty = &扎瓦罗特ZavarotCounty{}

func init() {
	f := CZavarot扎瓦罗特.(*扎瓦罗特ZavarotCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1828",
		Title:     "zavarot",
		TitleName: "扎瓦罗特",
		TitleCode: "c_zavarot",
		Baronies:  map[string]feud.Barony{},
	}

	f.安杰格Andeg = BAndeg安杰格
	f.安杰格Andeg.SetParent(f)
	
	f.洪古列伊Khongurey = BKhongurey洪古列伊
	f.洪古列伊Khongurey.SetParent(f)
	
	f.库亚Kuya = BKuya库亚
	f.库亚Kuya.SetParent(f)
	
	f.马卡罗沃Makarovo = BMakarovo马卡罗沃
	f.马卡罗沃Makarovo.SetParent(f)
	
	f.奥克西诺Oksino = BOksino奥克西诺
	f.奥克西诺Oksino.SetParent(f)
	
	f.佩列梅茨Pylemets = BPylemets佩列梅茨
	f.佩列梅茨Pylemets.SetParent(f)
	
	f.扎瓦罗特Zavarot = BZavarot扎瓦罗特
	f.扎瓦罗特Zavarot.SetParent(f)
	
}
