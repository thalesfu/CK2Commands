package lanka

import (
	"github.com/thalesfu/CK2Commands/earth/deccan/lanka/lanka/dakhina_desa"
	"github.com/thalesfu/CK2Commands/earth/deccan/lanka/lanka/rohana"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type LankaDuke interface {
    feud.Duke
    CDakhina_desa南域() 	dakhina_desa.Dakhina_desaCounty
    CRohana卢诃拿() 	rohana.RohanaCounty
}

type 楞迦LankaDuke struct {
	feud.BaseDuke
	南域Dakhina_desa 	dakhina_desa.Dakhina_desaCounty
	卢诃拿Rohana 	rohana.RohanaCounty
}

func (d *楞迦LankaDuke) CDakhina_desa南域() dakhina_desa.Dakhina_desaCounty {
	return d.南域Dakhina_desa
}
    
func (d *楞迦LankaDuke) CRohana卢诃拿() rohana.RohanaCounty {
	return d.卢诃拿Rohana
}
    
var DLanka楞迦 LankaDuke = &楞迦LankaDuke{}

func init() {
	f := DLanka楞迦.(*楞迦LankaDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "lanka",
		TitleName: "楞迦",
		TitleCode: "d_lanka",
		Counties:  map[string]feud.County{},
	}

	f.南域Dakhina_desa = dakhina_desa.CDakhina_desa南域
	f.南域Dakhina_desa.SetParent(f)
	
	f.卢诃拿Rohana = rohana.CRohana卢诃拿
	f.卢诃拿Rohana.SetParent(f)
	
}
