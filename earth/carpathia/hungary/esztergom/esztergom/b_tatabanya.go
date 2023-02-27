package esztergom

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 陶陶巴尼奥TatabanyaBarony struct {
	feud.BaseBarony
}

var BTatabanya陶陶巴尼奥 feud.Barony = &陶陶巴尼奥TatabanyaBarony{}

func init() {
    f := BTatabanya陶陶巴尼奥.(*陶陶巴尼奥TatabanyaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tatabanya",
		TitleName: "陶陶巴尼奥",
		TitleCode: "b_tatabanya",
	}
}
