package atbara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌姆迈迪UmmmardiBarony struct {
	feud.BaseBarony
}

var BUmmmardi乌姆迈迪 feud.Barony = &乌姆迈迪UmmmardiBarony{}

func init() {
    f := BUmmmardi乌姆迈迪.(*乌姆迈迪UmmmardiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ummmardi",
		TitleName: "乌姆迈迪",
		TitleCode: "b_ummmardi",
	}
}
