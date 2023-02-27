package sasaram

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 善陀JaundBarony struct {
	feud.BaseBarony
}

var BJaund善陀 feud.Barony = &善陀JaundBarony{}

func init() {
    f := BJaund善陀.(*善陀JaundBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jaund",
		TitleName: "善陀",
		TitleCode: "b_jaund",
	}
}
