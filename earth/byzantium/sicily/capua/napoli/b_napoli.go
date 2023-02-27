package napoli

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 那波利NapoliBarony struct {
	feud.BaseBarony
}

var BNapoli那波利 feud.Barony = &那波利NapoliBarony{}

func init() {
    f := BNapoli那波利.(*那波利NapoliBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "napoli",
		TitleName: "那波利",
		TitleCode: "b_napoli",
	}
}
