package teluch

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TeluchCounty interface {
    feud.County
    BGermanias日耳曼尼亚() 	feud.Barony
    BHajin哈金() 	feud.Barony
    BKapan卡盘() 	feud.Barony
    BKoksen科克森() 	feud.Barony
    BKomanal科曼阿拉() 	feud.Barony
    BPerre佩勒() 	feud.Barony
    BTavplur塔夫普鲁尔() 	feud.Barony
    BTeluch特卢克() 	feud.Barony
}

type 特卢克TeluchCounty struct {
	feud.BaseCounty
	日耳曼尼亚Germanias 	feud.Barony
	哈金Hajin 	feud.Barony
	卡盘Kapan 	feud.Barony
	科克森Koksen 	feud.Barony
	科曼阿拉Komanal 	feud.Barony
	佩勒Perre 	feud.Barony
	塔夫普鲁尔Tavplur 	feud.Barony
	特卢克Teluch 	feud.Barony
}

func (c *特卢克TeluchCounty) BGermanias日耳曼尼亚() feud.Barony {
	return c.日耳曼尼亚Germanias
}
    
func (c *特卢克TeluchCounty) BHajin哈金() feud.Barony {
	return c.哈金Hajin
}
    
func (c *特卢克TeluchCounty) BKapan卡盘() feud.Barony {
	return c.卡盘Kapan
}
    
func (c *特卢克TeluchCounty) BKoksen科克森() feud.Barony {
	return c.科克森Koksen
}
    
func (c *特卢克TeluchCounty) BKomanal科曼阿拉() feud.Barony {
	return c.科曼阿拉Komanal
}
    
func (c *特卢克TeluchCounty) BPerre佩勒() feud.Barony {
	return c.佩勒Perre
}
    
func (c *特卢克TeluchCounty) BTavplur塔夫普鲁尔() feud.Barony {
	return c.塔夫普鲁尔Tavplur
}
    
func (c *特卢克TeluchCounty) BTeluch特卢克() feud.Barony {
	return c.特卢克Teluch
}
    
var CTeluch特卢克 TeluchCounty = &特卢克TeluchCounty{}

func init() {
	f := CTeluch特卢克.(*特卢克TeluchCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "735",
		Title:     "teluch",
		TitleName: "特卢克",
		TitleCode: "c_teluch",
		Baronies:  map[string]feud.Barony{},
	}

	f.日耳曼尼亚Germanias = BGermanias日耳曼尼亚
	f.日耳曼尼亚Germanias.SetParent(f)
	
	f.哈金Hajin = BHajin哈金
	f.哈金Hajin.SetParent(f)
	
	f.卡盘Kapan = BKapan卡盘
	f.卡盘Kapan.SetParent(f)
	
	f.科克森Koksen = BKoksen科克森
	f.科克森Koksen.SetParent(f)
	
	f.科曼阿拉Komanal = BKomanal科曼阿拉
	f.科曼阿拉Komanal.SetParent(f)
	
	f.佩勒Perre = BPerre佩勒
	f.佩勒Perre.SetParent(f)
	
	f.塔夫普鲁尔Tavplur = BTavplur塔夫普鲁尔
	f.塔夫普鲁尔Tavplur.SetParent(f)
	
	f.特卢克Teluch = BTeluch特卢克
	f.特卢克Teluch.SetParent(f)
	
}
