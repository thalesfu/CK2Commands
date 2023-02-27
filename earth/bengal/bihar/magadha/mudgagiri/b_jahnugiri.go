package mudgagiri

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阇喜奴山JahnugiriBarony struct {
	feud.BaseBarony
}

var BJahnugiri阇喜奴山 feud.Barony = &阇喜奴山JahnugiriBarony{}

func init() {
    f := BJahnugiri阇喜奴山.(*阇喜奴山JahnugiriBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jahnugiri",
		TitleName: "阇喜奴山",
		TitleCode: "b_jahnugiri",
	}
}
