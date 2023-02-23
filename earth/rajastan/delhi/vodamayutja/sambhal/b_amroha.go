package sambhal

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 庵卢诃AmrohaBarony struct {
	feud.BaseBarony
}

var BAmroha庵卢诃 feud.Barony = &庵卢诃AmrohaBarony{}

func init() {
	f := BAmroha庵卢诃.(*庵卢诃AmrohaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "amroha",
		TitleName: "庵卢诃",
		TitleCode: "b_amroha",
	}
}
