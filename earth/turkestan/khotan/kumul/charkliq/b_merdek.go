package charkliq

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 麦德克MerdekBarony struct {
	feud.BaseBarony
}

var BMerdek麦德克 feud.Barony = &麦德克MerdekBarony{}

func init() {
    f := BMerdek麦德克.(*麦德克MerdekBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "merdek",
		TitleName: "麦德克",
		TitleCode: "b_merdek",
	}
}
