package leinster

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 利林LeighlinBarony struct {
	feud.BaseBarony
}

var BLeighlin利林 feud.Barony = &利林LeighlinBarony{}

func init() {
    f := BLeighlin利林.(*利林LeighlinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "leighlin",
		TitleName: "利林",
		TitleCode: "b_leighlin",
	}
}
