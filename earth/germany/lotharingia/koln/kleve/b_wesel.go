package kleve

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 韦塞尔WeselBarony struct {
	feud.BaseBarony
}

var BWesel韦塞尔 feud.Barony = &韦塞尔WeselBarony{}

func init() {
	f := BWesel韦塞尔.(*韦塞尔WeselBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "wesel",
		TitleName: "韦塞尔",
		TitleCode: "b_wesel",
	}
}
