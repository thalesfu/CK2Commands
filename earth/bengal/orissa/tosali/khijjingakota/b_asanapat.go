package khijjingakota

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿桑纳帕提AsanapatBarony struct {
	feud.BaseBarony
}

var BAsanapat阿桑纳帕提 feud.Barony = &阿桑纳帕提AsanapatBarony{}

func init() {
    f := BAsanapat阿桑纳帕提.(*阿桑纳帕提AsanapatBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "asanapat",
		TitleName: "阿桑纳帕提",
		TitleCode: "b_asanapat",
	}
}
