package naberezhnye

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type NaberezhnyeCounty interface {
    feud.County
    BAlmetyevsk阿尔梅季耶夫斯克() 	feud.Barony
    BAznakayevo阿兹纳卡耶沃() 	feud.Barony
    BBugulma布古利马() 	feud.Barony
    BChatyr_tau恰特尔_陶() 	feud.Barony
    BNaberezhnye纳别列日内耶() 	feud.Barony
    BNizhnekamsk下卡姆斯克() 	feud.Barony
    BTuymazy图伊马济() 	feud.Barony
}

type 纳别列日内耶NaberezhnyeCounty struct {
	feud.BaseCounty
	阿尔梅季耶夫斯克Almetyevsk 	feud.Barony
	阿兹纳卡耶沃Aznakayevo 	feud.Barony
	布古利马Bugulma 	feud.Barony
	恰特尔_陶Chatyr_tau 	feud.Barony
	纳别列日内耶Naberezhnye 	feud.Barony
	下卡姆斯克Nizhnekamsk 	feud.Barony
	图伊马济Tuymazy 	feud.Barony
}

func (c *纳别列日内耶NaberezhnyeCounty) BAlmetyevsk阿尔梅季耶夫斯克() feud.Barony {
	return c.阿尔梅季耶夫斯克Almetyevsk
}
    
func (c *纳别列日内耶NaberezhnyeCounty) BAznakayevo阿兹纳卡耶沃() feud.Barony {
	return c.阿兹纳卡耶沃Aznakayevo
}
    
func (c *纳别列日内耶NaberezhnyeCounty) BBugulma布古利马() feud.Barony {
	return c.布古利马Bugulma
}
    
func (c *纳别列日内耶NaberezhnyeCounty) BChatyr_tau恰特尔_陶() feud.Barony {
	return c.恰特尔_陶Chatyr_tau
}
    
func (c *纳别列日内耶NaberezhnyeCounty) BNaberezhnye纳别列日内耶() feud.Barony {
	return c.纳别列日内耶Naberezhnye
}
    
func (c *纳别列日内耶NaberezhnyeCounty) BNizhnekamsk下卡姆斯克() feud.Barony {
	return c.下卡姆斯克Nizhnekamsk
}
    
func (c *纳别列日内耶NaberezhnyeCounty) BTuymazy图伊马济() feud.Barony {
	return c.图伊马济Tuymazy
}
    
var CNaberezhnye纳别列日内耶 NaberezhnyeCounty = &纳别列日内耶NaberezhnyeCounty{}

func init() {
	f := CNaberezhnye纳别列日内耶.(*纳别列日内耶NaberezhnyeCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "436",
		Title:     "naberezhnye",
		TitleName: "纳别列日内耶",
		TitleCode: "c_naberezhnye",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿尔梅季耶夫斯克Almetyevsk = BAlmetyevsk阿尔梅季耶夫斯克
	f.阿尔梅季耶夫斯克Almetyevsk.SetParent(f)
	
	f.阿兹纳卡耶沃Aznakayevo = BAznakayevo阿兹纳卡耶沃
	f.阿兹纳卡耶沃Aznakayevo.SetParent(f)
	
	f.布古利马Bugulma = BBugulma布古利马
	f.布古利马Bugulma.SetParent(f)
	
	f.恰特尔_陶Chatyr_tau = BChatyr_tau恰特尔_陶
	f.恰特尔_陶Chatyr_tau.SetParent(f)
	
	f.纳别列日内耶Naberezhnye = BNaberezhnye纳别列日内耶
	f.纳别列日内耶Naberezhnye.SetParent(f)
	
	f.下卡姆斯克Nizhnekamsk = BNizhnekamsk下卡姆斯克
	f.下卡姆斯克Nizhnekamsk.SetParent(f)
	
	f.图伊马济Tuymazy = BTuymazy图伊马济
	f.图伊马济Tuymazy.SetParent(f)
	
}
