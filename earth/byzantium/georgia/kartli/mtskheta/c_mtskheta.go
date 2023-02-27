package mtskheta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MtskhetaCounty interface {
    feud.County
    BBebristsikhe贝布里斯齐赫() 	feud.Barony
    BDjvari德瓦力() 	feud.Barony
    BGori哥里() 	feud.Barony
    BMtskheta姆茨赫塔() 	feud.Barony
    BSamtavisi萨姆塔维西() 	feud.Barony
    BSioni锡奥尼() 	feud.Barony
    BUrbnisi乌尔布尼西() 	feud.Barony
}

type 姆茨赫塔MtskhetaCounty struct {
	feud.BaseCounty
	贝布里斯齐赫Bebristsikhe 	feud.Barony
	德瓦力Djvari 	feud.Barony
	哥里Gori 	feud.Barony
	姆茨赫塔Mtskheta 	feud.Barony
	萨姆塔维西Samtavisi 	feud.Barony
	锡奥尼Sioni 	feud.Barony
	乌尔布尼西Urbnisi 	feud.Barony
}

func (c *姆茨赫塔MtskhetaCounty) BBebristsikhe贝布里斯齐赫() feud.Barony {
	return c.贝布里斯齐赫Bebristsikhe
}
    
func (c *姆茨赫塔MtskhetaCounty) BDjvari德瓦力() feud.Barony {
	return c.德瓦力Djvari
}
    
func (c *姆茨赫塔MtskhetaCounty) BGori哥里() feud.Barony {
	return c.哥里Gori
}
    
func (c *姆茨赫塔MtskhetaCounty) BMtskheta姆茨赫塔() feud.Barony {
	return c.姆茨赫塔Mtskheta
}
    
func (c *姆茨赫塔MtskhetaCounty) BSamtavisi萨姆塔维西() feud.Barony {
	return c.萨姆塔维西Samtavisi
}
    
func (c *姆茨赫塔MtskhetaCounty) BSioni锡奥尼() feud.Barony {
	return c.锡奥尼Sioni
}
    
func (c *姆茨赫塔MtskhetaCounty) BUrbnisi乌尔布尼西() feud.Barony {
	return c.乌尔布尼西Urbnisi
}
    
var CMtskheta姆茨赫塔 MtskhetaCounty = &姆茨赫塔MtskhetaCounty{}

func init() {
	f := CMtskheta姆茨赫塔.(*姆茨赫塔MtskhetaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1808",
		Title:     "mtskheta",
		TitleName: "姆茨赫塔",
		TitleCode: "c_mtskheta",
		Baronies:  map[string]feud.Barony{},
	}

	f.贝布里斯齐赫Bebristsikhe = BBebristsikhe贝布里斯齐赫
	f.贝布里斯齐赫Bebristsikhe.SetParent(f)
	
	f.德瓦力Djvari = BDjvari德瓦力
	f.德瓦力Djvari.SetParent(f)
	
	f.哥里Gori = BGori哥里
	f.哥里Gori.SetParent(f)
	
	f.姆茨赫塔Mtskheta = BMtskheta姆茨赫塔
	f.姆茨赫塔Mtskheta.SetParent(f)
	
	f.萨姆塔维西Samtavisi = BSamtavisi萨姆塔维西
	f.萨姆塔维西Samtavisi.SetParent(f)
	
	f.锡奥尼Sioni = BSioni锡奥尼
	f.锡奥尼Sioni.SetParent(f)
	
	f.乌尔布尼西Urbnisi = BUrbnisi乌尔布尼西
	f.乌尔布尼西Urbnisi.SetParent(f)
	
}
