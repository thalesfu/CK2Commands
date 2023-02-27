package starodub

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type StarodubCounty interface {
    feud.County
    BAvdiika阿夫迪夫卡() 	feud.Barony
    BDesnianske杰斯尼扬西克() 	feud.Barony
    BObolonnya奥博隆尼亚() 	feud.Barony
    BPonornytsya波诺尔尼察() 	feud.Barony
    BStarodub斯塔罗杜布() 	feud.Barony
    BVyshenky维申基() 	feud.Barony
    BZelena_polyana泽列纳波利亚纳() 	feud.Barony
}

type 斯塔罗杜布StarodubCounty struct {
	feud.BaseCounty
	阿夫迪夫卡Avdiika 	feud.Barony
	杰斯尼扬西克Desnianske 	feud.Barony
	奥博隆尼亚Obolonnya 	feud.Barony
	波诺尔尼察Ponornytsya 	feud.Barony
	斯塔罗杜布Starodub 	feud.Barony
	维申基Vyshenky 	feud.Barony
	泽列纳波利亚纳Zelena_polyana 	feud.Barony
}

func (c *斯塔罗杜布StarodubCounty) BAvdiika阿夫迪夫卡() feud.Barony {
	return c.阿夫迪夫卡Avdiika
}
    
func (c *斯塔罗杜布StarodubCounty) BDesnianske杰斯尼扬西克() feud.Barony {
	return c.杰斯尼扬西克Desnianske
}
    
func (c *斯塔罗杜布StarodubCounty) BObolonnya奥博隆尼亚() feud.Barony {
	return c.奥博隆尼亚Obolonnya
}
    
func (c *斯塔罗杜布StarodubCounty) BPonornytsya波诺尔尼察() feud.Barony {
	return c.波诺尔尼察Ponornytsya
}
    
func (c *斯塔罗杜布StarodubCounty) BStarodub斯塔罗杜布() feud.Barony {
	return c.斯塔罗杜布Starodub
}
    
func (c *斯塔罗杜布StarodubCounty) BVyshenky维申基() feud.Barony {
	return c.维申基Vyshenky
}
    
func (c *斯塔罗杜布StarodubCounty) BZelena_polyana泽列纳波利亚纳() feud.Barony {
	return c.泽列纳波利亚纳Zelena_polyana
}
    
var CStarodub斯塔罗杜布 StarodubCounty = &斯塔罗杜布StarodubCounty{}

func init() {
	f := CStarodub斯塔罗杜布.(*斯塔罗杜布StarodubCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1656",
		Title:     "starodub",
		TitleName: "斯塔罗杜布",
		TitleCode: "c_starodub",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿夫迪夫卡Avdiika = BAvdiika阿夫迪夫卡
	f.阿夫迪夫卡Avdiika.SetParent(f)
	
	f.杰斯尼扬西克Desnianske = BDesnianske杰斯尼扬西克
	f.杰斯尼扬西克Desnianske.SetParent(f)
	
	f.奥博隆尼亚Obolonnya = BObolonnya奥博隆尼亚
	f.奥博隆尼亚Obolonnya.SetParent(f)
	
	f.波诺尔尼察Ponornytsya = BPonornytsya波诺尔尼察
	f.波诺尔尼察Ponornytsya.SetParent(f)
	
	f.斯塔罗杜布Starodub = BStarodub斯塔罗杜布
	f.斯塔罗杜布Starodub.SetParent(f)
	
	f.维申基Vyshenky = BVyshenky维申基
	f.维申基Vyshenky.SetParent(f)
	
	f.泽列纳波利亚纳Zelena_polyana = BZelena_polyana泽列纳波利亚纳
	f.泽列纳波利亚纳Zelena_polyana.SetParent(f)
	
}
