package abakan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 比里克丘尔BirikchulBarony struct {
	feud.BaseBarony
}

var BBirikchul比里克丘尔 feud.Barony = &比里克丘尔BirikchulBarony{}

func init() {
	f := BBirikchul比里克丘尔.(*比里克丘尔BirikchulBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "birikchul",
		TitleName: "比里克丘尔",
		TitleCode: "b_birikchul",
	}
}
