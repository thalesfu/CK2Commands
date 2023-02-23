package wagadougou

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 迪巴DibaBarony struct {
	feud.BaseBarony
}

var BDiba迪巴 feud.Barony = &迪巴DibaBarony{}

func init() {
	f := BDiba迪巴.(*迪巴DibaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "diba",
		TitleName: "迪巴",
		TitleCode: "b_diba",
	}
}
