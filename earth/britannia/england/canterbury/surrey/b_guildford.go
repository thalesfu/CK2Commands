package surrey

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 吉尔福德GuildfordBarony struct {
	feud.BaseBarony
}

var BGuildford吉尔福德 feud.Barony = &吉尔福德GuildfordBarony{}

func init() {
	f := BGuildford吉尔福德.(*吉尔福德GuildfordBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "guildford",
		TitleName: "吉尔福德",
		TitleCode: "b_guildford",
	}
}
