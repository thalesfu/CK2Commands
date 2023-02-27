package kanem

import (
	"github.com/thalesfu/CK2Commands/earth/kanem/kanem/bornu"
	"github.com/thalesfu/CK2Commands/earth/kanem/kanem/fezzan"
	"github.com/thalesfu/CK2Commands/earth/kanem/kanem/kanem"
	"github.com/thalesfu/CK2Commands/earth/kanem/kanem/wadai"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KanemKingdom interface {
    feud.Kingdom
    DBornu博尔努() 	bornu.BornuDuke
    DFezzan费赞() 	fezzan.FezzanDuke
    DKanem加涅姆() 	kanem.KanemDuke
    DWadai瓦达伊() 	wadai.WadaiDuke
}

type 加涅姆KanemKingdom struct {
	feud.BaseKingdom
	博尔努Bornu 	bornu.BornuDuke
	费赞Fezzan 	fezzan.FezzanDuke
	加涅姆Kanem 	kanem.KanemDuke
	瓦达伊Wadai 	wadai.WadaiDuke
}

func (k *加涅姆KanemKingdom) DBornu博尔努() bornu.BornuDuke {
	return k.博尔努Bornu
}
    
func (k *加涅姆KanemKingdom) DFezzan费赞() fezzan.FezzanDuke {
	return k.费赞Fezzan
}
    
func (k *加涅姆KanemKingdom) DKanem加涅姆() kanem.KanemDuke {
	return k.加涅姆Kanem
}
    
func (k *加涅姆KanemKingdom) DWadai瓦达伊() wadai.WadaiDuke {
	return k.瓦达伊Wadai
}
    
var KKanem加涅姆 KanemKingdom = &加涅姆KanemKingdom{}

func init() {
	f := KKanem加涅姆.(*加涅姆KanemKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "kanem",
		TitleName: "加涅姆",
		TitleCode: "k_kanem",
		Dukes:  map[string]feud.Duke{},
	}

	f.博尔努Bornu = bornu.DBornu博尔努
	f.博尔努Bornu.SetParent(f)
	
	f.费赞Fezzan = fezzan.DFezzan费赞
	f.费赞Fezzan.SetParent(f)
	
	f.加涅姆Kanem = kanem.DKanem加涅姆
	f.加涅姆Kanem.SetParent(f)
	
	f.瓦达伊Wadai = wadai.DWadai瓦达伊
	f.瓦达伊Wadai.SetParent(f)
	
}
