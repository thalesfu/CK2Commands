package jerusalem

import (
	"github.com/thalesfu/CK2Commands/earth/arabia/jerusalem/jerusalem/acre"
	"github.com/thalesfu/CK2Commands/earth/arabia/jerusalem/jerusalem/hebron"
	"github.com/thalesfu/CK2Commands/earth/arabia/jerusalem/jerusalem/jerusalem"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type JerusalemDuke interface {
    feud.Duke
    CAcre阿克() 	acre.AcreCounty
    CHebron希伯伦() 	hebron.HebronCounty
    CJerusalem耶路撒冷() 	jerusalem.JerusalemCounty
}

type 耶路撒冷JerusalemDuke struct {
	feud.BaseDuke
	阿克Acre 	acre.AcreCounty
	希伯伦Hebron 	hebron.HebronCounty
	耶路撒冷Jerusalem 	jerusalem.JerusalemCounty
}

func (d *耶路撒冷JerusalemDuke) CAcre阿克() acre.AcreCounty {
	return d.阿克Acre
}
    
func (d *耶路撒冷JerusalemDuke) CHebron希伯伦() hebron.HebronCounty {
	return d.希伯伦Hebron
}
    
func (d *耶路撒冷JerusalemDuke) CJerusalem耶路撒冷() jerusalem.JerusalemCounty {
	return d.耶路撒冷Jerusalem
}
    
var DJerusalem耶路撒冷 JerusalemDuke = &耶路撒冷JerusalemDuke{}

func init() {
	f := DJerusalem耶路撒冷.(*耶路撒冷JerusalemDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "jerusalem",
		TitleName: "耶路撒冷",
		TitleCode: "d_jerusalem",
		Counties:  map[string]feud.County{},
	}

	f.阿克Acre = acre.CAcre阿克
	f.阿克Acre.SetParent(f)
	
	f.希伯伦Hebron = hebron.CHebron希伯伦
	f.希伯伦Hebron.SetParent(f)
	
	f.耶路撒冷Jerusalem = jerusalem.CJerusalem耶路撒冷
	f.耶路撒冷Jerusalem.SetParent(f)
	
}
