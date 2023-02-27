package zhetysu

import (
	"github.com/thalesfu/CK2Commands/earth/turkestan/zhetysu/chuy"
	"github.com/thalesfu/CK2Commands/earth/turkestan/zhetysu/ili"
	"github.com/thalesfu/CK2Commands/earth/turkestan/zhetysu/zhetysu"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ZhetysuKingdom interface {
    feud.Kingdom
    DChuy垂河() 	chuy.ChuyDuke
    DIli伊丽() 	ili.IliDuke
    DZhetysu七河() 	zhetysu.ZhetysuDuke
}

type 七河ZhetysuKingdom struct {
	feud.BaseKingdom
	垂河Chuy 	chuy.ChuyDuke
	伊丽Ili 	ili.IliDuke
	七河Zhetysu 	zhetysu.ZhetysuDuke
}

func (k *七河ZhetysuKingdom) DChuy垂河() chuy.ChuyDuke {
	return k.垂河Chuy
}
    
func (k *七河ZhetysuKingdom) DIli伊丽() ili.IliDuke {
	return k.伊丽Ili
}
    
func (k *七河ZhetysuKingdom) DZhetysu七河() zhetysu.ZhetysuDuke {
	return k.七河Zhetysu
}
    
var KZhetysu七河 ZhetysuKingdom = &七河ZhetysuKingdom{}

func init() {
	f := KZhetysu七河.(*七河ZhetysuKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "zhetysu",
		TitleName: "七河",
		TitleCode: "k_zhetysu",
		Dukes:  map[string]feud.Duke{},
	}

	f.垂河Chuy = chuy.DChuy垂河
	f.垂河Chuy.SetParent(f)
	
	f.伊丽Ili = ili.DIli伊丽
	f.伊丽Ili.SetParent(f)
	
	f.七河Zhetysu = zhetysu.DZhetysu七河
	f.七河Zhetysu.SetParent(f)
	
}
