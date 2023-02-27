package tuggurt

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 图古尔特TuggurtBarony struct {
	feud.BaseBarony
}

var BTuggurt图古尔特 feud.Barony = &图古尔特TuggurtBarony{}

func init() {
    f := BTuggurt图古尔特.(*图古尔特TuggurtBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tuggurt",
		TitleName: "图古尔特",
		TitleCode: "b_tuggurt",
	}
}
