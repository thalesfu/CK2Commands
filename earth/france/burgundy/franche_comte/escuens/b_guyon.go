package escuens

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 吉永GuyonBarony struct {
	feud.BaseBarony
}

var BGuyon吉永 feud.Barony = &吉永GuyonBarony{}

func init() {
    f := BGuyon吉永.(*吉永GuyonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "guyon",
		TitleName: "吉永",
		TitleCode: "b_guyon",
	}
}
