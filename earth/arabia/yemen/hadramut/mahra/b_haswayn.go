package mahra

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈斯韦因HaswaynBarony struct {
	feud.BaseBarony
}

var BHaswayn哈斯韦因 feud.Barony = &哈斯韦因HaswaynBarony{}

func init() {
	f := BHaswayn哈斯韦因.(*哈斯韦因HaswaynBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "haswayn",
		TitleName: "哈斯韦因",
		TitleCode: "b_haswayn",
	}
}
