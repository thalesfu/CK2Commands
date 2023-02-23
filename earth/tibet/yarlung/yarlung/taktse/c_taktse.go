package taktse

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TaktseCounty interface {
	feud.County
	BComai措美() feud.Barony
	BKazhu卡珠() feud.Barony
	BLayu拉玉() feud.Barony
	BQoi曲沟() feud.Barony
	BQonggyai琼结() feud.Barony
	BTaktse达孜() feud.Barony
	BZhegu哲古() feud.Barony
}

type 达孜TaktseCounty struct {
	feud.BaseCounty
	措美Comai    feud.Barony
	卡珠Kazhu    feud.Barony
	拉玉Layu     feud.Barony
	曲沟Qoi      feud.Barony
	琼结Qonggyai feud.Barony
	达孜Taktse   feud.Barony
	哲古Zhegu    feud.Barony
}

func (c *达孜TaktseCounty) BComai措美() feud.Barony {
	return c.措美Comai
}

func (c *达孜TaktseCounty) BKazhu卡珠() feud.Barony {
	return c.卡珠Kazhu
}

func (c *达孜TaktseCounty) BLayu拉玉() feud.Barony {
	return c.拉玉Layu
}

func (c *达孜TaktseCounty) BQoi曲沟() feud.Barony {
	return c.曲沟Qoi
}

func (c *达孜TaktseCounty) BQonggyai琼结() feud.Barony {
	return c.琼结Qonggyai
}

func (c *达孜TaktseCounty) BTaktse达孜() feud.Barony {
	return c.达孜Taktse
}

func (c *达孜TaktseCounty) BZhegu哲古() feud.Barony {
	return c.哲古Zhegu
}

var CTaktse达孜 TaktseCounty = &达孜TaktseCounty{}

func init() {
	f := CTaktse达孜.(*达孜TaktseCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1496",
		Title:     "taktse",
		TitleName: "达孜",
		TitleCode: "c_taktse",
		Baronies:  map[string]feud.Barony{},
	}

	f.措美Comai = BComai措美
	f.措美Comai.SetParent(f)

	f.卡珠Kazhu = BKazhu卡珠
	f.卡珠Kazhu.SetParent(f)

	f.拉玉Layu = BLayu拉玉
	f.拉玉Layu.SetParent(f)

	f.曲沟Qoi = BQoi曲沟
	f.曲沟Qoi.SetParent(f)

	f.琼结Qonggyai = BQonggyai琼结
	f.琼结Qonggyai.SetParent(f)

	f.达孜Taktse = BTaktse达孜
	f.达孜Taktse.SetParent(f)

	f.哲古Zhegu = BZhegu哲古
	f.哲古Zhegu.SetParent(f)

}
