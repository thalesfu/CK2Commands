package beresty

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 别列斯捷BerestyBarony struct {
	feud.BaseBarony
}

var BBeresty别列斯捷 feud.Barony = &别列斯捷BerestyBarony{}

func init() {
    f := BBeresty别列斯捷.(*别列斯捷BerestyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "beresty",
		TitleName: "别列斯捷",
		TitleCode: "b_beresty",
	}
}
