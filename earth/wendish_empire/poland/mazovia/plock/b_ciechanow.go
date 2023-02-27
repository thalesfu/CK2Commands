package plock

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 切哈努夫CiechanowBarony struct {
	feud.BaseBarony
}

var BCiechanow切哈努夫 feud.Barony = &切哈努夫CiechanowBarony{}

func init() {
    f := BCiechanow切哈努夫.(*切哈努夫CiechanowBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ciechanow",
		TitleName: "切哈努夫",
		TitleCode: "b_ciechanow",
	}
}
