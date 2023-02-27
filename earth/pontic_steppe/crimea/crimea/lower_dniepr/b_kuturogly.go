package lower_dniepr

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 库图尔_奥格雷KuturoglyBarony struct {
	feud.BaseBarony
}

var BKuturogly库图尔_奥格雷 feud.Barony = &库图尔_奥格雷KuturoglyBarony{}

func init() {
    f := BKuturogly库图尔_奥格雷.(*库图尔_奥格雷KuturoglyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kuturogly",
		TitleName: "库图尔_奥格雷",
		TitleCode: "b_kuturogly",
	}
}
