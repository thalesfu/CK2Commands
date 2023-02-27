package brno

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 赖赫拉德RajhradBarony struct {
	feud.BaseBarony
}

var BRajhrad赖赫拉德 feud.Barony = &赖赫拉德RajhradBarony{}

func init() {
    f := BRajhrad赖赫拉德.(*赖赫拉德RajhradBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rajhrad",
		TitleName: "赖赫拉德",
		TitleCode: "b_rajhrad",
	}
}
