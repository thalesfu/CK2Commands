package otrar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 讹答剌OtrarBarony struct {
	feud.BaseBarony
}

var BOtrar讹答剌 feud.Barony = &讹答剌OtrarBarony{}

func init() {
    f := BOtrar讹答剌.(*讹答剌OtrarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "otrar",
		TitleName: "讹答剌",
		TitleCode: "b_otrar",
	}
}
