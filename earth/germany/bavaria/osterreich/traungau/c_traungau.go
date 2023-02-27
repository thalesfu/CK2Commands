package traungau

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TraungauCounty interface {
    feud.County
    BEbensee埃本塞() 	feud.Barony
    BEferding埃费丁() 	feud.Barony
    BEnisiburg埃尼西堡() 	feud.Barony
    BGmunden格蒙登() 	feud.Barony
    BHallstatt哈尔施塔特() 	feud.Barony
    BLinz林茨() 	feud.Barony
    BPons_veckelahe蓬斯韦克拉赫() 	feud.Barony
    BWels韦尔斯() 	feud.Barony
}

type 特劳恩高TraungauCounty struct {
	feud.BaseCounty
	埃本塞Ebensee 	feud.Barony
	埃费丁Eferding 	feud.Barony
	埃尼西堡Enisiburg 	feud.Barony
	格蒙登Gmunden 	feud.Barony
	哈尔施塔特Hallstatt 	feud.Barony
	林茨Linz 	feud.Barony
	蓬斯韦克拉赫Pons_veckelahe 	feud.Barony
	韦尔斯Wels 	feud.Barony
}

func (c *特劳恩高TraungauCounty) BEbensee埃本塞() feud.Barony {
	return c.埃本塞Ebensee
}
    
func (c *特劳恩高TraungauCounty) BEferding埃费丁() feud.Barony {
	return c.埃费丁Eferding
}
    
func (c *特劳恩高TraungauCounty) BEnisiburg埃尼西堡() feud.Barony {
	return c.埃尼西堡Enisiburg
}
    
func (c *特劳恩高TraungauCounty) BGmunden格蒙登() feud.Barony {
	return c.格蒙登Gmunden
}
    
func (c *特劳恩高TraungauCounty) BHallstatt哈尔施塔特() feud.Barony {
	return c.哈尔施塔特Hallstatt
}
    
func (c *特劳恩高TraungauCounty) BLinz林茨() feud.Barony {
	return c.林茨Linz
}
    
func (c *特劳恩高TraungauCounty) BPons_veckelahe蓬斯韦克拉赫() feud.Barony {
	return c.蓬斯韦克拉赫Pons_veckelahe
}
    
func (c *特劳恩高TraungauCounty) BWels韦尔斯() feud.Barony {
	return c.韦尔斯Wels
}
    
var CTraungau特劳恩高 TraungauCounty = &特劳恩高TraungauCounty{}

func init() {
	f := CTraungau特劳恩高.(*特劳恩高TraungauCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1690",
		Title:     "traungau",
		TitleName: "特劳恩高",
		TitleCode: "c_traungau",
		Baronies:  map[string]feud.Barony{},
	}

	f.埃本塞Ebensee = BEbensee埃本塞
	f.埃本塞Ebensee.SetParent(f)
	
	f.埃费丁Eferding = BEferding埃费丁
	f.埃费丁Eferding.SetParent(f)
	
	f.埃尼西堡Enisiburg = BEnisiburg埃尼西堡
	f.埃尼西堡Enisiburg.SetParent(f)
	
	f.格蒙登Gmunden = BGmunden格蒙登
	f.格蒙登Gmunden.SetParent(f)
	
	f.哈尔施塔特Hallstatt = BHallstatt哈尔施塔特
	f.哈尔施塔特Hallstatt.SetParent(f)
	
	f.林茨Linz = BLinz林茨
	f.林茨Linz.SetParent(f)
	
	f.蓬斯韦克拉赫Pons_veckelahe = BPons_veckelahe蓬斯韦克拉赫
	f.蓬斯韦克拉赫Pons_veckelahe.SetParent(f)
	
	f.韦尔斯Wels = BWels韦尔斯
	f.韦尔斯Wels.SetParent(f)
	
}
