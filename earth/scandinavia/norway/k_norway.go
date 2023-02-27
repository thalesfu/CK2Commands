package norway

import (
	"github.com/thalesfu/CK2Commands/earth/scandinavia/norway/agder"
	"github.com/thalesfu/CK2Commands/earth/scandinavia/norway/iceland"
	"github.com/thalesfu/CK2Commands/earth/scandinavia/norway/jamtland"
	"github.com/thalesfu/CK2Commands/earth/scandinavia/norway/oppland"
	"github.com/thalesfu/CK2Commands/earth/scandinavia/norway/orkney"
	"github.com/thalesfu/CK2Commands/earth/scandinavia/norway/ostlandet"
	"github.com/thalesfu/CK2Commands/earth/scandinavia/norway/trondelag"
	"github.com/thalesfu/CK2Commands/earth/scandinavia/norway/vestlandet"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type NorwayKingdom interface {
    feud.Kingdom
    DAgder阿格德尔() 	agder.AgderDuke
    DIceland冰岛() 	iceland.IcelandDuke
    DJamtland耶姆特兰() 	jamtland.JamtlandDuke
    DOppland奥普兰() 	oppland.OpplandDuke
    DOrkney奥克尼() 	orkney.OrkneyDuke
    DOstlandet维肯() 	ostlandet.OstlandetDuke
    DTrondelag尼达罗斯() 	trondelag.TrondelagDuke
    DVestlandet西挪威() 	vestlandet.VestlandetDuke
}

type 挪威NorwayKingdom struct {
	feud.BaseKingdom
	阿格德尔Agder 	agder.AgderDuke
	冰岛Iceland 	iceland.IcelandDuke
	耶姆特兰Jamtland 	jamtland.JamtlandDuke
	奥普兰Oppland 	oppland.OpplandDuke
	奥克尼Orkney 	orkney.OrkneyDuke
	维肯Ostlandet 	ostlandet.OstlandetDuke
	尼达罗斯Trondelag 	trondelag.TrondelagDuke
	西挪威Vestlandet 	vestlandet.VestlandetDuke
}

func (k *挪威NorwayKingdom) DAgder阿格德尔() agder.AgderDuke {
	return k.阿格德尔Agder
}
    
func (k *挪威NorwayKingdom) DIceland冰岛() iceland.IcelandDuke {
	return k.冰岛Iceland
}
    
func (k *挪威NorwayKingdom) DJamtland耶姆特兰() jamtland.JamtlandDuke {
	return k.耶姆特兰Jamtland
}
    
func (k *挪威NorwayKingdom) DOppland奥普兰() oppland.OpplandDuke {
	return k.奥普兰Oppland
}
    
func (k *挪威NorwayKingdom) DOrkney奥克尼() orkney.OrkneyDuke {
	return k.奥克尼Orkney
}
    
func (k *挪威NorwayKingdom) DOstlandet维肯() ostlandet.OstlandetDuke {
	return k.维肯Ostlandet
}
    
func (k *挪威NorwayKingdom) DTrondelag尼达罗斯() trondelag.TrondelagDuke {
	return k.尼达罗斯Trondelag
}
    
func (k *挪威NorwayKingdom) DVestlandet西挪威() vestlandet.VestlandetDuke {
	return k.西挪威Vestlandet
}
    
var KNorway挪威 NorwayKingdom = &挪威NorwayKingdom{}

func init() {
	f := KNorway挪威.(*挪威NorwayKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "norway",
		TitleName: "挪威",
		TitleCode: "k_norway",
		Dukes:  map[string]feud.Duke{},
	}

	f.阿格德尔Agder = agder.DAgder阿格德尔
	f.阿格德尔Agder.SetParent(f)
	
	f.冰岛Iceland = iceland.DIceland冰岛
	f.冰岛Iceland.SetParent(f)
	
	f.耶姆特兰Jamtland = jamtland.DJamtland耶姆特兰
	f.耶姆特兰Jamtland.SetParent(f)
	
	f.奥普兰Oppland = oppland.DOppland奥普兰
	f.奥普兰Oppland.SetParent(f)
	
	f.奥克尼Orkney = orkney.DOrkney奥克尼
	f.奥克尼Orkney.SetParent(f)
	
	f.维肯Ostlandet = ostlandet.DOstlandet维肯
	f.维肯Ostlandet.SetParent(f)
	
	f.尼达罗斯Trondelag = trondelag.DTrondelag尼达罗斯
	f.尼达罗斯Trondelag.SetParent(f)
	
	f.西挪威Vestlandet = vestlandet.DVestlandet西挪威
	f.西挪威Vestlandet.SetParent(f)
	
}
