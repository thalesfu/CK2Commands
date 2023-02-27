package gwynedd

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 大班戈Bangor_fawrBarony struct {
	feud.BaseBarony
}

var BBangor_fawr大班戈 feud.Barony = &大班戈Bangor_fawrBarony{}

func init() {
    f := BBangor_fawr大班戈.(*大班戈Bangor_fawrBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bangor_fawr",
		TitleName: "大班戈",
		TitleCode: "b_bangor_fawr",
	}
}
