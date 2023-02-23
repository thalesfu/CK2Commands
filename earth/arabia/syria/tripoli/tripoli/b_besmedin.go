package tripoli

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 比梅丁BesmedinBarony struct {
	feud.BaseBarony
}

var BBesmedin比梅丁 feud.Barony = &比梅丁BesmedinBarony{}

func init() {
	f := BBesmedin比梅丁.(*比梅丁BesmedinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "besmedin",
		TitleName: "比梅丁",
		TitleCode: "b_besmedin",
	}
}
