package cherson

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ChersonCounty interface {
    feud.County
    BCembalo琴巴洛() 	feud.Barony
    BCharax哈拉克斯() 	feud.Barony
    BDoros多罗斯() 	feud.Barony
    BKalamita卡拉米塔() 	feud.Barony
    BKerkinitis凯尔基尼提斯() 	feud.Barony
    BKherson赫尔松() 	feud.Barony
    BNeapol尼亚波尔() 	feud.Barony
    BSevastoupolis塞瓦斯图波利斯() 	feud.Barony
}

type 赫尔松ChersonCounty struct {
	feud.BaseCounty
	琴巴洛Cembalo 	feud.Barony
	哈拉克斯Charax 	feud.Barony
	多罗斯Doros 	feud.Barony
	卡拉米塔Kalamita 	feud.Barony
	凯尔基尼提斯Kerkinitis 	feud.Barony
	赫尔松Kherson 	feud.Barony
	尼亚波尔Neapol 	feud.Barony
	塞瓦斯图波利斯Sevastoupolis 	feud.Barony
}

func (c *赫尔松ChersonCounty) BCembalo琴巴洛() feud.Barony {
	return c.琴巴洛Cembalo
}
    
func (c *赫尔松ChersonCounty) BCharax哈拉克斯() feud.Barony {
	return c.哈拉克斯Charax
}
    
func (c *赫尔松ChersonCounty) BDoros多罗斯() feud.Barony {
	return c.多罗斯Doros
}
    
func (c *赫尔松ChersonCounty) BKalamita卡拉米塔() feud.Barony {
	return c.卡拉米塔Kalamita
}
    
func (c *赫尔松ChersonCounty) BKerkinitis凯尔基尼提斯() feud.Barony {
	return c.凯尔基尼提斯Kerkinitis
}
    
func (c *赫尔松ChersonCounty) BKherson赫尔松() feud.Barony {
	return c.赫尔松Kherson
}
    
func (c *赫尔松ChersonCounty) BNeapol尼亚波尔() feud.Barony {
	return c.尼亚波尔Neapol
}
    
func (c *赫尔松ChersonCounty) BSevastoupolis塞瓦斯图波利斯() feud.Barony {
	return c.塞瓦斯图波利斯Sevastoupolis
}
    
var CCherson赫尔松 ChersonCounty = &赫尔松ChersonCounty{}

func init() {
	f := CCherson赫尔松.(*赫尔松ChersonCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "560",
		Title:     "cherson",
		TitleName: "赫尔松",
		TitleCode: "c_cherson",
		Baronies:  map[string]feud.Barony{},
	}

	f.琴巴洛Cembalo = BCembalo琴巴洛
	f.琴巴洛Cembalo.SetParent(f)
	
	f.哈拉克斯Charax = BCharax哈拉克斯
	f.哈拉克斯Charax.SetParent(f)
	
	f.多罗斯Doros = BDoros多罗斯
	f.多罗斯Doros.SetParent(f)
	
	f.卡拉米塔Kalamita = BKalamita卡拉米塔
	f.卡拉米塔Kalamita.SetParent(f)
	
	f.凯尔基尼提斯Kerkinitis = BKerkinitis凯尔基尼提斯
	f.凯尔基尼提斯Kerkinitis.SetParent(f)
	
	f.赫尔松Kherson = BKherson赫尔松
	f.赫尔松Kherson.SetParent(f)
	
	f.尼亚波尔Neapol = BNeapol尼亚波尔
	f.尼亚波尔Neapol.SetParent(f)
	
	f.塞瓦斯图波利斯Sevastoupolis = BSevastoupolis塞瓦斯图波利斯
	f.塞瓦斯图波利斯Sevastoupolis.SetParent(f)
	
}
