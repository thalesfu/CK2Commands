package vidharba

import (
	"github.com/thalesfu/CK2Commands/earth/deccan/maharastra/vidharba/acalapura"
	"github.com/thalesfu/CK2Commands/earth/deccan/maharastra/vidharba/canda"
	"github.com/thalesfu/CK2Commands/earth/deccan/maharastra/vidharba/kherla"
	"github.com/thalesfu/CK2Commands/earth/deccan/maharastra/vidharba/ramagiri"
	"github.com/thalesfu/CK2Commands/earth/deccan/maharastra/vidharba/vairagara"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type VidharbaDuke interface {
    feud.Duke
    CAcalapura不动城() 	acalapura.AcalapuraCounty
    CCanda旃陀() 	canda.CandaCounty
    CKherla稽剌罗() 	kherla.KherlaCounty
    CRamagiri罗摩耆厘() 	ramagiri.RamagiriCounty
    CVairagara吠罗伽罗() 	vairagara.VairagaraCounty
}

type 毗达婆VidharbaDuke struct {
	feud.BaseDuke
	不动城Acalapura 	acalapura.AcalapuraCounty
	旃陀Canda 	canda.CandaCounty
	稽剌罗Kherla 	kherla.KherlaCounty
	罗摩耆厘Ramagiri 	ramagiri.RamagiriCounty
	吠罗伽罗Vairagara 	vairagara.VairagaraCounty
}

func (d *毗达婆VidharbaDuke) CAcalapura不动城() acalapura.AcalapuraCounty {
	return d.不动城Acalapura
}
    
func (d *毗达婆VidharbaDuke) CCanda旃陀() canda.CandaCounty {
	return d.旃陀Canda
}
    
func (d *毗达婆VidharbaDuke) CKherla稽剌罗() kherla.KherlaCounty {
	return d.稽剌罗Kherla
}
    
func (d *毗达婆VidharbaDuke) CRamagiri罗摩耆厘() ramagiri.RamagiriCounty {
	return d.罗摩耆厘Ramagiri
}
    
func (d *毗达婆VidharbaDuke) CVairagara吠罗伽罗() vairagara.VairagaraCounty {
	return d.吠罗伽罗Vairagara
}
    
var DVidharba毗达婆 VidharbaDuke = &毗达婆VidharbaDuke{}

func init() {
	f := DVidharba毗达婆.(*毗达婆VidharbaDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "vidharba",
		TitleName: "毗达婆",
		TitleCode: "d_vidharba",
		Counties:  map[string]feud.County{},
	}

	f.不动城Acalapura = acalapura.CAcalapura不动城
	f.不动城Acalapura.SetParent(f)
	
	f.旃陀Canda = canda.CCanda旃陀
	f.旃陀Canda.SetParent(f)
	
	f.稽剌罗Kherla = kherla.CKherla稽剌罗
	f.稽剌罗Kherla.SetParent(f)
	
	f.罗摩耆厘Ramagiri = ramagiri.CRamagiri罗摩耆厘
	f.罗摩耆厘Ramagiri.SetParent(f)
	
	f.吠罗伽罗Vairagara = vairagara.CVairagara吠罗伽罗
	f.吠罗伽罗Vairagara.SetParent(f)
	
}
