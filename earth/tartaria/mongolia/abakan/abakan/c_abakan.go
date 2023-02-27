package abakan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AbakanCounty interface {
    feud.County
    BAbakan阿巴坎() 	feud.Barony
    BBerenzhak别连扎克() 	feud.Barony
    BBirikchul比里克丘尔() 	feud.Barony
    BBiskamzha比斯卡姆扎() 	feud.Barony
    BKommunar科穆纳尔() 	feud.Barony
    BShira希拉() 	feud.Barony
    BTuim图伊姆() 	feud.Barony
}

type 阿巴坎AbakanCounty struct {
	feud.BaseCounty
	阿巴坎Abakan 	feud.Barony
	别连扎克Berenzhak 	feud.Barony
	比里克丘尔Birikchul 	feud.Barony
	比斯卡姆扎Biskamzha 	feud.Barony
	科穆纳尔Kommunar 	feud.Barony
	希拉Shira 	feud.Barony
	图伊姆Tuim 	feud.Barony
}

func (c *阿巴坎AbakanCounty) BAbakan阿巴坎() feud.Barony {
	return c.阿巴坎Abakan
}
    
func (c *阿巴坎AbakanCounty) BBerenzhak别连扎克() feud.Barony {
	return c.别连扎克Berenzhak
}
    
func (c *阿巴坎AbakanCounty) BBirikchul比里克丘尔() feud.Barony {
	return c.比里克丘尔Birikchul
}
    
func (c *阿巴坎AbakanCounty) BBiskamzha比斯卡姆扎() feud.Barony {
	return c.比斯卡姆扎Biskamzha
}
    
func (c *阿巴坎AbakanCounty) BKommunar科穆纳尔() feud.Barony {
	return c.科穆纳尔Kommunar
}
    
func (c *阿巴坎AbakanCounty) BShira希拉() feud.Barony {
	return c.希拉Shira
}
    
func (c *阿巴坎AbakanCounty) BTuim图伊姆() feud.Barony {
	return c.图伊姆Tuim
}
    
var CAbakan阿巴坎 AbakanCounty = &阿巴坎AbakanCounty{}

func init() {
	f := CAbakan阿巴坎.(*阿巴坎AbakanCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1899",
		Title:     "abakan",
		TitleName: "阿巴坎",
		TitleCode: "c_abakan",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿巴坎Abakan = BAbakan阿巴坎
	f.阿巴坎Abakan.SetParent(f)
	
	f.别连扎克Berenzhak = BBerenzhak别连扎克
	f.别连扎克Berenzhak.SetParent(f)
	
	f.比里克丘尔Birikchul = BBirikchul比里克丘尔
	f.比里克丘尔Birikchul.SetParent(f)
	
	f.比斯卡姆扎Biskamzha = BBiskamzha比斯卡姆扎
	f.比斯卡姆扎Biskamzha.SetParent(f)
	
	f.科穆纳尔Kommunar = BKommunar科穆纳尔
	f.科穆纳尔Kommunar.SetParent(f)
	
	f.希拉Shira = BShira希拉
	f.希拉Shira.SetParent(f)
	
	f.图伊姆Tuim = BTuim图伊姆
	f.图伊姆Tuim.SetParent(f)
	
}
