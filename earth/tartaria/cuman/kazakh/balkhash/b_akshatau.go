package balkhash

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿克沙套AkshatauBarony struct {
	feud.BaseBarony
}

var BAkshatau阿克沙套 feud.Barony = &阿克沙套AkshatauBarony{}

func init() {
	f := BAkshatau阿克沙套.(*阿克沙套AkshatauBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "akshatau",
		TitleName: "阿克沙套",
		TitleCode: "b_akshatau",
	}
}
