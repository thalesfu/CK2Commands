package bulgaria

import (
	"github.com/thalesfu/CK2Commands/earth/byzantium/bulgaria/karvuna"
	"github.com/thalesfu/CK2Commands/earth/byzantium/bulgaria/strymon"
	"github.com/thalesfu/CK2Commands/earth/byzantium/bulgaria/turnovo"
	"github.com/thalesfu/CK2Commands/earth/byzantium/bulgaria/vidin"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BulgariaKingdom interface {
	feud.Kingdom
	DKarvuna卡尔武纳() karvuna.KarvunaDuke
	DStrymon斯特里蒙() strymon.StrymonDuke
	DTurnovo默西亚() turnovo.TurnovoDuke
	DVidin维丁() vidin.VidinDuke
}

type 保加利亚BulgariaKingdom struct {
	feud.BaseKingdom
	卡尔武纳Karvuna karvuna.KarvunaDuke
	斯特里蒙Strymon strymon.StrymonDuke
	默西亚Turnovo  turnovo.TurnovoDuke
	维丁Vidin     vidin.VidinDuke
}

func (k *保加利亚BulgariaKingdom) DKarvuna卡尔武纳() karvuna.KarvunaDuke {
	return k.卡尔武纳Karvuna
}

func (k *保加利亚BulgariaKingdom) DStrymon斯特里蒙() strymon.StrymonDuke {
	return k.斯特里蒙Strymon
}

func (k *保加利亚BulgariaKingdom) DTurnovo默西亚() turnovo.TurnovoDuke {
	return k.默西亚Turnovo
}

func (k *保加利亚BulgariaKingdom) DVidin维丁() vidin.VidinDuke {
	return k.维丁Vidin
}

var KBulgaria保加利亚 BulgariaKingdom = &保加利亚BulgariaKingdom{}

func init() {
	f := KBulgaria保加利亚.(*保加利亚BulgariaKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "bulgaria",
		TitleName: "保加利亚",
		TitleCode: "k_bulgaria",
		Dukes:     map[string]feud.Duke{},
	}

	f.卡尔武纳Karvuna = karvuna.DKarvuna卡尔武纳
	f.卡尔武纳Karvuna.SetParent(f)

	f.斯特里蒙Strymon = strymon.DStrymon斯特里蒙
	f.斯特里蒙Strymon.SetParent(f)

	f.默西亚Turnovo = turnovo.DTurnovo默西亚
	f.默西亚Turnovo.SetParent(f)

	f.维丁Vidin = vidin.DVidin维丁
	f.维丁Vidin.SetParent(f)

}
