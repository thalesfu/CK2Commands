package jacwiez

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type JacwiezCounty interface {
    feud.County
    BGrodno格罗德诺() 	feud.Barony
    BIuje伊维耶() 	feud.Barony
    BJacwiez亚奇维日() 	feud.Barony
    BMir密尔() 	feud.Barony
    BNiasvizh涅斯维日() 	feud.Barony
    BNovogrudok新格鲁多克() 	feud.Barony
    BSkidziel斯基达利() 	feud.Barony
    BZirmuny日尔穆内() 	feud.Barony
}

type 格罗德诺JacwiezCounty struct {
	feud.BaseCounty
	格罗德诺Grodno 	feud.Barony
	伊维耶Iuje 	feud.Barony
	亚奇维日Jacwiez 	feud.Barony
	密尔Mir 	feud.Barony
	涅斯维日Niasvizh 	feud.Barony
	新格鲁多克Novogrudok 	feud.Barony
	斯基达利Skidziel 	feud.Barony
	日尔穆内Zirmuny 	feud.Barony
}

func (c *格罗德诺JacwiezCounty) BGrodno格罗德诺() feud.Barony {
	return c.格罗德诺Grodno
}
    
func (c *格罗德诺JacwiezCounty) BIuje伊维耶() feud.Barony {
	return c.伊维耶Iuje
}
    
func (c *格罗德诺JacwiezCounty) BJacwiez亚奇维日() feud.Barony {
	return c.亚奇维日Jacwiez
}
    
func (c *格罗德诺JacwiezCounty) BMir密尔() feud.Barony {
	return c.密尔Mir
}
    
func (c *格罗德诺JacwiezCounty) BNiasvizh涅斯维日() feud.Barony {
	return c.涅斯维日Niasvizh
}
    
func (c *格罗德诺JacwiezCounty) BNovogrudok新格鲁多克() feud.Barony {
	return c.新格鲁多克Novogrudok
}
    
func (c *格罗德诺JacwiezCounty) BSkidziel斯基达利() feud.Barony {
	return c.斯基达利Skidziel
}
    
func (c *格罗德诺JacwiezCounty) BZirmuny日尔穆内() feud.Barony {
	return c.日尔穆内Zirmuny
}
    
var CJacwiez格罗德诺 JacwiezCounty = &格罗德诺JacwiezCounty{}

func init() {
	f := CJacwiez格罗德诺.(*格罗德诺JacwiezCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "424",
		Title:     "jacwiez",
		TitleName: "格罗德诺",
		TitleCode: "c_jacwiez",
		Baronies:  map[string]feud.Barony{},
	}

	f.格罗德诺Grodno = BGrodno格罗德诺
	f.格罗德诺Grodno.SetParent(f)
	
	f.伊维耶Iuje = BIuje伊维耶
	f.伊维耶Iuje.SetParent(f)
	
	f.亚奇维日Jacwiez = BJacwiez亚奇维日
	f.亚奇维日Jacwiez.SetParent(f)
	
	f.密尔Mir = BMir密尔
	f.密尔Mir.SetParent(f)
	
	f.涅斯维日Niasvizh = BNiasvizh涅斯维日
	f.涅斯维日Niasvizh.SetParent(f)
	
	f.新格鲁多克Novogrudok = BNovogrudok新格鲁多克
	f.新格鲁多克Novogrudok.SetParent(f)
	
	f.斯基达利Skidziel = BSkidziel斯基达利
	f.斯基达利Skidziel.SetParent(f)
	
	f.日尔穆内Zirmuny = BZirmuny日尔穆内
	f.日尔穆内Zirmuny.SetParent(f)
	
}
