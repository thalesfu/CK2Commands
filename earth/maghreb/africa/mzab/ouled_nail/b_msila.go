package ouled_nail

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 姆西拉MsilaBarony struct {
	feud.BaseBarony
}

var BMsila姆西拉 feud.Barony = &姆西拉MsilaBarony{}

func init() {
    f := BMsila姆西拉.(*姆西拉MsilaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "msila",
		TitleName: "姆西拉",
		TitleCode: "b_msila",
	}
}
