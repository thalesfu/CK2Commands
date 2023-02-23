package nikopolis

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 普莱文PlevenBarony struct {
	feud.BaseBarony
}

var BPleven普莱文 feud.Barony = &普莱文PlevenBarony{}

func init() {
	f := BPleven普莱文.(*普莱文PlevenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pleven",
		TitleName: "普莱文",
		TitleCode: "b_pleven",
	}
}
