package medapata

import (
	"github.com/thalesfu/CK2Commands/earth/rajastan/rajputana/medapata/chitrakut"
	"github.com/thalesfu/CK2Commands/earth/rajastan/rajputana/medapata/kota"
	"github.com/thalesfu/CK2Commands/earth/rajastan/rajputana/medapata/medapata"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MedapataDuke interface {
    feud.Duke
    CChitrakut质多罗矩吒() 	chitrakut.ChitrakutCounty
    CKota拘吒() 	kota.KotaCounty
    CMedapata迷陀波吒() 	medapata.MedapataCounty
}

type 迷陀波吒MedapataDuke struct {
	feud.BaseDuke
	质多罗矩吒Chitrakut 	chitrakut.ChitrakutCounty
	拘吒Kota 	kota.KotaCounty
	迷陀波吒Medapata 	medapata.MedapataCounty
}

func (d *迷陀波吒MedapataDuke) CChitrakut质多罗矩吒() chitrakut.ChitrakutCounty {
	return d.质多罗矩吒Chitrakut
}
    
func (d *迷陀波吒MedapataDuke) CKota拘吒() kota.KotaCounty {
	return d.拘吒Kota
}
    
func (d *迷陀波吒MedapataDuke) CMedapata迷陀波吒() medapata.MedapataCounty {
	return d.迷陀波吒Medapata
}
    
var DMedapata迷陀波吒 MedapataDuke = &迷陀波吒MedapataDuke{}

func init() {
	f := DMedapata迷陀波吒.(*迷陀波吒MedapataDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "medapata",
		TitleName: "迷陀波吒",
		TitleCode: "d_medapata",
		Counties:  map[string]feud.County{},
	}

	f.质多罗矩吒Chitrakut = chitrakut.CChitrakut质多罗矩吒
	f.质多罗矩吒Chitrakut.SetParent(f)
	
	f.拘吒Kota = kota.CKota拘吒
	f.拘吒Kota.SetParent(f)
	
	f.迷陀波吒Medapata = medapata.CMedapata迷陀波吒
	f.迷陀波吒Medapata.SetParent(f)
	
}
