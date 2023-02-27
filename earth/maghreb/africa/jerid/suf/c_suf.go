package suf

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SufCounty interface {
    feud.County
    BBeni_dinhet贝尼迪内特() 	feud.Barony
    BBeni_kaltoum贝尼卡尔图姆() 	feud.Barony
    BBidane比当() 	feud.Barony
    BDar_abd_el_aziz达尔阿卜杜勒阿齐兹() 	feud.Barony
    BEl_aouina欧韦奈() 	feud.Barony
    BFatnassa法特纳萨() 	feud.Barony
    BSuf苏夫() 	feud.Barony
}

type 苏夫SufCounty struct {
	feud.BaseCounty
	贝尼迪内特Beni_dinhet 	feud.Barony
	贝尼卡尔图姆Beni_kaltoum 	feud.Barony
	比当Bidane 	feud.Barony
	达尔阿卜杜勒阿齐兹Dar_abd_el_aziz 	feud.Barony
	欧韦奈El_aouina 	feud.Barony
	法特纳萨Fatnassa 	feud.Barony
	苏夫Suf 	feud.Barony
}

func (c *苏夫SufCounty) BBeni_dinhet贝尼迪内特() feud.Barony {
	return c.贝尼迪内特Beni_dinhet
}
    
func (c *苏夫SufCounty) BBeni_kaltoum贝尼卡尔图姆() feud.Barony {
	return c.贝尼卡尔图姆Beni_kaltoum
}
    
func (c *苏夫SufCounty) BBidane比当() feud.Barony {
	return c.比当Bidane
}
    
func (c *苏夫SufCounty) BDar_abd_el_aziz达尔阿卜杜勒阿齐兹() feud.Barony {
	return c.达尔阿卜杜勒阿齐兹Dar_abd_el_aziz
}
    
func (c *苏夫SufCounty) BEl_aouina欧韦奈() feud.Barony {
	return c.欧韦奈El_aouina
}
    
func (c *苏夫SufCounty) BFatnassa法特纳萨() feud.Barony {
	return c.法特纳萨Fatnassa
}
    
func (c *苏夫SufCounty) BSuf苏夫() feud.Barony {
	return c.苏夫Suf
}
    
var CSuf苏夫 SufCounty = &苏夫SufCounty{}

func init() {
	f := CSuf苏夫.(*苏夫SufCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1725",
		Title:     "suf",
		TitleName: "苏夫",
		TitleCode: "c_suf",
		Baronies:  map[string]feud.Barony{},
	}

	f.贝尼迪内特Beni_dinhet = BBeni_dinhet贝尼迪内特
	f.贝尼迪内特Beni_dinhet.SetParent(f)
	
	f.贝尼卡尔图姆Beni_kaltoum = BBeni_kaltoum贝尼卡尔图姆
	f.贝尼卡尔图姆Beni_kaltoum.SetParent(f)
	
	f.比当Bidane = BBidane比当
	f.比当Bidane.SetParent(f)
	
	f.达尔阿卜杜勒阿齐兹Dar_abd_el_aziz = BDar_abd_el_aziz达尔阿卜杜勒阿齐兹
	f.达尔阿卜杜勒阿齐兹Dar_abd_el_aziz.SetParent(f)
	
	f.欧韦奈El_aouina = BEl_aouina欧韦奈
	f.欧韦奈El_aouina.SetParent(f)
	
	f.法特纳萨Fatnassa = BFatnassa法特纳萨
	f.法特纳萨Fatnassa.SetParent(f)
	
	f.苏夫Suf = BSuf苏夫
	f.苏夫Suf.SetParent(f)
	
}
