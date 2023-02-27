package overijssel

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 兹沃勒ZwolleBarony struct {
	feud.BaseBarony
}

var BZwolle兹沃勒 feud.Barony = &兹沃勒ZwolleBarony{}

func init() {
    f := BZwolle兹沃勒.(*兹沃勒ZwolleBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zwolle",
		TitleName: "兹沃勒",
		TitleCode: "b_zwolle",
	}
}
