package tannu_ola

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type Tannu_olaCounty interface {
    feud.County
    BBay_tal拜塔尔() 	feud.Barony
    BChadan_tannu_ola恰丹() 	feud.Barony
    BKara_khol卡拉霍尔() 	feud.Barony
    BSug_aksy苏格阿克瑟() 	feud.Barony
    BSut_khol苏特霍尔() 	feud.Barony
    BTannu_ola唐麓岭() 	feud.Barony
    BTeeli泰利() 	feud.Barony
}

type 唐麓岭Tannu_olaCounty struct {
	feud.BaseCounty
	拜塔尔Bay_tal 	feud.Barony
	恰丹Chadan_tannu_ola 	feud.Barony
	卡拉霍尔Kara_khol 	feud.Barony
	苏格阿克瑟Sug_aksy 	feud.Barony
	苏特霍尔Sut_khol 	feud.Barony
	唐麓岭Tannu_ola 	feud.Barony
	泰利Teeli 	feud.Barony
}

func (c *唐麓岭Tannu_olaCounty) BBay_tal拜塔尔() feud.Barony {
	return c.拜塔尔Bay_tal
}
    
func (c *唐麓岭Tannu_olaCounty) BChadan_tannu_ola恰丹() feud.Barony {
	return c.恰丹Chadan_tannu_ola
}
    
func (c *唐麓岭Tannu_olaCounty) BKara_khol卡拉霍尔() feud.Barony {
	return c.卡拉霍尔Kara_khol
}
    
func (c *唐麓岭Tannu_olaCounty) BSug_aksy苏格阿克瑟() feud.Barony {
	return c.苏格阿克瑟Sug_aksy
}
    
func (c *唐麓岭Tannu_olaCounty) BSut_khol苏特霍尔() feud.Barony {
	return c.苏特霍尔Sut_khol
}
    
func (c *唐麓岭Tannu_olaCounty) BTannu_ola唐麓岭() feud.Barony {
	return c.唐麓岭Tannu_ola
}
    
func (c *唐麓岭Tannu_olaCounty) BTeeli泰利() feud.Barony {
	return c.泰利Teeli
}
    
var CTannu_ola唐麓岭 Tannu_olaCounty = &唐麓岭Tannu_olaCounty{}

func init() {
	f := CTannu_ola唐麓岭.(*唐麓岭Tannu_olaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1905",
		Title:     "tannu_ola",
		TitleName: "唐麓岭",
		TitleCode: "c_tannu_ola",
		Baronies:  map[string]feud.Barony{},
	}

	f.拜塔尔Bay_tal = BBay_tal拜塔尔
	f.拜塔尔Bay_tal.SetParent(f)
	
	f.恰丹Chadan_tannu_ola = BChadan_tannu_ola恰丹
	f.恰丹Chadan_tannu_ola.SetParent(f)
	
	f.卡拉霍尔Kara_khol = BKara_khol卡拉霍尔
	f.卡拉霍尔Kara_khol.SetParent(f)
	
	f.苏格阿克瑟Sug_aksy = BSug_aksy苏格阿克瑟
	f.苏格阿克瑟Sug_aksy.SetParent(f)
	
	f.苏特霍尔Sut_khol = BSut_khol苏特霍尔
	f.苏特霍尔Sut_khol.SetParent(f)
	
	f.唐麓岭Tannu_ola = BTannu_ola唐麓岭
	f.唐麓岭Tannu_ola.SetParent(f)
	
	f.泰利Teeli = BTeeli泰利
	f.泰利Teeli.SetParent(f)
	
}
