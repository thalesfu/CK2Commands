package orleans

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 默恩MeungBarony struct {
	feud.BaseBarony
}

var BMeung默恩 feud.Barony = &默恩MeungBarony{}

func init() {
    f := BMeung默恩.(*默恩MeungBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "meung",
		TitleName: "默恩",
		TitleCode: "b_meung",
	}
}
