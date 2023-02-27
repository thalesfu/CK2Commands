package om

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 佐季诺ZotinoBarony struct {
	feud.BaseBarony
}

var BZotino佐季诺 feud.Barony = &佐季诺ZotinoBarony{}

func init() {
    f := BZotino佐季诺.(*佐季诺ZotinoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zotino",
		TitleName: "佐季诺",
		TitleCode: "b_zotino",
	}
}
