package bilino_polje

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 普罗佐尔ProzorBarony struct {
	feud.BaseBarony
}

var BProzor普罗佐尔 feud.Barony = &普罗佐尔ProzorBarony{}

func init() {
    f := BProzor普罗佐尔.(*普罗佐尔ProzorBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "prozor",
		TitleName: "普罗佐尔",
		TitleCode: "b_prozor",
	}
}
