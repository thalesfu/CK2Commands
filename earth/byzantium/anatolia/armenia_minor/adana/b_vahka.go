package adana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 费凯VahkaBarony struct {
	feud.BaseBarony
}

var BVahka费凯 feud.Barony = &费凯VahkaBarony{}

func init() {
    f := BVahka费凯.(*费凯VahkaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vahka",
		TitleName: "费凯",
		TitleCode: "b_vahka",
	}
}
