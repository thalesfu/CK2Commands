package paphlagonia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 博卢BoluBarony struct {
	feud.BaseBarony
}

var BBolu博卢 feud.Barony = &博卢BoluBarony{}

func init() {
    f := BBolu博卢.(*博卢BoluBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bolu",
		TitleName: "博卢",
		TitleCode: "b_bolu",
	}
}
