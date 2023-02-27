package istakhr

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊格利德IqlidBarony struct {
	feud.BaseBarony
}

var BIqlid伊格利德 feud.Barony = &伊格利德IqlidBarony{}

func init() {
    f := BIqlid伊格利德.(*伊格利德IqlidBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "iqlid",
		TitleName: "伊格利德",
		TitleCode: "b_iqlid",
	}
}
