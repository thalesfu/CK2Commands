package pamir

import (
	"github.com/thalesfu/CK2Commands/earth/tibet/kashmir/pamir/pamir"
	"github.com/thalesfu/CK2Commands/earth/tibet/kashmir/pamir/tashkurgan"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type PamirDuke interface {
    feud.Duke
    CPamir播密() 	pamir.PamirCounty
    CTashkurgan塔什库尔干() 	tashkurgan.TashkurganCounty
}

type 播密PamirDuke struct {
	feud.BaseDuke
	播密Pamir 	pamir.PamirCounty
	塔什库尔干Tashkurgan 	tashkurgan.TashkurganCounty
}

func (d *播密PamirDuke) CPamir播密() pamir.PamirCounty {
	return d.播密Pamir
}
    
func (d *播密PamirDuke) CTashkurgan塔什库尔干() tashkurgan.TashkurganCounty {
	return d.塔什库尔干Tashkurgan
}
    
var DPamir播密 PamirDuke = &播密PamirDuke{}

func init() {
	f := DPamir播密.(*播密PamirDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "pamir",
		TitleName: "播密",
		TitleCode: "d_pamir",
		Counties:  map[string]feud.County{},
	}

	f.播密Pamir = pamir.CPamir播密
	f.播密Pamir.SetParent(f)
	
	f.塔什库尔干Tashkurgan = tashkurgan.CTashkurgan塔什库尔干
	f.塔什库尔干Tashkurgan.SetParent(f)
	
}
