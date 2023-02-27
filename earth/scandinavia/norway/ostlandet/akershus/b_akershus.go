package akershus

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿克什胡斯AkershusBarony struct {
	feud.BaseBarony
}

var BAkershus阿克什胡斯 feud.Barony = &阿克什胡斯AkershusBarony{}

func init() {
    f := BAkershus阿克什胡斯.(*阿克什胡斯AkershusBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "akershus",
		TitleName: "阿克什胡斯",
		TitleCode: "b_akershus",
	}
}
