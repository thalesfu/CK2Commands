package sutai

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 吞黑勒TonkhilBarony struct {
	feud.BaseBarony
}

var BTonkhil吞黑勒 feud.Barony = &吞黑勒TonkhilBarony{}

func init() {
    f := BTonkhil吞黑勒.(*吞黑勒TonkhilBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tonkhil",
		TitleName: "吞黑勒",
		TitleCode: "b_tonkhil",
	}
}
