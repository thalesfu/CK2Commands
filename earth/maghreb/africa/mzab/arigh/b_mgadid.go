package arigh

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 姆加迪MgadidBarony struct {
	feud.BaseBarony
}

var BMgadid姆加迪 feud.Barony = &姆加迪MgadidBarony{}

func init() {
    f := BMgadid姆加迪.(*姆加迪MgadidBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mgadid",
		TitleName: "姆加迪",
		TitleCode: "b_mgadid",
	}
}
