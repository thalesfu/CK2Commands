package guines

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 敦刻尔克DunkerqueBarony struct {
	feud.BaseBarony
}

var BDunkerque敦刻尔克 feud.Barony = &敦刻尔克DunkerqueBarony{}

func init() {
    f := BDunkerque敦刻尔克.(*敦刻尔克DunkerqueBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dunkerque",
		TitleName: "敦刻尔克",
		TitleCode: "b_dunkerque",
	}
}
