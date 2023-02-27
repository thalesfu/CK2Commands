package pecheneg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 胡托尔卡KhutorkaBarony struct {
	feud.BaseBarony
}

var BKhutorka胡托尔卡 feud.Barony = &胡托尔卡KhutorkaBarony{}

func init() {
    f := BKhutorka胡托尔卡.(*胡托尔卡KhutorkaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khutorka",
		TitleName: "胡托尔卡",
		TitleCode: "b_khutorka",
	}
}
