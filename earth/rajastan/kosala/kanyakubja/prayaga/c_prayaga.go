package prayaga

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type PrayagaCounty interface {
	feud.County
	BAlopibagh阿卢毗婆迦() feud.Barony
	BChowk朝克() feud.Barony
	BDaraganj陀罗犍阇() feud.Barony
	BDariyabad达里耶巴德() feud.Barony
	BKara迦罗() feud.Barony
	BKausambi憍赏弥() feud.Barony
	BPrayaga钵罗耶伽() feud.Barony
}

type 钵罗耶伽PrayagaCounty struct {
	feud.BaseCounty
	阿卢毗婆迦Alopibagh feud.Barony
	朝克Chowk        feud.Barony
	陀罗犍阇Daraganj   feud.Barony
	达里耶巴德Dariyabad feud.Barony
	迦罗Kara         feud.Barony
	憍赏弥Kausambi    feud.Barony
	钵罗耶伽Prayaga    feud.Barony
}

func (c *钵罗耶伽PrayagaCounty) BAlopibagh阿卢毗婆迦() feud.Barony {
	return c.阿卢毗婆迦Alopibagh
}

func (c *钵罗耶伽PrayagaCounty) BChowk朝克() feud.Barony {
	return c.朝克Chowk
}

func (c *钵罗耶伽PrayagaCounty) BDaraganj陀罗犍阇() feud.Barony {
	return c.陀罗犍阇Daraganj
}

func (c *钵罗耶伽PrayagaCounty) BDariyabad达里耶巴德() feud.Barony {
	return c.达里耶巴德Dariyabad
}

func (c *钵罗耶伽PrayagaCounty) BKara迦罗() feud.Barony {
	return c.迦罗Kara
}

func (c *钵罗耶伽PrayagaCounty) BKausambi憍赏弥() feud.Barony {
	return c.憍赏弥Kausambi
}

func (c *钵罗耶伽PrayagaCounty) BPrayaga钵罗耶伽() feud.Barony {
	return c.钵罗耶伽Prayaga
}

var CPrayaga钵罗耶伽 PrayagaCounty = &钵罗耶伽PrayagaCounty{}

func init() {
	f := CPrayaga钵罗耶伽.(*钵罗耶伽PrayagaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1328",
		Title:     "prayaga",
		TitleName: "钵罗耶伽",
		TitleCode: "c_prayaga",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿卢毗婆迦Alopibagh = BAlopibagh阿卢毗婆迦
	f.阿卢毗婆迦Alopibagh.SetParent(f)

	f.朝克Chowk = BChowk朝克
	f.朝克Chowk.SetParent(f)

	f.陀罗犍阇Daraganj = BDaraganj陀罗犍阇
	f.陀罗犍阇Daraganj.SetParent(f)

	f.达里耶巴德Dariyabad = BDariyabad达里耶巴德
	f.达里耶巴德Dariyabad.SetParent(f)

	f.迦罗Kara = BKara迦罗
	f.迦罗Kara.SetParent(f)

	f.憍赏弥Kausambi = BKausambi憍赏弥
	f.憍赏弥Kausambi.SetParent(f)

	f.钵罗耶伽Prayaga = BPrayaga钵罗耶伽
	f.钵罗耶伽Prayaga.SetParent(f)

}
