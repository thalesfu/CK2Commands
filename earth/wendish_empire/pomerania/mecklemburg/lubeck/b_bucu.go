package lubeck

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布库BucuBarony struct {
	feud.BaseBarony
}

var BBucu布库 feud.Barony = &布库BucuBarony{}

func init() {
    f := BBucu布库.(*布库BucuBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bucu",
		TitleName: "布库",
		TitleCode: "b_bucu",
	}
}
