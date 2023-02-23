package birjand

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 博什鲁耶BoshruyehBarony struct {
	feud.BaseBarony
}

var BBoshruyeh博什鲁耶 feud.Barony = &博什鲁耶BoshruyehBarony{}

func init() {
	f := BBoshruyeh博什鲁耶.(*博什鲁耶BoshruyehBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "boshruyeh",
		TitleName: "博什鲁耶",
		TitleCode: "b_boshruyeh",
	}
}
