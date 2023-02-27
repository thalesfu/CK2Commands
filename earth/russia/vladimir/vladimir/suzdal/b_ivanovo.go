package suzdal

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊万诺沃IvanovoBarony struct {
	feud.BaseBarony
}

var BIvanovo伊万诺沃 feud.Barony = &伊万诺沃IvanovoBarony{}

func init() {
    f := BIvanovo伊万诺沃.(*伊万诺沃IvanovoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ivanovo",
		TitleName: "伊万诺沃",
		TitleCode: "b_ivanovo",
	}
}
