package kashgar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿克陶AktoBarony struct {
	feud.BaseBarony
}

var BAkto阿克陶 feud.Barony = &阿克陶AktoBarony{}

func init() {
	f := BAkto阿克陶.(*阿克陶AktoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "akto",
		TitleName: "阿克陶",
		TitleCode: "b_akto",
	}
}
