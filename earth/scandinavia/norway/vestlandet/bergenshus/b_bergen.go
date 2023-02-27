package bergenshus

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卑尔根BergenBarony struct {
	feud.BaseBarony
}

var BBergen卑尔根 feud.Barony = &卑尔根BergenBarony{}

func init() {
    f := BBergen卑尔根.(*卑尔根BergenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bergen",
		TitleName: "卑尔根",
		TitleCode: "b_bergen",
	}
}
