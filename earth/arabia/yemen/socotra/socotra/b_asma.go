package socotra

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿斯马AsmaBarony struct {
	feud.BaseBarony
}

var BAsma阿斯马 feud.Barony = &阿斯马AsmaBarony{}

func init() {
    f := BAsma阿斯马.(*阿斯马AsmaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "asma",
		TitleName: "阿斯马",
		TitleCode: "b_asma",
	}
}
