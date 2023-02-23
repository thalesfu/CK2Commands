package novogrodek

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 利达LidaBarony struct {
	feud.BaseBarony
}

var BLida利达 feud.Barony = &利达LidaBarony{}

func init() {
	f := BLida利达.(*利达LidaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lida",
		TitleName: "利达",
		TitleCode: "b_lida",
	}
}
