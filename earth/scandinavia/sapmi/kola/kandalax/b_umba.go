package kandalax

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 温巴UmbaBarony struct {
	feud.BaseBarony
}

var BUmba温巴 feud.Barony = &温巴UmbaBarony{}

func init() {
    f := BUmba温巴.(*温巴UmbaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "umba",
		TitleName: "温巴",
		TitleCode: "b_umba",
	}
}
