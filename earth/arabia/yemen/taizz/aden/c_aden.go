package aden

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AdenCounty interface {
    feud.County
    BAden亚丁() 	feud.Barony
    BAlarbadi阿尔巴迪() 	feud.Barony
    BAlawbhali奥扎利() 	feud.Barony
    BAlkawd考德() 	feud.Barony
    BAlmilah米拉() 	feud.Barony
    BDhala达利阿() 	feud.Barony
    BJaar杰阿尔() 	feud.Barony
    BLahej拉哈杰() 	feud.Barony
}

type 亚丁AdenCounty struct {
	feud.BaseCounty
	亚丁Aden 	feud.Barony
	阿尔巴迪Alarbadi 	feud.Barony
	奥扎利Alawbhali 	feud.Barony
	考德Alkawd 	feud.Barony
	米拉Almilah 	feud.Barony
	达利阿Dhala 	feud.Barony
	杰阿尔Jaar 	feud.Barony
	拉哈杰Lahej 	feud.Barony
}

func (c *亚丁AdenCounty) BAden亚丁() feud.Barony {
	return c.亚丁Aden
}
    
func (c *亚丁AdenCounty) BAlarbadi阿尔巴迪() feud.Barony {
	return c.阿尔巴迪Alarbadi
}
    
func (c *亚丁AdenCounty) BAlawbhali奥扎利() feud.Barony {
	return c.奥扎利Alawbhali
}
    
func (c *亚丁AdenCounty) BAlkawd考德() feud.Barony {
	return c.考德Alkawd
}
    
func (c *亚丁AdenCounty) BAlmilah米拉() feud.Barony {
	return c.米拉Almilah
}
    
func (c *亚丁AdenCounty) BDhala达利阿() feud.Barony {
	return c.达利阿Dhala
}
    
func (c *亚丁AdenCounty) BJaar杰阿尔() feud.Barony {
	return c.杰阿尔Jaar
}
    
func (c *亚丁AdenCounty) BLahej拉哈杰() feud.Barony {
	return c.拉哈杰Lahej
}
    
var CAden亚丁 AdenCounty = &亚丁AdenCounty{}

func init() {
	f := CAden亚丁.(*亚丁AdenCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "858",
		Title:     "aden",
		TitleName: "亚丁",
		TitleCode: "c_aden",
		Baronies:  map[string]feud.Barony{},
	}

	f.亚丁Aden = BAden亚丁
	f.亚丁Aden.SetParent(f)
	
	f.阿尔巴迪Alarbadi = BAlarbadi阿尔巴迪
	f.阿尔巴迪Alarbadi.SetParent(f)
	
	f.奥扎利Alawbhali = BAlawbhali奥扎利
	f.奥扎利Alawbhali.SetParent(f)
	
	f.考德Alkawd = BAlkawd考德
	f.考德Alkawd.SetParent(f)
	
	f.米拉Almilah = BAlmilah米拉
	f.米拉Almilah.SetParent(f)
	
	f.达利阿Dhala = BDhala达利阿
	f.达利阿Dhala.SetParent(f)
	
	f.杰阿尔Jaar = BJaar杰阿尔
	f.杰阿尔Jaar.SetParent(f)
	
	f.拉哈杰Lahej = BLahej拉哈杰
	f.拉哈杰Lahej.SetParent(f)
	
}
