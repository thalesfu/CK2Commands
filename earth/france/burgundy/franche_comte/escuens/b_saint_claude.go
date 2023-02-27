package escuens

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 圣克洛德Saint_claudeBarony struct {
	feud.BaseBarony
}

var BSaint_claude圣克洛德 feud.Barony = &圣克洛德Saint_claudeBarony{}

func init() {
    f := BSaint_claude圣克洛德.(*圣克洛德Saint_claudeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "saint_claude",
		TitleName: "圣克洛德",
		TitleCode: "b_saint_claude",
	}
}
