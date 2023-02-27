package medina

import (
	"github.com/thalesfu/CK2Commands/earth/arabia/arabia/medina/asir"
	"github.com/thalesfu/CK2Commands/earth/arabia/arabia/medina/jeddah"
	"github.com/thalesfu/CK2Commands/earth/arabia/arabia/medina/khaybar"
	"github.com/thalesfu/CK2Commands/earth/arabia/arabia/medina/mecca"
	"github.com/thalesfu/CK2Commands/earth/arabia/arabia/medina/medina"
	"github.com/thalesfu/CK2Commands/earth/arabia/arabia/medina/tihama"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MedinaDuke interface {
    feud.Duke
    CAsir阿西尔() 	asir.AsirCounty
    CJeddah吉达() 	jeddah.JeddahCounty
    CKhaybar海拜尔() 	khaybar.KhaybarCounty
    CMecca麦加() 	mecca.MeccaCounty
    CMedina麦地那() 	medina.MedinaCounty
    CTihama帖哈麦() 	tihama.TihamaCounty
}

type 汉志MedinaDuke struct {
	feud.BaseDuke
	阿西尔Asir 	asir.AsirCounty
	吉达Jeddah 	jeddah.JeddahCounty
	海拜尔Khaybar 	khaybar.KhaybarCounty
	麦加Mecca 	mecca.MeccaCounty
	麦地那Medina 	medina.MedinaCounty
	帖哈麦Tihama 	tihama.TihamaCounty
}

func (d *汉志MedinaDuke) CAsir阿西尔() asir.AsirCounty {
	return d.阿西尔Asir
}
    
func (d *汉志MedinaDuke) CJeddah吉达() jeddah.JeddahCounty {
	return d.吉达Jeddah
}
    
func (d *汉志MedinaDuke) CKhaybar海拜尔() khaybar.KhaybarCounty {
	return d.海拜尔Khaybar
}
    
func (d *汉志MedinaDuke) CMecca麦加() mecca.MeccaCounty {
	return d.麦加Mecca
}
    
func (d *汉志MedinaDuke) CMedina麦地那() medina.MedinaCounty {
	return d.麦地那Medina
}
    
func (d *汉志MedinaDuke) CTihama帖哈麦() tihama.TihamaCounty {
	return d.帖哈麦Tihama
}
    
var DMedina汉志 MedinaDuke = &汉志MedinaDuke{}

func init() {
	f := DMedina汉志.(*汉志MedinaDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "medina",
		TitleName: "汉志",
		TitleCode: "d_medina",
		Counties:  map[string]feud.County{},
	}

	f.阿西尔Asir = asir.CAsir阿西尔
	f.阿西尔Asir.SetParent(f)
	
	f.吉达Jeddah = jeddah.CJeddah吉达
	f.吉达Jeddah.SetParent(f)
	
	f.海拜尔Khaybar = khaybar.CKhaybar海拜尔
	f.海拜尔Khaybar.SetParent(f)
	
	f.麦加Mecca = mecca.CMecca麦加
	f.麦加Mecca.SetParent(f)
	
	f.麦地那Medina = medina.CMedina麦地那
	f.麦地那Medina.SetParent(f)
	
	f.帖哈麦Tihama = tihama.CTihama帖哈麦
	f.帖哈麦Tihama.SetParent(f)
	
}
