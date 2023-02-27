package soria

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SoriaCounty interface {
    feud.County
    BAlmazan阿尔马桑() 	feud.Barony
    BCastromoro卡斯特罗莫罗() 	feud.Barony
    BCovaleda科瓦莱达() 	feud.Barony
    BGormaz戈尔马斯() 	feud.Barony
    BMedinacelli梅迪纳塞利() 	feud.Barony
    BSanleonardodeyague圣莱昂纳多德亚圭() 	feud.Barony
    BSoria索里亚() 	feud.Barony
}

type 索里亚SoriaCounty struct {
	feud.BaseCounty
	阿尔马桑Almazan 	feud.Barony
	卡斯特罗莫罗Castromoro 	feud.Barony
	科瓦莱达Covaleda 	feud.Barony
	戈尔马斯Gormaz 	feud.Barony
	梅迪纳塞利Medinacelli 	feud.Barony
	圣莱昂纳多德亚圭Sanleonardodeyague 	feud.Barony
	索里亚Soria 	feud.Barony
}

func (c *索里亚SoriaCounty) BAlmazan阿尔马桑() feud.Barony {
	return c.阿尔马桑Almazan
}
    
func (c *索里亚SoriaCounty) BCastromoro卡斯特罗莫罗() feud.Barony {
	return c.卡斯特罗莫罗Castromoro
}
    
func (c *索里亚SoriaCounty) BCovaleda科瓦莱达() feud.Barony {
	return c.科瓦莱达Covaleda
}
    
func (c *索里亚SoriaCounty) BGormaz戈尔马斯() feud.Barony {
	return c.戈尔马斯Gormaz
}
    
func (c *索里亚SoriaCounty) BMedinacelli梅迪纳塞利() feud.Barony {
	return c.梅迪纳塞利Medinacelli
}
    
func (c *索里亚SoriaCounty) BSanleonardodeyague圣莱昂纳多德亚圭() feud.Barony {
	return c.圣莱昂纳多德亚圭Sanleonardodeyague
}
    
func (c *索里亚SoriaCounty) BSoria索里亚() feud.Barony {
	return c.索里亚Soria
}
    
var CSoria索里亚 SoriaCounty = &索里亚SoriaCounty{}

func init() {
	f := CSoria索里亚.(*索里亚SoriaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "200",
		Title:     "soria",
		TitleName: "索里亚",
		TitleCode: "c_soria",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿尔马桑Almazan = BAlmazan阿尔马桑
	f.阿尔马桑Almazan.SetParent(f)
	
	f.卡斯特罗莫罗Castromoro = BCastromoro卡斯特罗莫罗
	f.卡斯特罗莫罗Castromoro.SetParent(f)
	
	f.科瓦莱达Covaleda = BCovaleda科瓦莱达
	f.科瓦莱达Covaleda.SetParent(f)
	
	f.戈尔马斯Gormaz = BGormaz戈尔马斯
	f.戈尔马斯Gormaz.SetParent(f)
	
	f.梅迪纳塞利Medinacelli = BMedinacelli梅迪纳塞利
	f.梅迪纳塞利Medinacelli.SetParent(f)
	
	f.圣莱昂纳多德亚圭Sanleonardodeyague = BSanleonardodeyague圣莱昂纳多德亚圭
	f.圣莱昂纳多德亚圭Sanleonardodeyague.SetParent(f)
	
	f.索里亚Soria = BSoria索里亚
	f.索里亚Soria.SetParent(f)
	
}
