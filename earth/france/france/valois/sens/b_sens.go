package sens

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 桑斯SensBarony struct {
	feud.BaseBarony
}

var BSens桑斯 feud.Barony = &桑斯SensBarony{}

func init() {
	f := BSens桑斯.(*桑斯SensBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sens",
		TitleName: "桑斯",
		TitleCode: "b_sens",
	}
}
