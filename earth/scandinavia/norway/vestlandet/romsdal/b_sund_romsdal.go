package romsdal

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 松Sund_romsdalBarony struct {
	feud.BaseBarony
}

var BSund_romsdal松 feud.Barony = &松Sund_romsdalBarony{}

func init() {
    f := BSund_romsdal松.(*松Sund_romsdalBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sund_romsdal",
		TitleName: "松",
		TitleCode: "b_sund_romsdal",
	}
}
