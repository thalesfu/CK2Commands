package badajoz

import (
	"github.com/thalesfu/CK2Commands/earth/spain/badajoz/algarve"
	"github.com/thalesfu/CK2Commands/earth/spain/badajoz/badajoz"
	"github.com/thalesfu/CK2Commands/earth/spain/badajoz/balata"
	"github.com/thalesfu/CK2Commands/earth/spain/badajoz/beja"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BadajozKingdom interface {
    feud.Kingdom
    DAlgarve阿尔加维() 	algarve.AlgarveDuke
    DBadajoz巴达霍斯() 	badajoz.BadajozDuke
    DBalata里斯本() 	balata.BalataDuke
    DBeja贝雅() 	beja.BejaDuke
}

type 巴达霍斯BadajozKingdom struct {
	feud.BaseKingdom
	阿尔加维Algarve 	algarve.AlgarveDuke
	巴达霍斯Badajoz 	badajoz.BadajozDuke
	里斯本Balata 	balata.BalataDuke
	贝雅Beja 	beja.BejaDuke
}

func (k *巴达霍斯BadajozKingdom) DAlgarve阿尔加维() algarve.AlgarveDuke {
	return k.阿尔加维Algarve
}
    
func (k *巴达霍斯BadajozKingdom) DBadajoz巴达霍斯() badajoz.BadajozDuke {
	return k.巴达霍斯Badajoz
}
    
func (k *巴达霍斯BadajozKingdom) DBalata里斯本() balata.BalataDuke {
	return k.里斯本Balata
}
    
func (k *巴达霍斯BadajozKingdom) DBeja贝雅() beja.BejaDuke {
	return k.贝雅Beja
}
    
var KBadajoz巴达霍斯 BadajozKingdom = &巴达霍斯BadajozKingdom{}

func init() {
	f := KBadajoz巴达霍斯.(*巴达霍斯BadajozKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "badajoz",
		TitleName: "巴达霍斯",
		TitleCode: "k_badajoz",
		Dukes:  map[string]feud.Duke{},
	}

	f.阿尔加维Algarve = algarve.DAlgarve阿尔加维
	f.阿尔加维Algarve.SetParent(f)
	
	f.巴达霍斯Badajoz = badajoz.DBadajoz巴达霍斯
	f.巴达霍斯Badajoz.SetParent(f)
	
	f.里斯本Balata = balata.DBalata里斯本
	f.里斯本Balata.SetParent(f)
	
	f.贝雅Beja = beja.DBeja贝雅
	f.贝雅Beja.SetParent(f)
	
}
