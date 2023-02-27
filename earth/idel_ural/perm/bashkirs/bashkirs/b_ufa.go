package bashkirs

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌法UfaBarony struct {
	feud.BaseBarony
}

var BUfa乌法 feud.Barony = &乌法UfaBarony{}

func init() {
    f := BUfa乌法.(*乌法UfaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ufa",
		TitleName: "乌法",
		TitleCode: "b_ufa",
	}
}
