package mahoyadapuram

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伽罗迪KaladiBarony struct {
	feud.BaseBarony
}

var BKaladi伽罗迪 feud.Barony = &伽罗迪KaladiBarony{}

func init() {
    f := BKaladi伽罗迪.(*伽罗迪KaladiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kaladi",
		TitleName: "伽罗迪",
		TitleCode: "b_kaladi",
	}
}
