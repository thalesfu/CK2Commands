package poland

import (
	"github.com/thalesfu/CK2Commands/earth/wendish_empire/poland/greater_poland"
	"github.com/thalesfu/CK2Commands/earth/wendish_empire/poland/kuyavia"
	"github.com/thalesfu/CK2Commands/earth/wendish_empire/poland/lesser_poland"
	"github.com/thalesfu/CK2Commands/earth/wendish_empire/poland/mazovia"
	"github.com/thalesfu/CK2Commands/earth/wendish_empire/poland/silesia"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type PolandKingdom interface {
    feud.Kingdom
    DGreater_poland大波兰() 	greater_poland.Greater_polandDuke
    DKuyavia库亚维亚() 	kuyavia.KuyaviaDuke
    DLesser_poland小波兰() 	lesser_poland.Lesser_polandDuke
    DMazovia马佐维亚() 	mazovia.MazoviaDuke
    DSilesia西里西亚() 	silesia.SilesiaDuke
}

type 波兰PolandKingdom struct {
	feud.BaseKingdom
	大波兰Greater_poland 	greater_poland.Greater_polandDuke
	库亚维亚Kuyavia 	kuyavia.KuyaviaDuke
	小波兰Lesser_poland 	lesser_poland.Lesser_polandDuke
	马佐维亚Mazovia 	mazovia.MazoviaDuke
	西里西亚Silesia 	silesia.SilesiaDuke
}

func (k *波兰PolandKingdom) DGreater_poland大波兰() greater_poland.Greater_polandDuke {
	return k.大波兰Greater_poland
}
    
func (k *波兰PolandKingdom) DKuyavia库亚维亚() kuyavia.KuyaviaDuke {
	return k.库亚维亚Kuyavia
}
    
func (k *波兰PolandKingdom) DLesser_poland小波兰() lesser_poland.Lesser_polandDuke {
	return k.小波兰Lesser_poland
}
    
func (k *波兰PolandKingdom) DMazovia马佐维亚() mazovia.MazoviaDuke {
	return k.马佐维亚Mazovia
}
    
func (k *波兰PolandKingdom) DSilesia西里西亚() silesia.SilesiaDuke {
	return k.西里西亚Silesia
}
    
var KPoland波兰 PolandKingdom = &波兰PolandKingdom{}

func init() {
	f := KPoland波兰.(*波兰PolandKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "poland",
		TitleName: "波兰",
		TitleCode: "k_poland",
		Dukes:  map[string]feud.Duke{},
	}

	f.大波兰Greater_poland = greater_poland.DGreater_poland大波兰
	f.大波兰Greater_poland.SetParent(f)
	
	f.库亚维亚Kuyavia = kuyavia.DKuyavia库亚维亚
	f.库亚维亚Kuyavia.SetParent(f)
	
	f.小波兰Lesser_poland = lesser_poland.DLesser_poland小波兰
	f.小波兰Lesser_poland.SetParent(f)
	
	f.马佐维亚Mazovia = mazovia.DMazovia马佐维亚
	f.马佐维亚Mazovia.SetParent(f)
	
	f.西里西亚Silesia = silesia.DSilesia西里西亚
	f.西里西亚Silesia.SetParent(f)
	
}
