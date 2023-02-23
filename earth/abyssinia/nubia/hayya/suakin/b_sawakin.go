package suakin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨瓦金SawakinBarony struct {
	feud.BaseBarony
}

var BSawakin萨瓦金 feud.Barony = &萨瓦金SawakinBarony{}

func init() {
	f := BSawakin萨瓦金.(*萨瓦金SawakinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sawakin",
		TitleName: "萨瓦金",
		TitleCode: "b_sawakin",
	}
}
