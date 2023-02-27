package attaleia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 帕尼莫提库斯PanemotichusBarony struct {
	feud.BaseBarony
}

var BPanemotichus帕尼莫提库斯 feud.Barony = &帕尼莫提库斯PanemotichusBarony{}

func init() {
    f := BPanemotichus帕尼莫提库斯.(*帕尼莫提库斯PanemotichusBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "panemotichus",
		TitleName: "帕尼莫提库斯",
		TitleCode: "b_panemotichus",
	}
}
