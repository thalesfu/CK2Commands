package nasikya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 那私迦NasikyaBarony struct {
	feud.BaseBarony
}

var BNasikya那私迦 feud.Barony = &那私迦NasikyaBarony{}

func init() {
    f := BNasikya那私迦.(*那私迦NasikyaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nasikya",
		TitleName: "那私迦",
		TitleCode: "b_nasikya",
	}
}
