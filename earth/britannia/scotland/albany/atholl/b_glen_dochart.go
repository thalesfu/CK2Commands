package atholl

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 多赫特河谷Glen_dochartBarony struct {
	feud.BaseBarony
}

var BGlen_dochart多赫特河谷 feud.Barony = &多赫特河谷Glen_dochartBarony{}

func init() {
    f := BGlen_dochart多赫特河谷.(*多赫特河谷Glen_dochartBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "glen_dochart",
		TitleName: "多赫特河谷",
		TitleCode: "b_glen_dochart",
	}
}
