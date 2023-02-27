package zeta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ZetaCounty interface {
    feud.County
    BBar巴尔() 	feud.Barony
    BBudva布德瓦() 	feud.Barony
    BDanj达尼() 	feud.Barony
    BDrivast德里瓦斯特() 	feud.Barony
    BPodgorica波德戈里察() 	feud.Barony
    BSkadar斯卡达尔() 	feud.Barony
    BUlcinj乌尔齐尼() 	feud.Barony
}

type 泽塔ZetaCounty struct {
	feud.BaseCounty
	巴尔Bar 	feud.Barony
	布德瓦Budva 	feud.Barony
	达尼Danj 	feud.Barony
	德里瓦斯特Drivast 	feud.Barony
	波德戈里察Podgorica 	feud.Barony
	斯卡达尔Skadar 	feud.Barony
	乌尔齐尼Ulcinj 	feud.Barony
}

func (c *泽塔ZetaCounty) BBar巴尔() feud.Barony {
	return c.巴尔Bar
}
    
func (c *泽塔ZetaCounty) BBudva布德瓦() feud.Barony {
	return c.布德瓦Budva
}
    
func (c *泽塔ZetaCounty) BDanj达尼() feud.Barony {
	return c.达尼Danj
}
    
func (c *泽塔ZetaCounty) BDrivast德里瓦斯特() feud.Barony {
	return c.德里瓦斯特Drivast
}
    
func (c *泽塔ZetaCounty) BPodgorica波德戈里察() feud.Barony {
	return c.波德戈里察Podgorica
}
    
func (c *泽塔ZetaCounty) BSkadar斯卡达尔() feud.Barony {
	return c.斯卡达尔Skadar
}
    
func (c *泽塔ZetaCounty) BUlcinj乌尔齐尼() feud.Barony {
	return c.乌尔齐尼Ulcinj
}
    
var CZeta泽塔 ZetaCounty = &泽塔ZetaCounty{}

func init() {
	f := CZeta泽塔.(*泽塔ZetaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "469",
		Title:     "zeta",
		TitleName: "泽塔",
		TitleCode: "c_zeta",
		Baronies:  map[string]feud.Barony{},
	}

	f.巴尔Bar = BBar巴尔
	f.巴尔Bar.SetParent(f)
	
	f.布德瓦Budva = BBudva布德瓦
	f.布德瓦Budva.SetParent(f)
	
	f.达尼Danj = BDanj达尼
	f.达尼Danj.SetParent(f)
	
	f.德里瓦斯特Drivast = BDrivast德里瓦斯特
	f.德里瓦斯特Drivast.SetParent(f)
	
	f.波德戈里察Podgorica = BPodgorica波德戈里察
	f.波德戈里察Podgorica.SetParent(f)
	
	f.斯卡达尔Skadar = BSkadar斯卡达尔
	f.斯卡达尔Skadar.SetParent(f)
	
	f.乌尔齐尼Ulcinj = BUlcinj乌尔齐尼
	f.乌尔齐尼Ulcinj.SetParent(f)
	
}
