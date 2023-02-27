package snassen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌季达OujdaBarony struct {
	feud.BaseBarony
}

var BOujda乌季达 feud.Barony = &乌季达OujdaBarony{}

func init() {
    f := BOujda乌季达.(*乌季达OujdaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "oujda",
		TitleName: "乌季达",
		TitleCode: "b_oujda",
	}
}
