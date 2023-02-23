package ireland

import (
	"github.com/thalesfu/CK2Commands/earth/britannia/ireland/connacht"
	"github.com/thalesfu/CK2Commands/earth/britannia/ireland/leinster"
	"github.com/thalesfu/CK2Commands/earth/britannia/ireland/meath"
	"github.com/thalesfu/CK2Commands/earth/britannia/ireland/munster"
	"github.com/thalesfu/CK2Commands/earth/britannia/ireland/ulster"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type IrelandKingdom interface {
	feud.Kingdom
	DConnacht康诺特() connacht.ConnachtDuke
	DLeinster伦斯特() leinster.LeinsterDuke
	DMeath米斯() meath.MeathDuke
	DMunster芒斯特() munster.MunsterDuke
	DUlster阿尔斯特() ulster.UlsterDuke
}

type 爱尔兰IrelandKingdom struct {
	feud.BaseKingdom
	康诺特Connacht connacht.ConnachtDuke
	伦斯特Leinster leinster.LeinsterDuke
	米斯Meath     meath.MeathDuke
	芒斯特Munster  munster.MunsterDuke
	阿尔斯特Ulster  ulster.UlsterDuke
}

func (k *爱尔兰IrelandKingdom) DConnacht康诺特() connacht.ConnachtDuke {
	return k.康诺特Connacht
}

func (k *爱尔兰IrelandKingdom) DLeinster伦斯特() leinster.LeinsterDuke {
	return k.伦斯特Leinster
}

func (k *爱尔兰IrelandKingdom) DMeath米斯() meath.MeathDuke {
	return k.米斯Meath
}

func (k *爱尔兰IrelandKingdom) DMunster芒斯特() munster.MunsterDuke {
	return k.芒斯特Munster
}

func (k *爱尔兰IrelandKingdom) DUlster阿尔斯特() ulster.UlsterDuke {
	return k.阿尔斯特Ulster
}

var KIreland爱尔兰 IrelandKingdom = &爱尔兰IrelandKingdom{}

func init() {
	f := KIreland爱尔兰.(*爱尔兰IrelandKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "ireland",
		TitleName: "爱尔兰",
		TitleCode: "k_ireland",
		Dukes:     map[string]feud.Duke{},
	}

	f.康诺特Connacht = connacht.DConnacht康诺特
	f.康诺特Connacht.SetParent(f)

	f.伦斯特Leinster = leinster.DLeinster伦斯特
	f.伦斯特Leinster.SetParent(f)

	f.米斯Meath = meath.DMeath米斯
	f.米斯Meath.SetParent(f)

	f.芒斯特Munster = munster.DMunster芒斯特
	f.芒斯特Munster.SetParent(f)

	f.阿尔斯特Ulster = ulster.DUlster阿尔斯特
	f.阿尔斯特Ulster.SetParent(f)

}
