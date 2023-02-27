package matamma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伽罗GaroBarony struct {
	feud.BaseBarony
}

var BGaro伽罗 feud.Barony = &伽罗GaroBarony{}

func init() {
    f := BGaro伽罗.(*伽罗GaroBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "garo",
		TitleName: "伽罗",
		TitleCode: "b_garo",
	}
}
