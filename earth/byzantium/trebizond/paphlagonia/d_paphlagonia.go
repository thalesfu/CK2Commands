package paphlagonia

import (
	"github.com/thalesfu/CK2Commands/earth/byzantium/trebizond/paphlagonia/amastris"
	"github.com/thalesfu/CK2Commands/earth/byzantium/trebizond/paphlagonia/kastamon"
	"github.com/thalesfu/CK2Commands/earth/byzantium/trebizond/paphlagonia/paphlagonia"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type PaphlagoniaDuke interface {
	feud.Duke
	CAmastris阿马斯特里斯() amastris.AmastrisCounty
	CKastamon卡斯塔莫努() kastamon.KastamonCounty
	CPaphlagonia帕夫拉戈尼亚() paphlagonia.PaphlagoniaCounty
}

type 帕夫拉戈尼亚PaphlagoniaDuke struct {
	feud.BaseDuke
	阿马斯特里斯Amastris    amastris.AmastrisCounty
	卡斯塔莫努Kastamon     kastamon.KastamonCounty
	帕夫拉戈尼亚Paphlagonia paphlagonia.PaphlagoniaCounty
}

func (d *帕夫拉戈尼亚PaphlagoniaDuke) CAmastris阿马斯特里斯() amastris.AmastrisCounty {
	return d.阿马斯特里斯Amastris
}

func (d *帕夫拉戈尼亚PaphlagoniaDuke) CKastamon卡斯塔莫努() kastamon.KastamonCounty {
	return d.卡斯塔莫努Kastamon
}

func (d *帕夫拉戈尼亚PaphlagoniaDuke) CPaphlagonia帕夫拉戈尼亚() paphlagonia.PaphlagoniaCounty {
	return d.帕夫拉戈尼亚Paphlagonia
}

var DPaphlagonia帕夫拉戈尼亚 PaphlagoniaDuke = &帕夫拉戈尼亚PaphlagoniaDuke{}

func init() {
	f := DPaphlagonia帕夫拉戈尼亚.(*帕夫拉戈尼亚PaphlagoniaDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "paphlagonia",
		TitleName: "帕夫拉戈尼亚",
		TitleCode: "d_paphlagonia",
		Counties:  map[string]feud.County{},
	}

	f.阿马斯特里斯Amastris = amastris.CAmastris阿马斯特里斯
	f.阿马斯特里斯Amastris.SetParent(f)

	f.卡斯塔莫努Kastamon = kastamon.CKastamon卡斯塔莫努
	f.卡斯塔莫努Kastamon.SetParent(f)

	f.帕夫拉戈尼亚Paphlagonia = paphlagonia.CPaphlagonia帕夫拉戈尼亚
	f.帕夫拉戈尼亚Paphlagonia.SetParent(f)

}
