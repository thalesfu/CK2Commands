package prussia

import (
	"github.com/thalesfu/CK2Commands/earth/wendish_empire/lithuania/prussia/chelminskie"
	"github.com/thalesfu/CK2Commands/earth/wendish_empire/lithuania/prussia/galindia"
	"github.com/thalesfu/CK2Commands/earth/wendish_empire/lithuania/prussia/marienburg"
	"github.com/thalesfu/CK2Commands/earth/wendish_empire/lithuania/prussia/sambia"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type PrussiaDuke interface {
    feud.Duke
    CChelminskie海乌姆诺() 	chelminskie.ChelminskieCounty
    CGalindia加林迪亚() 	galindia.GalindiaCounty
    CMarienburg马林堡() 	marienburg.MarienburgCounty
    CSambia桑比亚() 	sambia.SambiaCounty
}

type 普鲁士PrussiaDuke struct {
	feud.BaseDuke
	海乌姆诺Chelminskie 	chelminskie.ChelminskieCounty
	加林迪亚Galindia 	galindia.GalindiaCounty
	马林堡Marienburg 	marienburg.MarienburgCounty
	桑比亚Sambia 	sambia.SambiaCounty
}

func (d *普鲁士PrussiaDuke) CChelminskie海乌姆诺() chelminskie.ChelminskieCounty {
	return d.海乌姆诺Chelminskie
}
    
func (d *普鲁士PrussiaDuke) CGalindia加林迪亚() galindia.GalindiaCounty {
	return d.加林迪亚Galindia
}
    
func (d *普鲁士PrussiaDuke) CMarienburg马林堡() marienburg.MarienburgCounty {
	return d.马林堡Marienburg
}
    
func (d *普鲁士PrussiaDuke) CSambia桑比亚() sambia.SambiaCounty {
	return d.桑比亚Sambia
}
    
var DPrussia普鲁士 PrussiaDuke = &普鲁士PrussiaDuke{}

func init() {
	f := DPrussia普鲁士.(*普鲁士PrussiaDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "prussia",
		TitleName: "普鲁士",
		TitleCode: "d_prussia",
		Counties:  map[string]feud.County{},
	}

	f.海乌姆诺Chelminskie = chelminskie.CChelminskie海乌姆诺
	f.海乌姆诺Chelminskie.SetParent(f)
	
	f.加林迪亚Galindia = galindia.CGalindia加林迪亚
	f.加林迪亚Galindia.SetParent(f)
	
	f.马林堡Marienburg = marienburg.CMarienburg马林堡
	f.马林堡Marienburg.SetParent(f)
	
	f.桑比亚Sambia = sambia.CSambia桑比亚
	f.桑比亚Sambia.SetParent(f)
	
}
