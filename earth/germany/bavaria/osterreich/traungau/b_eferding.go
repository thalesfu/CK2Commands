package traungau

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃费丁EferdingBarony struct {
	feud.BaseBarony
}

var BEferding埃费丁 feud.Barony = &埃费丁EferdingBarony{}

func init() {
	f := BEferding埃费丁.(*埃费丁EferdingBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "eferding",
		TitleName: "埃费丁",
		TitleCode: "b_eferding",
	}
}
