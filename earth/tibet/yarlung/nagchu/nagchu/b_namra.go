package nagchu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 那玛切NamraBarony struct {
	feud.BaseBarony
}

var BNamra那玛切 feud.Barony = &那玛切NamraBarony{}

func init() {
    f := BNamra那玛切.(*那玛切NamraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "namra",
		TitleName: "那玛切",
		TitleCode: "b_namra",
	}
}
