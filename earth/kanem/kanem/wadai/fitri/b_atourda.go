package fitri

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿图尔达AtourdaBarony struct {
	feud.BaseBarony
}

var BAtourda阿图尔达 feud.Barony = &阿图尔达AtourdaBarony{}

func init() {
	f := BAtourda阿图尔达.(*阿图尔达AtourdaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "atourda",
		TitleName: "阿图尔达",
		TitleCode: "b_atourda",
	}
}
