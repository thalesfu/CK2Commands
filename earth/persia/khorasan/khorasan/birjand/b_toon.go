package birjand

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 顿恩ToonBarony struct {
	feud.BaseBarony
}

var BToon顿恩 feud.Barony = &顿恩ToonBarony{}

func init() {
	f := BToon顿恩.(*顿恩ToonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "toon",
		TitleName: "顿恩",
		TitleCode: "b_toon",
	}
}
