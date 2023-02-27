package galaz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 穆尔杰尼MurgeniBarony struct {
	feud.BaseBarony
}

var BMurgeni穆尔杰尼 feud.Barony = &穆尔杰尼MurgeniBarony{}

func init() {
    f := BMurgeni穆尔杰尼.(*穆尔杰尼MurgeniBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "murgeni",
		TitleName: "穆尔杰尼",
		TitleCode: "b_murgeni",
	}
}
