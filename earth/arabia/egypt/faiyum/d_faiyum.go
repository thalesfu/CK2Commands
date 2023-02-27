package faiyum

import (
	"github.com/thalesfu/CK2Commands/earth/arabia/egypt/faiyum/buhairya"
	"github.com/thalesfu/CK2Commands/earth/arabia/egypt/faiyum/faiyum"
	"github.com/thalesfu/CK2Commands/earth/arabia/egypt/faiyum/farafra"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type FaiyumDuke interface {
    feud.Duke
    CBuhairyac_buhairya() 	buhairya.BuhairyaCounty
    CFaiyum法尤姆() 	faiyum.FaiyumCounty
    CFarafrac_farafra() 	farafra.FarafraCounty
}

type 法尤姆FaiyumDuke struct {
	feud.BaseDuke
	c_buhairyaBuhairya 	buhairya.BuhairyaCounty
	法尤姆Faiyum 	faiyum.FaiyumCounty
	c_farafraFarafra 	farafra.FarafraCounty
}

func (d *法尤姆FaiyumDuke) CBuhairyac_buhairya() buhairya.BuhairyaCounty {
	return d.c_buhairyaBuhairya
}
    
func (d *法尤姆FaiyumDuke) CFaiyum法尤姆() faiyum.FaiyumCounty {
	return d.法尤姆Faiyum
}
    
func (d *法尤姆FaiyumDuke) CFarafrac_farafra() farafra.FarafraCounty {
	return d.c_farafraFarafra
}
    
var DFaiyum法尤姆 FaiyumDuke = &法尤姆FaiyumDuke{}

func init() {
	f := DFaiyum法尤姆.(*法尤姆FaiyumDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "faiyum",
		TitleName: "法尤姆",
		TitleCode: "d_faiyum",
		Counties:  map[string]feud.County{},
	}

	f.c_buhairyaBuhairya = buhairya.CBuhairyac_buhairya
	f.c_buhairyaBuhairya.SetParent(f)
	
	f.法尤姆Faiyum = faiyum.CFaiyum法尤姆
	f.法尤姆Faiyum.SetParent(f)
	
	f.c_farafraFarafra = farafra.CFarafrac_farafra
	f.c_farafraFarafra.SetParent(f)
	
}
