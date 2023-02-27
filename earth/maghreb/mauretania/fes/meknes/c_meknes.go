package meknes

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MeknesCounty interface {
    feud.County
    BAmouguer阿穆盖尔() 	feud.Barony
    BIrhzi伊尔齐() 	feud.Barony
    BMeknes梅克内斯() 	feud.Barony
    BSefrou塞夫鲁() 	feud.Barony
    BSoughi苏吉() 	feud.Barony
    BTagragra塔格拉格拉() 	feud.Barony
    BTerhirha蒂尔希拉() 	feud.Barony
}

type 梅克内斯MeknesCounty struct {
	feud.BaseCounty
	阿穆盖尔Amouguer 	feud.Barony
	伊尔齐Irhzi 	feud.Barony
	梅克内斯Meknes 	feud.Barony
	塞夫鲁Sefrou 	feud.Barony
	苏吉Soughi 	feud.Barony
	塔格拉格拉Tagragra 	feud.Barony
	蒂尔希拉Terhirha 	feud.Barony
}

func (c *梅克内斯MeknesCounty) BAmouguer阿穆盖尔() feud.Barony {
	return c.阿穆盖尔Amouguer
}
    
func (c *梅克内斯MeknesCounty) BIrhzi伊尔齐() feud.Barony {
	return c.伊尔齐Irhzi
}
    
func (c *梅克内斯MeknesCounty) BMeknes梅克内斯() feud.Barony {
	return c.梅克内斯Meknes
}
    
func (c *梅克内斯MeknesCounty) BSefrou塞夫鲁() feud.Barony {
	return c.塞夫鲁Sefrou
}
    
func (c *梅克内斯MeknesCounty) BSoughi苏吉() feud.Barony {
	return c.苏吉Soughi
}
    
func (c *梅克内斯MeknesCounty) BTagragra塔格拉格拉() feud.Barony {
	return c.塔格拉格拉Tagragra
}
    
func (c *梅克内斯MeknesCounty) BTerhirha蒂尔希拉() feud.Barony {
	return c.蒂尔希拉Terhirha
}
    
var CMeknes梅克内斯 MeknesCounty = &梅克内斯MeknesCounty{}

func init() {
	f := CMeknes梅克内斯.(*梅克内斯MeknesCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1774",
		Title:     "meknes",
		TitleName: "梅克内斯",
		TitleCode: "c_meknes",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿穆盖尔Amouguer = BAmouguer阿穆盖尔
	f.阿穆盖尔Amouguer.SetParent(f)
	
	f.伊尔齐Irhzi = BIrhzi伊尔齐
	f.伊尔齐Irhzi.SetParent(f)
	
	f.梅克内斯Meknes = BMeknes梅克内斯
	f.梅克内斯Meknes.SetParent(f)
	
	f.塞夫鲁Sefrou = BSefrou塞夫鲁
	f.塞夫鲁Sefrou.SetParent(f)
	
	f.苏吉Soughi = BSoughi苏吉
	f.苏吉Soughi.SetParent(f)
	
	f.塔格拉格拉Tagragra = BTagragra塔格拉格拉
	f.塔格拉格拉Tagragra.SetParent(f)
	
	f.蒂尔希拉Terhirha = BTerhirha蒂尔希拉
	f.蒂尔希拉Terhirha.SetParent(f)
	
}
