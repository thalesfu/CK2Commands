package ankyra

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 特克托萨季TectosagiiBarony struct {
	feud.BaseBarony
}

var BTectosagii特克托萨季 feud.Barony = &特克托萨季TectosagiiBarony{}

func init() {
	f := BTectosagii特克托萨季.(*特克托萨季TectosagiiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tectosagii",
		TitleName: "特克托萨季",
		TitleCode: "b_tectosagii",
	}
}
