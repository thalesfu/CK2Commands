package druz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨姆吉SamjBarony struct {
	feud.BaseBarony
}

var BSamj萨姆吉 feud.Barony = &萨姆吉SamjBarony{}

func init() {
    f := BSamj萨姆吉.(*萨姆吉SamjBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "samj",
		TitleName: "萨姆吉",
		TitleCode: "b_samj",
	}
}
