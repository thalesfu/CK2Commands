package sarangpur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伽罗恩GagraunBarony struct {
	feud.BaseBarony
}

var BGagraun伽罗恩 feud.Barony = &伽罗恩GagraunBarony{}

func init() {
    f := BGagraun伽罗恩.(*伽罗恩GagraunBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gagraun",
		TitleName: "伽罗恩",
		TitleCode: "b_gagraun",
	}
}
