package dunbar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伯威克BerwickBarony struct {
	feud.BaseBarony
}

var BBerwick伯威克 feud.Barony = &伯威克BerwickBarony{}

func init() {
    f := BBerwick伯威克.(*伯威克BerwickBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "berwick",
		TitleName: "伯威克",
		TitleCode: "b_berwick",
	}
}
