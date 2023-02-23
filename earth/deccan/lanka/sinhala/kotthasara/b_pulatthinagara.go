package kotthasara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 补罗悉底耶那揭罗PulatthinagaraBarony struct {
	feud.BaseBarony
}

var BPulatthinagara补罗悉底耶那揭罗 feud.Barony = &补罗悉底耶那揭罗PulatthinagaraBarony{}

func init() {
	f := BPulatthinagara补罗悉底耶那揭罗.(*补罗悉底耶那揭罗PulatthinagaraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pulatthinagara",
		TitleName: "补罗悉底耶那揭罗",
		TitleCode: "b_pulatthinagara",
	}
}
