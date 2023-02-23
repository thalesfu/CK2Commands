package plasencia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 普拉森西亚PlasenciaBarony struct {
	feud.BaseBarony
}

var BPlasencia普拉森西亚 feud.Barony = &普拉森西亚PlasenciaBarony{}

func init() {
	f := BPlasencia普拉森西亚.(*普拉森西亚PlasenciaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "plasencia",
		TitleName: "普拉森西亚",
		TitleCode: "b_plasencia",
	}
}
