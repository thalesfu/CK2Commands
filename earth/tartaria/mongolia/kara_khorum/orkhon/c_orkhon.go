package orkhon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type OrkhonCounty interface {
    feud.County
    BBaruunburen巴伦布兰() 	feud.Barony
    BDarkhan_orkhon达尔汗() 	feud.Barony
    BKhushaat胡夏特() 	feud.Barony
    BKhutul呼图勒() 	feud.Barony
    BOrkhon嗢昆() 	feud.Barony
    BOrkhontuul鄂尔浑图勒() 	feud.Barony
    BSant桑特() 	feud.Barony
}

type 嗢昆OrkhonCounty struct {
	feud.BaseCounty
	巴伦布兰Baruunburen 	feud.Barony
	达尔汗Darkhan_orkhon 	feud.Barony
	胡夏特Khushaat 	feud.Barony
	呼图勒Khutul 	feud.Barony
	嗢昆Orkhon 	feud.Barony
	鄂尔浑图勒Orkhontuul 	feud.Barony
	桑特Sant 	feud.Barony
}

func (c *嗢昆OrkhonCounty) BBaruunburen巴伦布兰() feud.Barony {
	return c.巴伦布兰Baruunburen
}
    
func (c *嗢昆OrkhonCounty) BDarkhan_orkhon达尔汗() feud.Barony {
	return c.达尔汗Darkhan_orkhon
}
    
func (c *嗢昆OrkhonCounty) BKhushaat胡夏特() feud.Barony {
	return c.胡夏特Khushaat
}
    
func (c *嗢昆OrkhonCounty) BKhutul呼图勒() feud.Barony {
	return c.呼图勒Khutul
}
    
func (c *嗢昆OrkhonCounty) BOrkhon嗢昆() feud.Barony {
	return c.嗢昆Orkhon
}
    
func (c *嗢昆OrkhonCounty) BOrkhontuul鄂尔浑图勒() feud.Barony {
	return c.鄂尔浑图勒Orkhontuul
}
    
func (c *嗢昆OrkhonCounty) BSant桑特() feud.Barony {
	return c.桑特Sant
}
    
var COrkhon嗢昆 OrkhonCounty = &嗢昆OrkhonCounty{}

func init() {
	f := COrkhon嗢昆.(*嗢昆OrkhonCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1920",
		Title:     "orkhon",
		TitleName: "嗢昆",
		TitleCode: "c_orkhon",
		Baronies:  map[string]feud.Barony{},
	}

	f.巴伦布兰Baruunburen = BBaruunburen巴伦布兰
	f.巴伦布兰Baruunburen.SetParent(f)
	
	f.达尔汗Darkhan_orkhon = BDarkhan_orkhon达尔汗
	f.达尔汗Darkhan_orkhon.SetParent(f)
	
	f.胡夏特Khushaat = BKhushaat胡夏特
	f.胡夏特Khushaat.SetParent(f)
	
	f.呼图勒Khutul = BKhutul呼图勒
	f.呼图勒Khutul.SetParent(f)
	
	f.嗢昆Orkhon = BOrkhon嗢昆
	f.嗢昆Orkhon.SetParent(f)
	
	f.鄂尔浑图勒Orkhontuul = BOrkhontuul鄂尔浑图勒
	f.鄂尔浑图勒Orkhontuul.SetParent(f)
	
	f.桑特Sant = BSant桑特
	f.桑特Sant.SetParent(f)
	
}
