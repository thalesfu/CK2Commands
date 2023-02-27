package kangly

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 库利萨雷QulsaryBarony struct {
	feud.BaseBarony
}

var BQulsary库利萨雷 feud.Barony = &库利萨雷QulsaryBarony{}

func init() {
    f := BQulsary库利萨雷.(*库利萨雷QulsaryBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "qulsary",
		TitleName: "库利萨雷",
		TitleCode: "b_qulsary",
	}
}
