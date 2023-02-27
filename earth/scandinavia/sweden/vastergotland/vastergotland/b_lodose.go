package vastergotland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 勒德瑟LodoseBarony struct {
	feud.BaseBarony
}

var BLodose勒德瑟 feud.Barony = &勒德瑟LodoseBarony{}

func init() {
    f := BLodose勒德瑟.(*勒德瑟LodoseBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lodose",
		TitleName: "勒德瑟",
		TitleCode: "b_lodose",
	}
}
