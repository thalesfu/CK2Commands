package pinsk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊万诺沃IvanavaBarony struct {
	feud.BaseBarony
}

var BIvanava伊万诺沃 feud.Barony = &伊万诺沃IvanavaBarony{}

func init() {
	f := BIvanava伊万诺沃.(*伊万诺沃IvanavaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ivanava",
		TitleName: "伊万诺沃",
		TitleCode: "b_ivanava",
	}
}
