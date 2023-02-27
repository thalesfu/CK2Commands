package bolzano

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BolzanoCounty interface {
    feud.County
    BBozen博尔扎诺() 	feud.Barony
    BCortina科尔蒂纳() 	feud.Barony
    BDolladitz多拉迪兹() 	feud.Barony
    BEppan埃潘() 	feud.Barony
    BFlorian弗洛里安() 	feud.Barony
    BMeran梅兰() 	feud.Barony
    BSt_helena圣海伦娜() 	feud.Barony
    BSt_jacob圣雅各伯() 	feud.Barony
}

type 博尔扎诺BolzanoCounty struct {
	feud.BaseCounty
	博尔扎诺Bozen 	feud.Barony
	科尔蒂纳Cortina 	feud.Barony
	多拉迪兹Dolladitz 	feud.Barony
	埃潘Eppan 	feud.Barony
	弗洛里安Florian 	feud.Barony
	梅兰Meran 	feud.Barony
	圣海伦娜St_helena 	feud.Barony
	圣雅各伯St_jacob 	feud.Barony
}

func (c *博尔扎诺BolzanoCounty) BBozen博尔扎诺() feud.Barony {
	return c.博尔扎诺Bozen
}
    
func (c *博尔扎诺BolzanoCounty) BCortina科尔蒂纳() feud.Barony {
	return c.科尔蒂纳Cortina
}
    
func (c *博尔扎诺BolzanoCounty) BDolladitz多拉迪兹() feud.Barony {
	return c.多拉迪兹Dolladitz
}
    
func (c *博尔扎诺BolzanoCounty) BEppan埃潘() feud.Barony {
	return c.埃潘Eppan
}
    
func (c *博尔扎诺BolzanoCounty) BFlorian弗洛里安() feud.Barony {
	return c.弗洛里安Florian
}
    
func (c *博尔扎诺BolzanoCounty) BMeran梅兰() feud.Barony {
	return c.梅兰Meran
}
    
func (c *博尔扎诺BolzanoCounty) BSt_helena圣海伦娜() feud.Barony {
	return c.圣海伦娜St_helena
}
    
func (c *博尔扎诺BolzanoCounty) BSt_jacob圣雅各伯() feud.Barony {
	return c.圣雅各伯St_jacob
}
    
var CBolzano博尔扎诺 BolzanoCounty = &博尔扎诺BolzanoCounty{}

func init() {
	f := CBolzano博尔扎诺.(*博尔扎诺BolzanoCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1613",
		Title:     "bolzano",
		TitleName: "博尔扎诺",
		TitleCode: "c_bolzano",
		Baronies:  map[string]feud.Barony{},
	}

	f.博尔扎诺Bozen = BBozen博尔扎诺
	f.博尔扎诺Bozen.SetParent(f)
	
	f.科尔蒂纳Cortina = BCortina科尔蒂纳
	f.科尔蒂纳Cortina.SetParent(f)
	
	f.多拉迪兹Dolladitz = BDolladitz多拉迪兹
	f.多拉迪兹Dolladitz.SetParent(f)
	
	f.埃潘Eppan = BEppan埃潘
	f.埃潘Eppan.SetParent(f)
	
	f.弗洛里安Florian = BFlorian弗洛里安
	f.弗洛里安Florian.SetParent(f)
	
	f.梅兰Meran = BMeran梅兰
	f.梅兰Meran.SetParent(f)
	
	f.圣海伦娜St_helena = BSt_helena圣海伦娜
	f.圣海伦娜St_helena.SetParent(f)
	
	f.圣雅各伯St_jacob = BSt_jacob圣雅各伯
	f.圣雅各伯St_jacob.SetParent(f)
	
}
