package marmaros

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥克瑙舒高陶格AknasugatagBarony struct {
	feud.BaseBarony
}

var BAknasugatag奥克瑙舒高陶格 feud.Barony = &奥克瑙舒高陶格AknasugatagBarony{}

func init() {
	f := BAknasugatag奥克瑙舒高陶格.(*奥克瑙舒高陶格AknasugatagBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aknasugatag",
		TitleName: "奥克瑙舒高陶格",
		TitleCode: "b_aknasugatag",
	}
}
