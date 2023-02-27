package jersika

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马尔齐埃纳MarcienaBarony struct {
	feud.BaseBarony
}

var BMarciena马尔齐埃纳 feud.Barony = &马尔齐埃纳MarcienaBarony{}

func init() {
    f := BMarciena马尔齐埃纳.(*马尔齐埃纳MarcienaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "marciena",
		TitleName: "马尔齐埃纳",
		TitleCode: "b_marciena",
	}
}
