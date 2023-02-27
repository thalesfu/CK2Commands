package narva

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type NarvaCounty interface {
    feud.County
    BAgelinde阿格林黛() 	feud.Barony
    BAlentagh阿林塔加() 	feud.Barony
    BAskala埃斯凯莱() 	feud.Barony
    BNarva纳尔瓦() 	feud.Barony
    BPudiviru普迪维鲁() 	feud.Barony
    BRepel雷佩尔() 	feud.Barony
    BTelsborg特尔斯伯格() 	feud.Barony
    BWesenberg韦森贝格() 	feud.Barony
}

type 纳尔瓦NarvaCounty struct {
	feud.BaseCounty
	阿格林黛Agelinde 	feud.Barony
	阿林塔加Alentagh 	feud.Barony
	埃斯凯莱Askala 	feud.Barony
	纳尔瓦Narva 	feud.Barony
	普迪维鲁Pudiviru 	feud.Barony
	雷佩尔Repel 	feud.Barony
	特尔斯伯格Telsborg 	feud.Barony
	韦森贝格Wesenberg 	feud.Barony
}

func (c *纳尔瓦NarvaCounty) BAgelinde阿格林黛() feud.Barony {
	return c.阿格林黛Agelinde
}
    
func (c *纳尔瓦NarvaCounty) BAlentagh阿林塔加() feud.Barony {
	return c.阿林塔加Alentagh
}
    
func (c *纳尔瓦NarvaCounty) BAskala埃斯凯莱() feud.Barony {
	return c.埃斯凯莱Askala
}
    
func (c *纳尔瓦NarvaCounty) BNarva纳尔瓦() feud.Barony {
	return c.纳尔瓦Narva
}
    
func (c *纳尔瓦NarvaCounty) BPudiviru普迪维鲁() feud.Barony {
	return c.普迪维鲁Pudiviru
}
    
func (c *纳尔瓦NarvaCounty) BRepel雷佩尔() feud.Barony {
	return c.雷佩尔Repel
}
    
func (c *纳尔瓦NarvaCounty) BTelsborg特尔斯伯格() feud.Barony {
	return c.特尔斯伯格Telsborg
}
    
func (c *纳尔瓦NarvaCounty) BWesenberg韦森贝格() feud.Barony {
	return c.韦森贝格Wesenberg
}
    
var CNarva纳尔瓦 NarvaCounty = &纳尔瓦NarvaCounty{}

func init() {
	f := CNarva纳尔瓦.(*纳尔瓦NarvaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "380",
		Title:     "narva",
		TitleName: "纳尔瓦",
		TitleCode: "c_narva",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿格林黛Agelinde = BAgelinde阿格林黛
	f.阿格林黛Agelinde.SetParent(f)
	
	f.阿林塔加Alentagh = BAlentagh阿林塔加
	f.阿林塔加Alentagh.SetParent(f)
	
	f.埃斯凯莱Askala = BAskala埃斯凯莱
	f.埃斯凯莱Askala.SetParent(f)
	
	f.纳尔瓦Narva = BNarva纳尔瓦
	f.纳尔瓦Narva.SetParent(f)
	
	f.普迪维鲁Pudiviru = BPudiviru普迪维鲁
	f.普迪维鲁Pudiviru.SetParent(f)
	
	f.雷佩尔Repel = BRepel雷佩尔
	f.雷佩尔Repel.SetParent(f)
	
	f.特尔斯伯格Telsborg = BTelsborg特尔斯伯格
	f.特尔斯伯格Telsborg.SetParent(f)
	
	f.韦森贝格Wesenberg = BWesenberg韦森贝格
	f.韦森贝格Wesenberg.SetParent(f)
	
}
