package manupura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿提利比斯AthribisBarony struct {
	feud.BaseBarony
}

var BAthribis阿提利比斯 feud.Barony = &阿提利比斯AthribisBarony{}

func init() {
    f := BAthribis阿提利比斯.(*阿提利比斯AthribisBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "athribis",
		TitleName: "阿提利比斯",
		TitleCode: "b_athribis",
	}
}
