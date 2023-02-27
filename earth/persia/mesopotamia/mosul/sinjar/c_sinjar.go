package sinjar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SinjarCounty interface {
    feud.County
    BBaaj巴贾() 	feud.Barony
    BHatra哈特拉() 	feud.Barony
    BKouyunik阔玉尼克() 	feud.Barony
    BNabiyunus纳比尤努斯() 	feud.Barony
    BNineveh尼尼微() 	feud.Barony
    BSinjar辛贾尔() 	feud.Barony
    BTalkayf泰勒凯夫() 	feud.Barony
    BTelassar特拉萨() 	feud.Barony
}

type 辛贾尔SinjarCounty struct {
	feud.BaseCounty
	巴贾Baaj 	feud.Barony
	哈特拉Hatra 	feud.Barony
	阔玉尼克Kouyunik 	feud.Barony
	纳比尤努斯Nabiyunus 	feud.Barony
	尼尼微Nineveh 	feud.Barony
	辛贾尔Sinjar 	feud.Barony
	泰勒凯夫Talkayf 	feud.Barony
	特拉萨Telassar 	feud.Barony
}

func (c *辛贾尔SinjarCounty) BBaaj巴贾() feud.Barony {
	return c.巴贾Baaj
}
    
func (c *辛贾尔SinjarCounty) BHatra哈特拉() feud.Barony {
	return c.哈特拉Hatra
}
    
func (c *辛贾尔SinjarCounty) BKouyunik阔玉尼克() feud.Barony {
	return c.阔玉尼克Kouyunik
}
    
func (c *辛贾尔SinjarCounty) BNabiyunus纳比尤努斯() feud.Barony {
	return c.纳比尤努斯Nabiyunus
}
    
func (c *辛贾尔SinjarCounty) BNineveh尼尼微() feud.Barony {
	return c.尼尼微Nineveh
}
    
func (c *辛贾尔SinjarCounty) BSinjar辛贾尔() feud.Barony {
	return c.辛贾尔Sinjar
}
    
func (c *辛贾尔SinjarCounty) BTalkayf泰勒凯夫() feud.Barony {
	return c.泰勒凯夫Talkayf
}
    
func (c *辛贾尔SinjarCounty) BTelassar特拉萨() feud.Barony {
	return c.特拉萨Telassar
}
    
var CSinjar辛贾尔 SinjarCounty = &辛贾尔SinjarCounty{}

func init() {
	f := CSinjar辛贾尔.(*辛贾尔SinjarCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "711",
		Title:     "sinjar",
		TitleName: "辛贾尔",
		TitleCode: "c_sinjar",
		Baronies:  map[string]feud.Barony{},
	}

	f.巴贾Baaj = BBaaj巴贾
	f.巴贾Baaj.SetParent(f)
	
	f.哈特拉Hatra = BHatra哈特拉
	f.哈特拉Hatra.SetParent(f)
	
	f.阔玉尼克Kouyunik = BKouyunik阔玉尼克
	f.阔玉尼克Kouyunik.SetParent(f)
	
	f.纳比尤努斯Nabiyunus = BNabiyunus纳比尤努斯
	f.纳比尤努斯Nabiyunus.SetParent(f)
	
	f.尼尼微Nineveh = BNineveh尼尼微
	f.尼尼微Nineveh.SetParent(f)
	
	f.辛贾尔Sinjar = BSinjar辛贾尔
	f.辛贾尔Sinjar.SetParent(f)
	
	f.泰勒凯夫Talkayf = BTalkayf泰勒凯夫
	f.泰勒凯夫Talkayf.SetParent(f)
	
	f.特拉萨Telassar = BTelassar特拉萨
	f.特拉萨Telassar.SetParent(f)
	
}
