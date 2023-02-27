package egypt

import (
	"github.com/thalesfu/CK2Commands/earth/arabia/egypt/alexandria"
	"github.com/thalesfu/CK2Commands/earth/arabia/egypt/aswan"
	"github.com/thalesfu/CK2Commands/earth/arabia/egypt/asyut"
	"github.com/thalesfu/CK2Commands/earth/arabia/egypt/cairo"
	"github.com/thalesfu/CK2Commands/earth/arabia/egypt/damietta"
	"github.com/thalesfu/CK2Commands/earth/arabia/egypt/faiyum"
	"github.com/thalesfu/CK2Commands/earth/arabia/egypt/paraetonium"
	"github.com/thalesfu/CK2Commands/earth/arabia/egypt/sinai"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type EgyptKingdom interface {
    feud.Kingdom
    DAlexandria亚历山大() 	alexandria.AlexandriaDuke
    DAswan阿斯旺() 	aswan.AswanDuke
    DAsyut艾斯尤特() 	asyut.AsyutDuke
    DCairo开罗() 	cairo.CairoDuke
    DDamietta达米埃塔() 	damietta.DamiettaDuke
    DFaiyum法尤姆() 	faiyum.FaiyumDuke
    DParaetonium帕莱托尼翁() 	paraetonium.ParaetoniumDuke
    DSinai西奈() 	sinai.SinaiDuke
}

type 埃及EgyptKingdom struct {
	feud.BaseKingdom
	亚历山大Alexandria 	alexandria.AlexandriaDuke
	阿斯旺Aswan 	aswan.AswanDuke
	艾斯尤特Asyut 	asyut.AsyutDuke
	开罗Cairo 	cairo.CairoDuke
	达米埃塔Damietta 	damietta.DamiettaDuke
	法尤姆Faiyum 	faiyum.FaiyumDuke
	帕莱托尼翁Paraetonium 	paraetonium.ParaetoniumDuke
	西奈Sinai 	sinai.SinaiDuke
}

func (k *埃及EgyptKingdom) DAlexandria亚历山大() alexandria.AlexandriaDuke {
	return k.亚历山大Alexandria
}
    
func (k *埃及EgyptKingdom) DAswan阿斯旺() aswan.AswanDuke {
	return k.阿斯旺Aswan
}
    
func (k *埃及EgyptKingdom) DAsyut艾斯尤特() asyut.AsyutDuke {
	return k.艾斯尤特Asyut
}
    
func (k *埃及EgyptKingdom) DCairo开罗() cairo.CairoDuke {
	return k.开罗Cairo
}
    
func (k *埃及EgyptKingdom) DDamietta达米埃塔() damietta.DamiettaDuke {
	return k.达米埃塔Damietta
}
    
func (k *埃及EgyptKingdom) DFaiyum法尤姆() faiyum.FaiyumDuke {
	return k.法尤姆Faiyum
}
    
func (k *埃及EgyptKingdom) DParaetonium帕莱托尼翁() paraetonium.ParaetoniumDuke {
	return k.帕莱托尼翁Paraetonium
}
    
func (k *埃及EgyptKingdom) DSinai西奈() sinai.SinaiDuke {
	return k.西奈Sinai
}
    
var KEgypt埃及 EgyptKingdom = &埃及EgyptKingdom{}

func init() {
	f := KEgypt埃及.(*埃及EgyptKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "egypt",
		TitleName: "埃及",
		TitleCode: "k_egypt",
		Dukes:  map[string]feud.Duke{},
	}

	f.亚历山大Alexandria = alexandria.DAlexandria亚历山大
	f.亚历山大Alexandria.SetParent(f)
	
	f.阿斯旺Aswan = aswan.DAswan阿斯旺
	f.阿斯旺Aswan.SetParent(f)
	
	f.艾斯尤特Asyut = asyut.DAsyut艾斯尤特
	f.艾斯尤特Asyut.SetParent(f)
	
	f.开罗Cairo = cairo.DCairo开罗
	f.开罗Cairo.SetParent(f)
	
	f.达米埃塔Damietta = damietta.DDamietta达米埃塔
	f.达米埃塔Damietta.SetParent(f)
	
	f.法尤姆Faiyum = faiyum.DFaiyum法尤姆
	f.法尤姆Faiyum.SetParent(f)
	
	f.帕莱托尼翁Paraetonium = paraetonium.DParaetonium帕莱托尼翁
	f.帕莱托尼翁Paraetonium.SetParent(f)
	
	f.西奈Sinai = sinai.DSinai西奈
	f.西奈Sinai.SetParent(f)
	
}
