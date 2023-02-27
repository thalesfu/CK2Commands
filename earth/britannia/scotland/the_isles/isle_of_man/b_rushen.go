package isle_of_man

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 鲁森RushenBarony struct {
	feud.BaseBarony
}

var BRushen鲁森 feud.Barony = &鲁森RushenBarony{}

func init() {
    f := BRushen鲁森.(*鲁森RushenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rushen",
		TitleName: "鲁森",
		TitleCode: "b_rushen",
	}
}
