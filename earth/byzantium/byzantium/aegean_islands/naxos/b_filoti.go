package naxos

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 菲洛提FilotiBarony struct {
	feud.BaseBarony
}

var BFiloti菲洛提 feud.Barony = &菲洛提FilotiBarony{}

func init() {
    f := BFiloti菲洛提.(*菲洛提FilotiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "filoti",
		TitleName: "菲洛提",
		TitleCode: "b_filoti",
	}
}
