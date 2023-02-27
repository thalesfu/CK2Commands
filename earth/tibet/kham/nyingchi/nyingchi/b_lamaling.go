package nyingchi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 喇嘛林LamalingBarony struct {
	feud.BaseBarony
}

var BLamaling喇嘛林 feud.Barony = &喇嘛林LamalingBarony{}

func init() {
    f := BLamaling喇嘛林.(*喇嘛林LamalingBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lamaling",
		TitleName: "喇嘛林",
		TitleCode: "b_lamaling",
	}
}
