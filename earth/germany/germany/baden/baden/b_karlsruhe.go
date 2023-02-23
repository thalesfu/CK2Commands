package baden

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡尔斯鲁厄KarlsruheBarony struct {
	feud.BaseBarony
}

var BKarlsruhe卡尔斯鲁厄 feud.Barony = &卡尔斯鲁厄KarlsruheBarony{}

func init() {
	f := BKarlsruhe卡尔斯鲁厄.(*卡尔斯鲁厄KarlsruheBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "karlsruhe",
		TitleName: "卡尔斯鲁厄",
		TitleCode: "b_karlsruhe",
	}
}
