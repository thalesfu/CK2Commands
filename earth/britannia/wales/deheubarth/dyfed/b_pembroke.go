package dyfed

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 彭布罗克PembrokeBarony struct {
	feud.BaseBarony
}

var BPembroke彭布罗克 feud.Barony = &彭布罗克PembrokeBarony{}

func init() {
    f := BPembroke彭布罗克.(*彭布罗克PembrokeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pembroke",
		TitleName: "彭布罗克",
		TitleCode: "b_pembroke",
	}
}
