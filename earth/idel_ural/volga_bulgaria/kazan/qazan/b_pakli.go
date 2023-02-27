package qazan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 帕克利PakliBarony struct {
	feud.BaseBarony
}

var BPakli帕克利 feud.Barony = &帕克利PakliBarony{}

func init() {
    f := BPakli帕克利.(*帕克利PakliBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pakli",
		TitleName: "帕克利",
		TitleCode: "b_pakli",
	}
}
