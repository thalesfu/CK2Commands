package nafusa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 代什雷巴希林Dechret_bahirineBarony struct {
	feud.BaseBarony
}

var BDechret_bahirine代什雷巴希林 feud.Barony = &代什雷巴希林Dechret_bahirineBarony{}

func init() {
    f := BDechret_bahirine代什雷巴希林.(*代什雷巴希林Dechret_bahirineBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dechret_bahirine",
		TitleName: "代什雷巴希林",
		TitleCode: "b_dechret_bahirine",
	}
}
