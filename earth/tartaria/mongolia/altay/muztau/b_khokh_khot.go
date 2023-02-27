package muztau

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 呼和浩特Khokh_khotBarony struct {
	feud.BaseBarony
}

var BKhokh_khot呼和浩特 feud.Barony = &呼和浩特Khokh_khotBarony{}

func init() {
    f := BKhokh_khot呼和浩特.(*呼和浩特Khokh_khotBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khokh_khot",
		TitleName: "呼和浩特",
		TitleCode: "b_khokh_khot",
	}
}
