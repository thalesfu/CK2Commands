package qazwin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿卜耶克AbeyekBarony struct {
	feud.BaseBarony
}

var BAbeyek阿卜耶克 feud.Barony = &阿卜耶克AbeyekBarony{}

func init() {
	f := BAbeyek阿卜耶克.(*阿卜耶克AbeyekBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "abeyek",
		TitleName: "阿卜耶克",
		TitleCode: "b_abeyek",
	}
}
