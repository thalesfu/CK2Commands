package vodamayutja

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 祇伽GeghaBarony struct {
	feud.BaseBarony
}

var BGegha祇伽 feud.Barony = &祇伽GeghaBarony{}

func init() {
    f := BGegha祇伽.(*祇伽GeghaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gegha",
		TitleName: "祇伽",
		TitleCode: "b_gegha",
	}
}
