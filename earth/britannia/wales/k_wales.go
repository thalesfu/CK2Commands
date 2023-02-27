package wales

import (
	"github.com/thalesfu/CK2Commands/earth/britannia/wales/cornwall"
	"github.com/thalesfu/CK2Commands/earth/britannia/wales/deheubarth"
	"github.com/thalesfu/CK2Commands/earth/britannia/wales/gwent"
	"github.com/thalesfu/CK2Commands/earth/britannia/wales/gwynedd"
	"github.com/thalesfu/CK2Commands/earth/britannia/wales/powys"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type WalesKingdom interface {
    feud.Kingdom
    DCornwall康沃尔() 	cornwall.CornwallDuke
    DDeheubarth德赫巴思() 	deheubarth.DeheubarthDuke
    DGwent格温特() 	gwent.GwentDuke
    DGwynedd圭内思() 	gwynedd.GwyneddDuke
    DPowys波伊斯() 	powys.PowysDuke
}

type 威尔士WalesKingdom struct {
	feud.BaseKingdom
	康沃尔Cornwall 	cornwall.CornwallDuke
	德赫巴思Deheubarth 	deheubarth.DeheubarthDuke
	格温特Gwent 	gwent.GwentDuke
	圭内思Gwynedd 	gwynedd.GwyneddDuke
	波伊斯Powys 	powys.PowysDuke
}

func (k *威尔士WalesKingdom) DCornwall康沃尔() cornwall.CornwallDuke {
	return k.康沃尔Cornwall
}
    
func (k *威尔士WalesKingdom) DDeheubarth德赫巴思() deheubarth.DeheubarthDuke {
	return k.德赫巴思Deheubarth
}
    
func (k *威尔士WalesKingdom) DGwent格温特() gwent.GwentDuke {
	return k.格温特Gwent
}
    
func (k *威尔士WalesKingdom) DGwynedd圭内思() gwynedd.GwyneddDuke {
	return k.圭内思Gwynedd
}
    
func (k *威尔士WalesKingdom) DPowys波伊斯() powys.PowysDuke {
	return k.波伊斯Powys
}
    
var KWales威尔士 WalesKingdom = &威尔士WalesKingdom{}

func init() {
	f := KWales威尔士.(*威尔士WalesKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "wales",
		TitleName: "威尔士",
		TitleCode: "k_wales",
		Dukes:  map[string]feud.Duke{},
	}

	f.康沃尔Cornwall = cornwall.DCornwall康沃尔
	f.康沃尔Cornwall.SetParent(f)
	
	f.德赫巴思Deheubarth = deheubarth.DDeheubarth德赫巴思
	f.德赫巴思Deheubarth.SetParent(f)
	
	f.格温特Gwent = gwent.DGwent格温特
	f.格温特Gwent.SetParent(f)
	
	f.圭内思Gwynedd = gwynedd.DGwynedd圭内思
	f.圭内思Gwynedd.SetParent(f)
	
	f.波伊斯Powys = powys.DPowys波伊斯
	f.波伊斯Powys.SetParent(f)
	
}
