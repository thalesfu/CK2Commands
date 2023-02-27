package yatviags

import (
	"github.com/thalesfu/CK2Commands/earth/wendish_empire/lithuania/yatviags/jacwiez"
	"github.com/thalesfu/CK2Commands/earth/wendish_empire/lithuania/yatviags/sudovia"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type YatviagsDuke interface {
    feud.Duke
    CJacwiez格罗德诺() 	jacwiez.JacwiezCounty
    CSudovia苏多维亚() 	sudovia.SudoviaCounty
}

type 约特维恩格YatviagsDuke struct {
	feud.BaseDuke
	格罗德诺Jacwiez 	jacwiez.JacwiezCounty
	苏多维亚Sudovia 	sudovia.SudoviaCounty
}

func (d *约特维恩格YatviagsDuke) CJacwiez格罗德诺() jacwiez.JacwiezCounty {
	return d.格罗德诺Jacwiez
}
    
func (d *约特维恩格YatviagsDuke) CSudovia苏多维亚() sudovia.SudoviaCounty {
	return d.苏多维亚Sudovia
}
    
var DYatviags约特维恩格 YatviagsDuke = &约特维恩格YatviagsDuke{}

func init() {
	f := DYatviags约特维恩格.(*约特维恩格YatviagsDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "yatviags",
		TitleName: "约特维恩格",
		TitleCode: "d_yatviags",
		Counties:  map[string]feud.County{},
	}

	f.格罗德诺Jacwiez = jacwiez.CJacwiez格罗德诺
	f.格罗德诺Jacwiez.SetParent(f)
	
	f.苏多维亚Sudovia = sudovia.CSudovia苏多维亚
	f.苏多维亚Sudovia.SetParent(f)
	
}
