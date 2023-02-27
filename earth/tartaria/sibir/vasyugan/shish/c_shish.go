package shish

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ShishCounty interface {
    feud.County
    BAtirka阿季尔卡() 	feud.Barony
    BKnyazevka克尼亚泽夫卡() 	feud.Barony
    BSamsonovo萨姆索诺沃() 	feud.Barony
    BShish希什() 	feud.Barony
    BTevriz捷夫里兹() 	feud.Barony
    BVasiss瓦西斯() 	feud.Barony
    BVyatka维亚特卡() 	feud.Barony
}

type 希什ShishCounty struct {
	feud.BaseCounty
	阿季尔卡Atirka 	feud.Barony
	克尼亚泽夫卡Knyazevka 	feud.Barony
	萨姆索诺沃Samsonovo 	feud.Barony
	希什Shish 	feud.Barony
	捷夫里兹Tevriz 	feud.Barony
	瓦西斯Vasiss 	feud.Barony
	维亚特卡Vyatka 	feud.Barony
}

func (c *希什ShishCounty) BAtirka阿季尔卡() feud.Barony {
	return c.阿季尔卡Atirka
}
    
func (c *希什ShishCounty) BKnyazevka克尼亚泽夫卡() feud.Barony {
	return c.克尼亚泽夫卡Knyazevka
}
    
func (c *希什ShishCounty) BSamsonovo萨姆索诺沃() feud.Barony {
	return c.萨姆索诺沃Samsonovo
}
    
func (c *希什ShishCounty) BShish希什() feud.Barony {
	return c.希什Shish
}
    
func (c *希什ShishCounty) BTevriz捷夫里兹() feud.Barony {
	return c.捷夫里兹Tevriz
}
    
func (c *希什ShishCounty) BVasiss瓦西斯() feud.Barony {
	return c.瓦西斯Vasiss
}
    
func (c *希什ShishCounty) BVyatka维亚特卡() feud.Barony {
	return c.维亚特卡Vyatka
}
    
var CShish希什 ShishCounty = &希什ShishCounty{}

func init() {
	f := CShish希什.(*希什ShishCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1866",
		Title:     "shish",
		TitleName: "希什",
		TitleCode: "c_shish",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿季尔卡Atirka = BAtirka阿季尔卡
	f.阿季尔卡Atirka.SetParent(f)
	
	f.克尼亚泽夫卡Knyazevka = BKnyazevka克尼亚泽夫卡
	f.克尼亚泽夫卡Knyazevka.SetParent(f)
	
	f.萨姆索诺沃Samsonovo = BSamsonovo萨姆索诺沃
	f.萨姆索诺沃Samsonovo.SetParent(f)
	
	f.希什Shish = BShish希什
	f.希什Shish.SetParent(f)
	
	f.捷夫里兹Tevriz = BTevriz捷夫里兹
	f.捷夫里兹Tevriz.SetParent(f)
	
	f.瓦西斯Vasiss = BVasiss瓦西斯
	f.瓦西斯Vasiss.SetParent(f)
	
	f.维亚特卡Vyatka = BVyatka维亚特卡
	f.维亚特卡Vyatka.SetParent(f)
	
}
