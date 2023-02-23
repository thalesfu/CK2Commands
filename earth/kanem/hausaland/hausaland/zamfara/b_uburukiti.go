package zamfara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌布鲁基蒂UburukitiBarony struct {
	feud.BaseBarony
}

var BUburukiti乌布鲁基蒂 feud.Barony = &乌布鲁基蒂UburukitiBarony{}

func init() {
	f := BUburukiti乌布鲁基蒂.(*乌布鲁基蒂UburukitiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "uburukiti",
		TitleName: "乌布鲁基蒂",
		TitleCode: "b_uburukiti",
	}
}
