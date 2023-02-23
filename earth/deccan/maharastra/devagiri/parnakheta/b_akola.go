package parnakheta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿拘罗AkolaBarony struct {
	feud.BaseBarony
}

var BAkola阿拘罗 feud.Barony = &阿拘罗AkolaBarony{}

func init() {
	f := BAkola阿拘罗.(*阿拘罗AkolaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "akola",
		TitleName: "阿拘罗",
		TitleCode: "b_akola",
	}
}
