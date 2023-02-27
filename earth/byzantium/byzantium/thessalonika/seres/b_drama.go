package seres

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 兹拉马DramaBarony struct {
	feud.BaseBarony
}

var BDrama兹拉马 feud.Barony = &兹拉马DramaBarony{}

func init() {
    f := BDrama兹拉马.(*兹拉马DramaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "drama",
		TitleName: "兹拉马",
		TitleCode: "b_drama",
	}
}
