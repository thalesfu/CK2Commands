package kalpi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伽罗毗KalpiBarony struct {
	feud.BaseBarony
}

var BKalpi伽罗毗 feud.Barony = &伽罗毗KalpiBarony{}

func init() {
	f := BKalpi伽罗毗.(*伽罗毗KalpiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kalpi",
		TitleName: "伽罗毗",
		TitleCode: "b_kalpi",
	}
}
