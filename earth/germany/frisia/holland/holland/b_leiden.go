package holland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 雷登LeidenBarony struct {
	feud.BaseBarony
}

var BLeiden雷登 feud.Barony = &雷登LeidenBarony{}

func init() {
	f := BLeiden雷登.(*雷登LeidenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "leiden",
		TitleName: "雷登",
		TitleCode: "b_leiden",
	}
}
