package druz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 古哈利耶GhariyahBarony struct {
	feud.BaseBarony
}

var BGhariyah古哈利耶 feud.Barony = &古哈利耶GhariyahBarony{}

func init() {
    f := BGhariyah古哈利耶.(*古哈利耶GhariyahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ghariyah",
		TitleName: "古哈利耶",
		TitleCode: "b_ghariyah",
	}
}
