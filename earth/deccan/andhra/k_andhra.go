package andhra

import (
	"github.com/thalesfu/CK2Commands/earth/deccan/andhra/udayagiri"
	"github.com/thalesfu/CK2Commands/earth/deccan/andhra/vengi"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AndhraKingdom interface {
    feud.Kingdom
    DUdayagiri优陀耶耆厘() 	udayagiri.UdayagiriDuke
    DVengi瓶耆() 	vengi.VengiDuke
}

type 案达罗AndhraKingdom struct {
	feud.BaseKingdom
	优陀耶耆厘Udayagiri 	udayagiri.UdayagiriDuke
	瓶耆Vengi 	vengi.VengiDuke
}

func (k *案达罗AndhraKingdom) DUdayagiri优陀耶耆厘() udayagiri.UdayagiriDuke {
	return k.优陀耶耆厘Udayagiri
}
    
func (k *案达罗AndhraKingdom) DVengi瓶耆() vengi.VengiDuke {
	return k.瓶耆Vengi
}
    
var KAndhra案达罗 AndhraKingdom = &案达罗AndhraKingdom{}

func init() {
	f := KAndhra案达罗.(*案达罗AndhraKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "andhra",
		TitleName: "案达罗",
		TitleCode: "k_andhra",
		Dukes:  map[string]feud.Duke{},
	}

	f.优陀耶耆厘Udayagiri = udayagiri.DUdayagiri优陀耶耆厘
	f.优陀耶耆厘Udayagiri.SetParent(f)
	
	f.瓶耆Vengi = vengi.DVengi瓶耆
	f.瓶耆Vengi.SetParent(f)
	
}
