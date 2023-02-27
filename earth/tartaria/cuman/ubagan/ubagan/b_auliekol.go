package ubagan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥利耶科利AuliekolBarony struct {
	feud.BaseBarony
}

var BAuliekol奥利耶科利 feud.Barony = &奥利耶科利AuliekolBarony{}

func init() {
    f := BAuliekol奥利耶科利.(*奥利耶科利AuliekolBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "auliekol",
		TitleName: "奥利耶科利",
		TitleCode: "b_auliekol",
	}
}
