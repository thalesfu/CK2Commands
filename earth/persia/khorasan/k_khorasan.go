package khorasan

import (
	"github.com/thalesfu/CK2Commands/earth/persia/khorasan/balkh"
	"github.com/thalesfu/CK2Commands/earth/persia/khorasan/herat"
	"github.com/thalesfu/CK2Commands/earth/persia/khorasan/khorasan"
	"github.com/thalesfu/CK2Commands/earth/persia/khorasan/merv"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KhorasanKingdom interface {
    feud.Kingdom
    DBalkh缚喝() 	balkh.BalkhDuke
    DHerat亦鲁() 	herat.HeratDuke
    DKhorasan呼罗珊() 	khorasan.KhorasanDuke
    DMerv木鹿() 	merv.MervDuke
}

type 呼罗珊KhorasanKingdom struct {
	feud.BaseKingdom
	缚喝Balkh 	balkh.BalkhDuke
	亦鲁Herat 	herat.HeratDuke
	呼罗珊Khorasan 	khorasan.KhorasanDuke
	木鹿Merv 	merv.MervDuke
}

func (k *呼罗珊KhorasanKingdom) DBalkh缚喝() balkh.BalkhDuke {
	return k.缚喝Balkh
}
    
func (k *呼罗珊KhorasanKingdom) DHerat亦鲁() herat.HeratDuke {
	return k.亦鲁Herat
}
    
func (k *呼罗珊KhorasanKingdom) DKhorasan呼罗珊() khorasan.KhorasanDuke {
	return k.呼罗珊Khorasan
}
    
func (k *呼罗珊KhorasanKingdom) DMerv木鹿() merv.MervDuke {
	return k.木鹿Merv
}
    
var KKhorasan呼罗珊 KhorasanKingdom = &呼罗珊KhorasanKingdom{}

func init() {
	f := KKhorasan呼罗珊.(*呼罗珊KhorasanKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "khorasan",
		TitleName: "呼罗珊",
		TitleCode: "k_khorasan",
		Dukes:  map[string]feud.Duke{},
	}

	f.缚喝Balkh = balkh.DBalkh缚喝
	f.缚喝Balkh.SetParent(f)
	
	f.亦鲁Herat = herat.DHerat亦鲁
	f.亦鲁Herat.SetParent(f)
	
	f.呼罗珊Khorasan = khorasan.DKhorasan呼罗珊
	f.呼罗珊Khorasan.SetParent(f)
	
	f.木鹿Merv = merv.DMerv木鹿
	f.木鹿Merv.SetParent(f)
	
}
