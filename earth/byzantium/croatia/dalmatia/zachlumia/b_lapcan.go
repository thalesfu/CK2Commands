package zachlumia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉普钱LapcanBarony struct {
	feud.BaseBarony
}

var BLapcan拉普钱 feud.Barony = &拉普钱LapcanBarony{}

func init() {
    f := BLapcan拉普钱.(*拉普钱LapcanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lapcan",
		TitleName: "拉普钱",
		TitleCode: "b_lapcan",
	}
}
