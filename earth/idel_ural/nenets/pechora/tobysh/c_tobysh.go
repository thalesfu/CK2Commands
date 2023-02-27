package tobysh

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TobyshCounty interface {
    feud.County
    BKotkino科特基诺() 	feud.Barony
    BLabozhskoye拉博日斯科耶() 	feud.Barony
    BMedvezhka梅德韦日卡() 	feud.Barony
    BShchelino谢利诺() 	feud.Barony
    BSula_tobysh苏拉() 	feud.Barony
    BTobysh托贝什() 	feud.Barony
    BVelikovisochnoye大维索奇诺耶() 	feud.Barony
}

type 托贝什TobyshCounty struct {
	feud.BaseCounty
	科特基诺Kotkino 	feud.Barony
	拉博日斯科耶Labozhskoye 	feud.Barony
	梅德韦日卡Medvezhka 	feud.Barony
	谢利诺Shchelino 	feud.Barony
	苏拉Sula_tobysh 	feud.Barony
	托贝什Tobysh 	feud.Barony
	大维索奇诺耶Velikovisochnoye 	feud.Barony
}

func (c *托贝什TobyshCounty) BKotkino科特基诺() feud.Barony {
	return c.科特基诺Kotkino
}
    
func (c *托贝什TobyshCounty) BLabozhskoye拉博日斯科耶() feud.Barony {
	return c.拉博日斯科耶Labozhskoye
}
    
func (c *托贝什TobyshCounty) BMedvezhka梅德韦日卡() feud.Barony {
	return c.梅德韦日卡Medvezhka
}
    
func (c *托贝什TobyshCounty) BShchelino谢利诺() feud.Barony {
	return c.谢利诺Shchelino
}
    
func (c *托贝什TobyshCounty) BSula_tobysh苏拉() feud.Barony {
	return c.苏拉Sula_tobysh
}
    
func (c *托贝什TobyshCounty) BTobysh托贝什() feud.Barony {
	return c.托贝什Tobysh
}
    
func (c *托贝什TobyshCounty) BVelikovisochnoye大维索奇诺耶() feud.Barony {
	return c.大维索奇诺耶Velikovisochnoye
}
    
var CTobysh托贝什 TobyshCounty = &托贝什TobyshCounty{}

func init() {
	f := CTobysh托贝什.(*托贝什TobyshCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1832",
		Title:     "tobysh",
		TitleName: "托贝什",
		TitleCode: "c_tobysh",
		Baronies:  map[string]feud.Barony{},
	}

	f.科特基诺Kotkino = BKotkino科特基诺
	f.科特基诺Kotkino.SetParent(f)
	
	f.拉博日斯科耶Labozhskoye = BLabozhskoye拉博日斯科耶
	f.拉博日斯科耶Labozhskoye.SetParent(f)
	
	f.梅德韦日卡Medvezhka = BMedvezhka梅德韦日卡
	f.梅德韦日卡Medvezhka.SetParent(f)
	
	f.谢利诺Shchelino = BShchelino谢利诺
	f.谢利诺Shchelino.SetParent(f)
	
	f.苏拉Sula_tobysh = BSula_tobysh苏拉
	f.苏拉Sula_tobysh.SetParent(f)
	
	f.托贝什Tobysh = BTobysh托贝什
	f.托贝什Tobysh.SetParent(f)
	
	f.大维索奇诺耶Velikovisochnoye = BVelikovisochnoye大维索奇诺耶
	f.大维索奇诺耶Velikovisochnoye.SetParent(f)
	
}
