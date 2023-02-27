package kamarupanagara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 突阇耶DurjayaBarony struct {
	feud.BaseBarony
}

var BDurjaya突阇耶 feud.Barony = &突阇耶DurjayaBarony{}

func init() {
    f := BDurjaya突阇耶.(*突阇耶DurjayaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "durjaya",
		TitleName: "突阇耶",
		TitleCode: "b_durjaya",
	}
}
