package sicily

import (
	"github.com/thalesfu/CK2Commands/earth/byzantium/sicily/abruzzo"
	"github.com/thalesfu/CK2Commands/earth/byzantium/sicily/apulia"
	"github.com/thalesfu/CK2Commands/earth/byzantium/sicily/benevento"
	"github.com/thalesfu/CK2Commands/earth/byzantium/sicily/calabria"
	"github.com/thalesfu/CK2Commands/earth/byzantium/sicily/capua"
	"github.com/thalesfu/CK2Commands/earth/byzantium/sicily/salerno"
	"github.com/thalesfu/CK2Commands/earth/byzantium/sicily/sicily"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SicilyKingdom interface {
	feud.Kingdom
	DAbruzzo阿布鲁佐() abruzzo.AbruzzoDuke
	DApulia阿普利亚() apulia.ApuliaDuke
	DBenevento贝内文托() benevento.BeneventoDuke
	DCalabria卡拉布里亚() calabria.CalabriaDuke
	DCapua卡普阿() capua.CapuaDuke
	DSalerno萨莱诺() salerno.SalernoDuke
	DSicily西西里() sicily.SicilyDuke
}

type 西西里SicilyKingdom struct {
	feud.BaseKingdom
	阿布鲁佐Abruzzo   abruzzo.AbruzzoDuke
	阿普利亚Apulia    apulia.ApuliaDuke
	贝内文托Benevento benevento.BeneventoDuke
	卡拉布里亚Calabria calabria.CalabriaDuke
	卡普阿Capua      capua.CapuaDuke
	萨莱诺Salerno    salerno.SalernoDuke
	西西里Sicily     sicily.SicilyDuke
}

func (k *西西里SicilyKingdom) DAbruzzo阿布鲁佐() abruzzo.AbruzzoDuke {
	return k.阿布鲁佐Abruzzo
}

func (k *西西里SicilyKingdom) DApulia阿普利亚() apulia.ApuliaDuke {
	return k.阿普利亚Apulia
}

func (k *西西里SicilyKingdom) DBenevento贝内文托() benevento.BeneventoDuke {
	return k.贝内文托Benevento
}

func (k *西西里SicilyKingdom) DCalabria卡拉布里亚() calabria.CalabriaDuke {
	return k.卡拉布里亚Calabria
}

func (k *西西里SicilyKingdom) DCapua卡普阿() capua.CapuaDuke {
	return k.卡普阿Capua
}

func (k *西西里SicilyKingdom) DSalerno萨莱诺() salerno.SalernoDuke {
	return k.萨莱诺Salerno
}

func (k *西西里SicilyKingdom) DSicily西西里() sicily.SicilyDuke {
	return k.西西里Sicily
}

var KSicily西西里 SicilyKingdom = &西西里SicilyKingdom{}

func init() {
	f := KSicily西西里.(*西西里SicilyKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "sicily",
		TitleName: "西西里",
		TitleCode: "k_sicily",
		Dukes:     map[string]feud.Duke{},
	}

	f.阿布鲁佐Abruzzo = abruzzo.DAbruzzo阿布鲁佐
	f.阿布鲁佐Abruzzo.SetParent(f)

	f.阿普利亚Apulia = apulia.DApulia阿普利亚
	f.阿普利亚Apulia.SetParent(f)

	f.贝内文托Benevento = benevento.DBenevento贝内文托
	f.贝内文托Benevento.SetParent(f)

	f.卡拉布里亚Calabria = calabria.DCalabria卡拉布里亚
	f.卡拉布里亚Calabria.SetParent(f)

	f.卡普阿Capua = capua.DCapua卡普阿
	f.卡普阿Capua.SetParent(f)

	f.萨莱诺Salerno = salerno.DSalerno萨莱诺
	f.萨莱诺Salerno.SetParent(f)

	f.西西里Sicily = sicily.DSicily西西里
	f.西西里Sicily.SetParent(f)

}
