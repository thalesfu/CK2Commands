package guria

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴统BatumBarony struct {
	feud.BaseBarony
}

var BBatum巴统 feud.Barony = &巴统BatumBarony{}

func init() {
    f := BBatum巴统.(*巴统BatumBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "batum",
		TitleName: "巴统",
		TitleCode: "b_batum",
	}
}
