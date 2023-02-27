package tamilakam

import (
	"github.com/thalesfu/CK2Commands/earth/deccan/tamilakam/chera_nadu"
	"github.com/thalesfu/CK2Commands/earth/deccan/tamilakam/chola_nadu"
	"github.com/thalesfu/CK2Commands/earth/deccan/tamilakam/pandya_nadu"
	"github.com/thalesfu/CK2Commands/earth/deccan/tamilakam/tondai_nadu"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TamilakamKingdom interface {
    feud.Kingdom
    DChera_nadu支罗拿都() 	chera_nadu.Chera_naduDuke
    DChola_nadu朱罗拿柱() 	chola_nadu.Chola_naduDuke
    DPandya_nadu般提耶拿柱() 	pandya_nadu.Pandya_naduDuke
    DTondai_nadu旦榸拿柱() 	tondai_nadu.Tondai_naduDuke
}

type 昙密罗迦摩TamilakamKingdom struct {
	feud.BaseKingdom
	支罗拿都Chera_nadu 	chera_nadu.Chera_naduDuke
	朱罗拿柱Chola_nadu 	chola_nadu.Chola_naduDuke
	般提耶拿柱Pandya_nadu 	pandya_nadu.Pandya_naduDuke
	旦榸拿柱Tondai_nadu 	tondai_nadu.Tondai_naduDuke
}

func (k *昙密罗迦摩TamilakamKingdom) DChera_nadu支罗拿都() chera_nadu.Chera_naduDuke {
	return k.支罗拿都Chera_nadu
}
    
func (k *昙密罗迦摩TamilakamKingdom) DChola_nadu朱罗拿柱() chola_nadu.Chola_naduDuke {
	return k.朱罗拿柱Chola_nadu
}
    
func (k *昙密罗迦摩TamilakamKingdom) DPandya_nadu般提耶拿柱() pandya_nadu.Pandya_naduDuke {
	return k.般提耶拿柱Pandya_nadu
}
    
func (k *昙密罗迦摩TamilakamKingdom) DTondai_nadu旦榸拿柱() tondai_nadu.Tondai_naduDuke {
	return k.旦榸拿柱Tondai_nadu
}
    
var KTamilakam昙密罗迦摩 TamilakamKingdom = &昙密罗迦摩TamilakamKingdom{}

func init() {
	f := KTamilakam昙密罗迦摩.(*昙密罗迦摩TamilakamKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "tamilakam",
		TitleName: "昙密罗迦摩",
		TitleCode: "k_tamilakam",
		Dukes:  map[string]feud.Duke{},
	}

	f.支罗拿都Chera_nadu = chera_nadu.DChera_nadu支罗拿都
	f.支罗拿都Chera_nadu.SetParent(f)
	
	f.朱罗拿柱Chola_nadu = chola_nadu.DChola_nadu朱罗拿柱
	f.朱罗拿柱Chola_nadu.SetParent(f)
	
	f.般提耶拿柱Pandya_nadu = pandya_nadu.DPandya_nadu般提耶拿柱
	f.般提耶拿柱Pandya_nadu.SetParent(f)
	
	f.旦榸拿柱Tondai_nadu = tondai_nadu.DTondai_nadu旦榸拿柱
	f.旦榸拿柱Tondai_nadu.SetParent(f)
	
}
