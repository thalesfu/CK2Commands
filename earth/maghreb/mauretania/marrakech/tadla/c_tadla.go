package tadla

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TadlaCounty interface {
    feud.County
    BAfza阿夫扎() 	feud.Barony
    BAit_itab艾特伊塔布() 	feud.Barony
    BAit_iyat艾特伊亚特() 	feud.Barony
    BDai达伊() 	feud.Barony
    BTadla塔德拉() 	feud.Barony
    BWawmana瓦乌马纳() 	feud.Barony
    BWazeqqur维泽古尔() 	feud.Barony
}

type 塔德莱TadlaCounty struct {
	feud.BaseCounty
	阿夫扎Afza 	feud.Barony
	艾特伊塔布Ait_itab 	feud.Barony
	艾特伊亚特Ait_iyat 	feud.Barony
	达伊Dai 	feud.Barony
	塔德拉Tadla 	feud.Barony
	瓦乌马纳Wawmana 	feud.Barony
	维泽古尔Wazeqqur 	feud.Barony
}

func (c *塔德莱TadlaCounty) BAfza阿夫扎() feud.Barony {
	return c.阿夫扎Afza
}
    
func (c *塔德莱TadlaCounty) BAit_itab艾特伊塔布() feud.Barony {
	return c.艾特伊塔布Ait_itab
}
    
func (c *塔德莱TadlaCounty) BAit_iyat艾特伊亚特() feud.Barony {
	return c.艾特伊亚特Ait_iyat
}
    
func (c *塔德莱TadlaCounty) BDai达伊() feud.Barony {
	return c.达伊Dai
}
    
func (c *塔德莱TadlaCounty) BTadla塔德拉() feud.Barony {
	return c.塔德拉Tadla
}
    
func (c *塔德莱TadlaCounty) BWawmana瓦乌马纳() feud.Barony {
	return c.瓦乌马纳Wawmana
}
    
func (c *塔德莱TadlaCounty) BWazeqqur维泽古尔() feud.Barony {
	return c.维泽古尔Wazeqqur
}
    
var CTadla塔德莱 TadlaCounty = &塔德莱TadlaCounty{}

func init() {
	f := CTadla塔德莱.(*塔德莱TadlaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1777",
		Title:     "tadla",
		TitleName: "塔德莱",
		TitleCode: "c_tadla",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿夫扎Afza = BAfza阿夫扎
	f.阿夫扎Afza.SetParent(f)
	
	f.艾特伊塔布Ait_itab = BAit_itab艾特伊塔布
	f.艾特伊塔布Ait_itab.SetParent(f)
	
	f.艾特伊亚特Ait_iyat = BAit_iyat艾特伊亚特
	f.艾特伊亚特Ait_iyat.SetParent(f)
	
	f.达伊Dai = BDai达伊
	f.达伊Dai.SetParent(f)
	
	f.塔德拉Tadla = BTadla塔德拉
	f.塔德拉Tadla.SetParent(f)
	
	f.瓦乌马纳Wawmana = BWawmana瓦乌马纳
	f.瓦乌马纳Wawmana.SetParent(f)
	
	f.维泽古尔Wazeqqur = BWazeqqur维泽古尔
	f.维泽古尔Wazeqqur.SetParent(f)
	
}
