package nagchu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 那曲NagchuBarony struct {
	feud.BaseBarony
}

var BNagchu那曲 feud.Barony = &那曲NagchuBarony{}

func init() {
	f := BNagchu那曲.(*那曲NagchuBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nagchu",
		TitleName: "那曲",
		TitleCode: "b_nagchu",
	}
}
