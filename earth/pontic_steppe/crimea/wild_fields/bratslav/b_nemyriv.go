package bratslav

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 涅米罗夫NemyrivBarony struct {
	feud.BaseBarony
}

var BNemyriv涅米罗夫 feud.Barony = &涅米罗夫NemyrivBarony{}

func init() {
    f := BNemyriv涅米罗夫.(*涅米罗夫NemyrivBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nemyriv",
		TitleName: "涅米罗夫",
		TitleCode: "b_nemyriv",
	}
}
