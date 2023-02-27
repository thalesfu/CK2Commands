package cadota

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type CadotaCounty interface {
    feud.County
    BEndere安得悦() 	feud.Barony
    BGuqintang古琴塘() 	feud.Barony
    BHuangyangshan黄杨山() 	feud.Barony
    BNiya尼雅() 	feud.Barony
    BWogou窝沟() 	feud.Barony
    BXinjiancun辛间村() 	feud.Barony
    BZengwujing曾屋迳() 	feud.Barony
}

type 精绝CadotaCounty struct {
	feud.BaseCounty
	安得悦Endere 	feud.Barony
	古琴塘Guqintang 	feud.Barony
	黄杨山Huangyangshan 	feud.Barony
	尼雅Niya 	feud.Barony
	窝沟Wogou 	feud.Barony
	辛间村Xinjiancun 	feud.Barony
	曾屋迳Zengwujing 	feud.Barony
}

func (c *精绝CadotaCounty) BEndere安得悦() feud.Barony {
	return c.安得悦Endere
}
    
func (c *精绝CadotaCounty) BGuqintang古琴塘() feud.Barony {
	return c.古琴塘Guqintang
}
    
func (c *精绝CadotaCounty) BHuangyangshan黄杨山() feud.Barony {
	return c.黄杨山Huangyangshan
}
    
func (c *精绝CadotaCounty) BNiya尼雅() feud.Barony {
	return c.尼雅Niya
}
    
func (c *精绝CadotaCounty) BWogou窝沟() feud.Barony {
	return c.窝沟Wogou
}
    
func (c *精绝CadotaCounty) BXinjiancun辛间村() feud.Barony {
	return c.辛间村Xinjiancun
}
    
func (c *精绝CadotaCounty) BZengwujing曾屋迳() feud.Barony {
	return c.曾屋迳Zengwujing
}
    
var CCadota精绝 CadotaCounty = &精绝CadotaCounty{}

func init() {
	f := CCadota精绝.(*精绝CadotaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1522",
		Title:     "cadota",
		TitleName: "精绝",
		TitleCode: "c_cadota",
		Baronies:  map[string]feud.Barony{},
	}

	f.安得悦Endere = BEndere安得悦
	f.安得悦Endere.SetParent(f)
	
	f.古琴塘Guqintang = BGuqintang古琴塘
	f.古琴塘Guqintang.SetParent(f)
	
	f.黄杨山Huangyangshan = BHuangyangshan黄杨山
	f.黄杨山Huangyangshan.SetParent(f)
	
	f.尼雅Niya = BNiya尼雅
	f.尼雅Niya.SetParent(f)
	
	f.窝沟Wogou = BWogou窝沟
	f.窝沟Wogou.SetParent(f)
	
	f.辛间村Xinjiancun = BXinjiancun辛间村
	f.辛间村Xinjiancun.SetParent(f)
	
	f.曾屋迳Zengwujing = BZengwujing曾屋迳
	f.曾屋迳Zengwujing.SetParent(f)
	
}
