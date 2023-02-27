package uzens

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type UzensCounty interface {
    feud.County
    BAkkus阿克库斯() 	feud.Barony
    BAytbay艾特拜() 	feud.Barony
    BKaztal卡兹塔勒() 	feud.Barony
    BKhasan哈桑() 	feud.Barony
    BKosoba科索巴() 	feud.Barony
    BKrykuru克雷克_乌鲁() 	feud.Barony
    BKyzylasker克孜勒_阿斯克尔() 	feud.Barony
    BUshkuduk乌什库杜克() 	feud.Barony
}

type 乌津UzensCounty struct {
	feud.BaseCounty
	阿克库斯Akkus 	feud.Barony
	艾特拜Aytbay 	feud.Barony
	卡兹塔勒Kaztal 	feud.Barony
	哈桑Khasan 	feud.Barony
	科索巴Kosoba 	feud.Barony
	克雷克_乌鲁Krykuru 	feud.Barony
	克孜勒_阿斯克尔Kyzylasker 	feud.Barony
	乌什库杜克Ushkuduk 	feud.Barony
}

func (c *乌津UzensCounty) BAkkus阿克库斯() feud.Barony {
	return c.阿克库斯Akkus
}
    
func (c *乌津UzensCounty) BAytbay艾特拜() feud.Barony {
	return c.艾特拜Aytbay
}
    
func (c *乌津UzensCounty) BKaztal卡兹塔勒() feud.Barony {
	return c.卡兹塔勒Kaztal
}
    
func (c *乌津UzensCounty) BKhasan哈桑() feud.Barony {
	return c.哈桑Khasan
}
    
func (c *乌津UzensCounty) BKosoba科索巴() feud.Barony {
	return c.科索巴Kosoba
}
    
func (c *乌津UzensCounty) BKrykuru克雷克_乌鲁() feud.Barony {
	return c.克雷克_乌鲁Krykuru
}
    
func (c *乌津UzensCounty) BKyzylasker克孜勒_阿斯克尔() feud.Barony {
	return c.克孜勒_阿斯克尔Kyzylasker
}
    
func (c *乌津UzensCounty) BUshkuduk乌什库杜克() feud.Barony {
	return c.乌什库杜克Ushkuduk
}
    
var CUzens乌津 UzensCounty = &乌津UzensCounty{}

func init() {
	f := CUzens乌津.(*乌津UzensCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "617",
		Title:     "uzens",
		TitleName: "乌津",
		TitleCode: "c_uzens",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿克库斯Akkus = BAkkus阿克库斯
	f.阿克库斯Akkus.SetParent(f)
	
	f.艾特拜Aytbay = BAytbay艾特拜
	f.艾特拜Aytbay.SetParent(f)
	
	f.卡兹塔勒Kaztal = BKaztal卡兹塔勒
	f.卡兹塔勒Kaztal.SetParent(f)
	
	f.哈桑Khasan = BKhasan哈桑
	f.哈桑Khasan.SetParent(f)
	
	f.科索巴Kosoba = BKosoba科索巴
	f.科索巴Kosoba.SetParent(f)
	
	f.克雷克_乌鲁Krykuru = BKrykuru克雷克_乌鲁
	f.克雷克_乌鲁Krykuru.SetParent(f)
	
	f.克孜勒_阿斯克尔Kyzylasker = BKyzylasker克孜勒_阿斯克尔
	f.克孜勒_阿斯克尔Kyzylasker.SetParent(f)
	
	f.乌什库杜克Ushkuduk = BUshkuduk乌什库杜克
	f.乌什库杜克Ushkuduk.SetParent(f)
	
}
