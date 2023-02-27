package sakala

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 查谟JammuBarony struct {
	feud.BaseBarony
}

var BJammu查谟 feud.Barony = &查谟JammuBarony{}

func init() {
    f := BJammu查谟.(*查谟JammuBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jammu",
		TitleName: "查谟",
		TitleCode: "b_jammu",
	}
}
