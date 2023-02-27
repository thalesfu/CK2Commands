package kunggar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KunggarCounty interface {
    feud.County
    BDrigung必力工瓦() 	feud.Barony
    BGyama甲玛() 	feud.Barony
    BKunggar工卡() 	feud.Barony
    BMamba门巴() 	feud.Barony
    BNyimajangra尼玛江热() 	feud.Barony
    BTanggya塘加() 	feud.Barony
    BZaxoi扎雪() 	feud.Barony
}

type 工卡KunggarCounty struct {
	feud.BaseCounty
	必力工瓦Drigung 	feud.Barony
	甲玛Gyama 	feud.Barony
	工卡Kunggar 	feud.Barony
	门巴Mamba 	feud.Barony
	尼玛江热Nyimajangra 	feud.Barony
	塘加Tanggya 	feud.Barony
	扎雪Zaxoi 	feud.Barony
}

func (c *工卡KunggarCounty) BDrigung必力工瓦() feud.Barony {
	return c.必力工瓦Drigung
}
    
func (c *工卡KunggarCounty) BGyama甲玛() feud.Barony {
	return c.甲玛Gyama
}
    
func (c *工卡KunggarCounty) BKunggar工卡() feud.Barony {
	return c.工卡Kunggar
}
    
func (c *工卡KunggarCounty) BMamba门巴() feud.Barony {
	return c.门巴Mamba
}
    
func (c *工卡KunggarCounty) BNyimajangra尼玛江热() feud.Barony {
	return c.尼玛江热Nyimajangra
}
    
func (c *工卡KunggarCounty) BTanggya塘加() feud.Barony {
	return c.塘加Tanggya
}
    
func (c *工卡KunggarCounty) BZaxoi扎雪() feud.Barony {
	return c.扎雪Zaxoi
}
    
var CKunggar工卡 KunggarCounty = &工卡KunggarCounty{}

func init() {
	f := CKunggar工卡.(*工卡KunggarCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1557",
		Title:     "kunggar",
		TitleName: "工卡",
		TitleCode: "c_kunggar",
		Baronies:  map[string]feud.Barony{},
	}

	f.必力工瓦Drigung = BDrigung必力工瓦
	f.必力工瓦Drigung.SetParent(f)
	
	f.甲玛Gyama = BGyama甲玛
	f.甲玛Gyama.SetParent(f)
	
	f.工卡Kunggar = BKunggar工卡
	f.工卡Kunggar.SetParent(f)
	
	f.门巴Mamba = BMamba门巴
	f.门巴Mamba.SetParent(f)
	
	f.尼玛江热Nyimajangra = BNyimajangra尼玛江热
	f.尼玛江热Nyimajangra.SetParent(f)
	
	f.塘加Tanggya = BTanggya塘加
	f.塘加Tanggya.SetParent(f)
	
	f.扎雪Zaxoi = BZaxoi扎雪
	f.扎雪Zaxoi.SetParent(f)
	
}
