package jaen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈恩JaenBarony struct {
	feud.BaseBarony
}

var BJaen哈恩 feud.Barony = &哈恩JaenBarony{}

func init() {
    f := BJaen哈恩.(*哈恩JaenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jaen",
		TitleName: "哈恩",
		TitleCode: "b_jaen",
	}
}
