package cuman

import (
	"github.com/thalesfu/CK2Commands/earth/tartaria/cuman/irtysh"
	"github.com/thalesfu/CK2Commands/earth/tartaria/cuman/ishim"
	"github.com/thalesfu/CK2Commands/earth/tartaria/cuman/kazakh"
	"github.com/thalesfu/CK2Commands/earth/tartaria/cuman/ob"
	"github.com/thalesfu/CK2Commands/earth/tartaria/cuman/tarbagatai"
	"github.com/thalesfu/CK2Commands/earth/tartaria/cuman/turgay"
	"github.com/thalesfu/CK2Commands/earth/tartaria/cuman/ubagan"
	"github.com/thalesfu/CK2Commands/earth/tartaria/cuman/yaik"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type CumanKingdom interface {
    feud.Kingdom
    DIrtysh曳咥() 	irtysh.IrtyshDuke
    DIshim伊希姆() 	ishim.IshimDuke
    DKazakh哈萨克() 	kazakh.KazakhDuke
    DOb鄂毕() 	ob.ObDuke
    DTarbagatai塔尔巴哈台() 	tarbagatai.TarbagataiDuke
    DTurgay图尔盖() 	turgay.TurgayDuke
    DUbagan乌巴甘() 	ubagan.UbaganDuke
    DYaik亚伊克() 	yaik.YaikDuke
}

type 库曼CumanKingdom struct {
	feud.BaseKingdom
	曳咥Irtysh 	irtysh.IrtyshDuke
	伊希姆Ishim 	ishim.IshimDuke
	哈萨克Kazakh 	kazakh.KazakhDuke
	鄂毕Ob 	ob.ObDuke
	塔尔巴哈台Tarbagatai 	tarbagatai.TarbagataiDuke
	图尔盖Turgay 	turgay.TurgayDuke
	乌巴甘Ubagan 	ubagan.UbaganDuke
	亚伊克Yaik 	yaik.YaikDuke
}

func (k *库曼CumanKingdom) DIrtysh曳咥() irtysh.IrtyshDuke {
	return k.曳咥Irtysh
}
    
func (k *库曼CumanKingdom) DIshim伊希姆() ishim.IshimDuke {
	return k.伊希姆Ishim
}
    
func (k *库曼CumanKingdom) DKazakh哈萨克() kazakh.KazakhDuke {
	return k.哈萨克Kazakh
}
    
func (k *库曼CumanKingdom) DOb鄂毕() ob.ObDuke {
	return k.鄂毕Ob
}
    
func (k *库曼CumanKingdom) DTarbagatai塔尔巴哈台() tarbagatai.TarbagataiDuke {
	return k.塔尔巴哈台Tarbagatai
}
    
func (k *库曼CumanKingdom) DTurgay图尔盖() turgay.TurgayDuke {
	return k.图尔盖Turgay
}
    
func (k *库曼CumanKingdom) DUbagan乌巴甘() ubagan.UbaganDuke {
	return k.乌巴甘Ubagan
}
    
func (k *库曼CumanKingdom) DYaik亚伊克() yaik.YaikDuke {
	return k.亚伊克Yaik
}
    
var KCuman库曼 CumanKingdom = &库曼CumanKingdom{}

func init() {
	f := KCuman库曼.(*库曼CumanKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "cuman",
		TitleName: "库曼",
		TitleCode: "k_cuman",
		Dukes:  map[string]feud.Duke{},
	}

	f.曳咥Irtysh = irtysh.DIrtysh曳咥
	f.曳咥Irtysh.SetParent(f)
	
	f.伊希姆Ishim = ishim.DIshim伊希姆
	f.伊希姆Ishim.SetParent(f)
	
	f.哈萨克Kazakh = kazakh.DKazakh哈萨克
	f.哈萨克Kazakh.SetParent(f)
	
	f.鄂毕Ob = ob.DOb鄂毕
	f.鄂毕Ob.SetParent(f)
	
	f.塔尔巴哈台Tarbagatai = tarbagatai.DTarbagatai塔尔巴哈台
	f.塔尔巴哈台Tarbagatai.SetParent(f)
	
	f.图尔盖Turgay = turgay.DTurgay图尔盖
	f.图尔盖Turgay.SetParent(f)
	
	f.乌巴甘Ubagan = ubagan.DUbagan乌巴甘
	f.乌巴甘Ubagan.SetParent(f)
	
	f.亚伊克Yaik = yaik.DYaik亚伊克
	f.亚伊克Yaik.SetParent(f)
	
}
