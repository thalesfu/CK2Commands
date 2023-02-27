package bedford

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贝德福德BedfordBarony struct {
	feud.BaseBarony
}

var BBedford贝德福德 feud.Barony = &贝德福德BedfordBarony{}

func init() {
    f := BBedford贝德福德.(*贝德福德BedfordBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bedford",
		TitleName: "贝德福德",
		TitleCode: "b_bedford",
	}
}
