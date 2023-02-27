package stravani

import (
	"github.com/thalesfu/CK2Commands/earth/rajastan/rajputana/stravani/ludrava"
	"github.com/thalesfu/CK2Commands/earth/rajastan/rajputana/stravani/satyapura"
	"github.com/thalesfu/CK2Commands/earth/rajastan/rajputana/stravani/vijnot"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type StravaniDuke interface {
    feud.Duke
    CLudrava律陀罗婆() 	ludrava.LudravaCounty
    CSatyapura娑底也补罗() 	satyapura.SatyapuraCounty
    CVijnot毗者努多() 	vijnot.VijnotCounty
}

type 萨怛罗婆尼StravaniDuke struct {
	feud.BaseDuke
	律陀罗婆Ludrava 	ludrava.LudravaCounty
	娑底也补罗Satyapura 	satyapura.SatyapuraCounty
	毗者努多Vijnot 	vijnot.VijnotCounty
}

func (d *萨怛罗婆尼StravaniDuke) CLudrava律陀罗婆() ludrava.LudravaCounty {
	return d.律陀罗婆Ludrava
}
    
func (d *萨怛罗婆尼StravaniDuke) CSatyapura娑底也补罗() satyapura.SatyapuraCounty {
	return d.娑底也补罗Satyapura
}
    
func (d *萨怛罗婆尼StravaniDuke) CVijnot毗者努多() vijnot.VijnotCounty {
	return d.毗者努多Vijnot
}
    
var DStravani萨怛罗婆尼 StravaniDuke = &萨怛罗婆尼StravaniDuke{}

func init() {
	f := DStravani萨怛罗婆尼.(*萨怛罗婆尼StravaniDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "stravani",
		TitleName: "萨怛罗婆尼",
		TitleCode: "d_stravani",
		Counties:  map[string]feud.County{},
	}

	f.律陀罗婆Ludrava = ludrava.CLudrava律陀罗婆
	f.律陀罗婆Ludrava.SetParent(f)
	
	f.娑底也补罗Satyapura = satyapura.CSatyapura娑底也补罗
	f.娑底也补罗Satyapura.SetParent(f)
	
	f.毗者努多Vijnot = vijnot.CVijnot毗者努多
	f.毗者努多Vijnot.SetParent(f)
	
}
