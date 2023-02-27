package agder

import (
	"github.com/thalesfu/CK2Commands/earth/scandinavia/norway/agder/agder"
	"github.com/thalesfu/CK2Commands/earth/scandinavia/norway/agder/rogaland"
	"github.com/thalesfu/CK2Commands/earth/scandinavia/norway/agder/telemark"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AgderDuke interface {
    feud.Duke
    CAgder阿格德尔() 	agder.AgderCounty
    CRogaland罗加兰() 	rogaland.RogalandCounty
    CTelemark泰勒马克() 	telemark.TelemarkCounty
}

type 阿格德尔AgderDuke struct {
	feud.BaseDuke
	阿格德尔Agder 	agder.AgderCounty
	罗加兰Rogaland 	rogaland.RogalandCounty
	泰勒马克Telemark 	telemark.TelemarkCounty
}

func (d *阿格德尔AgderDuke) CAgder阿格德尔() agder.AgderCounty {
	return d.阿格德尔Agder
}
    
func (d *阿格德尔AgderDuke) CRogaland罗加兰() rogaland.RogalandCounty {
	return d.罗加兰Rogaland
}
    
func (d *阿格德尔AgderDuke) CTelemark泰勒马克() telemark.TelemarkCounty {
	return d.泰勒马克Telemark
}
    
var DAgder阿格德尔 AgderDuke = &阿格德尔AgderDuke{}

func init() {
	f := DAgder阿格德尔.(*阿格德尔AgderDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "agder",
		TitleName: "阿格德尔",
		TitleCode: "d_agder",
		Counties:  map[string]feud.County{},
	}

	f.阿格德尔Agder = agder.CAgder阿格德尔
	f.阿格德尔Agder.SetParent(f)
	
	f.罗加兰Rogaland = rogaland.CRogaland罗加兰
	f.罗加兰Rogaland.SetParent(f)
	
	f.泰勒马克Telemark = telemark.CTelemark泰勒马克
	f.泰勒马克Telemark.SetParent(f)
	
}
