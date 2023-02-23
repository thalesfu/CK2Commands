package pfalz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 施派尔SpeyerBarony struct {
	feud.BaseBarony
}

var BSpeyer施派尔 feud.Barony = &施派尔SpeyerBarony{}

func init() {
	f := BSpeyer施派尔.(*施派尔SpeyerBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "speyer",
		TitleName: "施派尔",
		TitleCode: "b_speyer",
	}
}
