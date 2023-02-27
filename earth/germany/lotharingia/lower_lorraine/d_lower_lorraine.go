package lower_lorraine

import (
	"github.com/thalesfu/CK2Commands/earth/germany/lotharingia/lower_lorraine/julich"
	"github.com/thalesfu/CK2Commands/earth/germany/lotharingia/lower_lorraine/liege"
	"github.com/thalesfu/CK2Commands/earth/germany/lotharingia/lower_lorraine/loon"
	"github.com/thalesfu/CK2Commands/earth/germany/lotharingia/lower_lorraine/luxembourg"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type Lower_lorraineDuke interface {
    feud.Duke
    CJulich于利希() 	julich.JulichCounty
    CLiege列日() 	liege.LiegeCounty
    CLoon洛翁() 	loon.LoonCounty
    CLuxembourg卢森堡() 	luxembourg.LuxembourgCounty
}

type 下洛林Lower_lorraineDuke struct {
	feud.BaseDuke
	于利希Julich 	julich.JulichCounty
	列日Liege 	liege.LiegeCounty
	洛翁Loon 	loon.LoonCounty
	卢森堡Luxembourg 	luxembourg.LuxembourgCounty
}

func (d *下洛林Lower_lorraineDuke) CJulich于利希() julich.JulichCounty {
	return d.于利希Julich
}
    
func (d *下洛林Lower_lorraineDuke) CLiege列日() liege.LiegeCounty {
	return d.列日Liege
}
    
func (d *下洛林Lower_lorraineDuke) CLoon洛翁() loon.LoonCounty {
	return d.洛翁Loon
}
    
func (d *下洛林Lower_lorraineDuke) CLuxembourg卢森堡() luxembourg.LuxembourgCounty {
	return d.卢森堡Luxembourg
}
    
var DLower_lorraine下洛林 Lower_lorraineDuke = &下洛林Lower_lorraineDuke{}

func init() {
	f := DLower_lorraine下洛林.(*下洛林Lower_lorraineDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "lower_lorraine",
		TitleName: "下洛林",
		TitleCode: "d_lower_lorraine",
		Counties:  map[string]feud.County{},
	}

	f.于利希Julich = julich.CJulich于利希
	f.于利希Julich.SetParent(f)
	
	f.列日Liege = liege.CLiege列日
	f.列日Liege.SetParent(f)
	
	f.洛翁Loon = loon.CLoon洛翁
	f.洛翁Loon.SetParent(f)
	
	f.卢森堡Luxembourg = luxembourg.CLuxembourg卢森堡
	f.卢森堡Luxembourg.SetParent(f)
	
}
