package munio

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 扎斯基亚TsaskiyaBarony struct {
	feud.BaseBarony
}

var BTsaskiya扎斯基亚 feud.Barony = &扎斯基亚TsaskiyaBarony{}

func init() {
	f := BTsaskiya扎斯基亚.(*扎斯基亚TsaskiyaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tsaskiya",
		TitleName: "扎斯基亚",
		TitleCode: "b_tsaskiya",
	}
}
