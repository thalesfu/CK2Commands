package fuqi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沙陀ShatuoBarony struct {
	feud.BaseBarony
}

var BShatuo沙陀 feud.Barony = &沙陀ShatuoBarony{}

func init() {
    f := BShatuo沙陀.(*沙陀ShatuoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shatuo",
		TitleName: "沙陀",
		TitleCode: "b_shatuo",
	}
}
