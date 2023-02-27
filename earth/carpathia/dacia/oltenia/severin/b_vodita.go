package severin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沃迪察VoditaBarony struct {
	feud.BaseBarony
}

var BVodita沃迪察 feud.Barony = &沃迪察VoditaBarony{}

func init() {
    f := BVodita沃迪察.(*沃迪察VoditaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vodita",
		TitleName: "沃迪察",
		TitleCode: "b_vodita",
	}
}
