package delhi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 那梨罗NarelaBarony struct {
	feud.BaseBarony
}

var BNarela那梨罗 feud.Barony = &那梨罗NarelaBarony{}

func init() {
    f := BNarela那梨罗.(*那梨罗NarelaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "narela",
		TitleName: "那梨罗",
		TitleCode: "b_narela",
	}
}
