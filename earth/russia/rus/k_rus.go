package rus

import (
	"github.com/thalesfu/CK2Commands/earth/russia/rus/beloozero"
	"github.com/thalesfu/CK2Commands/earth/russia/rus/novgorod"
	"github.com/thalesfu/CK2Commands/earth/russia/rus/pskov"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type RusKingdom interface {
    feud.Kingdom
    DBeloozero韦普斯() 	beloozero.BeloozeroDuke
    DNovgorod诺夫哥罗德() 	novgorod.NovgorodDuke
    DPskov普斯科夫() 	pskov.PskovDuke
}

type 诺夫哥罗德RusKingdom struct {
	feud.BaseKingdom
	韦普斯Beloozero 	beloozero.BeloozeroDuke
	诺夫哥罗德Novgorod 	novgorod.NovgorodDuke
	普斯科夫Pskov 	pskov.PskovDuke
}

func (k *诺夫哥罗德RusKingdom) DBeloozero韦普斯() beloozero.BeloozeroDuke {
	return k.韦普斯Beloozero
}
    
func (k *诺夫哥罗德RusKingdom) DNovgorod诺夫哥罗德() novgorod.NovgorodDuke {
	return k.诺夫哥罗德Novgorod
}
    
func (k *诺夫哥罗德RusKingdom) DPskov普斯科夫() pskov.PskovDuke {
	return k.普斯科夫Pskov
}
    
var KRus诺夫哥罗德 RusKingdom = &诺夫哥罗德RusKingdom{}

func init() {
	f := KRus诺夫哥罗德.(*诺夫哥罗德RusKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "rus",
		TitleName: "诺夫哥罗德",
		TitleCode: "k_rus",
		Dukes:  map[string]feud.Duke{},
	}

	f.韦普斯Beloozero = beloozero.DBeloozero韦普斯
	f.韦普斯Beloozero.SetParent(f)
	
	f.诺夫哥罗德Novgorod = novgorod.DNovgorod诺夫哥罗德
	f.诺夫哥罗德Novgorod.SetParent(f)
	
	f.普斯科夫Pskov = pskov.DPskov普斯科夫
	f.普斯科夫Pskov.SetParent(f)
	
}
