package lindisfarne

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 康希尔CornhillBarony struct {
	feud.BaseBarony
}

var BCornhill康希尔 feud.Barony = &康希尔CornhillBarony{}

func init() {
    f := BCornhill康希尔.(*康希尔CornhillBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cornhill",
		TitleName: "康希尔",
		TitleCode: "b_cornhill",
	}
}
