package sarangpur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿什多AshtaBarony struct {
	feud.BaseBarony
}

var BAshta阿什多 feud.Barony = &阿什多AshtaBarony{}

func init() {
    f := BAshta阿什多.(*阿什多AshtaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ashta",
		TitleName: "阿什多",
		TitleCode: "b_ashta",
	}
}
