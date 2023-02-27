package merya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 齐克迈TsikmaBarony struct {
	feud.BaseBarony
}

var BTsikma齐克迈 feud.Barony = &齐克迈TsikmaBarony{}

func init() {
    f := BTsikma齐克迈.(*齐克迈TsikmaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tsikma",
		TitleName: "齐克迈",
		TitleCode: "b_tsikma",
	}
}
