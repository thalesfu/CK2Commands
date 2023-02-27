package gauda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 谛梨阿迦罗TeliagarhiBarony struct {
	feud.BaseBarony
}

var BTeliagarhi谛梨阿迦罗 feud.Barony = &谛梨阿迦罗TeliagarhiBarony{}

func init() {
    f := BTeliagarhi谛梨阿迦罗.(*谛梨阿迦罗TeliagarhiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "teliagarhi",
		TitleName: "谛梨阿迦罗",
		TitleCode: "b_teliagarhi",
	}
}
