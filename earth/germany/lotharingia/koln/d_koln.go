package koln

import (
	"github.com/thalesfu/CK2Commands/earth/germany/lotharingia/koln/berg"
	"github.com/thalesfu/CK2Commands/earth/germany/lotharingia/koln/kleve"
	"github.com/thalesfu/CK2Commands/earth/germany/lotharingia/koln/koln"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KolnDuke interface {
    feud.Duke
    CBerg贝格() 	berg.BergCounty
    CKleve克莱沃() 	kleve.KleveCounty
    CKoln科隆() 	koln.KolnCounty
}

type 科隆KolnDuke struct {
	feud.BaseDuke
	贝格Berg 	berg.BergCounty
	克莱沃Kleve 	kleve.KleveCounty
	科隆Koln 	koln.KolnCounty
}

func (d *科隆KolnDuke) CBerg贝格() berg.BergCounty {
	return d.贝格Berg
}
    
func (d *科隆KolnDuke) CKleve克莱沃() kleve.KleveCounty {
	return d.克莱沃Kleve
}
    
func (d *科隆KolnDuke) CKoln科隆() koln.KolnCounty {
	return d.科隆Koln
}
    
var DKoln科隆 KolnDuke = &科隆KolnDuke{}

func init() {
	f := DKoln科隆.(*科隆KolnDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "koln",
		TitleName: "科隆",
		TitleCode: "d_koln",
		Counties:  map[string]feud.County{},
	}

	f.贝格Berg = berg.CBerg贝格
	f.贝格Berg.SetParent(f)
	
	f.克莱沃Kleve = kleve.CKleve克莱沃
	f.克莱沃Kleve.SetParent(f)
	
	f.科隆Koln = koln.CKoln科隆
	f.科隆Koln.SetParent(f)
	
}
