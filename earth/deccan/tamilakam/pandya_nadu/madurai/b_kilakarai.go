package madurai

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 吉罗伽罗KilakaraiBarony struct {
	feud.BaseBarony
}

var BKilakarai吉罗伽罗 feud.Barony = &吉罗伽罗KilakaraiBarony{}

func init() {
    f := BKilakarai吉罗伽罗.(*吉罗伽罗KilakaraiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kilakarai",
		TitleName: "吉罗伽罗",
		TitleCode: "b_kilakarai",
	}
}
