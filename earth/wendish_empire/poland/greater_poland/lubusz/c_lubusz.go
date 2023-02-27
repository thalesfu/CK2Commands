package lubusz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type LubuszCounty interface {
    feud.County
    BCedynia采迪尼亚() 	feud.Barony
    BKostrzyn科斯琴() 	feud.Barony
    BLubusz卢布什() 	feud.Barony
    BMysliborz梅希利布日() 	feud.Barony
    BNadodra纳多德拉() 	feud.Barony
    BPryzyce佩日采() 	feud.Barony
    BSulecia苏伦恰() 	feud.Barony
    BSulecin苏伦钦() 	feud.Barony
}

type 卢布什LubuszCounty struct {
	feud.BaseCounty
	采迪尼亚Cedynia 	feud.Barony
	科斯琴Kostrzyn 	feud.Barony
	卢布什Lubusz 	feud.Barony
	梅希利布日Mysliborz 	feud.Barony
	纳多德拉Nadodra 	feud.Barony
	佩日采Pryzyce 	feud.Barony
	苏伦恰Sulecia 	feud.Barony
	苏伦钦Sulecin 	feud.Barony
}

func (c *卢布什LubuszCounty) BCedynia采迪尼亚() feud.Barony {
	return c.采迪尼亚Cedynia
}
    
func (c *卢布什LubuszCounty) BKostrzyn科斯琴() feud.Barony {
	return c.科斯琴Kostrzyn
}
    
func (c *卢布什LubuszCounty) BLubusz卢布什() feud.Barony {
	return c.卢布什Lubusz
}
    
func (c *卢布什LubuszCounty) BMysliborz梅希利布日() feud.Barony {
	return c.梅希利布日Mysliborz
}
    
func (c *卢布什LubuszCounty) BNadodra纳多德拉() feud.Barony {
	return c.纳多德拉Nadodra
}
    
func (c *卢布什LubuszCounty) BPryzyce佩日采() feud.Barony {
	return c.佩日采Pryzyce
}
    
func (c *卢布什LubuszCounty) BSulecia苏伦恰() feud.Barony {
	return c.苏伦恰Sulecia
}
    
func (c *卢布什LubuszCounty) BSulecin苏伦钦() feud.Barony {
	return c.苏伦钦Sulecin
}
    
var CLubusz卢布什 LubuszCounty = &卢布什LubuszCounty{}

func init() {
	f := CLubusz卢布什.(*卢布什LubuszCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "430",
		Title:     "lubusz",
		TitleName: "卢布什",
		TitleCode: "c_lubusz",
		Baronies:  map[string]feud.Barony{},
	}

	f.采迪尼亚Cedynia = BCedynia采迪尼亚
	f.采迪尼亚Cedynia.SetParent(f)
	
	f.科斯琴Kostrzyn = BKostrzyn科斯琴
	f.科斯琴Kostrzyn.SetParent(f)
	
	f.卢布什Lubusz = BLubusz卢布什
	f.卢布什Lubusz.SetParent(f)
	
	f.梅希利布日Mysliborz = BMysliborz梅希利布日
	f.梅希利布日Mysliborz.SetParent(f)
	
	f.纳多德拉Nadodra = BNadodra纳多德拉
	f.纳多德拉Nadodra.SetParent(f)
	
	f.佩日采Pryzyce = BPryzyce佩日采
	f.佩日采Pryzyce.SetParent(f)
	
	f.苏伦恰Sulecia = BSulecia苏伦恰
	f.苏伦恰Sulecia.SetParent(f)
	
	f.苏伦钦Sulecin = BSulecin苏伦钦
	f.苏伦钦Sulecin.SetParent(f)
	
}
