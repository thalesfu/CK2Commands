package guria

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 戈尼奥GonioBarony struct {
	feud.BaseBarony
}

var BGonio戈尼奥 feud.Barony = &戈尼奥GonioBarony{}

func init() {
    f := BGonio戈尼奥.(*戈尼奥GonioBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gonio",
		TitleName: "戈尼奥",
		TitleCode: "b_gonio",
	}
}
