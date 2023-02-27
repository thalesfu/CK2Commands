package kolva

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 图尔潘TulpanBarony struct {
	feud.BaseBarony
}

var BTulpan图尔潘 feud.Barony = &图尔潘TulpanBarony{}

func init() {
    f := BTulpan图尔潘.(*图尔潘TulpanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tulpan",
		TitleName: "图尔潘",
		TitleCode: "b_tulpan",
	}
}
