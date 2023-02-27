package balkhash

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BalkhashCounty interface {
    feud.County
    BAkshatau阿克沙套() 	feud.Barony
    BAktogay阿克托盖() 	feud.Barony
    BBalkhash巴尔喀什() 	feud.Barony
    BGulshat古尔沙特() 	feud.Barony
    BKaragandy卡拉干达() 	feud.Barony
    BKarazhal卡拉扎尔() 	feud.Barony
    BShashubay沙舒拜() 	feud.Barony
}

type 巴尔喀什BalkhashCounty struct {
	feud.BaseCounty
	阿克沙套Akshatau 	feud.Barony
	阿克托盖Aktogay 	feud.Barony
	巴尔喀什Balkhash 	feud.Barony
	古尔沙特Gulshat 	feud.Barony
	卡拉干达Karagandy 	feud.Barony
	卡拉扎尔Karazhal 	feud.Barony
	沙舒拜Shashubay 	feud.Barony
}

func (c *巴尔喀什BalkhashCounty) BAkshatau阿克沙套() feud.Barony {
	return c.阿克沙套Akshatau
}
    
func (c *巴尔喀什BalkhashCounty) BAktogay阿克托盖() feud.Barony {
	return c.阿克托盖Aktogay
}
    
func (c *巴尔喀什BalkhashCounty) BBalkhash巴尔喀什() feud.Barony {
	return c.巴尔喀什Balkhash
}
    
func (c *巴尔喀什BalkhashCounty) BGulshat古尔沙特() feud.Barony {
	return c.古尔沙特Gulshat
}
    
func (c *巴尔喀什BalkhashCounty) BKaragandy卡拉干达() feud.Barony {
	return c.卡拉干达Karagandy
}
    
func (c *巴尔喀什BalkhashCounty) BKarazhal卡拉扎尔() feud.Barony {
	return c.卡拉扎尔Karazhal
}
    
func (c *巴尔喀什BalkhashCounty) BShashubay沙舒拜() feud.Barony {
	return c.沙舒拜Shashubay
}
    
var CBalkhash巴尔喀什 BalkhashCounty = &巴尔喀什BalkhashCounty{}

func init() {
	f := CBalkhash巴尔喀什.(*巴尔喀什BalkhashCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1434",
		Title:     "balkhash",
		TitleName: "巴尔喀什",
		TitleCode: "c_balkhash",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿克沙套Akshatau = BAkshatau阿克沙套
	f.阿克沙套Akshatau.SetParent(f)
	
	f.阿克托盖Aktogay = BAktogay阿克托盖
	f.阿克托盖Aktogay.SetParent(f)
	
	f.巴尔喀什Balkhash = BBalkhash巴尔喀什
	f.巴尔喀什Balkhash.SetParent(f)
	
	f.古尔沙特Gulshat = BGulshat古尔沙特
	f.古尔沙特Gulshat.SetParent(f)
	
	f.卡拉干达Karagandy = BKaragandy卡拉干达
	f.卡拉干达Karagandy.SetParent(f)
	
	f.卡拉扎尔Karazhal = BKarazhal卡拉扎尔
	f.卡拉扎尔Karazhal.SetParent(f)
	
	f.沙舒拜Shashubay = BShashubay沙舒拜
	f.沙舒拜Shashubay.SetParent(f)
	
}
