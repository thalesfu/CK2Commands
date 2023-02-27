package sirjan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 法拉盖FaraghehBarony struct {
	feud.BaseBarony
}

var BFaragheh法拉盖 feud.Barony = &法拉盖FaraghehBarony{}

func init() {
    f := BFaragheh法拉盖.(*法拉盖FaraghehBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "faragheh",
		TitleName: "法拉盖",
		TitleCode: "b_faragheh",
	}
}
