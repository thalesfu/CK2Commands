package narke

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿斯克松德AskersundBarony struct {
	feud.BaseBarony
}

var BAskersund阿斯克松德 feud.Barony = &阿斯克松德AskersundBarony{}

func init() {
    f := BAskersund阿斯克松德.(*阿斯克松德AskersundBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "askersund",
		TitleName: "阿斯克松德",
		TitleCode: "b_askersund",
	}
}
