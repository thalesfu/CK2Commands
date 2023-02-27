package pest

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type PestCounty interface {
    feud.County
    BAbrahamtelke阿布劳哈姆泰尔凯() 	feud.Barony
    BCegled采格莱德() 	feud.Barony
    BKecskemet凯奇凯梅特() 	feud.Barony
    BKiskoros小克勒什() 	feud.Barony
    BKiskunhalas基什孔豪洛什() 	feud.Barony
    BPest佩斯() 	feud.Barony
    BSzentendre圣安德烈() 	feud.Barony
    BVac瓦茨() 	feud.Barony
}

type 佩斯PestCounty struct {
	feud.BaseCounty
	阿布劳哈姆泰尔凯Abrahamtelke 	feud.Barony
	采格莱德Cegled 	feud.Barony
	凯奇凯梅特Kecskemet 	feud.Barony
	小克勒什Kiskoros 	feud.Barony
	基什孔豪洛什Kiskunhalas 	feud.Barony
	佩斯Pest 	feud.Barony
	圣安德烈Szentendre 	feud.Barony
	瓦茨Vac 	feud.Barony
}

func (c *佩斯PestCounty) BAbrahamtelke阿布劳哈姆泰尔凯() feud.Barony {
	return c.阿布劳哈姆泰尔凯Abrahamtelke
}
    
func (c *佩斯PestCounty) BCegled采格莱德() feud.Barony {
	return c.采格莱德Cegled
}
    
func (c *佩斯PestCounty) BKecskemet凯奇凯梅特() feud.Barony {
	return c.凯奇凯梅特Kecskemet
}
    
func (c *佩斯PestCounty) BKiskoros小克勒什() feud.Barony {
	return c.小克勒什Kiskoros
}
    
func (c *佩斯PestCounty) BKiskunhalas基什孔豪洛什() feud.Barony {
	return c.基什孔豪洛什Kiskunhalas
}
    
func (c *佩斯PestCounty) BPest佩斯() feud.Barony {
	return c.佩斯Pest
}
    
func (c *佩斯PestCounty) BSzentendre圣安德烈() feud.Barony {
	return c.圣安德烈Szentendre
}
    
func (c *佩斯PestCounty) BVac瓦茨() feud.Barony {
	return c.瓦茨Vac
}
    
var CPest佩斯 PestCounty = &佩斯PestCounty{}

func init() {
	f := CPest佩斯.(*佩斯PestCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "522",
		Title:     "pest",
		TitleName: "佩斯",
		TitleCode: "c_pest",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿布劳哈姆泰尔凯Abrahamtelke = BAbrahamtelke阿布劳哈姆泰尔凯
	f.阿布劳哈姆泰尔凯Abrahamtelke.SetParent(f)
	
	f.采格莱德Cegled = BCegled采格莱德
	f.采格莱德Cegled.SetParent(f)
	
	f.凯奇凯梅特Kecskemet = BKecskemet凯奇凯梅特
	f.凯奇凯梅特Kecskemet.SetParent(f)
	
	f.小克勒什Kiskoros = BKiskoros小克勒什
	f.小克勒什Kiskoros.SetParent(f)
	
	f.基什孔豪洛什Kiskunhalas = BKiskunhalas基什孔豪洛什
	f.基什孔豪洛什Kiskunhalas.SetParent(f)
	
	f.佩斯Pest = BPest佩斯
	f.佩斯Pest.SetParent(f)
	
	f.圣安德烈Szentendre = BSzentendre圣安德烈
	f.圣安德烈Szentendre.SetParent(f)
	
	f.瓦茨Vac = BVac瓦茨
	f.瓦茨Vac.SetParent(f)
	
}
