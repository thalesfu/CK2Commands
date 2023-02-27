package dege

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type DegeCounty interface {
    feud.County
    BBabang八邦() 	feud.Barony
    BBaiya白垭() 	feud.Barony
    BDama达马() 	feud.Barony
    BDege德格() 	feud.Barony
    BDzongsar宗萨() 	feud.Barony
    BGongya龚垭() 	feud.Barony
    BPalpung八邦() 	feud.Barony
}

type 德格DegeCounty struct {
	feud.BaseCounty
	八邦Babang 	feud.Barony
	白垭Baiya 	feud.Barony
	达马Dama 	feud.Barony
	德格Dege 	feud.Barony
	宗萨Dzongsar 	feud.Barony
	龚垭Gongya 	feud.Barony
	八邦Palpung 	feud.Barony
}

func (c *德格DegeCounty) BBabang八邦() feud.Barony {
	return c.八邦Babang
}
    
func (c *德格DegeCounty) BBaiya白垭() feud.Barony {
	return c.白垭Baiya
}
    
func (c *德格DegeCounty) BDama达马() feud.Barony {
	return c.达马Dama
}
    
func (c *德格DegeCounty) BDege德格() feud.Barony {
	return c.德格Dege
}
    
func (c *德格DegeCounty) BDzongsar宗萨() feud.Barony {
	return c.宗萨Dzongsar
}
    
func (c *德格DegeCounty) BGongya龚垭() feud.Barony {
	return c.龚垭Gongya
}
    
func (c *德格DegeCounty) BPalpung八邦() feud.Barony {
	return c.八邦Palpung
}
    
var CDege德格 DegeCounty = &德格DegeCounty{}

func init() {
	f := CDege德格.(*德格DegeCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1506",
		Title:     "dege",
		TitleName: "德格",
		TitleCode: "c_dege",
		Baronies:  map[string]feud.Barony{},
	}

	f.八邦Babang = BBabang八邦
	f.八邦Babang.SetParent(f)
	
	f.白垭Baiya = BBaiya白垭
	f.白垭Baiya.SetParent(f)
	
	f.达马Dama = BDama达马
	f.达马Dama.SetParent(f)
	
	f.德格Dege = BDege德格
	f.德格Dege.SetParent(f)
	
	f.宗萨Dzongsar = BDzongsar宗萨
	f.宗萨Dzongsar.SetParent(f)
	
	f.龚垭Gongya = BGongya龚垭
	f.龚垭Gongya.SetParent(f)
	
	f.八邦Palpung = BPalpung八邦
	f.八邦Palpung.SetParent(f)
	
}
