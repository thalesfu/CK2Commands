package crimea

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 库尔曼QurmanBarony struct {
	feud.BaseBarony
}

var BQurman库尔曼 feud.Barony = &库尔曼QurmanBarony{}

func init() {
    f := BQurman库尔曼.(*库尔曼QurmanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "qurman",
		TitleName: "库尔曼",
		TitleCode: "b_qurman",
	}
}
