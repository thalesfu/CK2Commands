package naissus

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布尔多BrdoBarony struct {
	feud.BaseBarony
}

var BBrdo布尔多 feud.Barony = &布尔多BrdoBarony{}

func init() {
    f := BBrdo布尔多.(*布尔多BrdoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "brdo",
		TitleName: "布尔多",
		TitleCode: "b_brdo",
	}
}
