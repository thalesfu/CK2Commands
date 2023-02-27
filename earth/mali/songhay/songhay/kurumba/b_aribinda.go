package kurumba

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿里宾达AribindaBarony struct {
	feud.BaseBarony
}

var BAribinda阿里宾达 feud.Barony = &阿里宾达AribindaBarony{}

func init() {
    f := BAribinda阿里宾达.(*阿里宾达AribindaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aribinda",
		TitleName: "阿里宾达",
		TitleCode: "b_aribinda",
	}
}
