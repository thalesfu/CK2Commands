package tara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TaraCounty interface {
	feud.County
	BChernovka切尔诺夫卡() feud.Barony
	BKarbyza卡尔贝扎() feud.Barony
	BKeyzes克伊泽斯() feud.Barony
	BNizovoye尼佐沃耶() feud.Barony
	BOkunevo奥库涅沃() feud.Barony
	BTara塔拉() feud.Barony
	BZalivino扎利维诺() feud.Barony
}

type 塔拉TaraCounty struct {
	feud.BaseCounty
	切尔诺夫卡Chernovka feud.Barony
	卡尔贝扎Karbyza    feud.Barony
	克伊泽斯Keyzes     feud.Barony
	尼佐沃耶Nizovoye   feud.Barony
	奥库涅沃Okunevo    feud.Barony
	塔拉Tara         feud.Barony
	扎利维诺Zalivino   feud.Barony
}

func (c *塔拉TaraCounty) BChernovka切尔诺夫卡() feud.Barony {
	return c.切尔诺夫卡Chernovka
}

func (c *塔拉TaraCounty) BKarbyza卡尔贝扎() feud.Barony {
	return c.卡尔贝扎Karbyza
}

func (c *塔拉TaraCounty) BKeyzes克伊泽斯() feud.Barony {
	return c.克伊泽斯Keyzes
}

func (c *塔拉TaraCounty) BNizovoye尼佐沃耶() feud.Barony {
	return c.尼佐沃耶Nizovoye
}

func (c *塔拉TaraCounty) BOkunevo奥库涅沃() feud.Barony {
	return c.奥库涅沃Okunevo
}

func (c *塔拉TaraCounty) BTara塔拉() feud.Barony {
	return c.塔拉Tara
}

func (c *塔拉TaraCounty) BZalivino扎利维诺() feud.Barony {
	return c.扎利维诺Zalivino
}

var CTara塔拉 TaraCounty = &塔拉TaraCounty{}

func init() {
	f := CTara塔拉.(*塔拉TaraCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1862",
		Title:     "tara",
		TitleName: "塔拉",
		TitleCode: "c_tara",
		Baronies:  map[string]feud.Barony{},
	}

	f.切尔诺夫卡Chernovka = BChernovka切尔诺夫卡
	f.切尔诺夫卡Chernovka.SetParent(f)

	f.卡尔贝扎Karbyza = BKarbyza卡尔贝扎
	f.卡尔贝扎Karbyza.SetParent(f)

	f.克伊泽斯Keyzes = BKeyzes克伊泽斯
	f.克伊泽斯Keyzes.SetParent(f)

	f.尼佐沃耶Nizovoye = BNizovoye尼佐沃耶
	f.尼佐沃耶Nizovoye.SetParent(f)

	f.奥库涅沃Okunevo = BOkunevo奥库涅沃
	f.奥库涅沃Okunevo.SetParent(f)

	f.塔拉Tara = BTara塔拉
	f.塔拉Tara.SetParent(f)

	f.扎利维诺Zalivino = BZalivino扎利维诺
	f.扎利维诺Zalivino.SetParent(f)

}
