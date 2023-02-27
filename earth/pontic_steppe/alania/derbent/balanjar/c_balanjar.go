package balanjar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BalanjarCounty interface {
    feud.County
    BArgun阿尔贡() 	feud.Barony
    BBalanjar巴兰贾尔() 	feud.Barony
    BKarata卡拉塔() 	feud.Barony
    BKiri基里() 	feud.Barony
    BShali沙利() 	feud.Barony
    BShatoy沙托伊() 	feud.Barony
    BVedeno韦杰诺() 	feud.Barony
}

type 巴兰贾尔BalanjarCounty struct {
	feud.BaseCounty
	阿尔贡Argun 	feud.Barony
	巴兰贾尔Balanjar 	feud.Barony
	卡拉塔Karata 	feud.Barony
	基里Kiri 	feud.Barony
	沙利Shali 	feud.Barony
	沙托伊Shatoy 	feud.Barony
	韦杰诺Vedeno 	feud.Barony
}

func (c *巴兰贾尔BalanjarCounty) BArgun阿尔贡() feud.Barony {
	return c.阿尔贡Argun
}
    
func (c *巴兰贾尔BalanjarCounty) BBalanjar巴兰贾尔() feud.Barony {
	return c.巴兰贾尔Balanjar
}
    
func (c *巴兰贾尔BalanjarCounty) BKarata卡拉塔() feud.Barony {
	return c.卡拉塔Karata
}
    
func (c *巴兰贾尔BalanjarCounty) BKiri基里() feud.Barony {
	return c.基里Kiri
}
    
func (c *巴兰贾尔BalanjarCounty) BShali沙利() feud.Barony {
	return c.沙利Shali
}
    
func (c *巴兰贾尔BalanjarCounty) BShatoy沙托伊() feud.Barony {
	return c.沙托伊Shatoy
}
    
func (c *巴兰贾尔BalanjarCounty) BVedeno韦杰诺() feud.Barony {
	return c.韦杰诺Vedeno
}
    
var CBalanjar巴兰贾尔 BalanjarCounty = &巴兰贾尔BalanjarCounty{}

func init() {
	f := CBalanjar巴兰贾尔.(*巴兰贾尔BalanjarCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1805",
		Title:     "balanjar",
		TitleName: "巴兰贾尔",
		TitleCode: "c_balanjar",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿尔贡Argun = BArgun阿尔贡
	f.阿尔贡Argun.SetParent(f)
	
	f.巴兰贾尔Balanjar = BBalanjar巴兰贾尔
	f.巴兰贾尔Balanjar.SetParent(f)
	
	f.卡拉塔Karata = BKarata卡拉塔
	f.卡拉塔Karata.SetParent(f)
	
	f.基里Kiri = BKiri基里
	f.基里Kiri.SetParent(f)
	
	f.沙利Shali = BShali沙利
	f.沙利Shali.SetParent(f)
	
	f.沙托伊Shatoy = BShatoy沙托伊
	f.沙托伊Shatoy.SetParent(f)
	
	f.韦杰诺Vedeno = BVedeno韦杰诺
	f.韦杰诺Vedeno.SetParent(f)
	
}
