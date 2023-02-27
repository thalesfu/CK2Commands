package hereford

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克利福德CliffordBarony struct {
	feud.BaseBarony
}

var BClifford克利福德 feud.Barony = &克利福德CliffordBarony{}

func init() {
    f := BClifford克利福德.(*克利福德CliffordBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "clifford",
		TitleName: "克利福德",
		TitleCode: "b_clifford",
	}
}
