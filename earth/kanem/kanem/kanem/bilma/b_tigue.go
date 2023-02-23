package bilma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蒂盖TigueBarony struct {
	feud.BaseBarony
}

var BTigue蒂盖 feud.Barony = &蒂盖TigueBarony{}

func init() {
	f := BTigue蒂盖.(*蒂盖TigueBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tigue",
		TitleName: "蒂盖",
		TitleCode: "b_tigue",
	}
}
