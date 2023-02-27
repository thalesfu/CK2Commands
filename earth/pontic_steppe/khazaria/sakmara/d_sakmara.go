package sakmara

import (
	"github.com/thalesfu/CK2Commands/earth/pontic_steppe/khazaria/sakmara/irtek"
	"github.com/thalesfu/CK2Commands/earth/pontic_steppe/khazaria/sakmara/sakmara"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SakmaraDuke interface {
    feud.Duke
    CIrtek伊尔捷克() 	irtek.IrtekCounty
    CSakmara萨克马拉() 	sakmara.SakmaraCounty
}

type 萨克马拉SakmaraDuke struct {
	feud.BaseDuke
	伊尔捷克Irtek 	irtek.IrtekCounty
	萨克马拉Sakmara 	sakmara.SakmaraCounty
}

func (d *萨克马拉SakmaraDuke) CIrtek伊尔捷克() irtek.IrtekCounty {
	return d.伊尔捷克Irtek
}
    
func (d *萨克马拉SakmaraDuke) CSakmara萨克马拉() sakmara.SakmaraCounty {
	return d.萨克马拉Sakmara
}
    
var DSakmara萨克马拉 SakmaraDuke = &萨克马拉SakmaraDuke{}

func init() {
	f := DSakmara萨克马拉.(*萨克马拉SakmaraDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "sakmara",
		TitleName: "萨克马拉",
		TitleCode: "d_sakmara",
		Counties:  map[string]feud.County{},
	}

	f.伊尔捷克Irtek = irtek.CIrtek伊尔捷克
	f.伊尔捷克Irtek.SetParent(f)
	
	f.萨克马拉Sakmara = sakmara.CSakmara萨克马拉
	f.萨克马拉Sakmara.SetParent(f)
	
}
