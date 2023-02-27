package tiberias

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贝特谢安BethsanBarony struct {
	feud.BaseBarony
}

var BBethsan贝特谢安 feud.Barony = &贝特谢安BethsanBarony{}

func init() {
    f := BBethsan贝特谢安.(*贝特谢安BethsanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bethsan",
		TitleName: "贝特谢安",
		TitleCode: "b_bethsan",
	}
}
