package viken

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯维尼比SvennebyBarony struct {
	feud.BaseBarony
}

var BSvenneby斯维尼比 feud.Barony = &斯维尼比SvennebyBarony{}

func init() {
    f := BSvenneby斯维尼比.(*斯维尼比SvennebyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "svenneby",
		TitleName: "斯维尼比",
		TitleCode: "b_svenneby",
	}
}
