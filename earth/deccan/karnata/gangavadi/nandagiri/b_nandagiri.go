package nandagiri

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 难陀山NandagiriBarony struct {
	feud.BaseBarony
}

var BNandagiri难陀山 feud.Barony = &难陀山NandagiriBarony{}

func init() {
	f := BNandagiri难陀山.(*难陀山NandagiriBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nandagiri",
		TitleName: "难陀山",
		TitleCode: "b_nandagiri",
	}
}
