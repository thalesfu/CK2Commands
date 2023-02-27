package potapi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 帝楼波底TirupatiBarony struct {
	feud.BaseBarony
}

var BTirupati帝楼波底 feud.Barony = &帝楼波底TirupatiBarony{}

func init() {
    f := BTirupati帝楼波底.(*帝楼波底TirupatiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tirupati",
		TitleName: "帝楼波底",
		TitleCode: "b_tirupati",
	}
}
