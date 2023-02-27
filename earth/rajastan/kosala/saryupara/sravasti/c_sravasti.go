package sravasti

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SravastiCounty interface {
    feud.County
    BAmorha阿没诃() 	feud.Barony
    BBarohiya婆卢醯耶() 	feud.Barony
    BGorakhnath_math瞿罗叉那他神庙() 	feud.Barony
    BGorakhpur瞿罗叉城() 	feud.Barony
    BJetavana祇园() 	feud.Barony
    BMaghar摩迦罗() 	feud.Barony
    BSravasti舍卫城() 	feud.Barony
}

type 舍卫城SravastiCounty struct {
	feud.BaseCounty
	阿没诃Amorha 	feud.Barony
	婆卢醯耶Barohiya 	feud.Barony
	瞿罗叉那他神庙Gorakhnath_math 	feud.Barony
	瞿罗叉城Gorakhpur 	feud.Barony
	祇园Jetavana 	feud.Barony
	摩迦罗Maghar 	feud.Barony
	舍卫城Sravasti 	feud.Barony
}

func (c *舍卫城SravastiCounty) BAmorha阿没诃() feud.Barony {
	return c.阿没诃Amorha
}
    
func (c *舍卫城SravastiCounty) BBarohiya婆卢醯耶() feud.Barony {
	return c.婆卢醯耶Barohiya
}
    
func (c *舍卫城SravastiCounty) BGorakhnath_math瞿罗叉那他神庙() feud.Barony {
	return c.瞿罗叉那他神庙Gorakhnath_math
}
    
func (c *舍卫城SravastiCounty) BGorakhpur瞿罗叉城() feud.Barony {
	return c.瞿罗叉城Gorakhpur
}
    
func (c *舍卫城SravastiCounty) BJetavana祇园() feud.Barony {
	return c.祇园Jetavana
}
    
func (c *舍卫城SravastiCounty) BMaghar摩迦罗() feud.Barony {
	return c.摩迦罗Maghar
}
    
func (c *舍卫城SravastiCounty) BSravasti舍卫城() feud.Barony {
	return c.舍卫城Sravasti
}
    
var CSravasti舍卫城 SravastiCounty = &舍卫城SravastiCounty{}

func init() {
	f := CSravasti舍卫城.(*舍卫城SravastiCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1421",
		Title:     "sravasti",
		TitleName: "舍卫城",
		TitleCode: "c_sravasti",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿没诃Amorha = BAmorha阿没诃
	f.阿没诃Amorha.SetParent(f)
	
	f.婆卢醯耶Barohiya = BBarohiya婆卢醯耶
	f.婆卢醯耶Barohiya.SetParent(f)
	
	f.瞿罗叉那他神庙Gorakhnath_math = BGorakhnath_math瞿罗叉那他神庙
	f.瞿罗叉那他神庙Gorakhnath_math.SetParent(f)
	
	f.瞿罗叉城Gorakhpur = BGorakhpur瞿罗叉城
	f.瞿罗叉城Gorakhpur.SetParent(f)
	
	f.祇园Jetavana = BJetavana祇园
	f.祇园Jetavana.SetParent(f)
	
	f.摩迦罗Maghar = BMaghar摩迦罗
	f.摩迦罗Maghar.SetParent(f)
	
	f.舍卫城Sravasti = BSravasti舍卫城
	f.舍卫城Sravasti.SetParent(f)
	
}
