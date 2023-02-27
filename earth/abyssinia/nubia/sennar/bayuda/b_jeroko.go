package bayuda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 杰罗科JerokoBarony struct {
	feud.BaseBarony
}

var BJeroko杰罗科 feud.Barony = &杰罗科JerokoBarony{}

func init() {
    f := BJeroko杰罗科.(*杰罗科JerokoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jeroko",
		TitleName: "杰罗科",
		TitleCode: "b_jeroko",
	}
}
