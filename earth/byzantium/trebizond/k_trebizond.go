package trebizond

import (
	"github.com/thalesfu/CK2Commands/earth/byzantium/trebizond/armeniacon"
	"github.com/thalesfu/CK2Commands/earth/byzantium/trebizond/paphlagonia"
	"github.com/thalesfu/CK2Commands/earth/byzantium/trebizond/trebizond"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TrebizondKingdom interface {
	feud.Kingdom
	DArmeniacon亚美尼亚坎() armeniacon.ArmeniaconDuke
	DPaphlagonia帕夫拉戈尼亚() paphlagonia.PaphlagoniaDuke
	DTrebizond特拉比松() trebizond.TrebizondDuke
}

type 特拉比松TrebizondKingdom struct {
	feud.BaseKingdom
	亚美尼亚坎Armeniacon   armeniacon.ArmeniaconDuke
	帕夫拉戈尼亚Paphlagonia paphlagonia.PaphlagoniaDuke
	特拉比松Trebizond     trebizond.TrebizondDuke
}

func (k *特拉比松TrebizondKingdom) DArmeniacon亚美尼亚坎() armeniacon.ArmeniaconDuke {
	return k.亚美尼亚坎Armeniacon
}

func (k *特拉比松TrebizondKingdom) DPaphlagonia帕夫拉戈尼亚() paphlagonia.PaphlagoniaDuke {
	return k.帕夫拉戈尼亚Paphlagonia
}

func (k *特拉比松TrebizondKingdom) DTrebizond特拉比松() trebizond.TrebizondDuke {
	return k.特拉比松Trebizond
}

var KTrebizond特拉比松 TrebizondKingdom = &特拉比松TrebizondKingdom{}

func init() {
	f := KTrebizond特拉比松.(*特拉比松TrebizondKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "trebizond",
		TitleName: "特拉比松",
		TitleCode: "k_trebizond",
		Dukes:     map[string]feud.Duke{},
	}

	f.亚美尼亚坎Armeniacon = armeniacon.DArmeniacon亚美尼亚坎
	f.亚美尼亚坎Armeniacon.SetParent(f)

	f.帕夫拉戈尼亚Paphlagonia = paphlagonia.DPaphlagonia帕夫拉戈尼亚
	f.帕夫拉戈尼亚Paphlagonia.SetParent(f)

	f.特拉比松Trebizond = trebizond.DTrebizond特拉比松
	f.特拉比松Trebizond.SetParent(f)

}
