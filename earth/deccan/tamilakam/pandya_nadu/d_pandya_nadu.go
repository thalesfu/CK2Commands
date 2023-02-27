package pandya_nadu

import (
	"github.com/thalesfu/CK2Commands/earth/deccan/tamilakam/pandya_nadu/kongu"
	"github.com/thalesfu/CK2Commands/earth/deccan/tamilakam/pandya_nadu/madurai"
	"github.com/thalesfu/CK2Commands/earth/deccan/tamilakam/pandya_nadu/tenkasi"
	"github.com/thalesfu/CK2Commands/earth/deccan/tamilakam/pandya_nadu/tirunelveli"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type Pandya_naduDuke interface {
    feud.Duke
    CKongu贡古() 	kongu.KonguCounty
    CMadurai摩菟来() 	madurai.MaduraiCounty
    CTenkasi谛那迦尸() 	tenkasi.TenkasiCounty
    CTirunelveli底卢内勒吠梨() 	tirunelveli.TirunelveliCounty
}

type 般提耶拿柱Pandya_naduDuke struct {
	feud.BaseDuke
	贡古Kongu 	kongu.KonguCounty
	摩菟来Madurai 	madurai.MaduraiCounty
	谛那迦尸Tenkasi 	tenkasi.TenkasiCounty
	底卢内勒吠梨Tirunelveli 	tirunelveli.TirunelveliCounty
}

func (d *般提耶拿柱Pandya_naduDuke) CKongu贡古() kongu.KonguCounty {
	return d.贡古Kongu
}
    
func (d *般提耶拿柱Pandya_naduDuke) CMadurai摩菟来() madurai.MaduraiCounty {
	return d.摩菟来Madurai
}
    
func (d *般提耶拿柱Pandya_naduDuke) CTenkasi谛那迦尸() tenkasi.TenkasiCounty {
	return d.谛那迦尸Tenkasi
}
    
func (d *般提耶拿柱Pandya_naduDuke) CTirunelveli底卢内勒吠梨() tirunelveli.TirunelveliCounty {
	return d.底卢内勒吠梨Tirunelveli
}
    
var DPandya_nadu般提耶拿柱 Pandya_naduDuke = &般提耶拿柱Pandya_naduDuke{}

func init() {
	f := DPandya_nadu般提耶拿柱.(*般提耶拿柱Pandya_naduDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "pandya_nadu",
		TitleName: "般提耶拿柱",
		TitleCode: "d_pandya_nadu",
		Counties:  map[string]feud.County{},
	}

	f.贡古Kongu = kongu.CKongu贡古
	f.贡古Kongu.SetParent(f)
	
	f.摩菟来Madurai = madurai.CMadurai摩菟来
	f.摩菟来Madurai.SetParent(f)
	
	f.谛那迦尸Tenkasi = tenkasi.CTenkasi谛那迦尸
	f.谛那迦尸Tenkasi.SetParent(f)
	
	f.底卢内勒吠梨Tirunelveli = tirunelveli.CTirunelveli底卢内勒吠梨
	f.底卢内勒吠梨Tirunelveli.SetParent(f)
	
}
