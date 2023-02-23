package alcala

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 锡古恩萨SegontiaBarony struct {
	feud.BaseBarony
}

var BSegontia锡古恩萨 feud.Barony = &锡古恩萨SegontiaBarony{}

func init() {
	f := BSegontia锡古恩萨.(*锡古恩萨SegontiaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "segontia",
		TitleName: "锡古恩萨",
		TitleCode: "b_segontia",
	}
}
