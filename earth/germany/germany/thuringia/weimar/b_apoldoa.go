package weimar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿波尔达ApoldoaBarony struct {
	feud.BaseBarony
}

var BApoldoa阿波尔达 feud.Barony = &阿波尔达ApoldoaBarony{}

func init() {
    f := BApoldoa阿波尔达.(*阿波尔达ApoldoaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "apoldoa",
		TitleName: "阿波尔达",
		TitleCode: "b_apoldoa",
	}
}
