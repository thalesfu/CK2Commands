package kamatapur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 那罗耶罗堡Nalrajar_garhBarony struct {
	feud.BaseBarony
}

var BNalrajar_garh那罗耶罗堡 feud.Barony = &那罗耶罗堡Nalrajar_garhBarony{}

func init() {
    f := BNalrajar_garh那罗耶罗堡.(*那罗耶罗堡Nalrajar_garhBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nalrajar_garh",
		TitleName: "那罗耶罗堡",
		TitleCode: "b_nalrajar_garh",
	}
}
