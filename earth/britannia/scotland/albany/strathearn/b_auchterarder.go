package strathearn

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥赫特拉德AuchterarderBarony struct {
	feud.BaseBarony
}

var BAuchterarder奥赫特拉德 feud.Barony = &奥赫特拉德AuchterarderBarony{}

func init() {
	f := BAuchterarder奥赫特拉德.(*奥赫特拉德AuchterarderBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "auchterarder",
		TitleName: "奥赫特拉德",
		TitleCode: "b_auchterarder",
	}
}
