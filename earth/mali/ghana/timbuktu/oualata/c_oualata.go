package oualata

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type OualataCounty interface {
    feud.County
    BBirnsara比尔恩萨拉() 	feud.Barony
    BBiru比如() 	feud.Barony
    BFassekra法瑟卡拉() 	feud.Barony
    BKonou科努() 	feud.Barony
    BNeiguiat内吉亚特() 	feud.Barony
    BOualata瓦拉塔() 	feud.Barony
    BTarhalet塔尔哈雷特() 	feud.Barony
    BTizert蒂泽尔特() 	feud.Barony
}

type 瓦拉塔OualataCounty struct {
	feud.BaseCounty
	比尔恩萨拉Birnsara 	feud.Barony
	比如Biru 	feud.Barony
	法瑟卡拉Fassekra 	feud.Barony
	科努Konou 	feud.Barony
	内吉亚特Neiguiat 	feud.Barony
	瓦拉塔Oualata 	feud.Barony
	塔尔哈雷特Tarhalet 	feud.Barony
	蒂泽尔特Tizert 	feud.Barony
}

func (c *瓦拉塔OualataCounty) BBirnsara比尔恩萨拉() feud.Barony {
	return c.比尔恩萨拉Birnsara
}
    
func (c *瓦拉塔OualataCounty) BBiru比如() feud.Barony {
	return c.比如Biru
}
    
func (c *瓦拉塔OualataCounty) BFassekra法瑟卡拉() feud.Barony {
	return c.法瑟卡拉Fassekra
}
    
func (c *瓦拉塔OualataCounty) BKonou科努() feud.Barony {
	return c.科努Konou
}
    
func (c *瓦拉塔OualataCounty) BNeiguiat内吉亚特() feud.Barony {
	return c.内吉亚特Neiguiat
}
    
func (c *瓦拉塔OualataCounty) BOualata瓦拉塔() feud.Barony {
	return c.瓦拉塔Oualata
}
    
func (c *瓦拉塔OualataCounty) BTarhalet塔尔哈雷特() feud.Barony {
	return c.塔尔哈雷特Tarhalet
}
    
func (c *瓦拉塔OualataCounty) BTizert蒂泽尔特() feud.Barony {
	return c.蒂泽尔特Tizert
}
    
var COualata瓦拉塔 OualataCounty = &瓦拉塔OualataCounty{}

func init() {
	f := COualata瓦拉塔.(*瓦拉塔OualataCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "919",
		Title:     "oualata",
		TitleName: "瓦拉塔",
		TitleCode: "c_oualata",
		Baronies:  map[string]feud.Barony{},
	}

	f.比尔恩萨拉Birnsara = BBirnsara比尔恩萨拉
	f.比尔恩萨拉Birnsara.SetParent(f)
	
	f.比如Biru = BBiru比如
	f.比如Biru.SetParent(f)
	
	f.法瑟卡拉Fassekra = BFassekra法瑟卡拉
	f.法瑟卡拉Fassekra.SetParent(f)
	
	f.科努Konou = BKonou科努
	f.科努Konou.SetParent(f)
	
	f.内吉亚特Neiguiat = BNeiguiat内吉亚特
	f.内吉亚特Neiguiat.SetParent(f)
	
	f.瓦拉塔Oualata = BOualata瓦拉塔
	f.瓦拉塔Oualata.SetParent(f)
	
	f.塔尔哈雷特Tarhalet = BTarhalet塔尔哈雷特
	f.塔尔哈雷特Tarhalet.SetParent(f)
	
	f.蒂泽尔特Tizert = BTizert蒂泽尔特
	f.蒂泽尔特Tizert.SetParent(f)
	
}
