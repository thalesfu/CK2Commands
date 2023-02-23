package altay

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 青格里QinggilBarony struct {
	feud.BaseBarony
}

var BQinggil青格里 feud.Barony = &青格里QinggilBarony{}

func init() {
	f := BQinggil青格里.(*青格里QinggilBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "qinggil",
		TitleName: "青格里",
		TitleCode: "b_qinggil",
	}
}
