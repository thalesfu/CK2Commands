package aukshayts

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AukshaytsCounty interface {
    feud.County
    BAlsenai阿尔塞奈() 	feud.Barony
    BAsmena阿斯梅纳() 	feud.Barony
    BEliskes埃利斯克斯() 	feud.Barony
    BKreva克列沃() 	feud.Barony
    BMedininkai梅迪宁凯() 	feud.Barony
    BTrakai特拉凯() 	feud.Barony
    BVilnius维尔纽斯() 	feud.Barony
}

type 维尔纽斯AukshaytsCounty struct {
	feud.BaseCounty
	阿尔塞奈Alsenai 	feud.Barony
	阿斯梅纳Asmena 	feud.Barony
	埃利斯克斯Eliskes 	feud.Barony
	克列沃Kreva 	feud.Barony
	梅迪宁凯Medininkai 	feud.Barony
	特拉凯Trakai 	feud.Barony
	维尔纽斯Vilnius 	feud.Barony
}

func (c *维尔纽斯AukshaytsCounty) BAlsenai阿尔塞奈() feud.Barony {
	return c.阿尔塞奈Alsenai
}
    
func (c *维尔纽斯AukshaytsCounty) BAsmena阿斯梅纳() feud.Barony {
	return c.阿斯梅纳Asmena
}
    
func (c *维尔纽斯AukshaytsCounty) BEliskes埃利斯克斯() feud.Barony {
	return c.埃利斯克斯Eliskes
}
    
func (c *维尔纽斯AukshaytsCounty) BKreva克列沃() feud.Barony {
	return c.克列沃Kreva
}
    
func (c *维尔纽斯AukshaytsCounty) BMedininkai梅迪宁凯() feud.Barony {
	return c.梅迪宁凯Medininkai
}
    
func (c *维尔纽斯AukshaytsCounty) BTrakai特拉凯() feud.Barony {
	return c.特拉凯Trakai
}
    
func (c *维尔纽斯AukshaytsCounty) BVilnius维尔纽斯() feud.Barony {
	return c.维尔纽斯Vilnius
}
    
var CAukshayts维尔纽斯 AukshaytsCounty = &维尔纽斯AukshaytsCounty{}

func init() {
	f := CAukshayts维尔纽斯.(*维尔纽斯AukshaytsCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "420",
		Title:     "aukshayts",
		TitleName: "维尔纽斯",
		TitleCode: "c_aukshayts",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿尔塞奈Alsenai = BAlsenai阿尔塞奈
	f.阿尔塞奈Alsenai.SetParent(f)
	
	f.阿斯梅纳Asmena = BAsmena阿斯梅纳
	f.阿斯梅纳Asmena.SetParent(f)
	
	f.埃利斯克斯Eliskes = BEliskes埃利斯克斯
	f.埃利斯克斯Eliskes.SetParent(f)
	
	f.克列沃Kreva = BKreva克列沃
	f.克列沃Kreva.SetParent(f)
	
	f.梅迪宁凯Medininkai = BMedininkai梅迪宁凯
	f.梅迪宁凯Medininkai.SetParent(f)
	
	f.特拉凯Trakai = BTrakai特拉凯
	f.特拉凯Trakai.SetParent(f)
	
	f.维尔纽斯Vilnius = BVilnius维尔纽斯
	f.维尔纽斯Vilnius.SetParent(f)
	
}
