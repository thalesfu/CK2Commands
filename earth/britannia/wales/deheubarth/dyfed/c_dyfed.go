package dyfed

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type DyfedCounty interface {
	feud.County
	BHaverford哈弗福德() feud.Barony
	BLlanstadwell兰斯塔德韦尔() feud.Barony
	BNarberth纳伯斯() feud.Barony
	BPembroke彭布罗克() feud.Barony
	BPenycwm佩尼库姆() feud.Barony
	BTenby滕比() feud.Barony
}

type 德维得DyfedCounty struct {
	feud.BaseCounty
	哈弗福德Haverford      feud.Barony
	兰斯塔德韦尔Llanstadwell feud.Barony
	纳伯斯Narberth        feud.Barony
	彭布罗克Pembroke       feud.Barony
	佩尼库姆Penycwm        feud.Barony
	滕比Tenby            feud.Barony
}

func (c *德维得DyfedCounty) BHaverford哈弗福德() feud.Barony {
	return c.哈弗福德Haverford
}

func (c *德维得DyfedCounty) BLlanstadwell兰斯塔德韦尔() feud.Barony {
	return c.兰斯塔德韦尔Llanstadwell
}

func (c *德维得DyfedCounty) BNarberth纳伯斯() feud.Barony {
	return c.纳伯斯Narberth
}

func (c *德维得DyfedCounty) BPembroke彭布罗克() feud.Barony {
	return c.彭布罗克Pembroke
}

func (c *德维得DyfedCounty) BPenycwm佩尼库姆() feud.Barony {
	return c.佩尼库姆Penycwm
}

func (c *德维得DyfedCounty) BTenby滕比() feud.Barony {
	return c.滕比Tenby
}

var CDyfed德维得 DyfedCounty = &德维得DyfedCounty{}

func init() {
	f := CDyfed德维得.(*德维得DyfedCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "18",
		Title:     "dyfed",
		TitleName: "德维得",
		TitleCode: "c_dyfed",
		Baronies:  map[string]feud.Barony{},
	}

	f.哈弗福德Haverford = BHaverford哈弗福德
	f.哈弗福德Haverford.SetParent(f)

	f.兰斯塔德韦尔Llanstadwell = BLlanstadwell兰斯塔德韦尔
	f.兰斯塔德韦尔Llanstadwell.SetParent(f)

	f.纳伯斯Narberth = BNarberth纳伯斯
	f.纳伯斯Narberth.SetParent(f)

	f.彭布罗克Pembroke = BPembroke彭布罗克
	f.彭布罗克Pembroke.SetParent(f)

	f.佩尼库姆Penycwm = BPenycwm佩尼库姆
	f.佩尼库姆Penycwm.SetParent(f)

	f.滕比Tenby = BTenby滕比
	f.滕比Tenby.SetParent(f)

}
