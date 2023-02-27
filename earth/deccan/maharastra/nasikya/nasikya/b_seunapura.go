package nasikya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 犀浮那补罗SeunapuraBarony struct {
	feud.BaseBarony
}

var BSeunapura犀浮那补罗 feud.Barony = &犀浮那补罗SeunapuraBarony{}

func init() {
    f := BSeunapura犀浮那补罗.(*犀浮那补罗SeunapuraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "seunapura",
		TitleName: "犀浮那补罗",
		TitleCode: "b_seunapura",
	}
}
