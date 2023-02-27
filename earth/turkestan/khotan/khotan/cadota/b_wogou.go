package cadota

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 窝沟WogouBarony struct {
	feud.BaseBarony
}

var BWogou窝沟 feud.Barony = &窝沟WogouBarony{}

func init() {
    f := BWogou窝沟.(*窝沟WogouBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "wogou",
		TitleName: "窝沟",
		TitleCode: "b_wogou",
	}
}
