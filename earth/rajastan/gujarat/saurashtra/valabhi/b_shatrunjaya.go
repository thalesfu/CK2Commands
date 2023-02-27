package valabhi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 设咄路阇耶ShatrunjayaBarony struct {
	feud.BaseBarony
}

var BShatrunjaya设咄路阇耶 feud.Barony = &设咄路阇耶ShatrunjayaBarony{}

func init() {
    f := BShatrunjaya设咄路阇耶.(*设咄路阇耶ShatrunjayaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shatrunjaya",
		TitleName: "设咄路阇耶",
		TitleCode: "b_shatrunjaya",
	}
}
