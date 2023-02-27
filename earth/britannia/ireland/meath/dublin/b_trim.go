package dublin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 特里姆TrimBarony struct {
	feud.BaseBarony
}

var BTrim特里姆 feud.Barony = &特里姆TrimBarony{}

func init() {
    f := BTrim特里姆.(*特里姆TrimBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "trim",
		TitleName: "特里姆",
		TitleCode: "b_trim",
	}
}
