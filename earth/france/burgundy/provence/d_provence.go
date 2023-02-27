package provence

import (
	"github.com/thalesfu/CK2Commands/earth/france/burgundy/provence/forcalquier"
	"github.com/thalesfu/CK2Commands/earth/france/burgundy/provence/nice"
	"github.com/thalesfu/CK2Commands/earth/france/burgundy/provence/provence"
	"github.com/thalesfu/CK2Commands/earth/france/burgundy/provence/venaissin"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ProvenceDuke interface {
    feud.Duke
    CForcalquier福卡尔基耶() 	forcalquier.ForcalquierCounty
    CNice尼斯() 	nice.NiceCounty
    CProvence普罗旺斯() 	provence.ProvenceCounty
    CVenaissin韦奈桑() 	venaissin.VenaissinCounty
}

type 普罗旺斯ProvenceDuke struct {
	feud.BaseDuke
	福卡尔基耶Forcalquier 	forcalquier.ForcalquierCounty
	尼斯Nice 	nice.NiceCounty
	普罗旺斯Provence 	provence.ProvenceCounty
	韦奈桑Venaissin 	venaissin.VenaissinCounty
}

func (d *普罗旺斯ProvenceDuke) CForcalquier福卡尔基耶() forcalquier.ForcalquierCounty {
	return d.福卡尔基耶Forcalquier
}
    
func (d *普罗旺斯ProvenceDuke) CNice尼斯() nice.NiceCounty {
	return d.尼斯Nice
}
    
func (d *普罗旺斯ProvenceDuke) CProvence普罗旺斯() provence.ProvenceCounty {
	return d.普罗旺斯Provence
}
    
func (d *普罗旺斯ProvenceDuke) CVenaissin韦奈桑() venaissin.VenaissinCounty {
	return d.韦奈桑Venaissin
}
    
var DProvence普罗旺斯 ProvenceDuke = &普罗旺斯ProvenceDuke{}

func init() {
	f := DProvence普罗旺斯.(*普罗旺斯ProvenceDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "provence",
		TitleName: "普罗旺斯",
		TitleCode: "d_provence",
		Counties:  map[string]feud.County{},
	}

	f.福卡尔基耶Forcalquier = forcalquier.CForcalquier福卡尔基耶
	f.福卡尔基耶Forcalquier.SetParent(f)
	
	f.尼斯Nice = nice.CNice尼斯
	f.尼斯Nice.SetParent(f)
	
	f.普罗旺斯Provence = provence.CProvence普罗旺斯
	f.普罗旺斯Provence.SetParent(f)
	
	f.韦奈桑Venaissin = venaissin.CVenaissin韦奈桑
	f.韦奈桑Venaissin.SetParent(f)
	
}
