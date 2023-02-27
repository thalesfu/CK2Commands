package maharastra

import (
	"github.com/thalesfu/CK2Commands/earth/deccan/maharastra/devagiri"
	"github.com/thalesfu/CK2Commands/earth/deccan/maharastra/konkana"
	"github.com/thalesfu/CK2Commands/earth/deccan/maharastra/nasikya"
	"github.com/thalesfu/CK2Commands/earth/deccan/maharastra/rattapadi"
	"github.com/thalesfu/CK2Commands/earth/deccan/maharastra/vidharba"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MaharastraKingdom interface {
    feud.Kingdom
    DDevagiri提婆耆厘() 	devagiri.DevagiriDuke
    DKonkana恭建那() 	konkana.KonkanaDuke
    DNasikya那私迦() 	nasikya.NasikyaDuke
    DRattapadi剌吒波地() 	rattapadi.RattapadiDuke
    DVidharba毗达婆() 	vidharba.VidharbaDuke
}

type 摩诃剌侘MaharastraKingdom struct {
	feud.BaseKingdom
	提婆耆厘Devagiri 	devagiri.DevagiriDuke
	恭建那Konkana 	konkana.KonkanaDuke
	那私迦Nasikya 	nasikya.NasikyaDuke
	剌吒波地Rattapadi 	rattapadi.RattapadiDuke
	毗达婆Vidharba 	vidharba.VidharbaDuke
}

func (k *摩诃剌侘MaharastraKingdom) DDevagiri提婆耆厘() devagiri.DevagiriDuke {
	return k.提婆耆厘Devagiri
}
    
func (k *摩诃剌侘MaharastraKingdom) DKonkana恭建那() konkana.KonkanaDuke {
	return k.恭建那Konkana
}
    
func (k *摩诃剌侘MaharastraKingdom) DNasikya那私迦() nasikya.NasikyaDuke {
	return k.那私迦Nasikya
}
    
func (k *摩诃剌侘MaharastraKingdom) DRattapadi剌吒波地() rattapadi.RattapadiDuke {
	return k.剌吒波地Rattapadi
}
    
func (k *摩诃剌侘MaharastraKingdom) DVidharba毗达婆() vidharba.VidharbaDuke {
	return k.毗达婆Vidharba
}
    
var KMaharastra摩诃剌侘 MaharastraKingdom = &摩诃剌侘MaharastraKingdom{}

func init() {
	f := KMaharastra摩诃剌侘.(*摩诃剌侘MaharastraKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "maharastra",
		TitleName: "摩诃剌侘",
		TitleCode: "k_maharastra",
		Dukes:  map[string]feud.Duke{},
	}

	f.提婆耆厘Devagiri = devagiri.DDevagiri提婆耆厘
	f.提婆耆厘Devagiri.SetParent(f)
	
	f.恭建那Konkana = konkana.DKonkana恭建那
	f.恭建那Konkana.SetParent(f)
	
	f.那私迦Nasikya = nasikya.DNasikya那私迦
	f.那私迦Nasikya.SetParent(f)
	
	f.剌吒波地Rattapadi = rattapadi.DRattapadi剌吒波地
	f.剌吒波地Rattapadi.SetParent(f)
	
	f.毗达婆Vidharba = vidharba.DVidharba毗达婆
	f.毗达婆Vidharba.SetParent(f)
	
}
