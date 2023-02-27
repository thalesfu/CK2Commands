package brandenburg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 柏林BerlinBarony struct {
	feud.BaseBarony
}

var BBerlin柏林 feud.Barony = &柏林BerlinBarony{}

func init() {
    f := BBerlin柏林.(*柏林BerlinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "berlin",
		TitleName: "柏林",
		TitleCode: "b_berlin",
	}
}
