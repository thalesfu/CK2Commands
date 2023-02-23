package tyrnovo

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 基利法雷夫斯基KilifarevskiBarony struct {
	feud.BaseBarony
}

var BKilifarevski基利法雷夫斯基 feud.Barony = &基利法雷夫斯基KilifarevskiBarony{}

func init() {
	f := BKilifarevski基利法雷夫斯基.(*基利法雷夫斯基KilifarevskiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kilifarevski",
		TitleName: "基利法雷夫斯基",
		TitleCode: "b_kilifarevski",
	}
}
