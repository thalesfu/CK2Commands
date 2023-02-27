package ifni

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 提德基特TidergitBarony struct {
	feud.BaseBarony
}

var BTidergit提德基特 feud.Barony = &提德基特TidergitBarony{}

func init() {
    f := BTidergit提德基特.(*提德基特TidergitBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tidergit",
		TitleName: "提德基特",
		TitleCode: "b_tidergit",
	}
}
