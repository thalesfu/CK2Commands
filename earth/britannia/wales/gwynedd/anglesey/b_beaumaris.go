package anglesey

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 博马里斯BeaumarisBarony struct {
	feud.BaseBarony
}

var BBeaumaris博马里斯 feud.Barony = &博马里斯BeaumarisBarony{}

func init() {
    f := BBeaumaris博马里斯.(*博马里斯BeaumarisBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "beaumaris",
		TitleName: "博马里斯",
		TitleCode: "b_beaumaris",
	}
}
