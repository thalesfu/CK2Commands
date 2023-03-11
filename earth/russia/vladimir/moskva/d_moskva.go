package moskva

import (
	"github.com/thalesfu/CK2Commands/earth/russia/vladimir/moskva/moskva"
	"github.com/thalesfu/CK2Commands/earth/russia/vladimir/moskva/naro_fominsk"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MoskvaDuke interface {
    feud.Duke
    CMoskva莫斯科() 	moskva.MoskvaCounty
    CNaro_fominsk纳罗福明斯克() 	naro_fominsk.Naro_fominskCounty
}

type 莫斯科MoskvaDuke struct {
	feud.BaseDuke
	莫斯科Moskva 	moskva.MoskvaCounty
	纳罗福明斯克Naro_fominsk 	naro_fominsk.Naro_fominskCounty
}

func (d *莫斯科MoskvaDuke) CMoskva莫斯科() moskva.MoskvaCounty {
	return d.莫斯科Moskva
}
    
func (d *莫斯科MoskvaDuke) CNaro_fominsk纳罗福明斯克() naro_fominsk.Naro_fominskCounty {
	return d.纳罗福明斯克Naro_fominsk
}
    
var DMoskva莫斯科 MoskvaDuke = &莫斯科MoskvaDuke{}

func init() {
	f := DMoskva莫斯科.(*莫斯科MoskvaDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "moskva",
		TitleName: "莫斯科",
		TitleCode: "d_moskva",
		Counties:  map[string]feud.County{},
	}

	f.莫斯科Moskva = moskva.CMoskva莫斯科
	f.莫斯科Moskva.SetParent(f)
	
	f.纳罗福明斯克Naro_fominsk = naro_fominsk.CNaro_fominsk纳罗福明斯克
	f.纳罗福明斯克Naro_fominsk.SetParent(f)
	
}
