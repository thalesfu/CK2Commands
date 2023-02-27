package karor

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沙迪亚ShadiaBarony struct {
	feud.BaseBarony
}

var BShadia沙迪亚 feud.Barony = &沙迪亚ShadiaBarony{}

func init() {
    f := BShadia沙迪亚.(*沙迪亚ShadiaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shadia",
		TitleName: "沙迪亚",
		TitleCode: "b_shadia",
	}
}
