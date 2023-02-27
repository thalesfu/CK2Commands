package ouadane

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瓦达内WadaneBarony struct {
	feud.BaseBarony
}

var BWadane瓦达内 feud.Barony = &瓦达内WadaneBarony{}

func init() {
    f := BWadane瓦达内.(*瓦达内WadaneBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "wadane",
		TitleName: "瓦达内",
		TitleCode: "b_wadane",
	}
}
