package ragusa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 那罗那NaronaBarony struct {
	feud.BaseBarony
}

var BNarona那罗那 feud.Barony = &那罗那NaronaBarony{}

func init() {
    f := BNarona那罗那.(*那罗那NaronaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "narona",
		TitleName: "那罗那",
		TitleCode: "b_narona",
	}
}
