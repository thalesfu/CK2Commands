package zaragoza

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ZaragozaCounty interface {
    feud.County
    BAlagon阿拉贡() 	feud.Barony
    BBorja博尔哈() 	feud.Barony
    BCaspe卡斯佩() 	feud.Barony
    BEjea埃赫阿() 	feud.Barony
    BEpila埃皮拉() 	feud.Barony
    BMedianadearagon梅迪亚纳德亚拉贡() 	feud.Barony
    BVeruela韦鲁埃拉() 	feud.Barony
    BZaragoza萨拉戈萨() 	feud.Barony
}

type 萨拉戈萨ZaragozaCounty struct {
	feud.BaseCounty
	阿拉贡Alagon 	feud.Barony
	博尔哈Borja 	feud.Barony
	卡斯佩Caspe 	feud.Barony
	埃赫阿Ejea 	feud.Barony
	埃皮拉Epila 	feud.Barony
	梅迪亚纳德亚拉贡Medianadearagon 	feud.Barony
	韦鲁埃拉Veruela 	feud.Barony
	萨拉戈萨Zaragoza 	feud.Barony
}

func (c *萨拉戈萨ZaragozaCounty) BAlagon阿拉贡() feud.Barony {
	return c.阿拉贡Alagon
}
    
func (c *萨拉戈萨ZaragozaCounty) BBorja博尔哈() feud.Barony {
	return c.博尔哈Borja
}
    
func (c *萨拉戈萨ZaragozaCounty) BCaspe卡斯佩() feud.Barony {
	return c.卡斯佩Caspe
}
    
func (c *萨拉戈萨ZaragozaCounty) BEjea埃赫阿() feud.Barony {
	return c.埃赫阿Ejea
}
    
func (c *萨拉戈萨ZaragozaCounty) BEpila埃皮拉() feud.Barony {
	return c.埃皮拉Epila
}
    
func (c *萨拉戈萨ZaragozaCounty) BMedianadearagon梅迪亚纳德亚拉贡() feud.Barony {
	return c.梅迪亚纳德亚拉贡Medianadearagon
}
    
func (c *萨拉戈萨ZaragozaCounty) BVeruela韦鲁埃拉() feud.Barony {
	return c.韦鲁埃拉Veruela
}
    
func (c *萨拉戈萨ZaragozaCounty) BZaragoza萨拉戈萨() feud.Barony {
	return c.萨拉戈萨Zaragoza
}
    
var CZaragoza萨拉戈萨 ZaragozaCounty = &萨拉戈萨ZaragozaCounty{}

func init() {
	f := CZaragoza萨拉戈萨.(*萨拉戈萨ZaragozaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "202",
		Title:     "zaragoza",
		TitleName: "萨拉戈萨",
		TitleCode: "c_zaragoza",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿拉贡Alagon = BAlagon阿拉贡
	f.阿拉贡Alagon.SetParent(f)
	
	f.博尔哈Borja = BBorja博尔哈
	f.博尔哈Borja.SetParent(f)
	
	f.卡斯佩Caspe = BCaspe卡斯佩
	f.卡斯佩Caspe.SetParent(f)
	
	f.埃赫阿Ejea = BEjea埃赫阿
	f.埃赫阿Ejea.SetParent(f)
	
	f.埃皮拉Epila = BEpila埃皮拉
	f.埃皮拉Epila.SetParent(f)
	
	f.梅迪亚纳德亚拉贡Medianadearagon = BMedianadearagon梅迪亚纳德亚拉贡
	f.梅迪亚纳德亚拉贡Medianadearagon.SetParent(f)
	
	f.韦鲁埃拉Veruela = BVeruela韦鲁埃拉
	f.韦鲁埃拉Veruela.SetParent(f)
	
	f.萨拉戈萨Zaragoza = BZaragoza萨拉戈萨
	f.萨拉戈萨Zaragoza.SetParent(f)
	
}
