package thessalonike

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塞萨洛尼基ThessalonikiBarony struct {
	feud.BaseBarony
}

var BThessaloniki塞萨洛尼基 feud.Barony = &塞萨洛尼基ThessalonikiBarony{}

func init() {
    f := BThessaloniki塞萨洛尼基.(*塞萨洛尼基ThessalonikiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "thessaloniki",
		TitleName: "塞萨洛尼基",
		TitleCode: "b_thessaloniki",
	}
}
