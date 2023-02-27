package northampton

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 彼得伯勒PeterboroughBarony struct {
	feud.BaseBarony
}

var BPeterborough彼得伯勒 feud.Barony = &彼得伯勒PeterboroughBarony{}

func init() {
    f := BPeterborough彼得伯勒.(*彼得伯勒PeterboroughBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "peterborough",
		TitleName: "彼得伯勒",
		TitleCode: "b_peterborough",
	}
}
