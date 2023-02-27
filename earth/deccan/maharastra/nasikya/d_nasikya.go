package nasikya

import (
	"github.com/thalesfu/CK2Commands/earth/deccan/maharastra/nasikya/kondana"
	"github.com/thalesfu/CK2Commands/earth/deccan/maharastra/nasikya/nandurbar"
	"github.com/thalesfu/CK2Commands/earth/deccan/maharastra/nasikya/nasikya"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type NasikyaDuke interface {
    feud.Duke
    CKondana军荼那() 	kondana.KondanaCounty
    CNandurbar难豆罗婆罗() 	nandurbar.NandurbarCounty
    CNasikya那私迦() 	nasikya.NasikyaCounty
}

type 那私迦NasikyaDuke struct {
	feud.BaseDuke
	军荼那Kondana 	kondana.KondanaCounty
	难豆罗婆罗Nandurbar 	nandurbar.NandurbarCounty
	那私迦Nasikya 	nasikya.NasikyaCounty
}

func (d *那私迦NasikyaDuke) CKondana军荼那() kondana.KondanaCounty {
	return d.军荼那Kondana
}
    
func (d *那私迦NasikyaDuke) CNandurbar难豆罗婆罗() nandurbar.NandurbarCounty {
	return d.难豆罗婆罗Nandurbar
}
    
func (d *那私迦NasikyaDuke) CNasikya那私迦() nasikya.NasikyaCounty {
	return d.那私迦Nasikya
}
    
var DNasikya那私迦 NasikyaDuke = &那私迦NasikyaDuke{}

func init() {
	f := DNasikya那私迦.(*那私迦NasikyaDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "nasikya",
		TitleName: "那私迦",
		TitleCode: "d_nasikya",
		Counties:  map[string]feud.County{},
	}

	f.军荼那Kondana = kondana.CKondana军荼那
	f.军荼那Kondana.SetParent(f)
	
	f.难豆罗婆罗Nandurbar = nandurbar.CNandurbar难豆罗婆罗
	f.难豆罗婆罗Nandurbar.SetParent(f)
	
	f.那私迦Nasikya = nasikya.CNasikya那私迦
	f.那私迦Nasikya.SetParent(f)
	
}
