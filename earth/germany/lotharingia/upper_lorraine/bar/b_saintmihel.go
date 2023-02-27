package bar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 圣米耶勒SaintmihelBarony struct {
	feud.BaseBarony
}

var BSaintmihel圣米耶勒 feud.Barony = &圣米耶勒SaintmihelBarony{}

func init() {
    f := BSaintmihel圣米耶勒.(*圣米耶勒SaintmihelBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "saintmihel",
		TitleName: "圣米耶勒",
		TitleCode: "b_saintmihel",
	}
}
