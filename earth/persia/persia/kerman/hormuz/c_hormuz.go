package hormuz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type HormuzCounty interface {
    feud.County
    BAbarkawan阿巴尔卡万() 	feud.Barony
    BAbu_musa阿布穆萨() 	feud.Barony
    BBam巴姆() 	feud.Barony
    BFin费恩() 	feud.Barony
    BGombroon贡布伦() 	feud.Barony
    BJiroth吉洛特() 	feud.Barony
    BKish基什() 	feud.Barony
    BMinab米纳卜() 	feud.Barony
}

type 霍尔木兹HormuzCounty struct {
	feud.BaseCounty
	阿巴尔卡万Abarkawan 	feud.Barony
	阿布穆萨Abu_musa 	feud.Barony
	巴姆Bam 	feud.Barony
	费恩Fin 	feud.Barony
	贡布伦Gombroon 	feud.Barony
	吉洛特Jiroth 	feud.Barony
	基什Kish 	feud.Barony
	米纳卜Minab 	feud.Barony
}

func (c *霍尔木兹HormuzCounty) BAbarkawan阿巴尔卡万() feud.Barony {
	return c.阿巴尔卡万Abarkawan
}
    
func (c *霍尔木兹HormuzCounty) BAbu_musa阿布穆萨() feud.Barony {
	return c.阿布穆萨Abu_musa
}
    
func (c *霍尔木兹HormuzCounty) BBam巴姆() feud.Barony {
	return c.巴姆Bam
}
    
func (c *霍尔木兹HormuzCounty) BFin费恩() feud.Barony {
	return c.费恩Fin
}
    
func (c *霍尔木兹HormuzCounty) BGombroon贡布伦() feud.Barony {
	return c.贡布伦Gombroon
}
    
func (c *霍尔木兹HormuzCounty) BJiroth吉洛特() feud.Barony {
	return c.吉洛特Jiroth
}
    
func (c *霍尔木兹HormuzCounty) BKish基什() feud.Barony {
	return c.基什Kish
}
    
func (c *霍尔木兹HormuzCounty) BMinab米纳卜() feud.Barony {
	return c.米纳卜Minab
}
    
var CHormuz霍尔木兹 HormuzCounty = &霍尔木兹HormuzCounty{}

func init() {
	f := CHormuz霍尔木兹.(*霍尔木兹HormuzCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "641",
		Title:     "hormuz",
		TitleName: "霍尔木兹",
		TitleCode: "c_hormuz",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿巴尔卡万Abarkawan = BAbarkawan阿巴尔卡万
	f.阿巴尔卡万Abarkawan.SetParent(f)
	
	f.阿布穆萨Abu_musa = BAbu_musa阿布穆萨
	f.阿布穆萨Abu_musa.SetParent(f)
	
	f.巴姆Bam = BBam巴姆
	f.巴姆Bam.SetParent(f)
	
	f.费恩Fin = BFin费恩
	f.费恩Fin.SetParent(f)
	
	f.贡布伦Gombroon = BGombroon贡布伦
	f.贡布伦Gombroon.SetParent(f)
	
	f.吉洛特Jiroth = BJiroth吉洛特
	f.吉洛特Jiroth.SetParent(f)
	
	f.基什Kish = BKish基什
	f.基什Kish.SetParent(f)
	
	f.米纳卜Minab = BMinab米纳卜
	f.米纳卜Minab.SetParent(f)
	
}
