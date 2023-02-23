package pinsk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 平斯克PinskBarony struct {
	feud.BaseBarony
}

var BPinsk平斯克 feud.Barony = &平斯克PinskBarony{}

func init() {
	f := BPinsk平斯克.(*平斯克PinskBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pinsk",
		TitleName: "平斯克",
		TitleCode: "b_pinsk",
	}
}
