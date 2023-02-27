package frisia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 多克姆DokkumBarony struct {
	feud.BaseBarony
}

var BDokkum多克姆 feud.Barony = &多克姆DokkumBarony{}

func init() {
    f := BDokkum多克姆.(*多克姆DokkumBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dokkum",
		TitleName: "多克姆",
		TitleCode: "b_dokkum",
	}
}
