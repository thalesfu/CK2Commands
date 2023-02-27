package jharkand

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 房毗耶BhambiyaBarony struct {
	feud.BaseBarony
}

var BBhambiya房毗耶 feud.Barony = &房毗耶BhambiyaBarony{}

func init() {
    f := BBhambiya房毗耶.(*房毗耶BhambiyaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bhambiya",
		TitleName: "房毗耶",
		TitleCode: "b_bhambiya",
	}
}
