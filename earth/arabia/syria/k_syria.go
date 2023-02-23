package syria

import (
	"github.com/thalesfu/CK2Commands/earth/arabia/syria/aleppo"
	"github.com/thalesfu/CK2Commands/earth/arabia/syria/antioch"
	"github.com/thalesfu/CK2Commands/earth/arabia/syria/damascus"
	"github.com/thalesfu/CK2Commands/earth/arabia/syria/palmyra"
	"github.com/thalesfu/CK2Commands/earth/arabia/syria/tripoli"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SyriaKingdom interface {
	feud.Kingdom
	DAleppo阿勒颇() aleppo.AleppoDuke
	DAntioch安条克() antioch.AntiochDuke
	DDamascus大马士革() damascus.DamascusDuke
	DPalmyra巴尔米拉() palmyra.PalmyraDuke
	DTripoli的黎波里() tripoli.TripoliDuke
}

type 叙利亚SyriaKingdom struct {
	feud.BaseKingdom
	阿勒颇Aleppo    aleppo.AleppoDuke
	安条克Antioch   antioch.AntiochDuke
	大马士革Damascus damascus.DamascusDuke
	巴尔米拉Palmyra  palmyra.PalmyraDuke
	的黎波里Tripoli  tripoli.TripoliDuke
}

func (k *叙利亚SyriaKingdom) DAleppo阿勒颇() aleppo.AleppoDuke {
	return k.阿勒颇Aleppo
}

func (k *叙利亚SyriaKingdom) DAntioch安条克() antioch.AntiochDuke {
	return k.安条克Antioch
}

func (k *叙利亚SyriaKingdom) DDamascus大马士革() damascus.DamascusDuke {
	return k.大马士革Damascus
}

func (k *叙利亚SyriaKingdom) DPalmyra巴尔米拉() palmyra.PalmyraDuke {
	return k.巴尔米拉Palmyra
}

func (k *叙利亚SyriaKingdom) DTripoli的黎波里() tripoli.TripoliDuke {
	return k.的黎波里Tripoli
}

var KSyria叙利亚 SyriaKingdom = &叙利亚SyriaKingdom{}

func init() {
	f := KSyria叙利亚.(*叙利亚SyriaKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "syria",
		TitleName: "叙利亚",
		TitleCode: "k_syria",
		Dukes:     map[string]feud.Duke{},
	}

	f.阿勒颇Aleppo = aleppo.DAleppo阿勒颇
	f.阿勒颇Aleppo.SetParent(f)

	f.安条克Antioch = antioch.DAntioch安条克
	f.安条克Antioch.SetParent(f)

	f.大马士革Damascus = damascus.DDamascus大马士革
	f.大马士革Damascus.SetParent(f)

	f.巴尔米拉Palmyra = palmyra.DPalmyra巴尔米拉
	f.巴尔米拉Palmyra.SetParent(f)

	f.的黎波里Tripoli = tripoli.DTripoli的黎波里
	f.的黎波里Tripoli.SetParent(f)

}
