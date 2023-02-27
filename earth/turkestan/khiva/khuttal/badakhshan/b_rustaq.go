package badakhshan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 鲁斯塔克RustaqBarony struct {
	feud.BaseBarony
}

var BRustaq鲁斯塔克 feud.Barony = &鲁斯塔克RustaqBarony{}

func init() {
    f := BRustaq鲁斯塔克.(*鲁斯塔克RustaqBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rustaq",
		TitleName: "鲁斯塔克",
		TitleCode: "b_rustaq",
	}
}
