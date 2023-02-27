package zagreb

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 兹林ZrinBarony struct {
	feud.BaseBarony
}

var BZrin兹林 feud.Barony = &兹林ZrinBarony{}

func init() {
    f := BZrin兹林.(*兹林ZrinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zrin",
		TitleName: "兹林",
		TitleCode: "b_zrin",
	}
}
