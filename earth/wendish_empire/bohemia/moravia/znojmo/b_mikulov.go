package znojmo

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 米库洛夫MikulovBarony struct {
	feud.BaseBarony
}

var BMikulov米库洛夫 feud.Barony = &米库洛夫MikulovBarony{}

func init() {
    f := BMikulov米库洛夫.(*米库洛夫MikulovBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mikulov",
		TitleName: "米库洛夫",
		TitleCode: "b_mikulov",
	}
}
