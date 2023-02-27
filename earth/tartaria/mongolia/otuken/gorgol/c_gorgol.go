package gorgol

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type GorgolCounty interface {
    feud.County
    BAbaza阿巴扎() 	feud.Barony
    BAltan_talbar阿勒坦塔勒巴尔() 	feud.Barony
    BBelcheeriin贝尔彻林() 	feud.Barony
    BGorgol戈尔戈勒() 	feud.Barony
    BGuunii_suu古尼苏() 	feud.Barony
    BTashtyp塔什特普() 	feud.Barony
    BTsetseg_moriin齐齐格穆林() 	feud.Barony
}

type 叶尼塞GorgolCounty struct {
	feud.BaseCounty
	阿巴扎Abaza 	feud.Barony
	阿勒坦塔勒巴尔Altan_talbar 	feud.Barony
	贝尔彻林Belcheeriin 	feud.Barony
	戈尔戈勒Gorgol 	feud.Barony
	古尼苏Guunii_suu 	feud.Barony
	塔什特普Tashtyp 	feud.Barony
	齐齐格穆林Tsetseg_moriin 	feud.Barony
}

func (c *叶尼塞GorgolCounty) BAbaza阿巴扎() feud.Barony {
	return c.阿巴扎Abaza
}
    
func (c *叶尼塞GorgolCounty) BAltan_talbar阿勒坦塔勒巴尔() feud.Barony {
	return c.阿勒坦塔勒巴尔Altan_talbar
}
    
func (c *叶尼塞GorgolCounty) BBelcheeriin贝尔彻林() feud.Barony {
	return c.贝尔彻林Belcheeriin
}
    
func (c *叶尼塞GorgolCounty) BGorgol戈尔戈勒() feud.Barony {
	return c.戈尔戈勒Gorgol
}
    
func (c *叶尼塞GorgolCounty) BGuunii_suu古尼苏() feud.Barony {
	return c.古尼苏Guunii_suu
}
    
func (c *叶尼塞GorgolCounty) BTashtyp塔什特普() feud.Barony {
	return c.塔什特普Tashtyp
}
    
func (c *叶尼塞GorgolCounty) BTsetseg_moriin齐齐格穆林() feud.Barony {
	return c.齐齐格穆林Tsetseg_moriin
}
    
var CGorgol叶尼塞 GorgolCounty = &叶尼塞GorgolCounty{}

func init() {
	f := CGorgol叶尼塞.(*叶尼塞GorgolCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1462",
		Title:     "gorgol",
		TitleName: "叶尼塞",
		TitleCode: "c_gorgol",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿巴扎Abaza = BAbaza阿巴扎
	f.阿巴扎Abaza.SetParent(f)
	
	f.阿勒坦塔勒巴尔Altan_talbar = BAltan_talbar阿勒坦塔勒巴尔
	f.阿勒坦塔勒巴尔Altan_talbar.SetParent(f)
	
	f.贝尔彻林Belcheeriin = BBelcheeriin贝尔彻林
	f.贝尔彻林Belcheeriin.SetParent(f)
	
	f.戈尔戈勒Gorgol = BGorgol戈尔戈勒
	f.戈尔戈勒Gorgol.SetParent(f)
	
	f.古尼苏Guunii_suu = BGuunii_suu古尼苏
	f.古尼苏Guunii_suu.SetParent(f)
	
	f.塔什特普Tashtyp = BTashtyp塔什特普
	f.塔什特普Tashtyp.SetParent(f)
	
	f.齐齐格穆林Tsetseg_moriin = BTsetseg_moriin齐齐格穆林
	f.齐齐格穆林Tsetseg_moriin.SetParent(f)
	
}
