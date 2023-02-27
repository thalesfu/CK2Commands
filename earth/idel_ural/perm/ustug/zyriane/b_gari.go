package zyriane

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 加里GariBarony struct {
	feud.BaseBarony
}

var BGari加里 feud.Barony = &加里GariBarony{}

func init() {
    f := BGari加里.(*加里GariBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gari",
		TitleName: "加里",
		TitleCode: "b_gari",
	}
}
