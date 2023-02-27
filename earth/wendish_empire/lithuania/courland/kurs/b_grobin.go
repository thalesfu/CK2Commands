package kurs

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 格罗宾GrobinBarony struct {
	feud.BaseBarony
}

var BGrobin格罗宾 feud.Barony = &格罗宾GrobinBarony{}

func init() {
    f := BGrobin格罗宾.(*格罗宾GrobinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "grobin",
		TitleName: "格罗宾",
		TitleCode: "b_grobin",
	}
}
