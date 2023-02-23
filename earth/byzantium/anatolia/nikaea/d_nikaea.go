package nikaea

import (
	"github.com/thalesfu/CK2Commands/earth/byzantium/anatolia/nikaea/dorylaion"
	"github.com/thalesfu/CK2Commands/earth/byzantium/anatolia/nikaea/nikaea"
	"github.com/thalesfu/CK2Commands/earth/byzantium/anatolia/nikaea/prusa"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type NikaeaDuke interface {
	feud.Duke
	CDorylaion多律莱翁() dorylaion.DorylaionCounty
	CNikaea尼西亚() nikaea.NikaeaCounty
	CPrusa普鲁萨() prusa.PrusaCounty
}

type 奥普西金NikaeaDuke struct {
	feud.BaseDuke
	多律莱翁Dorylaion dorylaion.DorylaionCounty
	尼西亚Nikaea     nikaea.NikaeaCounty
	普鲁萨Prusa      prusa.PrusaCounty
}

func (d *奥普西金NikaeaDuke) CDorylaion多律莱翁() dorylaion.DorylaionCounty {
	return d.多律莱翁Dorylaion
}

func (d *奥普西金NikaeaDuke) CNikaea尼西亚() nikaea.NikaeaCounty {
	return d.尼西亚Nikaea
}

func (d *奥普西金NikaeaDuke) CPrusa普鲁萨() prusa.PrusaCounty {
	return d.普鲁萨Prusa
}

var DNikaea奥普西金 NikaeaDuke = &奥普西金NikaeaDuke{}

func init() {
	f := DNikaea奥普西金.(*奥普西金NikaeaDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "nikaea",
		TitleName: "奥普西金",
		TitleCode: "d_nikaea",
		Counties:  map[string]feud.County{},
	}

	f.多律莱翁Dorylaion = dorylaion.CDorylaion多律莱翁
	f.多律莱翁Dorylaion.SetParent(f)

	f.尼西亚Nikaea = nikaea.CNikaea尼西亚
	f.尼西亚Nikaea.SetParent(f)

	f.普鲁萨Prusa = prusa.CPrusa普鲁萨
	f.普鲁萨Prusa.SetParent(f)

}
