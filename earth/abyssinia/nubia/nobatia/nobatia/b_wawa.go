package nobatia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瓦瓦WawaBarony struct {
	feud.BaseBarony
}

var BWawa瓦瓦 feud.Barony = &瓦瓦WawaBarony{}

func init() {
    f := BWawa瓦瓦.(*瓦瓦WawaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "wawa",
		TitleName: "瓦瓦",
		TitleCode: "b_wawa",
	}
}
