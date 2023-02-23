package torzhok

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥斯塔什科夫OstashkovBarony struct {
	feud.BaseBarony
}

var BOstashkov奥斯塔什科夫 feud.Barony = &奥斯塔什科夫OstashkovBarony{}

func init() {
	f := BOstashkov奥斯塔什科夫.(*奥斯塔什科夫OstashkovBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ostashkov",
		TitleName: "奥斯塔什科夫",
		TitleCode: "b_ostashkov",
	}
}
