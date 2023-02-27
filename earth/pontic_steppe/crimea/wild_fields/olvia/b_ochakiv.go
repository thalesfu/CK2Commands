package olvia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥恰科夫OchakivBarony struct {
	feud.BaseBarony
}

var BOchakiv奥恰科夫 feud.Barony = &奥恰科夫OchakivBarony{}

func init() {
    f := BOchakiv奥恰科夫.(*奥恰科夫OchakivBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ochakiv",
		TitleName: "奥恰科夫",
		TitleCode: "b_ochakiv",
	}
}
