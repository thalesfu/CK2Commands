package banbar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沙丁SadengBarony struct {
	feud.BaseBarony
}

var BSadeng沙丁 feud.Barony = &沙丁SadengBarony{}

func init() {
    f := BSadeng沙丁.(*沙丁SadengBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sadeng",
		TitleName: "沙丁",
		TitleCode: "b_sadeng",
	}
}
