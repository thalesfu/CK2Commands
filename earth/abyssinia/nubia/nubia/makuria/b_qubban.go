package makuria

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 古班QubbanBarony struct {
	feud.BaseBarony
}

var BQubban古班 feud.Barony = &古班QubbanBarony{}

func init() {
    f := BQubban古班.(*古班QubbanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "qubban",
		TitleName: "古班",
		TitleCode: "b_qubban",
	}
}
