package khaylam

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KhaylamCounty interface {
    feud.County
    BAkhsikath西鞬() 	feud.Barony
    BArdalankath阿达兰卡() 	feud.Barony
    BKasan渴塞() 	feud.Barony
    BKhaylam海拉姆() 	feud.Barony
    BMiyanrudhan米岩鲁汉() 	feud.Barony
    BNajm纳吉姆() 	feud.Barony
    BWankath万卡思() 	feud.Barony
}

type 海拉姆KhaylamCounty struct {
	feud.BaseCounty
	西鞬Akhsikath 	feud.Barony
	阿达兰卡Ardalankath 	feud.Barony
	渴塞Kasan 	feud.Barony
	海拉姆Khaylam 	feud.Barony
	米岩鲁汉Miyanrudhan 	feud.Barony
	纳吉姆Najm 	feud.Barony
	万卡思Wankath 	feud.Barony
}

func (c *海拉姆KhaylamCounty) BAkhsikath西鞬() feud.Barony {
	return c.西鞬Akhsikath
}
    
func (c *海拉姆KhaylamCounty) BArdalankath阿达兰卡() feud.Barony {
	return c.阿达兰卡Ardalankath
}
    
func (c *海拉姆KhaylamCounty) BKasan渴塞() feud.Barony {
	return c.渴塞Kasan
}
    
func (c *海拉姆KhaylamCounty) BKhaylam海拉姆() feud.Barony {
	return c.海拉姆Khaylam
}
    
func (c *海拉姆KhaylamCounty) BMiyanrudhan米岩鲁汉() feud.Barony {
	return c.米岩鲁汉Miyanrudhan
}
    
func (c *海拉姆KhaylamCounty) BNajm纳吉姆() feud.Barony {
	return c.纳吉姆Najm
}
    
func (c *海拉姆KhaylamCounty) BWankath万卡思() feud.Barony {
	return c.万卡思Wankath
}
    
var CKhaylam海拉姆 KhaylamCounty = &海拉姆KhaylamCounty{}

func init() {
	f := CKhaylam海拉姆.(*海拉姆KhaylamCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1548",
		Title:     "khaylam",
		TitleName: "海拉姆",
		TitleCode: "c_khaylam",
		Baronies:  map[string]feud.Barony{},
	}

	f.西鞬Akhsikath = BAkhsikath西鞬
	f.西鞬Akhsikath.SetParent(f)
	
	f.阿达兰卡Ardalankath = BArdalankath阿达兰卡
	f.阿达兰卡Ardalankath.SetParent(f)
	
	f.渴塞Kasan = BKasan渴塞
	f.渴塞Kasan.SetParent(f)
	
	f.海拉姆Khaylam = BKhaylam海拉姆
	f.海拉姆Khaylam.SetParent(f)
	
	f.米岩鲁汉Miyanrudhan = BMiyanrudhan米岩鲁汉
	f.米岩鲁汉Miyanrudhan.SetParent(f)
	
	f.纳吉姆Najm = BNajm纳吉姆
	f.纳吉姆Najm.SetParent(f)
	
	f.万卡思Wankath = BWankath万卡思
	f.万卡思Wankath.SetParent(f)
	
}
