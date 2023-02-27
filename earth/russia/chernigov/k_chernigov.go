package chernigov

import (
	"github.com/thalesfu/CK2Commands/earth/russia/chernigov/chernigov"
	"github.com/thalesfu/CK2Commands/earth/russia/chernigov/karachev"
	"github.com/thalesfu/CK2Commands/earth/russia/chernigov/murom"
	"github.com/thalesfu/CK2Commands/earth/russia/chernigov/novosil"
	"github.com/thalesfu/CK2Commands/earth/russia/chernigov/ryazan"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ChernigovKingdom interface {
    feud.Kingdom
    DChernigov切尔尼戈夫() 	chernigov.ChernigovDuke
    DKarachev卡拉切夫() 	karachev.KarachevDuke
    DMurom穆罗姆() 	murom.MuromDuke
    DNovosil诺沃西利() 	novosil.NovosilDuke
    DRyazan梁赞() 	ryazan.RyazanDuke
}

type 切尔尼戈夫ChernigovKingdom struct {
	feud.BaseKingdom
	切尔尼戈夫Chernigov 	chernigov.ChernigovDuke
	卡拉切夫Karachev 	karachev.KarachevDuke
	穆罗姆Murom 	murom.MuromDuke
	诺沃西利Novosil 	novosil.NovosilDuke
	梁赞Ryazan 	ryazan.RyazanDuke
}

func (k *切尔尼戈夫ChernigovKingdom) DChernigov切尔尼戈夫() chernigov.ChernigovDuke {
	return k.切尔尼戈夫Chernigov
}
    
func (k *切尔尼戈夫ChernigovKingdom) DKarachev卡拉切夫() karachev.KarachevDuke {
	return k.卡拉切夫Karachev
}
    
func (k *切尔尼戈夫ChernigovKingdom) DMurom穆罗姆() murom.MuromDuke {
	return k.穆罗姆Murom
}
    
func (k *切尔尼戈夫ChernigovKingdom) DNovosil诺沃西利() novosil.NovosilDuke {
	return k.诺沃西利Novosil
}
    
func (k *切尔尼戈夫ChernigovKingdom) DRyazan梁赞() ryazan.RyazanDuke {
	return k.梁赞Ryazan
}
    
var KChernigov切尔尼戈夫 ChernigovKingdom = &切尔尼戈夫ChernigovKingdom{}

func init() {
	f := KChernigov切尔尼戈夫.(*切尔尼戈夫ChernigovKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "chernigov",
		TitleName: "切尔尼戈夫",
		TitleCode: "k_chernigov",
		Dukes:  map[string]feud.Duke{},
	}

	f.切尔尼戈夫Chernigov = chernigov.DChernigov切尔尼戈夫
	f.切尔尼戈夫Chernigov.SetParent(f)
	
	f.卡拉切夫Karachev = karachev.DKarachev卡拉切夫
	f.卡拉切夫Karachev.SetParent(f)
	
	f.穆罗姆Murom = murom.DMurom穆罗姆
	f.穆罗姆Murom.SetParent(f)
	
	f.诺沃西利Novosil = novosil.DNovosil诺沃西利
	f.诺沃西利Novosil.SetParent(f)
	
	f.梁赞Ryazan = ryazan.DRyazan梁赞
	f.梁赞Ryazan.SetParent(f)
	
}
