package heves

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type HevesCounty interface {
    feud.County
    BEger埃格尔() 	feud.Barony
    BGyongyos珍珠市() 	feud.Barony
    BHatvan豪特翁() 	feud.Barony
    BHeves赫维什() 	feud.Barony
    BMezokovesd迈泽克韦什德() 	feud.Barony
    BMiskolc米什科尔茨() 	feud.Barony
    BPetervasara彼得瓦沙劳() 	feud.Barony
    BTiszafured蒂绍菲赖德() 	feud.Barony
}

type 赫维什HevesCounty struct {
	feud.BaseCounty
	埃格尔Eger 	feud.Barony
	珍珠市Gyongyos 	feud.Barony
	豪特翁Hatvan 	feud.Barony
	赫维什Heves 	feud.Barony
	迈泽克韦什德Mezokovesd 	feud.Barony
	米什科尔茨Miskolc 	feud.Barony
	彼得瓦沙劳Petervasara 	feud.Barony
	蒂绍菲赖德Tiszafured 	feud.Barony
}

func (c *赫维什HevesCounty) BEger埃格尔() feud.Barony {
	return c.埃格尔Eger
}
    
func (c *赫维什HevesCounty) BGyongyos珍珠市() feud.Barony {
	return c.珍珠市Gyongyos
}
    
func (c *赫维什HevesCounty) BHatvan豪特翁() feud.Barony {
	return c.豪特翁Hatvan
}
    
func (c *赫维什HevesCounty) BHeves赫维什() feud.Barony {
	return c.赫维什Heves
}
    
func (c *赫维什HevesCounty) BMezokovesd迈泽克韦什德() feud.Barony {
	return c.迈泽克韦什德Mezokovesd
}
    
func (c *赫维什HevesCounty) BMiskolc米什科尔茨() feud.Barony {
	return c.米什科尔茨Miskolc
}
    
func (c *赫维什HevesCounty) BPetervasara彼得瓦沙劳() feud.Barony {
	return c.彼得瓦沙劳Petervasara
}
    
func (c *赫维什HevesCounty) BTiszafured蒂绍菲赖德() feud.Barony {
	return c.蒂绍菲赖德Tiszafured
}
    
var CHeves赫维什 HevesCounty = &赫维什HevesCounty{}

func init() {
	f := CHeves赫维什.(*赫维什HevesCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "523",
		Title:     "heves",
		TitleName: "赫维什",
		TitleCode: "c_heves",
		Baronies:  map[string]feud.Barony{},
	}

	f.埃格尔Eger = BEger埃格尔
	f.埃格尔Eger.SetParent(f)
	
	f.珍珠市Gyongyos = BGyongyos珍珠市
	f.珍珠市Gyongyos.SetParent(f)
	
	f.豪特翁Hatvan = BHatvan豪特翁
	f.豪特翁Hatvan.SetParent(f)
	
	f.赫维什Heves = BHeves赫维什
	f.赫维什Heves.SetParent(f)
	
	f.迈泽克韦什德Mezokovesd = BMezokovesd迈泽克韦什德
	f.迈泽克韦什德Mezokovesd.SetParent(f)
	
	f.米什科尔茨Miskolc = BMiskolc米什科尔茨
	f.米什科尔茨Miskolc.SetParent(f)
	
	f.彼得瓦沙劳Petervasara = BPetervasara彼得瓦沙劳
	f.彼得瓦沙劳Petervasara.SetParent(f)
	
	f.蒂绍菲赖德Tiszafured = BTiszafured蒂绍菲赖德
	f.蒂绍菲赖德Tiszafured.SetParent(f)
	
}
