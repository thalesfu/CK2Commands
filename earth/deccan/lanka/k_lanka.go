package lanka

import (
	"github.com/thalesfu/CK2Commands/earth/deccan/lanka/lanka"
	"github.com/thalesfu/CK2Commands/earth/deccan/lanka/sinhala"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type LankaKingdom interface {
	feud.Kingdom
	DLanka楞迦() lanka.LankaDuke
	DSinhala僧伽罗() sinhala.SinhalaDuke
}

type 楞迦LankaKingdom struct {
	feud.BaseKingdom
	楞迦Lanka    lanka.LankaDuke
	僧伽罗Sinhala sinhala.SinhalaDuke
}

func (k *楞迦LankaKingdom) DLanka楞迦() lanka.LankaDuke {
	return k.楞迦Lanka
}

func (k *楞迦LankaKingdom) DSinhala僧伽罗() sinhala.SinhalaDuke {
	return k.僧伽罗Sinhala
}

var KLanka楞迦 LankaKingdom = &楞迦LankaKingdom{}

func init() {
	f := KLanka楞迦.(*楞迦LankaKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "lanka",
		TitleName: "楞迦",
		TitleCode: "k_lanka",
		Dukes:     map[string]feud.Duke{},
	}

	f.楞迦Lanka = lanka.DLanka楞迦
	f.楞迦Lanka.SetParent(f)

	f.僧伽罗Sinhala = sinhala.DSinhala僧伽罗
	f.僧伽罗Sinhala.SetParent(f)

}
