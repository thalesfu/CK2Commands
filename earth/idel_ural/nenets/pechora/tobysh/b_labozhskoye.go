package tobysh

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉博日斯科耶LabozhskoyeBarony struct {
	feud.BaseBarony
}

var BLabozhskoye拉博日斯科耶 feud.Barony = &拉博日斯科耶LabozhskoyeBarony{}

func init() {
    f := BLabozhskoye拉博日斯科耶.(*拉博日斯科耶LabozhskoyeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "labozhskoye",
		TitleName: "拉博日斯科耶",
		TitleCode: "b_labozhskoye",
	}
}
