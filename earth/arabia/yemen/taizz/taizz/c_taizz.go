package taizz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TaizzCounty interface {
    feud.County
    BDhisufal济苏法勒() 	feud.Barony
    BIbb伊卜() 	feud.Barony
    BJibla吉卜拉() 	feud.Barony
    BMocha摩卡() 	feud.Barony
    BPerim丕林() 	feud.Barony
    BShuqra舒拉() 	feud.Barony
    BTaizz塔伊兹() 	feud.Barony
    BZafar宰费尔() 	feud.Barony
}

type 塔伊兹TaizzCounty struct {
	feud.BaseCounty
	济苏法勒Dhisufal 	feud.Barony
	伊卜Ibb 	feud.Barony
	吉卜拉Jibla 	feud.Barony
	摩卡Mocha 	feud.Barony
	丕林Perim 	feud.Barony
	舒拉Shuqra 	feud.Barony
	塔伊兹Taizz 	feud.Barony
	宰费尔Zafar 	feud.Barony
}

func (c *塔伊兹TaizzCounty) BDhisufal济苏法勒() feud.Barony {
	return c.济苏法勒Dhisufal
}
    
func (c *塔伊兹TaizzCounty) BIbb伊卜() feud.Barony {
	return c.伊卜Ibb
}
    
func (c *塔伊兹TaizzCounty) BJibla吉卜拉() feud.Barony {
	return c.吉卜拉Jibla
}
    
func (c *塔伊兹TaizzCounty) BMocha摩卡() feud.Barony {
	return c.摩卡Mocha
}
    
func (c *塔伊兹TaizzCounty) BPerim丕林() feud.Barony {
	return c.丕林Perim
}
    
func (c *塔伊兹TaizzCounty) BShuqra舒拉() feud.Barony {
	return c.舒拉Shuqra
}
    
func (c *塔伊兹TaizzCounty) BTaizz塔伊兹() feud.Barony {
	return c.塔伊兹Taizz
}
    
func (c *塔伊兹TaizzCounty) BZafar宰费尔() feud.Barony {
	return c.宰费尔Zafar
}
    
var CTaizz塔伊兹 TaizzCounty = &塔伊兹TaizzCounty{}

func init() {
	f := CTaizz塔伊兹.(*塔伊兹TaizzCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "859",
		Title:     "taizz",
		TitleName: "塔伊兹",
		TitleCode: "c_taizz",
		Baronies:  map[string]feud.Barony{},
	}

	f.济苏法勒Dhisufal = BDhisufal济苏法勒
	f.济苏法勒Dhisufal.SetParent(f)
	
	f.伊卜Ibb = BIbb伊卜
	f.伊卜Ibb.SetParent(f)
	
	f.吉卜拉Jibla = BJibla吉卜拉
	f.吉卜拉Jibla.SetParent(f)
	
	f.摩卡Mocha = BMocha摩卡
	f.摩卡Mocha.SetParent(f)
	
	f.丕林Perim = BPerim丕林
	f.丕林Perim.SetParent(f)
	
	f.舒拉Shuqra = BShuqra舒拉
	f.舒拉Shuqra.SetParent(f)
	
	f.塔伊兹Taizz = BTaizz塔伊兹
	f.塔伊兹Taizz.SetParent(f)
	
	f.宰费尔Zafar = BZafar宰费尔
	f.宰费尔Zafar.SetParent(f)
	
}
