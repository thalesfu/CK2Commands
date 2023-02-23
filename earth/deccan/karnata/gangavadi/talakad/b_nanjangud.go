package talakad

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 难若那瞿荼NanjangudBarony struct {
	feud.BaseBarony
}

var BNanjangud难若那瞿荼 feud.Barony = &难若那瞿荼NanjangudBarony{}

func init() {
	f := BNanjangud难若那瞿荼.(*难若那瞿荼NanjangudBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nanjangud",
		TitleName: "难若那瞿荼",
		TitleCode: "b_nanjangud",
	}
}
