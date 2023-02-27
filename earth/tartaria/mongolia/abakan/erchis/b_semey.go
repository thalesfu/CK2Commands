package erchis

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塞米伊SemeyBarony struct {
	feud.BaseBarony
}

var BSemey塞米伊 feud.Barony = &塞米伊SemeyBarony{}

func init() {
    f := BSemey塞米伊.(*塞米伊SemeyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "semey",
		TitleName: "塞米伊",
		TitleCode: "b_semey",
	}
}
