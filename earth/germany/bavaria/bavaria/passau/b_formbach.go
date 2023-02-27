package passau

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 福尔姆巴赫FormbachBarony struct {
	feud.BaseBarony
}

var BFormbach福尔姆巴赫 feud.Barony = &福尔姆巴赫FormbachBarony{}

func init() {
    f := BFormbach福尔姆巴赫.(*福尔姆巴赫FormbachBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "formbach",
		TitleName: "福尔姆巴赫",
		TitleCode: "b_formbach",
	}
}
