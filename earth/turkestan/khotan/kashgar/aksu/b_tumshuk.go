package aksu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 图木舒克TumshukBarony struct {
	feud.BaseBarony
}

var BTumshuk图木舒克 feud.Barony = &图木舒克TumshukBarony{}

func init() {
    f := BTumshuk图木舒克.(*图木舒克TumshukBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tumshuk",
		TitleName: "图木舒克",
		TitleCode: "b_tumshuk",
	}
}
