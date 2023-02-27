package az_zarqa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈拉巴特堡QasralhallabatBarony struct {
	feud.BaseBarony
}

var BQasralhallabat哈拉巴特堡 feud.Barony = &哈拉巴特堡QasralhallabatBarony{}

func init() {
    f := BQasralhallabat哈拉巴特堡.(*哈拉巴特堡QasralhallabatBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "qasralhallabat",
		TitleName: "哈拉巴特堡",
		TitleCode: "b_qasralhallabat",
	}
}
