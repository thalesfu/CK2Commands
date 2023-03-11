package ponthieu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type PonthieuCounty interface {
    feud.County
    BAbbeville阿布维尔() 	feud.Barony
    BAult欧村() 	feud.Barony
    BFlixecourt弗利克斯库尔() 	feud.Barony
    BLe_crotoy勒克罗图瓦() 	feud.Barony
    BLe_treport勒特雷波尔() 	feud.Barony
    BSaint_valery圣瓦莱里() 	feud.Barony
}

type 蓬蒂耶PonthieuCounty struct {
	feud.BaseCounty
	阿布维尔Abbeville 	feud.Barony
	欧村Ault 	feud.Barony
	弗利克斯库尔Flixecourt 	feud.Barony
	勒克罗图瓦Le_crotoy 	feud.Barony
	勒特雷波尔Le_treport 	feud.Barony
	圣瓦莱里Saint_valery 	feud.Barony
}

func (c *蓬蒂耶PonthieuCounty) BAbbeville阿布维尔() feud.Barony {
	return c.阿布维尔Abbeville
}
    
func (c *蓬蒂耶PonthieuCounty) BAult欧村() feud.Barony {
	return c.欧村Ault
}
    
func (c *蓬蒂耶PonthieuCounty) BFlixecourt弗利克斯库尔() feud.Barony {
	return c.弗利克斯库尔Flixecourt
}
    
func (c *蓬蒂耶PonthieuCounty) BLe_crotoy勒克罗图瓦() feud.Barony {
	return c.勒克罗图瓦Le_crotoy
}
    
func (c *蓬蒂耶PonthieuCounty) BLe_treport勒特雷波尔() feud.Barony {
	return c.勒特雷波尔Le_treport
}
    
func (c *蓬蒂耶PonthieuCounty) BSaint_valery圣瓦莱里() feud.Barony {
	return c.圣瓦莱里Saint_valery
}
    
var CPonthieu蓬蒂耶 PonthieuCounty = &蓬蒂耶PonthieuCounty{}

func init() {
	f := CPonthieu蓬蒂耶.(*蓬蒂耶PonthieuCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1962",
		Title:     "ponthieu",
		TitleName: "蓬蒂耶",
		TitleCode: "c_ponthieu",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿布维尔Abbeville = BAbbeville阿布维尔
	f.阿布维尔Abbeville.SetParent(f)
	
	f.欧村Ault = BAult欧村
	f.欧村Ault.SetParent(f)
	
	f.弗利克斯库尔Flixecourt = BFlixecourt弗利克斯库尔
	f.弗利克斯库尔Flixecourt.SetParent(f)
	
	f.勒克罗图瓦Le_crotoy = BLe_crotoy勒克罗图瓦
	f.勒克罗图瓦Le_crotoy.SetParent(f)
	
	f.勒特雷波尔Le_treport = BLe_treport勒特雷波尔
	f.勒特雷波尔Le_treport.SetParent(f)
	
	f.圣瓦莱里Saint_valery = BSaint_valery圣瓦莱里
	f.圣瓦莱里Saint_valery.SetParent(f)
	
}
