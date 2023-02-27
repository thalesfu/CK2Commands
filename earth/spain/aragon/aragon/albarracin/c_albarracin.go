package albarracin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AlbarracinCounty interface {
    feud.County
    BAlbarracin阿尔瓦拉辛() 	feud.Barony
    BAlcaniz阿尔卡尼斯() 	feud.Barony
    BCalamocha卡拉莫查() 	feud.Barony
    BCalanda卡兰达() 	feud.Barony
    BHijar伊哈尔() 	feud.Barony
    BMontalban蒙塔尔万() 	feud.Barony
    BTeruel特鲁埃尔() 	feud.Barony
    BUtrillas乌特里利亚斯() 	feud.Barony
}

type 阿尔瓦拉辛AlbarracinCounty struct {
	feud.BaseCounty
	阿尔瓦拉辛Albarracin 	feud.Barony
	阿尔卡尼斯Alcaniz 	feud.Barony
	卡拉莫查Calamocha 	feud.Barony
	卡兰达Calanda 	feud.Barony
	伊哈尔Hijar 	feud.Barony
	蒙塔尔万Montalban 	feud.Barony
	特鲁埃尔Teruel 	feud.Barony
	乌特里利亚斯Utrillas 	feud.Barony
}

func (c *阿尔瓦拉辛AlbarracinCounty) BAlbarracin阿尔瓦拉辛() feud.Barony {
	return c.阿尔瓦拉辛Albarracin
}
    
func (c *阿尔瓦拉辛AlbarracinCounty) BAlcaniz阿尔卡尼斯() feud.Barony {
	return c.阿尔卡尼斯Alcaniz
}
    
func (c *阿尔瓦拉辛AlbarracinCounty) BCalamocha卡拉莫查() feud.Barony {
	return c.卡拉莫查Calamocha
}
    
func (c *阿尔瓦拉辛AlbarracinCounty) BCalanda卡兰达() feud.Barony {
	return c.卡兰达Calanda
}
    
func (c *阿尔瓦拉辛AlbarracinCounty) BHijar伊哈尔() feud.Barony {
	return c.伊哈尔Hijar
}
    
func (c *阿尔瓦拉辛AlbarracinCounty) BMontalban蒙塔尔万() feud.Barony {
	return c.蒙塔尔万Montalban
}
    
func (c *阿尔瓦拉辛AlbarracinCounty) BTeruel特鲁埃尔() feud.Barony {
	return c.特鲁埃尔Teruel
}
    
func (c *阿尔瓦拉辛AlbarracinCounty) BUtrillas乌特里利亚斯() feud.Barony {
	return c.乌特里利亚斯Utrillas
}
    
var CAlbarracin阿尔瓦拉辛 AlbarracinCounty = &阿尔瓦拉辛AlbarracinCounty{}

func init() {
	f := CAlbarracin阿尔瓦拉辛.(*阿尔瓦拉辛AlbarracinCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "174",
		Title:     "albarracin",
		TitleName: "阿尔瓦拉辛",
		TitleCode: "c_albarracin",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿尔瓦拉辛Albarracin = BAlbarracin阿尔瓦拉辛
	f.阿尔瓦拉辛Albarracin.SetParent(f)
	
	f.阿尔卡尼斯Alcaniz = BAlcaniz阿尔卡尼斯
	f.阿尔卡尼斯Alcaniz.SetParent(f)
	
	f.卡拉莫查Calamocha = BCalamocha卡拉莫查
	f.卡拉莫查Calamocha.SetParent(f)
	
	f.卡兰达Calanda = BCalanda卡兰达
	f.卡兰达Calanda.SetParent(f)
	
	f.伊哈尔Hijar = BHijar伊哈尔
	f.伊哈尔Hijar.SetParent(f)
	
	f.蒙塔尔万Montalban = BMontalban蒙塔尔万
	f.蒙塔尔万Montalban.SetParent(f)
	
	f.特鲁埃尔Teruel = BTeruel特鲁埃尔
	f.特鲁埃尔Teruel.SetParent(f)
	
	f.乌特里利亚斯Utrillas = BUtrillas乌特里利亚斯
	f.乌特里利亚斯Utrillas.SetParent(f)
	
}
