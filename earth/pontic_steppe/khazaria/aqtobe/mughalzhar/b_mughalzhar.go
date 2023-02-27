package mughalzhar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 穆加尔扎尔MughalzharBarony struct {
	feud.BaseBarony
}

var BMughalzhar穆加尔扎尔 feud.Barony = &穆加尔扎尔MughalzharBarony{}

func init() {
    f := BMughalzhar穆加尔扎尔.(*穆加尔扎尔MughalzharBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mughalzhar",
		TitleName: "穆加尔扎尔",
		TitleCode: "b_mughalzhar",
	}
}
