package ennedi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沙拉法Esh_sharafaBarony struct {
	feud.BaseBarony
}

var BEsh_sharafa沙拉法 feud.Barony = &沙拉法Esh_sharafaBarony{}

func init() {
    f := BEsh_sharafa沙拉法.(*沙拉法Esh_sharafaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "esh_sharafa",
		TitleName: "沙拉法",
		TitleCode: "b_esh_sharafa",
	}
}
