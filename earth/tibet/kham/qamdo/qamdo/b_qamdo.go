package qamdo

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 察木多QamdoBarony struct {
	feud.BaseBarony
}

var BQamdo察木多 feud.Barony = &察木多QamdoBarony{}

func init() {
	f := BQamdo察木多.(*察木多QamdoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "qamdo",
		TitleName: "察木多",
		TitleCode: "b_qamdo",
	}
}
