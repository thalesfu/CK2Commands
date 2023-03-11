package peremyshl

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 佩列梅什利PeremyshlBarony struct {
	feud.BaseBarony
}

var BPeremyshl佩列梅什利 feud.Barony = &佩列梅什利PeremyshlBarony{}

func init() {
    f := BPeremyshl佩列梅什利.(*佩列梅什利PeremyshlBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "peremyshl",
		TitleName: "佩列梅什利",
		TitleCode: "b_peremyshl",
	}
}
