package rogaland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 内尔伯NaerboBarony struct {
	feud.BaseBarony
}

var BNaerbo内尔伯 feud.Barony = &内尔伯NaerboBarony{}

func init() {
    f := BNaerbo内尔伯.(*内尔伯NaerboBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "naerbo",
		TitleName: "内尔伯",
		TitleCode: "b_naerbo",
	}
}
