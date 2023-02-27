package retz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type RetzCounty interface {
    feud.County
    BClisson克利松() 	feud.Barony
    BDeas代阿斯() 	feud.Barony
    BLa_bernerie拉贝尔讷里() 	feud.Barony
    BMachecoul马什库勒() 	feud.Barony
    BPornic波尔尼克() 	feud.Barony
    BReze勒泽() 	feud.Barony
    BSaint_brevin圣布雷万() 	feud.Barony
}

type 雷茨RetzCounty struct {
	feud.BaseCounty
	克利松Clisson 	feud.Barony
	代阿斯Deas 	feud.Barony
	拉贝尔讷里La_bernerie 	feud.Barony
	马什库勒Machecoul 	feud.Barony
	波尔尼克Pornic 	feud.Barony
	勒泽Reze 	feud.Barony
	圣布雷万Saint_brevin 	feud.Barony
}

func (c *雷茨RetzCounty) BClisson克利松() feud.Barony {
	return c.克利松Clisson
}
    
func (c *雷茨RetzCounty) BDeas代阿斯() feud.Barony {
	return c.代阿斯Deas
}
    
func (c *雷茨RetzCounty) BLa_bernerie拉贝尔讷里() feud.Barony {
	return c.拉贝尔讷里La_bernerie
}
    
func (c *雷茨RetzCounty) BMachecoul马什库勒() feud.Barony {
	return c.马什库勒Machecoul
}
    
func (c *雷茨RetzCounty) BPornic波尔尼克() feud.Barony {
	return c.波尔尼克Pornic
}
    
func (c *雷茨RetzCounty) BReze勒泽() feud.Barony {
	return c.勒泽Reze
}
    
func (c *雷茨RetzCounty) BSaint_brevin圣布雷万() feud.Barony {
	return c.圣布雷万Saint_brevin
}
    
var CRetz雷茨 RetzCounty = &雷茨RetzCounty{}

func init() {
	f := CRetz雷茨.(*雷茨RetzCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1957",
		Title:     "retz",
		TitleName: "雷茨",
		TitleCode: "c_retz",
		Baronies:  map[string]feud.Barony{},
	}

	f.克利松Clisson = BClisson克利松
	f.克利松Clisson.SetParent(f)
	
	f.代阿斯Deas = BDeas代阿斯
	f.代阿斯Deas.SetParent(f)
	
	f.拉贝尔讷里La_bernerie = BLa_bernerie拉贝尔讷里
	f.拉贝尔讷里La_bernerie.SetParent(f)
	
	f.马什库勒Machecoul = BMachecoul马什库勒
	f.马什库勒Machecoul.SetParent(f)
	
	f.波尔尼克Pornic = BPornic波尔尼克
	f.波尔尼克Pornic.SetParent(f)
	
	f.勒泽Reze = BReze勒泽
	f.勒泽Reze.SetParent(f)
	
	f.圣布雷万Saint_brevin = BSaint_brevin圣布雷万
	f.圣布雷万Saint_brevin.SetParent(f)
	
}
