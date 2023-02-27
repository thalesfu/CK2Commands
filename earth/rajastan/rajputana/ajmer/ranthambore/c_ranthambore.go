package ranthambore

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type RanthamboreCounty interface {
    feud.County
    BGangapur恒伽补罗() 	feud.Barony
    BMandrael曼荼罗耶罗() 	feud.Barony
    BRanthambore罗那萨担婆补罗() 	feud.Barony
    BSendamarani森陀摩罗尼() 	feud.Barony
    BSheopur谢奥补罗() 	feud.Barony
    BTonkra通卡() 	feud.Barony
    BUtgir乌德吉尔() 	feud.Barony
}

type 罗那萨担婆补罗RanthamboreCounty struct {
	feud.BaseCounty
	恒伽补罗Gangapur 	feud.Barony
	曼荼罗耶罗Mandrael 	feud.Barony
	罗那萨担婆补罗Ranthambore 	feud.Barony
	森陀摩罗尼Sendamarani 	feud.Barony
	谢奥补罗Sheopur 	feud.Barony
	通卡Tonkra 	feud.Barony
	乌德吉尔Utgir 	feud.Barony
}

func (c *罗那萨担婆补罗RanthamboreCounty) BGangapur恒伽补罗() feud.Barony {
	return c.恒伽补罗Gangapur
}
    
func (c *罗那萨担婆补罗RanthamboreCounty) BMandrael曼荼罗耶罗() feud.Barony {
	return c.曼荼罗耶罗Mandrael
}
    
func (c *罗那萨担婆补罗RanthamboreCounty) BRanthambore罗那萨担婆补罗() feud.Barony {
	return c.罗那萨担婆补罗Ranthambore
}
    
func (c *罗那萨担婆补罗RanthamboreCounty) BSendamarani森陀摩罗尼() feud.Barony {
	return c.森陀摩罗尼Sendamarani
}
    
func (c *罗那萨担婆补罗RanthamboreCounty) BSheopur谢奥补罗() feud.Barony {
	return c.谢奥补罗Sheopur
}
    
func (c *罗那萨担婆补罗RanthamboreCounty) BTonkra通卡() feud.Barony {
	return c.通卡Tonkra
}
    
func (c *罗那萨担婆补罗RanthamboreCounty) BUtgir乌德吉尔() feud.Barony {
	return c.乌德吉尔Utgir
}
    
var CRanthambore罗那萨担婆补罗 RanthamboreCounty = &罗那萨担婆补罗RanthamboreCounty{}

func init() {
	f := CRanthambore罗那萨担婆补罗.(*罗那萨担婆补罗RanthamboreCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1355",
		Title:     "ranthambore",
		TitleName: "罗那萨担婆补罗",
		TitleCode: "c_ranthambore",
		Baronies:  map[string]feud.Barony{},
	}

	f.恒伽补罗Gangapur = BGangapur恒伽补罗
	f.恒伽补罗Gangapur.SetParent(f)
	
	f.曼荼罗耶罗Mandrael = BMandrael曼荼罗耶罗
	f.曼荼罗耶罗Mandrael.SetParent(f)
	
	f.罗那萨担婆补罗Ranthambore = BRanthambore罗那萨担婆补罗
	f.罗那萨担婆补罗Ranthambore.SetParent(f)
	
	f.森陀摩罗尼Sendamarani = BSendamarani森陀摩罗尼
	f.森陀摩罗尼Sendamarani.SetParent(f)
	
	f.谢奥补罗Sheopur = BSheopur谢奥补罗
	f.谢奥补罗Sheopur.SetParent(f)
	
	f.通卡Tonkra = BTonkra通卡
	f.通卡Tonkra.SetParent(f)
	
	f.乌德吉尔Utgir = BUtgir乌德吉尔
	f.乌德吉尔Utgir.SetParent(f)
	
}
