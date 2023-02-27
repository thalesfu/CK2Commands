package miass

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 辛古利SingulBarony struct {
	feud.BaseBarony
}

var BSingul辛古利 feud.Barony = &辛古利SingulBarony{}

func init() {
    f := BSingul辛古利.(*辛古利SingulBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "singul",
		TitleName: "辛古利",
		TitleCode: "b_singul",
	}
}
