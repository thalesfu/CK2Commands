package dioclea

import (
	"github.com/thalesfu/CK2Commands/earth/byzantium/serbia/dioclea/ragusa"
	"github.com/thalesfu/CK2Commands/earth/byzantium/serbia/dioclea/travunia"
	"github.com/thalesfu/CK2Commands/earth/byzantium/serbia/dioclea/zeta"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type DiocleaDuke interface {
    feud.Duke
    CRagusa拉古萨() 	ragusa.RagusaCounty
    CTravunia特拉武尼亚() 	travunia.TravuniaCounty
    CZeta泽塔() 	zeta.ZetaCounty
}

type 迪奥克勒亚DiocleaDuke struct {
	feud.BaseDuke
	拉古萨Ragusa 	ragusa.RagusaCounty
	特拉武尼亚Travunia 	travunia.TravuniaCounty
	泽塔Zeta 	zeta.ZetaCounty
}

func (d *迪奥克勒亚DiocleaDuke) CRagusa拉古萨() ragusa.RagusaCounty {
	return d.拉古萨Ragusa
}
    
func (d *迪奥克勒亚DiocleaDuke) CTravunia特拉武尼亚() travunia.TravuniaCounty {
	return d.特拉武尼亚Travunia
}
    
func (d *迪奥克勒亚DiocleaDuke) CZeta泽塔() zeta.ZetaCounty {
	return d.泽塔Zeta
}
    
var DDioclea迪奥克勒亚 DiocleaDuke = &迪奥克勒亚DiocleaDuke{}

func init() {
	f := DDioclea迪奥克勒亚.(*迪奥克勒亚DiocleaDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "dioclea",
		TitleName: "迪奥克勒亚",
		TitleCode: "d_dioclea",
		Counties:  map[string]feud.County{},
	}

	f.拉古萨Ragusa = ragusa.CRagusa拉古萨
	f.拉古萨Ragusa.SetParent(f)
	
	f.特拉武尼亚Travunia = travunia.CTravunia特拉武尼亚
	f.特拉武尼亚Travunia.SetParent(f)
	
	f.泽塔Zeta = zeta.CZeta泽塔
	f.泽塔Zeta.SetParent(f)
	
}
