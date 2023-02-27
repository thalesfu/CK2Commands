package thessalia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 纽佩特拉NeopetraBarony struct {
	feud.BaseBarony
}

var BNeopetra纽佩特拉 feud.Barony = &纽佩特拉NeopetraBarony{}

func init() {
    f := BNeopetra纽佩特拉.(*纽佩特拉NeopetraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "neopetra",
		TitleName: "纽佩特拉",
		TitleCode: "b_neopetra",
	}
}
