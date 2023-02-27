package khlynov

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KhlynovCounty interface {
    feud.County
    BArsk阿尔斯克() 	feud.Barony
    BDubyazy杜布亚济() 	feud.Barony
    BKabachischi卡巴奇希() 	feud.Barony
    BKhlynov赫雷诺夫() 	feud.Barony
    BKukmor库克莫尔() 	feud.Barony
    BUchel乌切尔() 	feud.Barony
    BVoljsk沃尔日斯克() 	feud.Barony
}

type 赫雷诺夫KhlynovCounty struct {
	feud.BaseCounty
	阿尔斯克Arsk 	feud.Barony
	杜布亚济Dubyazy 	feud.Barony
	卡巴奇希Kabachischi 	feud.Barony
	赫雷诺夫Khlynov 	feud.Barony
	库克莫尔Kukmor 	feud.Barony
	乌切尔Uchel 	feud.Barony
	沃尔日斯克Voljsk 	feud.Barony
}

func (c *赫雷诺夫KhlynovCounty) BArsk阿尔斯克() feud.Barony {
	return c.阿尔斯克Arsk
}
    
func (c *赫雷诺夫KhlynovCounty) BDubyazy杜布亚济() feud.Barony {
	return c.杜布亚济Dubyazy
}
    
func (c *赫雷诺夫KhlynovCounty) BKabachischi卡巴奇希() feud.Barony {
	return c.卡巴奇希Kabachischi
}
    
func (c *赫雷诺夫KhlynovCounty) BKhlynov赫雷诺夫() feud.Barony {
	return c.赫雷诺夫Khlynov
}
    
func (c *赫雷诺夫KhlynovCounty) BKukmor库克莫尔() feud.Barony {
	return c.库克莫尔Kukmor
}
    
func (c *赫雷诺夫KhlynovCounty) BUchel乌切尔() feud.Barony {
	return c.乌切尔Uchel
}
    
func (c *赫雷诺夫KhlynovCounty) BVoljsk沃尔日斯克() feud.Barony {
	return c.沃尔日斯克Voljsk
}
    
var CKhlynov赫雷诺夫 KhlynovCounty = &赫雷诺夫KhlynovCounty{}

func init() {
	f := CKhlynov赫雷诺夫.(*赫雷诺夫KhlynovCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1718",
		Title:     "khlynov",
		TitleName: "赫雷诺夫",
		TitleCode: "c_khlynov",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿尔斯克Arsk = BArsk阿尔斯克
	f.阿尔斯克Arsk.SetParent(f)
	
	f.杜布亚济Dubyazy = BDubyazy杜布亚济
	f.杜布亚济Dubyazy.SetParent(f)
	
	f.卡巴奇希Kabachischi = BKabachischi卡巴奇希
	f.卡巴奇希Kabachischi.SetParent(f)
	
	f.赫雷诺夫Khlynov = BKhlynov赫雷诺夫
	f.赫雷诺夫Khlynov.SetParent(f)
	
	f.库克莫尔Kukmor = BKukmor库克莫尔
	f.库克莫尔Kukmor.SetParent(f)
	
	f.乌切尔Uchel = BUchel乌切尔
	f.乌切尔Uchel.SetParent(f)
	
	f.沃尔日斯克Voljsk = BVoljsk沃尔日斯克
	f.沃尔日斯克Voljsk.SetParent(f)
	
}
