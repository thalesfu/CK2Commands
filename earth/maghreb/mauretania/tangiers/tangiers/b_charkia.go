package tangiers

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈拉克亚CharkiaBarony struct {
	feud.BaseBarony
}

var BCharkia哈拉克亚 feud.Barony = &哈拉克亚CharkiaBarony{}

func init() {
    f := BCharkia哈拉克亚.(*哈拉克亚CharkiaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "charkia",
		TitleName: "哈拉克亚",
		TitleCode: "b_charkia",
	}
}
