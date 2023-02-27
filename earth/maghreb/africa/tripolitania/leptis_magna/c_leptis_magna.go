package leptis_magna

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type Leptis_magnaCounty interface {
    feud.County
    BAbunujaym阿布努贾() 	feud.Barony
    BAlghiran艾格赫拉() 	feud.Barony
    BAlhasun艾哈苏() 	feud.Barony
    BGioda吉欧德() 	feud.Barony
    BLeptismagna大莱波蒂斯() 	feud.Barony
    BMisratah米苏拉塔() 	feud.Barony
    BNaimah奈伊迈() 	feud.Barony
    BQirdah吉拉达赫() 	feud.Barony
}

type 大莱波蒂斯Leptis_magnaCounty struct {
	feud.BaseCounty
	阿布努贾Abunujaym 	feud.Barony
	艾格赫拉Alghiran 	feud.Barony
	艾哈苏Alhasun 	feud.Barony
	吉欧德Gioda 	feud.Barony
	大莱波蒂斯Leptismagna 	feud.Barony
	米苏拉塔Misratah 	feud.Barony
	奈伊迈Naimah 	feud.Barony
	吉拉达赫Qirdah 	feud.Barony
}

func (c *大莱波蒂斯Leptis_magnaCounty) BAbunujaym阿布努贾() feud.Barony {
	return c.阿布努贾Abunujaym
}
    
func (c *大莱波蒂斯Leptis_magnaCounty) BAlghiran艾格赫拉() feud.Barony {
	return c.艾格赫拉Alghiran
}
    
func (c *大莱波蒂斯Leptis_magnaCounty) BAlhasun艾哈苏() feud.Barony {
	return c.艾哈苏Alhasun
}
    
func (c *大莱波蒂斯Leptis_magnaCounty) BGioda吉欧德() feud.Barony {
	return c.吉欧德Gioda
}
    
func (c *大莱波蒂斯Leptis_magnaCounty) BLeptismagna大莱波蒂斯() feud.Barony {
	return c.大莱波蒂斯Leptismagna
}
    
func (c *大莱波蒂斯Leptis_magnaCounty) BMisratah米苏拉塔() feud.Barony {
	return c.米苏拉塔Misratah
}
    
func (c *大莱波蒂斯Leptis_magnaCounty) BNaimah奈伊迈() feud.Barony {
	return c.奈伊迈Naimah
}
    
func (c *大莱波蒂斯Leptis_magnaCounty) BQirdah吉拉达赫() feud.Barony {
	return c.吉拉达赫Qirdah
}
    
var CLeptis_magna大莱波蒂斯 Leptis_magnaCounty = &大莱波蒂斯Leptis_magnaCounty{}

func init() {
	f := CLeptis_magna大莱波蒂斯.(*大莱波蒂斯Leptis_magnaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "810",
		Title:     "leptis_magna",
		TitleName: "大莱波蒂斯",
		TitleCode: "c_leptis_magna",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿布努贾Abunujaym = BAbunujaym阿布努贾
	f.阿布努贾Abunujaym.SetParent(f)
	
	f.艾格赫拉Alghiran = BAlghiran艾格赫拉
	f.艾格赫拉Alghiran.SetParent(f)
	
	f.艾哈苏Alhasun = BAlhasun艾哈苏
	f.艾哈苏Alhasun.SetParent(f)
	
	f.吉欧德Gioda = BGioda吉欧德
	f.吉欧德Gioda.SetParent(f)
	
	f.大莱波蒂斯Leptismagna = BLeptismagna大莱波蒂斯
	f.大莱波蒂斯Leptismagna.SetParent(f)
	
	f.米苏拉塔Misratah = BMisratah米苏拉塔
	f.米苏拉塔Misratah.SetParent(f)
	
	f.奈伊迈Naimah = BNaimah奈伊迈
	f.奈伊迈Naimah.SetParent(f)
	
	f.吉拉达赫Qirdah = BQirdah吉拉达赫
	f.吉拉达赫Qirdah.SetParent(f)
	
}
