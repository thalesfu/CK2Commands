package menorca

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MenorcaCounty interface {
    feud.County
    BAlaior阿莱奥尔() 	feud.Barony
    BCiutadella休塔德利亚() 	feud.Barony
    BEscastell埃斯卡斯特利() 	feud.Barony
    BEsmercadal埃斯梅卡达尔() 	feud.Barony
    BFerreries费雷列斯() 	feud.Barony
    BMahon马奥() 	feud.Barony
    BSantaagueda圣阿格达() 	feud.Barony
    BSantlluis圣路易斯() 	feud.Barony
}

type 梅诺卡MenorcaCounty struct {
	feud.BaseCounty
	阿莱奥尔Alaior 	feud.Barony
	休塔德利亚Ciutadella 	feud.Barony
	埃斯卡斯特利Escastell 	feud.Barony
	埃斯梅卡达尔Esmercadal 	feud.Barony
	费雷列斯Ferreries 	feud.Barony
	马奥Mahon 	feud.Barony
	圣阿格达Santaagueda 	feud.Barony
	圣路易斯Santlluis 	feud.Barony
}

func (c *梅诺卡MenorcaCounty) BAlaior阿莱奥尔() feud.Barony {
	return c.阿莱奥尔Alaior
}
    
func (c *梅诺卡MenorcaCounty) BCiutadella休塔德利亚() feud.Barony {
	return c.休塔德利亚Ciutadella
}
    
func (c *梅诺卡MenorcaCounty) BEscastell埃斯卡斯特利() feud.Barony {
	return c.埃斯卡斯特利Escastell
}
    
func (c *梅诺卡MenorcaCounty) BEsmercadal埃斯梅卡达尔() feud.Barony {
	return c.埃斯梅卡达尔Esmercadal
}
    
func (c *梅诺卡MenorcaCounty) BFerreries费雷列斯() feud.Barony {
	return c.费雷列斯Ferreries
}
    
func (c *梅诺卡MenorcaCounty) BMahon马奥() feud.Barony {
	return c.马奥Mahon
}
    
func (c *梅诺卡MenorcaCounty) BSantaagueda圣阿格达() feud.Barony {
	return c.圣阿格达Santaagueda
}
    
func (c *梅诺卡MenorcaCounty) BSantlluis圣路易斯() feud.Barony {
	return c.圣路易斯Santlluis
}
    
var CMenorca梅诺卡 MenorcaCounty = &梅诺卡MenorcaCounty{}

func init() {
	f := CMenorca梅诺卡.(*梅诺卡MenorcaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "826",
		Title:     "menorca",
		TitleName: "梅诺卡",
		TitleCode: "c_menorca",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿莱奥尔Alaior = BAlaior阿莱奥尔
	f.阿莱奥尔Alaior.SetParent(f)
	
	f.休塔德利亚Ciutadella = BCiutadella休塔德利亚
	f.休塔德利亚Ciutadella.SetParent(f)
	
	f.埃斯卡斯特利Escastell = BEscastell埃斯卡斯特利
	f.埃斯卡斯特利Escastell.SetParent(f)
	
	f.埃斯梅卡达尔Esmercadal = BEsmercadal埃斯梅卡达尔
	f.埃斯梅卡达尔Esmercadal.SetParent(f)
	
	f.费雷列斯Ferreries = BFerreries费雷列斯
	f.费雷列斯Ferreries.SetParent(f)
	
	f.马奥Mahon = BMahon马奥
	f.马奥Mahon.SetParent(f)
	
	f.圣阿格达Santaagueda = BSantaagueda圣阿格达
	f.圣阿格达Santaagueda.SetParent(f)
	
	f.圣路易斯Santlluis = BSantlluis圣路易斯
	f.圣路易斯Santlluis.SetParent(f)
	
}
