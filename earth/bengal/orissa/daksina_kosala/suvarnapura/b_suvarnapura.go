package suvarnapura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 苏伐剌那补罗SuvarnapuraBarony struct {
	feud.BaseBarony
}

var BSuvarnapura苏伐剌那补罗 feud.Barony = &苏伐剌那补罗SuvarnapuraBarony{}

func init() {
    f := BSuvarnapura苏伐剌那补罗.(*苏伐剌那补罗SuvarnapuraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "suvarnapura",
		TitleName: "苏伐剌那补罗",
		TitleCode: "b_suvarnapura",
	}
}
