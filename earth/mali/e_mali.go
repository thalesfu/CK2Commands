package mali

import (
	"github.com/thalesfu/CK2Commands/earth/mali/ghana"
	"github.com/thalesfu/CK2Commands/earth/mali/mali"
	"github.com/thalesfu/CK2Commands/earth/mali/songhay"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MaliEmpire interface {
    feud.Empire
    KGhana加纳() 	ghana.GhanaKingdom
    KMali马里() 	mali.MaliKingdom
    KSonghay桑海() 	songhay.SonghayKingdom
}

type 马里MaliEmpire struct {
	feud.BaseEmpire
	加纳Ghana 	ghana.GhanaKingdom
	马里Mali 	mali.MaliKingdom
	桑海Songhay 	songhay.SonghayKingdom
}

func (e *马里MaliEmpire) KGhana加纳() ghana.GhanaKingdom {
	return e.加纳Ghana
}
    
func (e *马里MaliEmpire) KMali马里() mali.MaliKingdom {
	return e.马里Mali
}
    
func (e *马里MaliEmpire) KSonghay桑海() songhay.SonghayKingdom {
	return e.桑海Songhay
}
    
var EMali马里 MaliEmpire = &马里MaliEmpire{}

func init() {
	f := EMali马里.(*马里MaliEmpire)
	f.BaseEmpire = feud.BaseEmpire{
		Title:     "mali",
		TitleName: "马里",
		TitleCode: "e_mali",
		Kingdoms:  map[string]feud.Kingdom{},
	}
	f.加纳Ghana = ghana.KGhana加纳
	f.加纳Ghana.SetParent(f)
	f.马里Mali = mali.KMali马里
	f.马里Mali.SetParent(f)
	f.桑海Songhay = songhay.KSonghay桑海
	f.桑海Songhay.SetParent(f)
}
