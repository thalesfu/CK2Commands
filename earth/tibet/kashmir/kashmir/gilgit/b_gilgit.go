package gilgit

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 吉尔吉特GilgitBarony struct {
	feud.BaseBarony
}

var BGilgit吉尔吉特 feud.Barony = &吉尔吉特GilgitBarony{}

func init() {
    f := BGilgit吉尔吉特.(*吉尔吉特GilgitBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gilgit",
		TitleName: "吉尔吉特",
		TitleCode: "b_gilgit",
	}
}
