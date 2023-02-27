package neuchatel

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 纳沙泰尔NeuchatelBarony struct {
	feud.BaseBarony
}

var BNeuchatel纳沙泰尔 feud.Barony = &纳沙泰尔NeuchatelBarony{}

func init() {
    f := BNeuchatel纳沙泰尔.(*纳沙泰尔NeuchatelBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "neuchatel",
		TitleName: "纳沙泰尔",
		TitleCode: "b_neuchatel",
	}
}
