package bilino_polje

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type Bilino_poljeCounty interface {
    feud.County
    BGorazde戈拉日代() 	feud.Barony
    BKonjic科尼茨() 	feud.Barony
    BKresevo克雷舍沃() 	feud.Barony
    BProzor普罗佐尔() 	feud.Barony
    BVisegrad维舍格勒() 	feud.Barony
    BVisoki维索基() 	feud.Barony
    BVrhbosna弗尔赫波斯纳() 	feud.Barony
}

type 弗尔赫波斯纳Bilino_poljeCounty struct {
	feud.BaseCounty
	戈拉日代Gorazde 	feud.Barony
	科尼茨Konjic 	feud.Barony
	克雷舍沃Kresevo 	feud.Barony
	普罗佐尔Prozor 	feud.Barony
	维舍格勒Visegrad 	feud.Barony
	维索基Visoki 	feud.Barony
	弗尔赫波斯纳Vrhbosna 	feud.Barony
}

func (c *弗尔赫波斯纳Bilino_poljeCounty) BGorazde戈拉日代() feud.Barony {
	return c.戈拉日代Gorazde
}
    
func (c *弗尔赫波斯纳Bilino_poljeCounty) BKonjic科尼茨() feud.Barony {
	return c.科尼茨Konjic
}
    
func (c *弗尔赫波斯纳Bilino_poljeCounty) BKresevo克雷舍沃() feud.Barony {
	return c.克雷舍沃Kresevo
}
    
func (c *弗尔赫波斯纳Bilino_poljeCounty) BProzor普罗佐尔() feud.Barony {
	return c.普罗佐尔Prozor
}
    
func (c *弗尔赫波斯纳Bilino_poljeCounty) BVisegrad维舍格勒() feud.Barony {
	return c.维舍格勒Visegrad
}
    
func (c *弗尔赫波斯纳Bilino_poljeCounty) BVisoki维索基() feud.Barony {
	return c.维索基Visoki
}
    
func (c *弗尔赫波斯纳Bilino_poljeCounty) BVrhbosna弗尔赫波斯纳() feud.Barony {
	return c.弗尔赫波斯纳Vrhbosna
}
    
var CBilino_polje弗尔赫波斯纳 Bilino_poljeCounty = &弗尔赫波斯纳Bilino_poljeCounty{}

func init() {
	f := CBilino_polje弗尔赫波斯纳.(*弗尔赫波斯纳Bilino_poljeCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1971",
		Title:     "bilino_polje",
		TitleName: "弗尔赫波斯纳",
		TitleCode: "c_bilino_polje",
		Baronies:  map[string]feud.Barony{},
	}

	f.戈拉日代Gorazde = BGorazde戈拉日代
	f.戈拉日代Gorazde.SetParent(f)
	
	f.科尼茨Konjic = BKonjic科尼茨
	f.科尼茨Konjic.SetParent(f)
	
	f.克雷舍沃Kresevo = BKresevo克雷舍沃
	f.克雷舍沃Kresevo.SetParent(f)
	
	f.普罗佐尔Prozor = BProzor普罗佐尔
	f.普罗佐尔Prozor.SetParent(f)
	
	f.维舍格勒Visegrad = BVisegrad维舍格勒
	f.维舍格勒Visegrad.SetParent(f)
	
	f.维索基Visoki = BVisoki维索基
	f.维索基Visoki.SetParent(f)
	
	f.弗尔赫波斯纳Vrhbosna = BVrhbosna弗尔赫波斯纳
	f.弗尔赫波斯纳Vrhbosna.SetParent(f)
	
}
