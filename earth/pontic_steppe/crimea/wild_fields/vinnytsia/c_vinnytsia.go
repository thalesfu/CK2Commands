package vinnytsia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type VinnytsiaCounty interface {
    feud.County
    BBershad伯沙德() 	feud.Barony
    BHaisyn海辛() 	feud.Barony
    BIllintsi伊林齐() 	feud.Barony
    BLadyzhyn拉德任() 	feud.Barony
    BPervomaisk佩尔沃迈斯克() 	feud.Barony
    BRov罗夫() 	feud.Barony
    BVinnytsia文尼察() 	feud.Barony
}

type 文尼察VinnytsiaCounty struct {
	feud.BaseCounty
	伯沙德Bershad 	feud.Barony
	海辛Haisyn 	feud.Barony
	伊林齐Illintsi 	feud.Barony
	拉德任Ladyzhyn 	feud.Barony
	佩尔沃迈斯克Pervomaisk 	feud.Barony
	罗夫Rov 	feud.Barony
	文尼察Vinnytsia 	feud.Barony
}

func (c *文尼察VinnytsiaCounty) BBershad伯沙德() feud.Barony {
	return c.伯沙德Bershad
}
    
func (c *文尼察VinnytsiaCounty) BHaisyn海辛() feud.Barony {
	return c.海辛Haisyn
}
    
func (c *文尼察VinnytsiaCounty) BIllintsi伊林齐() feud.Barony {
	return c.伊林齐Illintsi
}
    
func (c *文尼察VinnytsiaCounty) BLadyzhyn拉德任() feud.Barony {
	return c.拉德任Ladyzhyn
}
    
func (c *文尼察VinnytsiaCounty) BPervomaisk佩尔沃迈斯克() feud.Barony {
	return c.佩尔沃迈斯克Pervomaisk
}
    
func (c *文尼察VinnytsiaCounty) BRov罗夫() feud.Barony {
	return c.罗夫Rov
}
    
func (c *文尼察VinnytsiaCounty) BVinnytsia文尼察() feud.Barony {
	return c.文尼察Vinnytsia
}
    
var CVinnytsia文尼察 VinnytsiaCounty = &文尼察VinnytsiaCounty{}

func init() {
	f := CVinnytsia文尼察.(*文尼察VinnytsiaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1647",
		Title:     "vinnytsia",
		TitleName: "文尼察",
		TitleCode: "c_vinnytsia",
		Baronies:  map[string]feud.Barony{},
	}

	f.伯沙德Bershad = BBershad伯沙德
	f.伯沙德Bershad.SetParent(f)
	
	f.海辛Haisyn = BHaisyn海辛
	f.海辛Haisyn.SetParent(f)
	
	f.伊林齐Illintsi = BIllintsi伊林齐
	f.伊林齐Illintsi.SetParent(f)
	
	f.拉德任Ladyzhyn = BLadyzhyn拉德任
	f.拉德任Ladyzhyn.SetParent(f)
	
	f.佩尔沃迈斯克Pervomaisk = BPervomaisk佩尔沃迈斯克
	f.佩尔沃迈斯克Pervomaisk.SetParent(f)
	
	f.罗夫Rov = BRov罗夫
	f.罗夫Rov.SetParent(f)
	
	f.文尼察Vinnytsia = BVinnytsia文尼察
	f.文尼察Vinnytsia.SetParent(f)
	
}
