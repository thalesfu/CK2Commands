package rohana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 迦阇罗伽摩KajaragamaBarony struct {
	feud.BaseBarony
}

var BKajaragama迦阇罗伽摩 feud.Barony = &迦阇罗伽摩KajaragamaBarony{}

func init() {
    f := BKajaragama迦阇罗伽摩.(*迦阇罗伽摩KajaragamaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kajaragama",
		TitleName: "迦阇罗伽摩",
		TitleCode: "b_kajaragama",
	}
}
