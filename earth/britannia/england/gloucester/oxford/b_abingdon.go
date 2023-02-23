package oxford

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿宾登AbingdonBarony struct {
	feud.BaseBarony
}

var BAbingdon阿宾登 feud.Barony = &阿宾登AbingdonBarony{}

func init() {
	f := BAbingdon阿宾登.(*阿宾登AbingdonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "abingdon",
		TitleName: "阿宾登",
		TitleCode: "b_abingdon",
	}
}
