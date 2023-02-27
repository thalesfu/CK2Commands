package north_dvina

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 尼科洛_科列尔斯基NikolokorelskiBarony struct {
	feud.BaseBarony
}

var BNikolokorelski尼科洛_科列尔斯基 feud.Barony = &尼科洛_科列尔斯基NikolokorelskiBarony{}

func init() {
    f := BNikolokorelski尼科洛_科列尔斯基.(*尼科洛_科列尔斯基NikolokorelskiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nikolokorelski",
		TitleName: "尼科洛_科列尔斯基",
		TitleCode: "b_nikolokorelski",
	}
}
