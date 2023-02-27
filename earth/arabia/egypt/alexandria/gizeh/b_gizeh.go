package gizeh

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 吉萨GizehBarony struct {
	feud.BaseBarony
}

var BGizeh吉萨 feud.Barony = &吉萨GizehBarony{}

func init() {
    f := BGizeh吉萨.(*吉萨GizehBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gizeh",
		TitleName: "吉萨",
		TitleCode: "b_gizeh",
	}
}
