package hesse

import (
	"github.com/thalesfu/CK2Commands/earth/germany/germany/hesse/frankfurt"
	"github.com/thalesfu/CK2Commands/earth/germany/germany/hesse/hesse"
	"github.com/thalesfu/CK2Commands/earth/germany/germany/hesse/nassau"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type HesseDuke interface {
    feud.Duke
    CFrankfurt法兰克福() 	frankfurt.FrankfurtCounty
    CHesse马堡() 	hesse.HesseCounty
    CNassau弗里茨拉尔() 	nassau.NassauCounty
}

type 黑森HesseDuke struct {
	feud.BaseDuke
	法兰克福Frankfurt 	frankfurt.FrankfurtCounty
	马堡Hesse 	hesse.HesseCounty
	弗里茨拉尔Nassau 	nassau.NassauCounty
}

func (d *黑森HesseDuke) CFrankfurt法兰克福() frankfurt.FrankfurtCounty {
	return d.法兰克福Frankfurt
}
    
func (d *黑森HesseDuke) CHesse马堡() hesse.HesseCounty {
	return d.马堡Hesse
}
    
func (d *黑森HesseDuke) CNassau弗里茨拉尔() nassau.NassauCounty {
	return d.弗里茨拉尔Nassau
}
    
var DHesse黑森 HesseDuke = &黑森HesseDuke{}

func init() {
	f := DHesse黑森.(*黑森HesseDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "hesse",
		TitleName: "黑森",
		TitleCode: "d_hesse",
		Counties:  map[string]feud.County{},
	}

	f.法兰克福Frankfurt = frankfurt.CFrankfurt法兰克福
	f.法兰克福Frankfurt.SetParent(f)
	
	f.马堡Hesse = hesse.CHesse马堡
	f.马堡Hesse.SetParent(f)
	
	f.弗里茨拉尔Nassau = nassau.CNassau弗里茨拉尔
	f.弗里茨拉尔Nassau.SetParent(f)
	
}
