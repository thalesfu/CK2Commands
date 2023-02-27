package naxos

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type NaxosCounty interface {
    feud.County
    BAndros安德罗斯() 	feud.Barony
    BFiloti菲洛提() 	feud.Barony
    BHermoupolis埃尔穆波利() 	feud.Barony
    BKastraki卡斯特拉基() 	feud.Barony
    BMykonos米科诺斯() 	feud.Barony
    BNaxos纳克索斯() 	feud.Barony
    BSantorini圣托里尼() 	feud.Barony
    BTinos蒂诺() 	feud.Barony
}

type 纳克索斯NaxosCounty struct {
	feud.BaseCounty
	安德罗斯Andros 	feud.Barony
	菲洛提Filoti 	feud.Barony
	埃尔穆波利Hermoupolis 	feud.Barony
	卡斯特拉基Kastraki 	feud.Barony
	米科诺斯Mykonos 	feud.Barony
	纳克索斯Naxos 	feud.Barony
	圣托里尼Santorini 	feud.Barony
	蒂诺Tinos 	feud.Barony
}

func (c *纳克索斯NaxosCounty) BAndros安德罗斯() feud.Barony {
	return c.安德罗斯Andros
}
    
func (c *纳克索斯NaxosCounty) BFiloti菲洛提() feud.Barony {
	return c.菲洛提Filoti
}
    
func (c *纳克索斯NaxosCounty) BHermoupolis埃尔穆波利() feud.Barony {
	return c.埃尔穆波利Hermoupolis
}
    
func (c *纳克索斯NaxosCounty) BKastraki卡斯特拉基() feud.Barony {
	return c.卡斯特拉基Kastraki
}
    
func (c *纳克索斯NaxosCounty) BMykonos米科诺斯() feud.Barony {
	return c.米科诺斯Mykonos
}
    
func (c *纳克索斯NaxosCounty) BNaxos纳克索斯() feud.Barony {
	return c.纳克索斯Naxos
}
    
func (c *纳克索斯NaxosCounty) BSantorini圣托里尼() feud.Barony {
	return c.圣托里尼Santorini
}
    
func (c *纳克索斯NaxosCounty) BTinos蒂诺() feud.Barony {
	return c.蒂诺Tinos
}
    
var CNaxos纳克索斯 NaxosCounty = &纳克索斯NaxosCounty{}

func init() {
	f := CNaxos纳克索斯.(*纳克索斯NaxosCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "484",
		Title:     "naxos",
		TitleName: "纳克索斯",
		TitleCode: "c_naxos",
		Baronies:  map[string]feud.Barony{},
	}

	f.安德罗斯Andros = BAndros安德罗斯
	f.安德罗斯Andros.SetParent(f)
	
	f.菲洛提Filoti = BFiloti菲洛提
	f.菲洛提Filoti.SetParent(f)
	
	f.埃尔穆波利Hermoupolis = BHermoupolis埃尔穆波利
	f.埃尔穆波利Hermoupolis.SetParent(f)
	
	f.卡斯特拉基Kastraki = BKastraki卡斯特拉基
	f.卡斯特拉基Kastraki.SetParent(f)
	
	f.米科诺斯Mykonos = BMykonos米科诺斯
	f.米科诺斯Mykonos.SetParent(f)
	
	f.纳克索斯Naxos = BNaxos纳克索斯
	f.纳克索斯Naxos.SetParent(f)
	
	f.圣托里尼Santorini = BSantorini圣托里尼
	f.圣托里尼Santorini.SetParent(f)
	
	f.蒂诺Tinos = BTinos蒂诺
	f.蒂诺Tinos.SetParent(f)
	
}
