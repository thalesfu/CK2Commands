package treviso

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 帕埃塞PaeseBarony struct {
	feud.BaseBarony
}

var BPaese帕埃塞 feud.Barony = &帕埃塞PaeseBarony{}

func init() {
	f := BPaese帕埃塞.(*帕埃塞PaeseBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "paese",
		TitleName: "帕埃塞",
		TitleCode: "b_paese",
	}
}
