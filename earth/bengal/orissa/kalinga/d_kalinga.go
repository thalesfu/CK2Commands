package kalinga

import (
	"github.com/thalesfu/CK2Commands/earth/bengal/orissa/kalinga/kalinganagar"
	"github.com/thalesfu/CK2Commands/earth/bengal/orissa/kalinga/puri"
	"github.com/thalesfu/CK2Commands/earth/bengal/orissa/kalinga/vizagipatam"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KalingaDuke interface {
	feud.Duke
	CKalinganagar羯陵伽那揭罗() kalinganagar.KalinganagarCounty
	CPuri补梨() puri.PuriCounty
	CVizagipatam毗舍佉波吒南() vizagipatam.VizagipatamCounty
}

type 羯陵伽KalingaDuke struct {
	feud.BaseDuke
	羯陵伽那揭罗Kalinganagar kalinganagar.KalinganagarCounty
	补梨Puri             puri.PuriCounty
	毗舍佉波吒南Vizagipatam  vizagipatam.VizagipatamCounty
}

func (d *羯陵伽KalingaDuke) CKalinganagar羯陵伽那揭罗() kalinganagar.KalinganagarCounty {
	return d.羯陵伽那揭罗Kalinganagar
}

func (d *羯陵伽KalingaDuke) CPuri补梨() puri.PuriCounty {
	return d.补梨Puri
}

func (d *羯陵伽KalingaDuke) CVizagipatam毗舍佉波吒南() vizagipatam.VizagipatamCounty {
	return d.毗舍佉波吒南Vizagipatam
}

var DKalinga羯陵伽 KalingaDuke = &羯陵伽KalingaDuke{}

func init() {
	f := DKalinga羯陵伽.(*羯陵伽KalingaDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "kalinga",
		TitleName: "羯陵伽",
		TitleCode: "d_kalinga",
		Counties:  map[string]feud.County{},
	}

	f.羯陵伽那揭罗Kalinganagar = kalinganagar.CKalinganagar羯陵伽那揭罗
	f.羯陵伽那揭罗Kalinganagar.SetParent(f)

	f.补梨Puri = puri.CPuri补梨
	f.补梨Puri.SetParent(f)

	f.毗舍佉波吒南Vizagipatam = vizagipatam.CVizagipatam毗舍佉波吒南
	f.毗舍佉波吒南Vizagipatam.SetParent(f)

}
