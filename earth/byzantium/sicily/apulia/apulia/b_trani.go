package apulia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 特拉尼TraniBarony struct {
	feud.BaseBarony
}

var BTrani特拉尼 feud.Barony = &特拉尼TraniBarony{}

func init() {
	f := BTrani特拉尼.(*特拉尼TraniBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "trani",
		TitleName: "特拉尼",
		TitleCode: "b_trani",
	}
}
