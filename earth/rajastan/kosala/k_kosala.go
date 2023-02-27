package kosala

import (
	"github.com/thalesfu/CK2Commands/earth/rajastan/kosala/jejakabhukti"
	"github.com/thalesfu/CK2Commands/earth/rajastan/kosala/kanyakubja"
	"github.com/thalesfu/CK2Commands/earth/rajastan/kosala/saryupara"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KosalaKingdom interface {
    feud.Kingdom
    DJejakabhukti掷枳陀() 	jejakabhukti.JejakabhuktiDuke
    DKanyakubja曲女城() 	kanyakubja.KanyakubjaDuke
    DSaryupara萨罗瑜波罗() 	saryupara.SaryuparaDuke
}

type 憍萨罗KosalaKingdom struct {
	feud.BaseKingdom
	掷枳陀Jejakabhukti 	jejakabhukti.JejakabhuktiDuke
	曲女城Kanyakubja 	kanyakubja.KanyakubjaDuke
	萨罗瑜波罗Saryupara 	saryupara.SaryuparaDuke
}

func (k *憍萨罗KosalaKingdom) DJejakabhukti掷枳陀() jejakabhukti.JejakabhuktiDuke {
	return k.掷枳陀Jejakabhukti
}
    
func (k *憍萨罗KosalaKingdom) DKanyakubja曲女城() kanyakubja.KanyakubjaDuke {
	return k.曲女城Kanyakubja
}
    
func (k *憍萨罗KosalaKingdom) DSaryupara萨罗瑜波罗() saryupara.SaryuparaDuke {
	return k.萨罗瑜波罗Saryupara
}
    
var KKosala憍萨罗 KosalaKingdom = &憍萨罗KosalaKingdom{}

func init() {
	f := KKosala憍萨罗.(*憍萨罗KosalaKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "kosala",
		TitleName: "憍萨罗",
		TitleCode: "k_kosala",
		Dukes:  map[string]feud.Duke{},
	}

	f.掷枳陀Jejakabhukti = jejakabhukti.DJejakabhukti掷枳陀
	f.掷枳陀Jejakabhukti.SetParent(f)
	
	f.曲女城Kanyakubja = kanyakubja.DKanyakubja曲女城
	f.曲女城Kanyakubja.SetParent(f)
	
	f.萨罗瑜波罗Saryupara = saryupara.DSaryupara萨罗瑜波罗
	f.萨罗瑜波罗Saryupara.SetParent(f)
	
}
