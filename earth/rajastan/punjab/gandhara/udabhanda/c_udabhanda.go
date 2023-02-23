package udabhanda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type UdabhandaCounty interface {
	feud.County
	BAttak阿塔克() feud.Barony
	BCharsada濮达() feud.Barony
	BMankiala门吉亚罗() feud.Barony
	BOddiyana乌苌() feud.Barony
	BOhind奥希德() feud.Barony
	BRawal拉瓦尔() feud.Barony
	BUdabhanda乌铎迦汉荼() feud.Barony
}

type 乌铎迦汉荼UdabhandaCounty struct {
	feud.BaseCounty
	阿塔克Attak       feud.Barony
	濮达Charsada     feud.Barony
	门吉亚罗Mankiala   feud.Barony
	乌苌Oddiyana     feud.Barony
	奥希德Ohind       feud.Barony
	拉瓦尔Rawal       feud.Barony
	乌铎迦汉荼Udabhanda feud.Barony
}

func (c *乌铎迦汉荼UdabhandaCounty) BAttak阿塔克() feud.Barony {
	return c.阿塔克Attak
}

func (c *乌铎迦汉荼UdabhandaCounty) BCharsada濮达() feud.Barony {
	return c.濮达Charsada
}

func (c *乌铎迦汉荼UdabhandaCounty) BMankiala门吉亚罗() feud.Barony {
	return c.门吉亚罗Mankiala
}

func (c *乌铎迦汉荼UdabhandaCounty) BOddiyana乌苌() feud.Barony {
	return c.乌苌Oddiyana
}

func (c *乌铎迦汉荼UdabhandaCounty) BOhind奥希德() feud.Barony {
	return c.奥希德Ohind
}

func (c *乌铎迦汉荼UdabhandaCounty) BRawal拉瓦尔() feud.Barony {
	return c.拉瓦尔Rawal
}

func (c *乌铎迦汉荼UdabhandaCounty) BUdabhanda乌铎迦汉荼() feud.Barony {
	return c.乌铎迦汉荼Udabhanda
}

var CUdabhanda乌铎迦汉荼 UdabhandaCounty = &乌铎迦汉荼UdabhandaCounty{}

func init() {
	f := CUdabhanda乌铎迦汉荼.(*乌铎迦汉荼UdabhandaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1180",
		Title:     "udabhanda",
		TitleName: "乌铎迦汉荼",
		TitleCode: "c_udabhanda",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿塔克Attak = BAttak阿塔克
	f.阿塔克Attak.SetParent(f)

	f.濮达Charsada = BCharsada濮达
	f.濮达Charsada.SetParent(f)

	f.门吉亚罗Mankiala = BMankiala门吉亚罗
	f.门吉亚罗Mankiala.SetParent(f)

	f.乌苌Oddiyana = BOddiyana乌苌
	f.乌苌Oddiyana.SetParent(f)

	f.奥希德Ohind = BOhind奥希德
	f.奥希德Ohind.SetParent(f)

	f.拉瓦尔Rawal = BRawal拉瓦尔
	f.拉瓦尔Rawal.SetParent(f)

	f.乌铎迦汉荼Udabhanda = BUdabhanda乌铎迦汉荼
	f.乌铎迦汉荼Udabhanda.SetParent(f)

}
