package hausaland

import (
	"github.com/thalesfu/CK2Commands/earth/kanem/hausaland/air"
	"github.com/thalesfu/CK2Commands/earth/kanem/hausaland/hausaland"
	"github.com/thalesfu/CK2Commands/earth/kanem/hausaland/kebbi"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type HausalandKingdom interface {
    feud.Kingdom
    DAir艾尔() 	air.AirDuke
    DHausaland豪萨() 	hausaland.HausalandDuke
    DKebbi凯比() 	kebbi.KebbiDuke
}

type 豪萨HausalandKingdom struct {
	feud.BaseKingdom
	艾尔Air 	air.AirDuke
	豪萨Hausaland 	hausaland.HausalandDuke
	凯比Kebbi 	kebbi.KebbiDuke
}

func (k *豪萨HausalandKingdom) DAir艾尔() air.AirDuke {
	return k.艾尔Air
}
    
func (k *豪萨HausalandKingdom) DHausaland豪萨() hausaland.HausalandDuke {
	return k.豪萨Hausaland
}
    
func (k *豪萨HausalandKingdom) DKebbi凯比() kebbi.KebbiDuke {
	return k.凯比Kebbi
}
    
var KHausaland豪萨 HausalandKingdom = &豪萨HausalandKingdom{}

func init() {
	f := KHausaland豪萨.(*豪萨HausalandKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "hausaland",
		TitleName: "豪萨",
		TitleCode: "k_hausaland",
		Dukes:  map[string]feud.Duke{},
	}

	f.艾尔Air = air.DAir艾尔
	f.艾尔Air.SetParent(f)
	
	f.豪萨Hausaland = hausaland.DHausaland豪萨
	f.豪萨Hausaland.SetParent(f)
	
	f.凯比Kebbi = kebbi.DKebbi凯比
	f.凯比Kebbi.SetParent(f)
	
}
