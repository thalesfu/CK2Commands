package bezichy

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BezichyCounty interface {
	feud.County
	BBezichy别日奇() feud.Barony
	BKlyuchevaya克柳切瓦亚() feud.Barony
	BKnyazhikha克尼亚日哈() feud.Barony
	BMaksatikha马克萨季哈() feud.Barony
	BRameshki季杰里诺() feud.Barony
	BTolmachi托尔马奇() feud.Barony
	BTorby托尔贝() feud.Barony
}

type 别日奇BezichyCounty struct {
	feud.BaseCounty
	别日奇Bezichy       feud.Barony
	克柳切瓦亚Klyuchevaya feud.Barony
	克尼亚日哈Knyazhikha  feud.Barony
	马克萨季哈Maksatikha  feud.Barony
	季杰里诺Rameshki     feud.Barony
	托尔马奇Tolmachi     feud.Barony
	托尔贝Torby         feud.Barony
}

func (c *别日奇BezichyCounty) BBezichy别日奇() feud.Barony {
	return c.别日奇Bezichy
}

func (c *别日奇BezichyCounty) BKlyuchevaya克柳切瓦亚() feud.Barony {
	return c.克柳切瓦亚Klyuchevaya
}

func (c *别日奇BezichyCounty) BKnyazhikha克尼亚日哈() feud.Barony {
	return c.克尼亚日哈Knyazhikha
}

func (c *别日奇BezichyCounty) BMaksatikha马克萨季哈() feud.Barony {
	return c.马克萨季哈Maksatikha
}

func (c *别日奇BezichyCounty) BRameshki季杰里诺() feud.Barony {
	return c.季杰里诺Rameshki
}

func (c *别日奇BezichyCounty) BTolmachi托尔马奇() feud.Barony {
	return c.托尔马奇Tolmachi
}

func (c *别日奇BezichyCounty) BTorby托尔贝() feud.Barony {
	return c.托尔贝Torby
}

var CBezichy别日奇 BezichyCounty = &别日奇BezichyCounty{}

func init() {
	f := CBezichy别日奇.(*别日奇BezichyCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1670",
		Title:     "bezichy",
		TitleName: "别日奇",
		TitleCode: "c_bezichy",
		Baronies:  map[string]feud.Barony{},
	}

	f.别日奇Bezichy = BBezichy别日奇
	f.别日奇Bezichy.SetParent(f)

	f.克柳切瓦亚Klyuchevaya = BKlyuchevaya克柳切瓦亚
	f.克柳切瓦亚Klyuchevaya.SetParent(f)

	f.克尼亚日哈Knyazhikha = BKnyazhikha克尼亚日哈
	f.克尼亚日哈Knyazhikha.SetParent(f)

	f.马克萨季哈Maksatikha = BMaksatikha马克萨季哈
	f.马克萨季哈Maksatikha.SetParent(f)

	f.季杰里诺Rameshki = BRameshki季杰里诺
	f.季杰里诺Rameshki.SetParent(f)

	f.托尔马奇Tolmachi = BTolmachi托尔马奇
	f.托尔马奇Tolmachi.SetParent(f)

	f.托尔贝Torby = BTorby托尔贝
	f.托尔贝Torby.SetParent(f)

}
