package sweden

import (
	"github.com/thalesfu/CK2Commands/earth/scandinavia/sweden/bergslagen"
	"github.com/thalesfu/CK2Commands/earth/scandinavia/sweden/gotland"
	"github.com/thalesfu/CK2Commands/earth/scandinavia/sweden/norrland"
	"github.com/thalesfu/CK2Commands/earth/scandinavia/sweden/ostergotland"
	"github.com/thalesfu/CK2Commands/earth/scandinavia/sweden/smaland"
	"github.com/thalesfu/CK2Commands/earth/scandinavia/sweden/tioharad"
	"github.com/thalesfu/CK2Commands/earth/scandinavia/sweden/uppland"
	"github.com/thalesfu/CK2Commands/earth/scandinavia/sweden/vastergotland"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SwedenKingdom interface {
    feud.Kingdom
    DBergslagen贝里斯拉根() 	bergslagen.BergslagenDuke
    DGotland哥特兰() 	gotland.GotlandDuke
    DNorrland海尔辛兰() 	norrland.NorrlandDuke
    DOstergotland东约特兰() 	ostergotland.OstergotlandDuke
    DSmaland斯莫兰() 	smaland.SmalandDuke
    DTioharad韦伦德() 	tioharad.TioharadDuke
    DUppland乌普兰() 	uppland.UpplandDuke
    DVastergotland西约特兰() 	vastergotland.VastergotlandDuke
}

type 瑞典SwedenKingdom struct {
	feud.BaseKingdom
	贝里斯拉根Bergslagen 	bergslagen.BergslagenDuke
	哥特兰Gotland 	gotland.GotlandDuke
	海尔辛兰Norrland 	norrland.NorrlandDuke
	东约特兰Ostergotland 	ostergotland.OstergotlandDuke
	斯莫兰Smaland 	smaland.SmalandDuke
	韦伦德Tioharad 	tioharad.TioharadDuke
	乌普兰Uppland 	uppland.UpplandDuke
	西约特兰Vastergotland 	vastergotland.VastergotlandDuke
}

func (k *瑞典SwedenKingdom) DBergslagen贝里斯拉根() bergslagen.BergslagenDuke {
	return k.贝里斯拉根Bergslagen
}
    
func (k *瑞典SwedenKingdom) DGotland哥特兰() gotland.GotlandDuke {
	return k.哥特兰Gotland
}
    
func (k *瑞典SwedenKingdom) DNorrland海尔辛兰() norrland.NorrlandDuke {
	return k.海尔辛兰Norrland
}
    
func (k *瑞典SwedenKingdom) DOstergotland东约特兰() ostergotland.OstergotlandDuke {
	return k.东约特兰Ostergotland
}
    
func (k *瑞典SwedenKingdom) DSmaland斯莫兰() smaland.SmalandDuke {
	return k.斯莫兰Smaland
}
    
func (k *瑞典SwedenKingdom) DTioharad韦伦德() tioharad.TioharadDuke {
	return k.韦伦德Tioharad
}
    
func (k *瑞典SwedenKingdom) DUppland乌普兰() uppland.UpplandDuke {
	return k.乌普兰Uppland
}
    
func (k *瑞典SwedenKingdom) DVastergotland西约特兰() vastergotland.VastergotlandDuke {
	return k.西约特兰Vastergotland
}
    
var KSweden瑞典 SwedenKingdom = &瑞典SwedenKingdom{}

func init() {
	f := KSweden瑞典.(*瑞典SwedenKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "sweden",
		TitleName: "瑞典",
		TitleCode: "k_sweden",
		Dukes:  map[string]feud.Duke{},
	}

	f.贝里斯拉根Bergslagen = bergslagen.DBergslagen贝里斯拉根
	f.贝里斯拉根Bergslagen.SetParent(f)
	
	f.哥特兰Gotland = gotland.DGotland哥特兰
	f.哥特兰Gotland.SetParent(f)
	
	f.海尔辛兰Norrland = norrland.DNorrland海尔辛兰
	f.海尔辛兰Norrland.SetParent(f)
	
	f.东约特兰Ostergotland = ostergotland.DOstergotland东约特兰
	f.东约特兰Ostergotland.SetParent(f)
	
	f.斯莫兰Smaland = smaland.DSmaland斯莫兰
	f.斯莫兰Smaland.SetParent(f)
	
	f.韦伦德Tioharad = tioharad.DTioharad韦伦德
	f.韦伦德Tioharad.SetParent(f)
	
	f.乌普兰Uppland = uppland.DUppland乌普兰
	f.乌普兰Uppland.SetParent(f)
	
	f.西约特兰Vastergotland = vastergotland.DVastergotland西约特兰
	f.西约特兰Vastergotland.SetParent(f)
	
}
