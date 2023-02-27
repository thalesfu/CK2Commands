package oleshye

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type OleshyeCounty interface {
    feud.County
    BDnieprkherson赫尔松() 	feud.Barony
    BDnipro第聂伯() 	feud.Barony
    BKamiansk卡缅斯克() 	feud.Barony
    BKryvyi_rih克里维里赫() 	feud.Barony
    BMarhanets马尔加涅茨() 	feud.Barony
    BMykolaiv尼古拉耶夫() 	feud.Barony
    BNikopol尼科波尔() 	feud.Barony
}

type 下第聂伯OleshyeCounty struct {
	feud.BaseCounty
	赫尔松Dnieprkherson 	feud.Barony
	第聂伯Dnipro 	feud.Barony
	卡缅斯克Kamiansk 	feud.Barony
	克里维里赫Kryvyi_rih 	feud.Barony
	马尔加涅茨Marhanets 	feud.Barony
	尼古拉耶夫Mykolaiv 	feud.Barony
	尼科波尔Nikopol 	feud.Barony
}

func (c *下第聂伯OleshyeCounty) BDnieprkherson赫尔松() feud.Barony {
	return c.赫尔松Dnieprkherson
}
    
func (c *下第聂伯OleshyeCounty) BDnipro第聂伯() feud.Barony {
	return c.第聂伯Dnipro
}
    
func (c *下第聂伯OleshyeCounty) BKamiansk卡缅斯克() feud.Barony {
	return c.卡缅斯克Kamiansk
}
    
func (c *下第聂伯OleshyeCounty) BKryvyi_rih克里维里赫() feud.Barony {
	return c.克里维里赫Kryvyi_rih
}
    
func (c *下第聂伯OleshyeCounty) BMarhanets马尔加涅茨() feud.Barony {
	return c.马尔加涅茨Marhanets
}
    
func (c *下第聂伯OleshyeCounty) BMykolaiv尼古拉耶夫() feud.Barony {
	return c.尼古拉耶夫Mykolaiv
}
    
func (c *下第聂伯OleshyeCounty) BNikopol尼科波尔() feud.Barony {
	return c.尼科波尔Nikopol
}
    
var COleshye下第聂伯 OleshyeCounty = &下第聂伯OleshyeCounty{}

func init() {
	f := COleshye下第聂伯.(*下第聂伯OleshyeCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "543",
		Title:     "oleshye",
		TitleName: "下第聂伯",
		TitleCode: "c_oleshye",
		Baronies:  map[string]feud.Barony{},
	}

	f.赫尔松Dnieprkherson = BDnieprkherson赫尔松
	f.赫尔松Dnieprkherson.SetParent(f)
	
	f.第聂伯Dnipro = BDnipro第聂伯
	f.第聂伯Dnipro.SetParent(f)
	
	f.卡缅斯克Kamiansk = BKamiansk卡缅斯克
	f.卡缅斯克Kamiansk.SetParent(f)
	
	f.克里维里赫Kryvyi_rih = BKryvyi_rih克里维里赫
	f.克里维里赫Kryvyi_rih.SetParent(f)
	
	f.马尔加涅茨Marhanets = BMarhanets马尔加涅茨
	f.马尔加涅茨Marhanets.SetParent(f)
	
	f.尼古拉耶夫Mykolaiv = BMykolaiv尼古拉耶夫
	f.尼古拉耶夫Mykolaiv.SetParent(f)
	
	f.尼科波尔Nikopol = BNikopol尼科波尔
	f.尼科波尔Nikopol.SetParent(f)
	
}
