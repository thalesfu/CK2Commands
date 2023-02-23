package koln

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 霍赫施塔登HochstadenBarony struct {
	feud.BaseBarony
}

var BHochstaden霍赫施塔登 feud.Barony = &霍赫施塔登HochstadenBarony{}

func init() {
	f := BHochstaden霍赫施塔登.(*霍赫施塔登HochstadenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hochstaden",
		TitleName: "霍赫施塔登",
		TitleCode: "b_hochstaden",
	}
}
