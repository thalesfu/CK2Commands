package talakad

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 苏摩那他补罗SomanathapuraBarony struct {
	feud.BaseBarony
}

var BSomanathapura苏摩那他补罗 feud.Barony = &苏摩那他补罗SomanathapuraBarony{}

func init() {
    f := BSomanathapura苏摩那他补罗.(*苏摩那他补罗SomanathapuraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "somanathapura",
		TitleName: "苏摩那他补罗",
		TitleCode: "b_somanathapura",
	}
}
