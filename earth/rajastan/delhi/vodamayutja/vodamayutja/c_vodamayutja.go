package vodamayutja

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type VodamayutjaCounty interface {
    feud.County
    BAhichatra恶醯掣呾逻() 	feud.Barony
    BAnwala阿瓦罗() 	feud.Barony
    BBareily婆雷利() 	feud.Barony
    BGegha祇伽() 	feud.Barony
    BSheikhupur舍侯补罗() 	feud.Barony
    BTilokpur帝卢迦补罗() 	feud.Barony
    BVodamayutja菩陀摩瑜多() 	feud.Barony
}

type 菩陀摩瑜多VodamayutjaCounty struct {
	feud.BaseCounty
	恶醯掣呾逻Ahichatra 	feud.Barony
	阿瓦罗Anwala 	feud.Barony
	婆雷利Bareily 	feud.Barony
	祇伽Gegha 	feud.Barony
	舍侯补罗Sheikhupur 	feud.Barony
	帝卢迦补罗Tilokpur 	feud.Barony
	菩陀摩瑜多Vodamayutja 	feud.Barony
}

func (c *菩陀摩瑜多VodamayutjaCounty) BAhichatra恶醯掣呾逻() feud.Barony {
	return c.恶醯掣呾逻Ahichatra
}
    
func (c *菩陀摩瑜多VodamayutjaCounty) BAnwala阿瓦罗() feud.Barony {
	return c.阿瓦罗Anwala
}
    
func (c *菩陀摩瑜多VodamayutjaCounty) BBareily婆雷利() feud.Barony {
	return c.婆雷利Bareily
}
    
func (c *菩陀摩瑜多VodamayutjaCounty) BGegha祇伽() feud.Barony {
	return c.祇伽Gegha
}
    
func (c *菩陀摩瑜多VodamayutjaCounty) BSheikhupur舍侯补罗() feud.Barony {
	return c.舍侯补罗Sheikhupur
}
    
func (c *菩陀摩瑜多VodamayutjaCounty) BTilokpur帝卢迦补罗() feud.Barony {
	return c.帝卢迦补罗Tilokpur
}
    
func (c *菩陀摩瑜多VodamayutjaCounty) BVodamayutja菩陀摩瑜多() feud.Barony {
	return c.菩陀摩瑜多Vodamayutja
}
    
var CVodamayutja菩陀摩瑜多 VodamayutjaCounty = &菩陀摩瑜多VodamayutjaCounty{}

func init() {
	f := CVodamayutja菩陀摩瑜多.(*菩陀摩瑜多VodamayutjaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1358",
		Title:     "vodamayutja",
		TitleName: "菩陀摩瑜多",
		TitleCode: "c_vodamayutja",
		Baronies:  map[string]feud.Barony{},
	}

	f.恶醯掣呾逻Ahichatra = BAhichatra恶醯掣呾逻
	f.恶醯掣呾逻Ahichatra.SetParent(f)
	
	f.阿瓦罗Anwala = BAnwala阿瓦罗
	f.阿瓦罗Anwala.SetParent(f)
	
	f.婆雷利Bareily = BBareily婆雷利
	f.婆雷利Bareily.SetParent(f)
	
	f.祇伽Gegha = BGegha祇伽
	f.祇伽Gegha.SetParent(f)
	
	f.舍侯补罗Sheikhupur = BSheikhupur舍侯补罗
	f.舍侯补罗Sheikhupur.SetParent(f)
	
	f.帝卢迦补罗Tilokpur = BTilokpur帝卢迦补罗
	f.帝卢迦补罗Tilokpur.SetParent(f)
	
	f.菩陀摩瑜多Vodamayutja = BVodamayutja菩陀摩瑜多
	f.菩陀摩瑜多Vodamayutja.SetParent(f)
	
}
