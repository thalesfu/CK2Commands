package ziz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布德尼卜BoudnibBarony struct {
	feud.BaseBarony
}

var BBoudnib布德尼卜 feud.Barony = &布德尼卜BoudnibBarony{}

func init() {
	f := BBoudnib布德尼卜.(*布德尼卜BoudnibBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "boudnib",
		TitleName: "布德尼卜",
		TitleCode: "b_boudnib",
	}
}
