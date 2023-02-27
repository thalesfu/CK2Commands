package ile_de_france

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type Ile_de_franceCounty interface {
    feud.County
    BCorbeil科尔贝() 	feud.Barony
    BEtampes埃唐普() 	feud.Barony
    BMelun默伦() 	feud.Barony
    BMontfortlamaury蒙福尔拉莫里() 	feud.Barony
    BMontlhery蒙莱里() 	feud.Barony
    BParis巴黎() 	feud.Barony
    BStdenis圣但尼() 	feud.Barony
}

type 法兰西岛Ile_de_franceCounty struct {
	feud.BaseCounty
	科尔贝Corbeil 	feud.Barony
	埃唐普Etampes 	feud.Barony
	默伦Melun 	feud.Barony
	蒙福尔拉莫里Montfortlamaury 	feud.Barony
	蒙莱里Montlhery 	feud.Barony
	巴黎Paris 	feud.Barony
	圣但尼Stdenis 	feud.Barony
}

func (c *法兰西岛Ile_de_franceCounty) BCorbeil科尔贝() feud.Barony {
	return c.科尔贝Corbeil
}
    
func (c *法兰西岛Ile_de_franceCounty) BEtampes埃唐普() feud.Barony {
	return c.埃唐普Etampes
}
    
func (c *法兰西岛Ile_de_franceCounty) BMelun默伦() feud.Barony {
	return c.默伦Melun
}
    
func (c *法兰西岛Ile_de_franceCounty) BMontfortlamaury蒙福尔拉莫里() feud.Barony {
	return c.蒙福尔拉莫里Montfortlamaury
}
    
func (c *法兰西岛Ile_de_franceCounty) BMontlhery蒙莱里() feud.Barony {
	return c.蒙莱里Montlhery
}
    
func (c *法兰西岛Ile_de_franceCounty) BParis巴黎() feud.Barony {
	return c.巴黎Paris
}
    
func (c *法兰西岛Ile_de_franceCounty) BStdenis圣但尼() feud.Barony {
	return c.圣但尼Stdenis
}
    
var CIle_de_france法兰西岛 Ile_de_franceCounty = &法兰西岛Ile_de_franceCounty{}

func init() {
	f := CIle_de_france法兰西岛.(*法兰西岛Ile_de_franceCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "112",
		Title:     "ile_de_france",
		TitleName: "法兰西岛",
		TitleCode: "c_ile_de_france",
		Baronies:  map[string]feud.Barony{},
	}

	f.科尔贝Corbeil = BCorbeil科尔贝
	f.科尔贝Corbeil.SetParent(f)
	
	f.埃唐普Etampes = BEtampes埃唐普
	f.埃唐普Etampes.SetParent(f)
	
	f.默伦Melun = BMelun默伦
	f.默伦Melun.SetParent(f)
	
	f.蒙福尔拉莫里Montfortlamaury = BMontfortlamaury蒙福尔拉莫里
	f.蒙福尔拉莫里Montfortlamaury.SetParent(f)
	
	f.蒙莱里Montlhery = BMontlhery蒙莱里
	f.蒙莱里Montlhery.SetParent(f)
	
	f.巴黎Paris = BParis巴黎
	f.巴黎Paris.SetParent(f)
	
	f.圣但尼Stdenis = BStdenis圣但尼
	f.圣但尼Stdenis.SetParent(f)
	
}
