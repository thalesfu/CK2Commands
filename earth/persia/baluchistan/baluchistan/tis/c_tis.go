package tis

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TisCounty interface {
    feud.County
    BChabahar哈赫巴哈尔() 	feud.Barony
    BKursar库雷斯萨() 	feud.Barony
    BParak帕拉克() 	feud.Barony
    BPozm_machchan帕姆马处汗() 	feud.Barony
    BRegedan雷德丹() 	feud.Barony
    BSergen舍尔甘() 	feud.Barony
    BTis蒂斯() 	feud.Barony
}

type 蒂斯TisCounty struct {
	feud.BaseCounty
	哈赫巴哈尔Chabahar 	feud.Barony
	库雷斯萨Kursar 	feud.Barony
	帕拉克Parak 	feud.Barony
	帕姆马处汗Pozm_machchan 	feud.Barony
	雷德丹Regedan 	feud.Barony
	舍尔甘Sergen 	feud.Barony
	蒂斯Tis 	feud.Barony
}

func (c *蒂斯TisCounty) BChabahar哈赫巴哈尔() feud.Barony {
	return c.哈赫巴哈尔Chabahar
}
    
func (c *蒂斯TisCounty) BKursar库雷斯萨() feud.Barony {
	return c.库雷斯萨Kursar
}
    
func (c *蒂斯TisCounty) BParak帕拉克() feud.Barony {
	return c.帕拉克Parak
}
    
func (c *蒂斯TisCounty) BPozm_machchan帕姆马处汗() feud.Barony {
	return c.帕姆马处汗Pozm_machchan
}
    
func (c *蒂斯TisCounty) BRegedan雷德丹() feud.Barony {
	return c.雷德丹Regedan
}
    
func (c *蒂斯TisCounty) BSergen舍尔甘() feud.Barony {
	return c.舍尔甘Sergen
}
    
func (c *蒂斯TisCounty) BTis蒂斯() feud.Barony {
	return c.蒂斯Tis
}
    
var CTis蒂斯 TisCounty = &蒂斯TisCounty{}

func init() {
	f := CTis蒂斯.(*蒂斯TisCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "899",
		Title:     "tis",
		TitleName: "蒂斯",
		TitleCode: "c_tis",
		Baronies:  map[string]feud.Barony{},
	}

	f.哈赫巴哈尔Chabahar = BChabahar哈赫巴哈尔
	f.哈赫巴哈尔Chabahar.SetParent(f)
	
	f.库雷斯萨Kursar = BKursar库雷斯萨
	f.库雷斯萨Kursar.SetParent(f)
	
	f.帕拉克Parak = BParak帕拉克
	f.帕拉克Parak.SetParent(f)
	
	f.帕姆马处汗Pozm_machchan = BPozm_machchan帕姆马处汗
	f.帕姆马处汗Pozm_machchan.SetParent(f)
	
	f.雷德丹Regedan = BRegedan雷德丹
	f.雷德丹Regedan.SetParent(f)
	
	f.舍尔甘Sergen = BSergen舍尔甘
	f.舍尔甘Sergen.SetParent(f)
	
	f.蒂斯Tis = BTis蒂斯
	f.蒂斯Tis.SetParent(f)
	
}
