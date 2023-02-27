package tavda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塔利察TalitsaBarony struct {
	feud.BaseBarony
}

var BTalitsa塔利察 feud.Barony = &塔利察TalitsaBarony{}

func init() {
    f := BTalitsa塔利察.(*塔利察TalitsaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "talitsa",
		TitleName: "塔利察",
		TitleCode: "b_talitsa",
	}
}
