package az_zarqa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 鲁赛法RusseifaBarony struct {
	feud.BaseBarony
}

var BRusseifa鲁赛法 feud.Barony = &鲁赛法RusseifaBarony{}

func init() {
    f := BRusseifa鲁赛法.(*鲁赛法RusseifaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "russeifa",
		TitleName: "鲁赛法",
		TitleCode: "b_russeifa",
	}
}
