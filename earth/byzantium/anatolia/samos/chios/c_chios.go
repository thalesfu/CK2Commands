package chios

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ChiosCounty interface {
    feud.County
    BChios希俄斯() 	feud.Barony
    BChrysostomos赫里索斯托莫斯() 	feud.Barony
    BFourni富尔尼() 	feud.Barony
    BIkaria伊卡利亚() 	feud.Barony
    BMarathokampos马拉多堪波斯() 	feud.Barony
    BPagondas帕贡达斯() 	feud.Barony
    BSamos萨摩斯() 	feud.Barony
    BTigani提加尼() 	feud.Barony
}

type 希俄斯ChiosCounty struct {
	feud.BaseCounty
	希俄斯Chios 	feud.Barony
	赫里索斯托莫斯Chrysostomos 	feud.Barony
	富尔尼Fourni 	feud.Barony
	伊卡利亚Ikaria 	feud.Barony
	马拉多堪波斯Marathokampos 	feud.Barony
	帕贡达斯Pagondas 	feud.Barony
	萨摩斯Samos 	feud.Barony
	提加尼Tigani 	feud.Barony
}

func (c *希俄斯ChiosCounty) BChios希俄斯() feud.Barony {
	return c.希俄斯Chios
}
    
func (c *希俄斯ChiosCounty) BChrysostomos赫里索斯托莫斯() feud.Barony {
	return c.赫里索斯托莫斯Chrysostomos
}
    
func (c *希俄斯ChiosCounty) BFourni富尔尼() feud.Barony {
	return c.富尔尼Fourni
}
    
func (c *希俄斯ChiosCounty) BIkaria伊卡利亚() feud.Barony {
	return c.伊卡利亚Ikaria
}
    
func (c *希俄斯ChiosCounty) BMarathokampos马拉多堪波斯() feud.Barony {
	return c.马拉多堪波斯Marathokampos
}
    
func (c *希俄斯ChiosCounty) BPagondas帕贡达斯() feud.Barony {
	return c.帕贡达斯Pagondas
}
    
func (c *希俄斯ChiosCounty) BSamos萨摩斯() feud.Barony {
	return c.萨摩斯Samos
}
    
func (c *希俄斯ChiosCounty) BTigani提加尼() feud.Barony {
	return c.提加尼Tigani
}
    
var CChios希俄斯 ChiosCounty = &希俄斯ChiosCounty{}

func init() {
	f := CChios希俄斯.(*希俄斯ChiosCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "486",
		Title:     "chios",
		TitleName: "希俄斯",
		TitleCode: "c_chios",
		Baronies:  map[string]feud.Barony{},
	}

	f.希俄斯Chios = BChios希俄斯
	f.希俄斯Chios.SetParent(f)
	
	f.赫里索斯托莫斯Chrysostomos = BChrysostomos赫里索斯托莫斯
	f.赫里索斯托莫斯Chrysostomos.SetParent(f)
	
	f.富尔尼Fourni = BFourni富尔尼
	f.富尔尼Fourni.SetParent(f)
	
	f.伊卡利亚Ikaria = BIkaria伊卡利亚
	f.伊卡利亚Ikaria.SetParent(f)
	
	f.马拉多堪波斯Marathokampos = BMarathokampos马拉多堪波斯
	f.马拉多堪波斯Marathokampos.SetParent(f)
	
	f.帕贡达斯Pagondas = BPagondas帕贡达斯
	f.帕贡达斯Pagondas.SetParent(f)
	
	f.萨摩斯Samos = BSamos萨摩斯
	f.萨摩斯Samos.SetParent(f)
	
	f.提加尼Tigani = BTigani提加尼
	f.提加尼Tigani.SetParent(f)
	
}
