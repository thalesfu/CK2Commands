package janakpur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 提婆伽吒DevghatBarony struct {
	feud.BaseBarony
}

var BDevghat提婆伽吒 feud.Barony = &提婆伽吒DevghatBarony{}

func init() {
    f := BDevghat提婆伽吒.(*提婆伽吒DevghatBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "devghat",
		TitleName: "提婆伽吒",
		TitleCode: "b_devghat",
	}
}
