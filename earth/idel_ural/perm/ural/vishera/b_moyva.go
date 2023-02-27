package vishera

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莫伊瓦MoyvaBarony struct {
	feud.BaseBarony
}

var BMoyva莫伊瓦 feud.Barony = &莫伊瓦MoyvaBarony{}

func init() {
    f := BMoyva莫伊瓦.(*莫伊瓦MoyvaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "moyva",
		TitleName: "莫伊瓦",
		TitleCode: "b_moyva",
	}
}
