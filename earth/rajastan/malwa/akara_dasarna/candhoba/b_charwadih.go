package candhoba

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 遮罗婆提诃CharwadihBarony struct {
	feud.BaseBarony
}

var BCharwadih遮罗婆提诃 feud.Barony = &遮罗婆提诃CharwadihBarony{}

func init() {
    f := BCharwadih遮罗婆提诃.(*遮罗婆提诃CharwadihBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "charwadih",
		TitleName: "遮罗婆提诃",
		TitleCode: "b_charwadih",
	}
}
