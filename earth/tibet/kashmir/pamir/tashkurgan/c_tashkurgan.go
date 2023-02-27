package tashkurgan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TashkurganCounty interface {
    feud.County
    BBulongke布隆克() 	feud.Barony
    BHangdi行地() 	feud.Barony
    BJianggala将嘎拉() 	feud.Barony
    BKaladong卡拉懂() 	feud.Barony
    BQiate恰特() 	feud.Barony
    BTashkurgan塔什库尔干() 	feud.Barony
    BZankan赞坎() 	feud.Barony
}

type 塔什库尔干TashkurganCounty struct {
	feud.BaseCounty
	布隆克Bulongke 	feud.Barony
	行地Hangdi 	feud.Barony
	将嘎拉Jianggala 	feud.Barony
	卡拉懂Kaladong 	feud.Barony
	恰特Qiate 	feud.Barony
	塔什库尔干Tashkurgan 	feud.Barony
	赞坎Zankan 	feud.Barony
}

func (c *塔什库尔干TashkurganCounty) BBulongke布隆克() feud.Barony {
	return c.布隆克Bulongke
}
    
func (c *塔什库尔干TashkurganCounty) BHangdi行地() feud.Barony {
	return c.行地Hangdi
}
    
func (c *塔什库尔干TashkurganCounty) BJianggala将嘎拉() feud.Barony {
	return c.将嘎拉Jianggala
}
    
func (c *塔什库尔干TashkurganCounty) BKaladong卡拉懂() feud.Barony {
	return c.卡拉懂Kaladong
}
    
func (c *塔什库尔干TashkurganCounty) BQiate恰特() feud.Barony {
	return c.恰特Qiate
}
    
func (c *塔什库尔干TashkurganCounty) BTashkurgan塔什库尔干() feud.Barony {
	return c.塔什库尔干Tashkurgan
}
    
func (c *塔什库尔干TashkurganCounty) BZankan赞坎() feud.Barony {
	return c.赞坎Zankan
}
    
var CTashkurgan塔什库尔干 TashkurganCounty = &塔什库尔干TashkurganCounty{}

func init() {
	f := CTashkurgan塔什库尔干.(*塔什库尔干TashkurganCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1465",
		Title:     "tashkurgan",
		TitleName: "塔什库尔干",
		TitleCode: "c_tashkurgan",
		Baronies:  map[string]feud.Barony{},
	}

	f.布隆克Bulongke = BBulongke布隆克
	f.布隆克Bulongke.SetParent(f)
	
	f.行地Hangdi = BHangdi行地
	f.行地Hangdi.SetParent(f)
	
	f.将嘎拉Jianggala = BJianggala将嘎拉
	f.将嘎拉Jianggala.SetParent(f)
	
	f.卡拉懂Kaladong = BKaladong卡拉懂
	f.卡拉懂Kaladong.SetParent(f)
	
	f.恰特Qiate = BQiate恰特
	f.恰特Qiate.SetParent(f)
	
	f.塔什库尔干Tashkurgan = BTashkurgan塔什库尔干
	f.塔什库尔干Tashkurgan.SetParent(f)
	
	f.赞坎Zankan = BZankan赞坎
	f.赞坎Zankan.SetParent(f)
	
}
