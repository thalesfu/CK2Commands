package zarma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 祖鲁ZuruBarony struct {
	feud.BaseBarony
}

var BZuru祖鲁 feud.Barony = &祖鲁ZuruBarony{}

func init() {
	f := BZuru祖鲁.(*祖鲁ZuruBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zuru",
		TitleName: "祖鲁",
		TitleCode: "b_zuru",
	}
}
