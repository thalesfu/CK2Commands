package khiva

import (
	"github.com/thalesfu/CK2Commands/earth/turkestan/khiva/ferghana"
	"github.com/thalesfu/CK2Commands/earth/turkestan/khiva/khiva"
	"github.com/thalesfu/CK2Commands/earth/turkestan/khiva/khuttal"
	"github.com/thalesfu/CK2Commands/earth/turkestan/khiva/samarkand"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KhivaKingdom interface {
    feud.Kingdom
    DFerghana费尔干纳() 	ferghana.FerghanaDuke
    DKhiva希瓦() 	khiva.KhivaDuke
    DKhuttal珂咄罗() 	khuttal.KhuttalDuke
    DSamarkand撒马尔罕() 	samarkand.SamarkandDuke
}

type 河中KhivaKingdom struct {
	feud.BaseKingdom
	费尔干纳Ferghana 	ferghana.FerghanaDuke
	希瓦Khiva 	khiva.KhivaDuke
	珂咄罗Khuttal 	khuttal.KhuttalDuke
	撒马尔罕Samarkand 	samarkand.SamarkandDuke
}

func (k *河中KhivaKingdom) DFerghana费尔干纳() ferghana.FerghanaDuke {
	return k.费尔干纳Ferghana
}
    
func (k *河中KhivaKingdom) DKhiva希瓦() khiva.KhivaDuke {
	return k.希瓦Khiva
}
    
func (k *河中KhivaKingdom) DKhuttal珂咄罗() khuttal.KhuttalDuke {
	return k.珂咄罗Khuttal
}
    
func (k *河中KhivaKingdom) DSamarkand撒马尔罕() samarkand.SamarkandDuke {
	return k.撒马尔罕Samarkand
}
    
var KKhiva河中 KhivaKingdom = &河中KhivaKingdom{}

func init() {
	f := KKhiva河中.(*河中KhivaKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "khiva",
		TitleName: "河中",
		TitleCode: "k_khiva",
		Dukes:  map[string]feud.Duke{},
	}

	f.费尔干纳Ferghana = ferghana.DFerghana费尔干纳
	f.费尔干纳Ferghana.SetParent(f)
	
	f.希瓦Khiva = khiva.DKhiva希瓦
	f.希瓦Khiva.SetParent(f)
	
	f.珂咄罗Khuttal = khuttal.DKhuttal珂咄罗
	f.珂咄罗Khuttal.SetParent(f)
	
	f.撒马尔罕Samarkand = samarkand.DSamarkand撒马尔罕
	f.撒马尔罕Samarkand.SetParent(f)
	
}
