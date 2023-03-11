package rylsk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 格卢什科沃GlushkvoBarony struct {
	feud.BaseBarony
}

var BGlushkvo格卢什科沃 feud.Barony = &格卢什科沃GlushkvoBarony{}

func init() {
    f := BGlushkvo格卢什科沃.(*格卢什科沃GlushkvoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "glushkvo",
		TitleName: "格卢什科沃",
		TitleCode: "b_glushkvo",
	}
}
