package zurichgau

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ZurichgauCounty interface {
    feud.County
    BAarau阿劳() 	feud.Barony
    BKappel卡珀尔() 	feud.Barony
    BKyburg基堡() 	feud.Barony
    BRapperswil拉珀斯维尔() 	feud.Barony
    BSwiss_baden巴登() 	feud.Barony
    BToggenburg吐根堡() 	feud.Barony
    BWinterthur温特图尔() 	feud.Barony
    BZug楚格() 	feud.Barony
    BZurich苏黎世() 	feud.Barony
}

type 苏黎世ZurichgauCounty struct {
	feud.BaseCounty
	阿劳Aarau 	feud.Barony
	卡珀尔Kappel 	feud.Barony
	基堡Kyburg 	feud.Barony
	拉珀斯维尔Rapperswil 	feud.Barony
	巴登Swiss_baden 	feud.Barony
	吐根堡Toggenburg 	feud.Barony
	温特图尔Winterthur 	feud.Barony
	楚格Zug 	feud.Barony
	苏黎世Zurich 	feud.Barony
}

func (c *苏黎世ZurichgauCounty) BAarau阿劳() feud.Barony {
	return c.阿劳Aarau
}
    
func (c *苏黎世ZurichgauCounty) BKappel卡珀尔() feud.Barony {
	return c.卡珀尔Kappel
}
    
func (c *苏黎世ZurichgauCounty) BKyburg基堡() feud.Barony {
	return c.基堡Kyburg
}
    
func (c *苏黎世ZurichgauCounty) BRapperswil拉珀斯维尔() feud.Barony {
	return c.拉珀斯维尔Rapperswil
}
    
func (c *苏黎世ZurichgauCounty) BSwiss_baden巴登() feud.Barony {
	return c.巴登Swiss_baden
}
    
func (c *苏黎世ZurichgauCounty) BToggenburg吐根堡() feud.Barony {
	return c.吐根堡Toggenburg
}
    
func (c *苏黎世ZurichgauCounty) BWinterthur温特图尔() feud.Barony {
	return c.温特图尔Winterthur
}
    
func (c *苏黎世ZurichgauCounty) BZug楚格() feud.Barony {
	return c.楚格Zug
}
    
func (c *苏黎世ZurichgauCounty) BZurich苏黎世() feud.Barony {
	return c.苏黎世Zurich
}
    
var CZurichgau苏黎世 ZurichgauCounty = &苏黎世ZurichgauCounty{}

func init() {
	f := CZurichgau苏黎世.(*苏黎世ZurichgauCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "245",
		Title:     "zurichgau",
		TitleName: "苏黎世",
		TitleCode: "c_zurichgau",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿劳Aarau = BAarau阿劳
	f.阿劳Aarau.SetParent(f)
	
	f.卡珀尔Kappel = BKappel卡珀尔
	f.卡珀尔Kappel.SetParent(f)
	
	f.基堡Kyburg = BKyburg基堡
	f.基堡Kyburg.SetParent(f)
	
	f.拉珀斯维尔Rapperswil = BRapperswil拉珀斯维尔
	f.拉珀斯维尔Rapperswil.SetParent(f)
	
	f.巴登Swiss_baden = BSwiss_baden巴登
	f.巴登Swiss_baden.SetParent(f)
	
	f.吐根堡Toggenburg = BToggenburg吐根堡
	f.吐根堡Toggenburg.SetParent(f)
	
	f.温特图尔Winterthur = BWinterthur温特图尔
	f.温特图尔Winterthur.SetParent(f)
	
	f.楚格Zug = BZug楚格
	f.楚格Zug.SetParent(f)
	
	f.苏黎世Zurich = BZurich苏黎世
	f.苏黎世Zurich.SetParent(f)
	
}
