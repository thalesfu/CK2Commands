package terek

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TerekCounty interface {
    feud.County
    BChanta昌塔() 	feud.Barony
    BGornyy戈尔内() 	feud.Barony
    BKurskaya库尔斯卡亚() 	feud.Barony
    BMajar马贾尔() 	feud.Barony
    BStepnoye斯捷普诺耶() 	feud.Barony
    BTerek捷列克() 	feud.Barony
    BTermita捷尔米塔() 	feud.Barony
}

type 捷列克TerekCounty struct {
	feud.BaseCounty
	昌塔Chanta 	feud.Barony
	戈尔内Gornyy 	feud.Barony
	库尔斯卡亚Kurskaya 	feud.Barony
	马贾尔Majar 	feud.Barony
	斯捷普诺耶Stepnoye 	feud.Barony
	捷列克Terek 	feud.Barony
	捷尔米塔Termita 	feud.Barony
}

func (c *捷列克TerekCounty) BChanta昌塔() feud.Barony {
	return c.昌塔Chanta
}
    
func (c *捷列克TerekCounty) BGornyy戈尔内() feud.Barony {
	return c.戈尔内Gornyy
}
    
func (c *捷列克TerekCounty) BKurskaya库尔斯卡亚() feud.Barony {
	return c.库尔斯卡亚Kurskaya
}
    
func (c *捷列克TerekCounty) BMajar马贾尔() feud.Barony {
	return c.马贾尔Majar
}
    
func (c *捷列克TerekCounty) BStepnoye斯捷普诺耶() feud.Barony {
	return c.斯捷普诺耶Stepnoye
}
    
func (c *捷列克TerekCounty) BTerek捷列克() feud.Barony {
	return c.捷列克Terek
}
    
func (c *捷列克TerekCounty) BTermita捷尔米塔() feud.Barony {
	return c.捷尔米塔Termita
}
    
var CTerek捷列克 TerekCounty = &捷列克TerekCounty{}

func init() {
	f := CTerek捷列克.(*捷列克TerekCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1810",
		Title:     "terek",
		TitleName: "捷列克",
		TitleCode: "c_terek",
		Baronies:  map[string]feud.Barony{},
	}

	f.昌塔Chanta = BChanta昌塔
	f.昌塔Chanta.SetParent(f)
	
	f.戈尔内Gornyy = BGornyy戈尔内
	f.戈尔内Gornyy.SetParent(f)
	
	f.库尔斯卡亚Kurskaya = BKurskaya库尔斯卡亚
	f.库尔斯卡亚Kurskaya.SetParent(f)
	
	f.马贾尔Majar = BMajar马贾尔
	f.马贾尔Majar.SetParent(f)
	
	f.斯捷普诺耶Stepnoye = BStepnoye斯捷普诺耶
	f.斯捷普诺耶Stepnoye.SetParent(f)
	
	f.捷列克Terek = BTerek捷列克
	f.捷列克Terek.SetParent(f)
	
	f.捷尔米塔Termita = BTermita捷尔米塔
	f.捷尔米塔Termita.SetParent(f)
	
}
