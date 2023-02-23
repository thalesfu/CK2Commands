package almeria

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贝尔哈BerjaBarony struct {
	feud.BaseBarony
}

var BBerja贝尔哈 feud.Barony = &贝尔哈BerjaBarony{}

func init() {
	f := BBerja贝尔哈.(*贝尔哈BerjaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "berja",
		TitleName: "贝尔哈",
		TitleCode: "b_berja",
	}
}
