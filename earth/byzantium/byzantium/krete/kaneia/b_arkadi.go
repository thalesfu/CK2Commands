package kaneia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尔卡季ArkadiBarony struct {
	feud.BaseBarony
}

var BArkadi阿尔卡季 feud.Barony = &阿尔卡季ArkadiBarony{}

func init() {
	f := BArkadi阿尔卡季.(*阿尔卡季ArkadiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "arkadi",
		TitleName: "阿尔卡季",
		TitleCode: "b_arkadi",
	}
}
