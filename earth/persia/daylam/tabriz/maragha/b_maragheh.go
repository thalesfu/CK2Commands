package maragha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马拉盖MaraghehBarony struct {
	feud.BaseBarony
}

var BMaragheh马拉盖 feud.Barony = &马拉盖MaraghehBarony{}

func init() {
    f := BMaragheh马拉盖.(*马拉盖MaraghehBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "maragheh",
		TitleName: "马拉盖",
		TitleCode: "b_maragheh",
	}
}
