package limbuwan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊拉姆Ilam_limbuwanBarony struct {
	feud.BaseBarony
}

var BIlam_limbuwan伊拉姆 feud.Barony = &伊拉姆Ilam_limbuwanBarony{}

func init() {
    f := BIlam_limbuwan伊拉姆.(*伊拉姆Ilam_limbuwanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ilam_limbuwan",
		TitleName: "伊拉姆",
		TitleCode: "b_ilam_limbuwan",
	}
}
