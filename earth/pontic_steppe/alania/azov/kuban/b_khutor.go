package kuban

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 胡托尔KhutorBarony struct {
	feud.BaseBarony
}

var BKhutor胡托尔 feud.Barony = &胡托尔KhutorBarony{}

func init() {
    f := BKhutor胡托尔.(*胡托尔KhutorBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khutor",
		TitleName: "胡托尔",
		TitleCode: "b_khutor",
	}
}
