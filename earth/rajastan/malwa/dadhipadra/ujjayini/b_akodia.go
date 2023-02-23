package ujjayini

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿戈迪亚AkodiaBarony struct {
	feud.BaseBarony
}

var BAkodia阿戈迪亚 feud.Barony = &阿戈迪亚AkodiaBarony{}

func init() {
	f := BAkodia阿戈迪亚.(*阿戈迪亚AkodiaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "akodia",
		TitleName: "阿戈迪亚",
		TitleCode: "b_akodia",
	}
}
