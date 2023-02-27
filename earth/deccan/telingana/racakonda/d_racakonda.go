package racakonda

import (
	"github.com/thalesfu/CK2Commands/earth/deccan/telingana/racakonda/kollipake"
	"github.com/thalesfu/CK2Commands/earth/deccan/telingana/racakonda/medak"
	"github.com/thalesfu/CK2Commands/earth/deccan/telingana/racakonda/nilagiri"
	"github.com/thalesfu/CK2Commands/earth/deccan/telingana/racakonda/pannagallu"
	"github.com/thalesfu/CK2Commands/earth/deccan/telingana/racakonda/racakonda"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type RacakondaDuke interface {
    feud.Duke
    CKollipake拘梨波稽() 	kollipake.KollipakeCounty
    CMedak弥陀迦() 	medak.MedakCounty
    CNilagiri尼罗耆厘() 	nilagiri.NilagiriCounty
    CPannagallu般那伽楼() 	pannagallu.PannagalluCounty
    CRacakonda罗遮军荼() 	racakonda.RacakondaCounty
}

type 罗遮军荼RacakondaDuke struct {
	feud.BaseDuke
	拘梨波稽Kollipake 	kollipake.KollipakeCounty
	弥陀迦Medak 	medak.MedakCounty
	尼罗耆厘Nilagiri 	nilagiri.NilagiriCounty
	般那伽楼Pannagallu 	pannagallu.PannagalluCounty
	罗遮军荼Racakonda 	racakonda.RacakondaCounty
}

func (d *罗遮军荼RacakondaDuke) CKollipake拘梨波稽() kollipake.KollipakeCounty {
	return d.拘梨波稽Kollipake
}
    
func (d *罗遮军荼RacakondaDuke) CMedak弥陀迦() medak.MedakCounty {
	return d.弥陀迦Medak
}
    
func (d *罗遮军荼RacakondaDuke) CNilagiri尼罗耆厘() nilagiri.NilagiriCounty {
	return d.尼罗耆厘Nilagiri
}
    
func (d *罗遮军荼RacakondaDuke) CPannagallu般那伽楼() pannagallu.PannagalluCounty {
	return d.般那伽楼Pannagallu
}
    
func (d *罗遮军荼RacakondaDuke) CRacakonda罗遮军荼() racakonda.RacakondaCounty {
	return d.罗遮军荼Racakonda
}
    
var DRacakonda罗遮军荼 RacakondaDuke = &罗遮军荼RacakondaDuke{}

func init() {
	f := DRacakonda罗遮军荼.(*罗遮军荼RacakondaDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "racakonda",
		TitleName: "罗遮军荼",
		TitleCode: "d_racakonda",
		Counties:  map[string]feud.County{},
	}

	f.拘梨波稽Kollipake = kollipake.CKollipake拘梨波稽
	f.拘梨波稽Kollipake.SetParent(f)
	
	f.弥陀迦Medak = medak.CMedak弥陀迦
	f.弥陀迦Medak.SetParent(f)
	
	f.尼罗耆厘Nilagiri = nilagiri.CNilagiri尼罗耆厘
	f.尼罗耆厘Nilagiri.SetParent(f)
	
	f.般那伽楼Pannagallu = pannagallu.CPannagallu般那伽楼
	f.般那伽楼Pannagallu.SetParent(f)
	
	f.罗遮军荼Racakonda = racakonda.CRacakonda罗遮军荼
	f.罗遮军荼Racakonda.SetParent(f)
	
}
