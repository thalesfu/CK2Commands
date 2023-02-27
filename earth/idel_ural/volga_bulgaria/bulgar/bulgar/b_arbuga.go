package bulgar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尔布加ArbugaBarony struct {
	feud.BaseBarony
}

var BArbuga阿尔布加 feud.Barony = &阿尔布加ArbugaBarony{}

func init() {
    f := BArbuga阿尔布加.(*阿尔布加ArbugaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "arbuga",
		TitleName: "阿尔布加",
		TitleCode: "b_arbuga",
	}
}
