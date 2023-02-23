package emba

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 穆库尔MukurBarony struct {
	feud.BaseBarony
}

var BMukur穆库尔 feud.Barony = &穆库尔MukurBarony{}

func init() {
	f := BMukur穆库尔.(*穆库尔MukurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mukur",
		TitleName: "穆库尔",
		TitleCode: "b_mukur",
	}
}
