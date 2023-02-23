package nyima

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 琼宗KhyungdzongBarony struct {
	feud.BaseBarony
}

var BKhyungdzong琼宗 feud.Barony = &琼宗KhyungdzongBarony{}

func init() {
	f := BKhyungdzong琼宗.(*琼宗KhyungdzongBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khyungdzong",
		TitleName: "琼宗",
		TitleCode: "b_khyungdzong",
	}
}
