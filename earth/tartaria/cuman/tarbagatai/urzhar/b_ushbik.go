package urzhar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌什比克UshbikBarony struct {
	feud.BaseBarony
}

var BUshbik乌什比克 feud.Barony = &乌什比克UshbikBarony{}

func init() {
    f := BUshbik乌什比克.(*乌什比克UshbikBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ushbik",
		TitleName: "乌什比克",
		TitleCode: "b_ushbik",
	}
}
