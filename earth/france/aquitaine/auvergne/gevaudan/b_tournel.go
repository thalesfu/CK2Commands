package gevaudan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 图尔内尔TournelBarony struct {
	feud.BaseBarony
}

var BTournel图尔内尔 feud.Barony = &图尔内尔TournelBarony{}

func init() {
	f := BTournel图尔内尔.(*图尔内尔TournelBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tournel",
		TitleName: "图尔内尔",
		TitleCode: "b_tournel",
	}
}
