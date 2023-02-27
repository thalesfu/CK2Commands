package zavarot

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 佩列梅茨PylemetsBarony struct {
	feud.BaseBarony
}

var BPylemets佩列梅茨 feud.Barony = &佩列梅茨PylemetsBarony{}

func init() {
    f := BPylemets佩列梅茨.(*佩列梅茨PylemetsBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pylemets",
		TitleName: "佩列梅茨",
		TitleCode: "b_pylemets",
	}
}
