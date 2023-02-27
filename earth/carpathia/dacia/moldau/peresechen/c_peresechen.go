package peresechen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type PeresechenCounty interface {
    feud.County
    BBalti伯尔兹() 	feud.Barony
    BButuceni布图切尼() 	feud.Barony
    BFalesti弗莱什蒂() 	feud.Barony
    BOrhei奥尔海伊() 	feud.Barony
    BRautu鲁特() 	feud.Barony
    BSoroca索罗卡() 	feud.Barony
    BTuzara图扎拉() 	feud.Barony
}

type 佩列塞琴PeresechenCounty struct {
	feud.BaseCounty
	伯尔兹Balti 	feud.Barony
	布图切尼Butuceni 	feud.Barony
	弗莱什蒂Falesti 	feud.Barony
	奥尔海伊Orhei 	feud.Barony
	鲁特Rautu 	feud.Barony
	索罗卡Soroca 	feud.Barony
	图扎拉Tuzara 	feud.Barony
}

func (c *佩列塞琴PeresechenCounty) BBalti伯尔兹() feud.Barony {
	return c.伯尔兹Balti
}
    
func (c *佩列塞琴PeresechenCounty) BButuceni布图切尼() feud.Barony {
	return c.布图切尼Butuceni
}
    
func (c *佩列塞琴PeresechenCounty) BFalesti弗莱什蒂() feud.Barony {
	return c.弗莱什蒂Falesti
}
    
func (c *佩列塞琴PeresechenCounty) BOrhei奥尔海伊() feud.Barony {
	return c.奥尔海伊Orhei
}
    
func (c *佩列塞琴PeresechenCounty) BRautu鲁特() feud.Barony {
	return c.鲁特Rautu
}
    
func (c *佩列塞琴PeresechenCounty) BSoroca索罗卡() feud.Barony {
	return c.索罗卡Soroca
}
    
func (c *佩列塞琴PeresechenCounty) BTuzara图扎拉() feud.Barony {
	return c.图扎拉Tuzara
}
    
var CPeresechen佩列塞琴 PeresechenCounty = &佩列塞琴PeresechenCounty{}

func init() {
	f := CPeresechen佩列塞琴.(*佩列塞琴PeresechenCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "541",
		Title:     "peresechen",
		TitleName: "佩列塞琴",
		TitleCode: "c_peresechen",
		Baronies:  map[string]feud.Barony{},
	}

	f.伯尔兹Balti = BBalti伯尔兹
	f.伯尔兹Balti.SetParent(f)
	
	f.布图切尼Butuceni = BButuceni布图切尼
	f.布图切尼Butuceni.SetParent(f)
	
	f.弗莱什蒂Falesti = BFalesti弗莱什蒂
	f.弗莱什蒂Falesti.SetParent(f)
	
	f.奥尔海伊Orhei = BOrhei奥尔海伊
	f.奥尔海伊Orhei.SetParent(f)
	
	f.鲁特Rautu = BRautu鲁特
	f.鲁特Rautu.SetParent(f)
	
	f.索罗卡Soroca = BSoroca索罗卡
	f.索罗卡Soroca.SetParent(f)
	
	f.图扎拉Tuzara = BTuzara图扎拉
	f.图扎拉Tuzara.SetParent(f)
	
}
