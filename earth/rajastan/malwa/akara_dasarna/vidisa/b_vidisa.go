package vidisa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卑提写VidisaBarony struct {
	feud.BaseBarony
}

var BVidisa卑提写 feud.Barony = &卑提写VidisaBarony{}

func init() {
    f := BVidisa卑提写.(*卑提写VidisaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vidisa",
		TitleName: "卑提写",
		TitleCode: "b_vidisa",
	}
}
