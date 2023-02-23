package chagai

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ChagaiCounty interface {
	feud.County
	BBochra布什拉() feud.Barony
	BChachoke遮朱计() feud.Barony
	BDad达德() feud.Barony
	BKallarwala卡拉瓦拉() feud.Barony
	BKandar坎达尔() feud.Barony
	BNushki努什基() feud.Barony
	BPeshai贝希() feud.Barony
}

type 查盖ChagaiCounty struct {
	feud.BaseCounty
	布什拉Bochra      feud.Barony
	遮朱计Chachoke    feud.Barony
	达德Dad          feud.Barony
	卡拉瓦拉Kallarwala feud.Barony
	坎达尔Kandar      feud.Barony
	努什基Nushki      feud.Barony
	贝希Peshai       feud.Barony
}

func (c *查盖ChagaiCounty) BBochra布什拉() feud.Barony {
	return c.布什拉Bochra
}

func (c *查盖ChagaiCounty) BChachoke遮朱计() feud.Barony {
	return c.遮朱计Chachoke
}

func (c *查盖ChagaiCounty) BDad达德() feud.Barony {
	return c.达德Dad
}

func (c *查盖ChagaiCounty) BKallarwala卡拉瓦拉() feud.Barony {
	return c.卡拉瓦拉Kallarwala
}

func (c *查盖ChagaiCounty) BKandar坎达尔() feud.Barony {
	return c.坎达尔Kandar
}

func (c *查盖ChagaiCounty) BNushki努什基() feud.Barony {
	return c.努什基Nushki
}

func (c *查盖ChagaiCounty) BPeshai贝希() feud.Barony {
	return c.贝希Peshai
}

var CChagai查盖 ChagaiCounty = &查盖ChagaiCounty{}

func init() {
	f := CChagai查盖.(*查盖ChagaiCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1113",
		Title:     "chagai",
		TitleName: "查盖",
		TitleCode: "c_chagai",
		Baronies:  map[string]feud.Barony{},
	}

	f.布什拉Bochra = BBochra布什拉
	f.布什拉Bochra.SetParent(f)

	f.遮朱计Chachoke = BChachoke遮朱计
	f.遮朱计Chachoke.SetParent(f)

	f.达德Dad = BDad达德
	f.达德Dad.SetParent(f)

	f.卡拉瓦拉Kallarwala = BKallarwala卡拉瓦拉
	f.卡拉瓦拉Kallarwala.SetParent(f)

	f.坎达尔Kandar = BKandar坎达尔
	f.坎达尔Kandar.SetParent(f)

	f.努什基Nushki = BNushki努什基
	f.努什基Nushki.SetParent(f)

	f.贝希Peshai = BPeshai贝希
	f.贝希Peshai.SetParent(f)

}
