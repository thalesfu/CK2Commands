package york

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 圣彼得St_petersBarony struct {
	feud.BaseBarony
}

var BSt_peters圣彼得 feud.Barony = &圣彼得St_petersBarony{}

func init() {
    f := BSt_peters圣彼得.(*圣彼得St_petersBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "st_peters",
		TitleName: "圣彼得",
		TitleCode: "b_st_peters",
	}
}
