package kastamon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 庞贝波利斯PompeiopolisBarony struct {
	feud.BaseBarony
}

var BPompeiopolis庞贝波利斯 feud.Barony = &庞贝波利斯PompeiopolisBarony{}

func init() {
	f := BPompeiopolis庞贝波利斯.(*庞贝波利斯PompeiopolisBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pompeiopolis",
		TitleName: "庞贝波利斯",
		TitleCode: "b_pompeiopolis",
	}
}
