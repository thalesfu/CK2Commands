package tharasset

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TharassetCounty interface {
    feud.County
    BBelkassem贝尔卡塞姆() 	feud.Barony
    BBiourga比乌格拉() 	feud.Barony
    BGuelmim盖勒敏() 	feud.Barony
    BIgilliz伊吉利兹() 	feud.Barony
    BMihrleft米勒福特() 	feud.Barony
    BOuaouteit瓦乌泰特() 	feud.Barony
    BRgab拉格布() 	feud.Barony
    BTaroudaut塔鲁丹特() 	feud.Barony
    BTata塔塔() 	feud.Barony
    BTiznit提兹尼特() 	feud.Barony
}

type 苏斯TharassetCounty struct {
	feud.BaseCounty
	贝尔卡塞姆Belkassem 	feud.Barony
	比乌格拉Biourga 	feud.Barony
	盖勒敏Guelmim 	feud.Barony
	伊吉利兹Igilliz 	feud.Barony
	米勒福特Mihrleft 	feud.Barony
	瓦乌泰特Ouaouteit 	feud.Barony
	拉格布Rgab 	feud.Barony
	塔鲁丹特Taroudaut 	feud.Barony
	塔塔Tata 	feud.Barony
	提兹尼特Tiznit 	feud.Barony
}

func (c *苏斯TharassetCounty) BBelkassem贝尔卡塞姆() feud.Barony {
	return c.贝尔卡塞姆Belkassem
}
    
func (c *苏斯TharassetCounty) BBiourga比乌格拉() feud.Barony {
	return c.比乌格拉Biourga
}
    
func (c *苏斯TharassetCounty) BGuelmim盖勒敏() feud.Barony {
	return c.盖勒敏Guelmim
}
    
func (c *苏斯TharassetCounty) BIgilliz伊吉利兹() feud.Barony {
	return c.伊吉利兹Igilliz
}
    
func (c *苏斯TharassetCounty) BMihrleft米勒福特() feud.Barony {
	return c.米勒福特Mihrleft
}
    
func (c *苏斯TharassetCounty) BOuaouteit瓦乌泰特() feud.Barony {
	return c.瓦乌泰特Ouaouteit
}
    
func (c *苏斯TharassetCounty) BRgab拉格布() feud.Barony {
	return c.拉格布Rgab
}
    
func (c *苏斯TharassetCounty) BTaroudaut塔鲁丹特() feud.Barony {
	return c.塔鲁丹特Taroudaut
}
    
func (c *苏斯TharassetCounty) BTata塔塔() feud.Barony {
	return c.塔塔Tata
}
    
func (c *苏斯TharassetCounty) BTiznit提兹尼特() feud.Barony {
	return c.提兹尼特Tiznit
}
    
var CTharasset苏斯 TharassetCounty = &苏斯TharassetCounty{}

func init() {
	f := CTharasset苏斯.(*苏斯TharassetCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "847",
		Title:     "tharasset",
		TitleName: "苏斯",
		TitleCode: "c_tharasset",
		Baronies:  map[string]feud.Barony{},
	}

	f.贝尔卡塞姆Belkassem = BBelkassem贝尔卡塞姆
	f.贝尔卡塞姆Belkassem.SetParent(f)
	
	f.比乌格拉Biourga = BBiourga比乌格拉
	f.比乌格拉Biourga.SetParent(f)
	
	f.盖勒敏Guelmim = BGuelmim盖勒敏
	f.盖勒敏Guelmim.SetParent(f)
	
	f.伊吉利兹Igilliz = BIgilliz伊吉利兹
	f.伊吉利兹Igilliz.SetParent(f)
	
	f.米勒福特Mihrleft = BMihrleft米勒福特
	f.米勒福特Mihrleft.SetParent(f)
	
	f.瓦乌泰特Ouaouteit = BOuaouteit瓦乌泰特
	f.瓦乌泰特Ouaouteit.SetParent(f)
	
	f.拉格布Rgab = BRgab拉格布
	f.拉格布Rgab.SetParent(f)
	
	f.塔鲁丹特Taroudaut = BTaroudaut塔鲁丹特
	f.塔鲁丹特Taroudaut.SetParent(f)
	
	f.塔塔Tata = BTata塔塔
	f.塔塔Tata.SetParent(f)
	
	f.提兹尼特Tiznit = BTiznit提兹尼特
	f.提兹尼特Tiznit.SetParent(f)
	
}
