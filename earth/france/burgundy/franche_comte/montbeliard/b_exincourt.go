package montbeliard

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃克桑库尔ExincourtBarony struct {
	feud.BaseBarony
}

var BExincourt埃克桑库尔 feud.Barony = &埃克桑库尔ExincourtBarony{}

func init() {
    f := BExincourt埃克桑库尔.(*埃克桑库尔ExincourtBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "exincourt",
		TitleName: "埃克桑库尔",
		TitleCode: "b_exincourt",
	}
}
