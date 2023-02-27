package palermo

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type PalermoCounty interface {
    feud.County
    BCaltavuturo卡尔塔武图罗() 	feud.Barony
    BCefalu切法卢() 	feud.Barony
    BGratteri格拉泰里() 	feud.Barony
    BMisilmeri米西尔梅里() 	feud.Barony
    BMistretta米斯特雷塔() 	feud.Barony
    BMonreale蒙雷阿莱() 	feud.Barony
    BPalermo巴勒莫() 	feud.Barony
    BPetralia佩特拉利亚() 	feud.Barony
}

type 巴勒莫PalermoCounty struct {
	feud.BaseCounty
	卡尔塔武图罗Caltavuturo 	feud.Barony
	切法卢Cefalu 	feud.Barony
	格拉泰里Gratteri 	feud.Barony
	米西尔梅里Misilmeri 	feud.Barony
	米斯特雷塔Mistretta 	feud.Barony
	蒙雷阿莱Monreale 	feud.Barony
	巴勒莫Palermo 	feud.Barony
	佩特拉利亚Petralia 	feud.Barony
}

func (c *巴勒莫PalermoCounty) BCaltavuturo卡尔塔武图罗() feud.Barony {
	return c.卡尔塔武图罗Caltavuturo
}
    
func (c *巴勒莫PalermoCounty) BCefalu切法卢() feud.Barony {
	return c.切法卢Cefalu
}
    
func (c *巴勒莫PalermoCounty) BGratteri格拉泰里() feud.Barony {
	return c.格拉泰里Gratteri
}
    
func (c *巴勒莫PalermoCounty) BMisilmeri米西尔梅里() feud.Barony {
	return c.米西尔梅里Misilmeri
}
    
func (c *巴勒莫PalermoCounty) BMistretta米斯特雷塔() feud.Barony {
	return c.米斯特雷塔Mistretta
}
    
func (c *巴勒莫PalermoCounty) BMonreale蒙雷阿莱() feud.Barony {
	return c.蒙雷阿莱Monreale
}
    
func (c *巴勒莫PalermoCounty) BPalermo巴勒莫() feud.Barony {
	return c.巴勒莫Palermo
}
    
func (c *巴勒莫PalermoCounty) BPetralia佩特拉利亚() feud.Barony {
	return c.佩特拉利亚Petralia
}
    
var CPalermo巴勒莫 PalermoCounty = &巴勒莫PalermoCounty{}

func init() {
	f := CPalermo巴勒莫.(*巴勒莫PalermoCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "340",
		Title:     "palermo",
		TitleName: "巴勒莫",
		TitleCode: "c_palermo",
		Baronies:  map[string]feud.Barony{},
	}

	f.卡尔塔武图罗Caltavuturo = BCaltavuturo卡尔塔武图罗
	f.卡尔塔武图罗Caltavuturo.SetParent(f)
	
	f.切法卢Cefalu = BCefalu切法卢
	f.切法卢Cefalu.SetParent(f)
	
	f.格拉泰里Gratteri = BGratteri格拉泰里
	f.格拉泰里Gratteri.SetParent(f)
	
	f.米西尔梅里Misilmeri = BMisilmeri米西尔梅里
	f.米西尔梅里Misilmeri.SetParent(f)
	
	f.米斯特雷塔Mistretta = BMistretta米斯特雷塔
	f.米斯特雷塔Mistretta.SetParent(f)
	
	f.蒙雷阿莱Monreale = BMonreale蒙雷阿莱
	f.蒙雷阿莱Monreale.SetParent(f)
	
	f.巴勒莫Palermo = BPalermo巴勒莫
	f.巴勒莫Palermo.SetParent(f)
	
	f.佩特拉利亚Petralia = BPetralia佩特拉利亚
	f.佩特拉利亚Petralia.SetParent(f)
	
}
