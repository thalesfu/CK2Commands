package perigord

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贝尔热拉克BergeracBarony struct {
	feud.BaseBarony
}

var BBergerac贝尔热拉克 feud.Barony = &贝尔热拉克BergeracBarony{}

func init() {
	f := BBergerac贝尔热拉克.(*贝尔热拉克BergeracBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bergerac",
		TitleName: "贝尔热拉克",
		TitleCode: "b_bergerac",
	}
}
