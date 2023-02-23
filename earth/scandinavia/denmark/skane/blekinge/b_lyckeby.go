package blekinge

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 吕克比LyckebyBarony struct {
	feud.BaseBarony
}

var BLyckeby吕克比 feud.Barony = &吕克比LyckebyBarony{}

func init() {
	f := BLyckeby吕克比.(*吕克比LyckebyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lyckeby",
		TitleName: "吕克比",
		TitleCode: "b_lyckeby",
	}
}
