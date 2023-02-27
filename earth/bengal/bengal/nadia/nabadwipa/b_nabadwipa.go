package nabadwipa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 那婆提鞞波NabadwipaBarony struct {
	feud.BaseBarony
}

var BNabadwipa那婆提鞞波 feud.Barony = &那婆提鞞波NabadwipaBarony{}

func init() {
    f := BNabadwipa那婆提鞞波.(*那婆提鞞波NabadwipaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nabadwipa",
		TitleName: "那婆提鞞波",
		TitleCode: "b_nabadwipa",
	}
}
