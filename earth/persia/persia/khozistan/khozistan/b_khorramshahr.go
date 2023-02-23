package khozistan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 霍拉姆沙赫尔KhorramshahrBarony struct {
	feud.BaseBarony
}

var BKhorramshahr霍拉姆沙赫尔 feud.Barony = &霍拉姆沙赫尔KhorramshahrBarony{}

func init() {
	f := BKhorramshahr霍拉姆沙赫尔.(*霍拉姆沙赫尔KhorramshahrBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khorramshahr",
		TitleName: "霍拉姆沙赫尔",
		TitleCode: "b_khorramshahr",
	}
}
