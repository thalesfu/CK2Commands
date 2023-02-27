package karmanta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 优那拘胝UnakotiBarony struct {
	feud.BaseBarony
}

var BUnakoti优那拘胝 feud.Barony = &优那拘胝UnakotiBarony{}

func init() {
    f := BUnakoti优那拘胝.(*优那拘胝UnakotiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "unakoti",
		TitleName: "优那拘胝",
		TitleCode: "b_unakoti",
	}
}
