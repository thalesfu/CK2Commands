package tell_bashir

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 杰拉布卢斯JarabulusBarony struct {
	feud.BaseBarony
}

var BJarabulus杰拉布卢斯 feud.Barony = &杰拉布卢斯JarabulusBarony{}

func init() {
    f := BJarabulus杰拉布卢斯.(*杰拉布卢斯JarabulusBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jarabulus",
		TitleName: "杰拉布卢斯",
		TitleCode: "b_jarabulus",
	}
}
