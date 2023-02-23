package lettigalians

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塞斯韦根SeswegenBarony struct {
	feud.BaseBarony
}

var BSeswegen塞斯韦根 feud.Barony = &塞斯韦根SeswegenBarony{}

func init() {
	f := BSeswegen塞斯韦根.(*塞斯韦根SeswegenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "seswegen",
		TitleName: "塞斯韦根",
		TitleCode: "b_seswegen",
	}
}
