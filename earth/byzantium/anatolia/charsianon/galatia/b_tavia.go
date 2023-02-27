package galatia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 泰维亚TaviaBarony struct {
	feud.BaseBarony
}

var BTavia泰维亚 feud.Barony = &泰维亚TaviaBarony{}

func init() {
    f := BTavia泰维亚.(*泰维亚TaviaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tavia",
		TitleName: "泰维亚",
		TitleCode: "b_tavia",
	}
}
