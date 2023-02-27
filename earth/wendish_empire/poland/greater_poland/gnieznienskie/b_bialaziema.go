package gnieznienskie

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 比亚瓦杰马BialaziemaBarony struct {
	feud.BaseBarony
}

var BBialaziema比亚瓦杰马 feud.Barony = &比亚瓦杰马BialaziemaBarony{}

func init() {
    f := BBialaziema比亚瓦杰马.(*比亚瓦杰马BialaziemaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bialaziema",
		TitleName: "比亚瓦杰马",
		TitleCode: "b_bialaziema",
	}
}
