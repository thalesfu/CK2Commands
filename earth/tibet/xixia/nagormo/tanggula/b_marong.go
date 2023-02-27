package tanggula

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 玛绒MarongBarony struct {
	feud.BaseBarony
}

var BMarong玛绒 feud.Barony = &玛绒MarongBarony{}

func init() {
    f := BMarong玛绒.(*玛绒MarongBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "marong",
		TitleName: "玛绒",
		TitleCode: "b_marong",
	}
}
