package beersheb

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 底璧DebirBarony struct {
	feud.BaseBarony
}

var BDebir底璧 feud.Barony = &底璧DebirBarony{}

func init() {
    f := BDebir底璧.(*底璧DebirBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "debir",
		TitleName: "底璧",
		TitleCode: "b_debir",
	}
}
