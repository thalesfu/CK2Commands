package gaya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伽耶GayaBarony struct {
	feud.BaseBarony
}

var BGaya伽耶 feud.Barony = &伽耶GayaBarony{}

func init() {
    f := BGaya伽耶.(*伽耶GayaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gaya",
		TitleName: "伽耶",
		TitleCode: "b_gaya",
	}
}
