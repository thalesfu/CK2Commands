package naupactus

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type NaupactusCounty interface {
    feud.County
    BAgrinio阿格里尼翁() 	feud.Barony
    BAlyzeia阿吕齐亚() 	feud.Barony
    BLidoriki利佐里基() 	feud.Barony
    BMesolongi迈索隆吉() 	feud.Barony
    BNaupaktos纳夫帕克托斯() 	feud.Barony
    BParavola帕拉沃拉() 	feud.Barony
    BSollion索利翁() 	feud.Barony
}

type 纳夫帕克托斯NaupactusCounty struct {
	feud.BaseCounty
	阿格里尼翁Agrinio 	feud.Barony
	阿吕齐亚Alyzeia 	feud.Barony
	利佐里基Lidoriki 	feud.Barony
	迈索隆吉Mesolongi 	feud.Barony
	纳夫帕克托斯Naupaktos 	feud.Barony
	帕拉沃拉Paravola 	feud.Barony
	索利翁Sollion 	feud.Barony
}

func (c *纳夫帕克托斯NaupactusCounty) BAgrinio阿格里尼翁() feud.Barony {
	return c.阿格里尼翁Agrinio
}
    
func (c *纳夫帕克托斯NaupactusCounty) BAlyzeia阿吕齐亚() feud.Barony {
	return c.阿吕齐亚Alyzeia
}
    
func (c *纳夫帕克托斯NaupactusCounty) BLidoriki利佐里基() feud.Barony {
	return c.利佐里基Lidoriki
}
    
func (c *纳夫帕克托斯NaupactusCounty) BMesolongi迈索隆吉() feud.Barony {
	return c.迈索隆吉Mesolongi
}
    
func (c *纳夫帕克托斯NaupactusCounty) BNaupaktos纳夫帕克托斯() feud.Barony {
	return c.纳夫帕克托斯Naupaktos
}
    
func (c *纳夫帕克托斯NaupactusCounty) BParavola帕拉沃拉() feud.Barony {
	return c.帕拉沃拉Paravola
}
    
func (c *纳夫帕克托斯NaupactusCounty) BSollion索利翁() feud.Barony {
	return c.索利翁Sollion
}
    
var CNaupactus纳夫帕克托斯 NaupactusCounty = &纳夫帕克托斯NaupactusCounty{}

func init() {
	f := CNaupactus纳夫帕克托斯.(*纳夫帕克托斯NaupactusCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1881",
		Title:     "naupactus",
		TitleName: "纳夫帕克托斯",
		TitleCode: "c_naupactus",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿格里尼翁Agrinio = BAgrinio阿格里尼翁
	f.阿格里尼翁Agrinio.SetParent(f)
	
	f.阿吕齐亚Alyzeia = BAlyzeia阿吕齐亚
	f.阿吕齐亚Alyzeia.SetParent(f)
	
	f.利佐里基Lidoriki = BLidoriki利佐里基
	f.利佐里基Lidoriki.SetParent(f)
	
	f.迈索隆吉Mesolongi = BMesolongi迈索隆吉
	f.迈索隆吉Mesolongi.SetParent(f)
	
	f.纳夫帕克托斯Naupaktos = BNaupaktos纳夫帕克托斯
	f.纳夫帕克托斯Naupaktos.SetParent(f)
	
	f.帕拉沃拉Paravola = BParavola帕拉沃拉
	f.帕拉沃拉Paravola.SetParent(f)
	
	f.索利翁Sollion = BSollion索利翁
	f.索利翁Sollion.SetParent(f)
	
}
