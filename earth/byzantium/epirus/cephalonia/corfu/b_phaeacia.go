package corfu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 费阿刻斯PhaeaciaBarony struct {
	feud.BaseBarony
}

var BPhaeacia费阿刻斯 feud.Barony = &费阿刻斯PhaeaciaBarony{}

func init() {
	f := BPhaeacia费阿刻斯.(*费阿刻斯PhaeaciaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "phaeacia",
		TitleName: "费阿刻斯",
		TitleCode: "b_phaeacia",
	}
}
