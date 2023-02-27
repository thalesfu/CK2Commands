package samoyeds

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SamoyedsCounty interface {
    feud.County
    BArkhipovo阿尔希波沃() 	feud.Barony
    BKiya基亚() 	feud.Barony
    BSukhanikha苏哈尼哈() 	feud.Barony
    BTarasova塔拉索瓦() 	feud.Barony
    BVerkhmgla韦尔赫姆格拉() 	feud.Barony
    BVizhas维扎斯() 	feud.Barony
    BYazhma亚日马() 	feud.Barony
}

type 斯诺帕SamoyedsCounty struct {
	feud.BaseCounty
	阿尔希波沃Arkhipovo 	feud.Barony
	基亚Kiya 	feud.Barony
	苏哈尼哈Sukhanikha 	feud.Barony
	塔拉索瓦Tarasova 	feud.Barony
	韦尔赫姆格拉Verkhmgla 	feud.Barony
	维扎斯Vizhas 	feud.Barony
	亚日马Yazhma 	feud.Barony
}

func (c *斯诺帕SamoyedsCounty) BArkhipovo阿尔希波沃() feud.Barony {
	return c.阿尔希波沃Arkhipovo
}
    
func (c *斯诺帕SamoyedsCounty) BKiya基亚() feud.Barony {
	return c.基亚Kiya
}
    
func (c *斯诺帕SamoyedsCounty) BSukhanikha苏哈尼哈() feud.Barony {
	return c.苏哈尼哈Sukhanikha
}
    
func (c *斯诺帕SamoyedsCounty) BTarasova塔拉索瓦() feud.Barony {
	return c.塔拉索瓦Tarasova
}
    
func (c *斯诺帕SamoyedsCounty) BVerkhmgla韦尔赫姆格拉() feud.Barony {
	return c.韦尔赫姆格拉Verkhmgla
}
    
func (c *斯诺帕SamoyedsCounty) BVizhas维扎斯() feud.Barony {
	return c.维扎斯Vizhas
}
    
func (c *斯诺帕SamoyedsCounty) BYazhma亚日马() feud.Barony {
	return c.亚日马Yazhma
}
    
var CSamoyeds斯诺帕 SamoyedsCounty = &斯诺帕SamoyedsCounty{}

func init() {
	f := CSamoyeds斯诺帕.(*斯诺帕SamoyedsCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "397",
		Title:     "samoyeds",
		TitleName: "斯诺帕",
		TitleCode: "c_samoyeds",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿尔希波沃Arkhipovo = BArkhipovo阿尔希波沃
	f.阿尔希波沃Arkhipovo.SetParent(f)
	
	f.基亚Kiya = BKiya基亚
	f.基亚Kiya.SetParent(f)
	
	f.苏哈尼哈Sukhanikha = BSukhanikha苏哈尼哈
	f.苏哈尼哈Sukhanikha.SetParent(f)
	
	f.塔拉索瓦Tarasova = BTarasova塔拉索瓦
	f.塔拉索瓦Tarasova.SetParent(f)
	
	f.韦尔赫姆格拉Verkhmgla = BVerkhmgla韦尔赫姆格拉
	f.韦尔赫姆格拉Verkhmgla.SetParent(f)
	
	f.维扎斯Vizhas = BVizhas维扎斯
	f.维扎斯Vizhas.SetParent(f)
	
	f.亚日马Yazhma = BYazhma亚日马
	f.亚日马Yazhma.SetParent(f)
	
}
