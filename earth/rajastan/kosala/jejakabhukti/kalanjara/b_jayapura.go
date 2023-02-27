package kalanjara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阇耶城JayapuraBarony struct {
	feud.BaseBarony
}

var BJayapura阇耶城 feud.Barony = &阇耶城JayapuraBarony{}

func init() {
    f := BJayapura阇耶城.(*阇耶城JayapuraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jayapura",
		TitleName: "阇耶城",
		TitleCode: "b_jayapura",
	}
}
