package tsaparang

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 芒嘎ManggaBarony struct {
	feud.BaseBarony
}

var BMangga芒嘎 feud.Barony = &芒嘎ManggaBarony{}

func init() {
    f := BMangga芒嘎.(*芒嘎ManggaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mangga",
		TitleName: "芒嘎",
		TitleCode: "b_mangga",
	}
}
