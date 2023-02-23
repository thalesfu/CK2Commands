package badghis

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马林MalinBarony struct {
	feud.BaseBarony
}

var BMalin马林 feud.Barony = &马林MalinBarony{}

func init() {
	f := BMalin马林.(*马林MalinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "malin",
		TitleName: "马林",
		TitleCode: "b_malin",
	}
}
