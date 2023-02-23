package narbonne

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 皮塞吉耶PuisserguierBarony struct {
	feud.BaseBarony
}

var BPuisserguier皮塞吉耶 feud.Barony = &皮塞吉耶PuisserguierBarony{}

func init() {
	f := BPuisserguier皮塞吉耶.(*皮塞吉耶PuisserguierBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "puisserguier",
		TitleName: "皮塞吉耶",
		TitleCode: "b_puisserguier",
	}
}
