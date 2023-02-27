package sortavala

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉赫坚波希亚LakhdenpokhiaBarony struct {
	feud.BaseBarony
}

var BLakhdenpokhia拉赫坚波希亚 feud.Barony = &拉赫坚波希亚LakhdenpokhiaBarony{}

func init() {
    f := BLakhdenpokhia拉赫坚波希亚.(*拉赫坚波希亚LakhdenpokhiaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lakhdenpokhia",
		TitleName: "拉赫坚波希亚",
		TitleCode: "b_lakhdenpokhia",
	}
}
