package oberbayern

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃伯斯贝格EbersbergBarony struct {
	feud.BaseBarony
}

var BEbersberg埃伯斯贝格 feud.Barony = &埃伯斯贝格EbersbergBarony{}

func init() {
	f := BEbersberg埃伯斯贝格.(*埃伯斯贝格EbersbergBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ebersberg",
		TitleName: "埃伯斯贝格",
		TitleCode: "b_ebersberg",
	}
}
