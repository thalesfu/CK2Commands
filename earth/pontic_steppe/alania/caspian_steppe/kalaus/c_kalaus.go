package kalaus

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KalausCounty interface {
    feud.County
    BArmavir阿尔马维尔() 	feud.Barony
    BBazovyy巴佐维() 	feud.Barony
    BDemino杰米诺() 	feud.Barony
    BKalaus卡劳斯() 	feud.Barony
    BKugulta库古利塔() 	feud.Barony
    BRozliv罗兹利夫() 	feud.Barony
    BYankul扬库尔() 	feud.Barony
}

type 卡劳斯KalausCounty struct {
	feud.BaseCounty
	阿尔马维尔Armavir 	feud.Barony
	巴佐维Bazovyy 	feud.Barony
	杰米诺Demino 	feud.Barony
	卡劳斯Kalaus 	feud.Barony
	库古利塔Kugulta 	feud.Barony
	罗兹利夫Rozliv 	feud.Barony
	扬库尔Yankul 	feud.Barony
}

func (c *卡劳斯KalausCounty) BArmavir阿尔马维尔() feud.Barony {
	return c.阿尔马维尔Armavir
}
    
func (c *卡劳斯KalausCounty) BBazovyy巴佐维() feud.Barony {
	return c.巴佐维Bazovyy
}
    
func (c *卡劳斯KalausCounty) BDemino杰米诺() feud.Barony {
	return c.杰米诺Demino
}
    
func (c *卡劳斯KalausCounty) BKalaus卡劳斯() feud.Barony {
	return c.卡劳斯Kalaus
}
    
func (c *卡劳斯KalausCounty) BKugulta库古利塔() feud.Barony {
	return c.库古利塔Kugulta
}
    
func (c *卡劳斯KalausCounty) BRozliv罗兹利夫() feud.Barony {
	return c.罗兹利夫Rozliv
}
    
func (c *卡劳斯KalausCounty) BYankul扬库尔() feud.Barony {
	return c.扬库尔Yankul
}
    
var CKalaus卡劳斯 KalausCounty = &卡劳斯KalausCounty{}

func init() {
	f := CKalaus卡劳斯.(*卡劳斯KalausCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1811",
		Title:     "kalaus",
		TitleName: "卡劳斯",
		TitleCode: "c_kalaus",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿尔马维尔Armavir = BArmavir阿尔马维尔
	f.阿尔马维尔Armavir.SetParent(f)
	
	f.巴佐维Bazovyy = BBazovyy巴佐维
	f.巴佐维Bazovyy.SetParent(f)
	
	f.杰米诺Demino = BDemino杰米诺
	f.杰米诺Demino.SetParent(f)
	
	f.卡劳斯Kalaus = BKalaus卡劳斯
	f.卡劳斯Kalaus.SetParent(f)
	
	f.库古利塔Kugulta = BKugulta库古利塔
	f.库古利塔Kugulta.SetParent(f)
	
	f.罗兹利夫Rozliv = BRozliv罗兹利夫
	f.罗兹利夫Rozliv.SetParent(f)
	
	f.扬库尔Yankul = BYankul扬库尔
	f.扬库尔Yankul.SetParent(f)
	
}
