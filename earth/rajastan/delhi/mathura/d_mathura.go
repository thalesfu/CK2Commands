package mathura

import (
	"github.com/thalesfu/CK2Commands/earth/rajastan/delhi/mathura/kol"
	"github.com/thalesfu/CK2Commands/earth/rajastan/delhi/mathura/mathura"
	"github.com/thalesfu/CK2Commands/earth/rajastan/delhi/mathura/sripatha"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MathuraDuke interface {
    feud.Duke
    CKol拘罗() 	kol.KolCounty
    CMathura秣菟罗() 	mathura.MathuraCounty
    CSripatha室利波他() 	sripatha.SripathaCounty
}

type 秣菟罗MathuraDuke struct {
	feud.BaseDuke
	拘罗Kol 	kol.KolCounty
	秣菟罗Mathura 	mathura.MathuraCounty
	室利波他Sripatha 	sripatha.SripathaCounty
}

func (d *秣菟罗MathuraDuke) CKol拘罗() kol.KolCounty {
	return d.拘罗Kol
}
    
func (d *秣菟罗MathuraDuke) CMathura秣菟罗() mathura.MathuraCounty {
	return d.秣菟罗Mathura
}
    
func (d *秣菟罗MathuraDuke) CSripatha室利波他() sripatha.SripathaCounty {
	return d.室利波他Sripatha
}
    
var DMathura秣菟罗 MathuraDuke = &秣菟罗MathuraDuke{}

func init() {
	f := DMathura秣菟罗.(*秣菟罗MathuraDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "mathura",
		TitleName: "秣菟罗",
		TitleCode: "d_mathura",
		Counties:  map[string]feud.County{},
	}

	f.拘罗Kol = kol.CKol拘罗
	f.拘罗Kol.SetParent(f)
	
	f.秣菟罗Mathura = mathura.CMathura秣菟罗
	f.秣菟罗Mathura.SetParent(f)
	
	f.室利波他Sripatha = sripatha.CSripatha室利波他
	f.室利波他Sripatha.SetParent(f)
	
}
