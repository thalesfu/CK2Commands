package bira

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 比雷BiraBarony struct {
	feud.BaseBarony
}

var BBira比雷 feud.Barony = &比雷BiraBarony{}

func init() {
    f := BBira比雷.(*比雷BiraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bira",
		TitleName: "比雷",
		TitleCode: "b_bira",
	}
}
