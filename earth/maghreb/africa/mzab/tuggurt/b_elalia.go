package tuggurt

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿利耶ElaliaBarony struct {
	feud.BaseBarony
}

var BElalia阿利耶 feud.Barony = &阿利耶ElaliaBarony{}

func init() {
	f := BElalia阿利耶.(*阿利耶ElaliaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "elalia",
		TitleName: "阿利耶",
		TitleCode: "b_elalia",
	}
}
