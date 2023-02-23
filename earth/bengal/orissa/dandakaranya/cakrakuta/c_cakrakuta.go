package cakrakuta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type CakrakutaCounty interface {
	feud.County
	BCakrakuta遮迦罗矩吒() feud.Barony
	BJagdalpur阇迦陀罗城() feud.Barony
	BKanker建艺罗() feud.Barony
	BPudunagram富度那揭罗() feud.Barony
	BPurbbadulki富波婆菟吉() feud.Barony
	BRaigarh罗耶姞利呬() feud.Barony
	BSankarpur商羯罗补罗() feud.Barony
}

type 遮迦罗矩吒CakrakutaCounty struct {
	feud.BaseCounty
	遮迦罗矩吒Cakrakuta   feud.Barony
	阇迦陀罗城Jagdalpur   feud.Barony
	建艺罗Kanker        feud.Barony
	富度那揭罗Pudunagram  feud.Barony
	富波婆菟吉Purbbadulki feud.Barony
	罗耶姞利呬Raigarh     feud.Barony
	商羯罗补罗Sankarpur   feud.Barony
}

func (c *遮迦罗矩吒CakrakutaCounty) BCakrakuta遮迦罗矩吒() feud.Barony {
	return c.遮迦罗矩吒Cakrakuta
}

func (c *遮迦罗矩吒CakrakutaCounty) BJagdalpur阇迦陀罗城() feud.Barony {
	return c.阇迦陀罗城Jagdalpur
}

func (c *遮迦罗矩吒CakrakutaCounty) BKanker建艺罗() feud.Barony {
	return c.建艺罗Kanker
}

func (c *遮迦罗矩吒CakrakutaCounty) BPudunagram富度那揭罗() feud.Barony {
	return c.富度那揭罗Pudunagram
}

func (c *遮迦罗矩吒CakrakutaCounty) BPurbbadulki富波婆菟吉() feud.Barony {
	return c.富波婆菟吉Purbbadulki
}

func (c *遮迦罗矩吒CakrakutaCounty) BRaigarh罗耶姞利呬() feud.Barony {
	return c.罗耶姞利呬Raigarh
}

func (c *遮迦罗矩吒CakrakutaCounty) BSankarpur商羯罗补罗() feud.Barony {
	return c.商羯罗补罗Sankarpur
}

var CCakrakuta遮迦罗矩吒 CakrakutaCounty = &遮迦罗矩吒CakrakutaCounty{}

func init() {
	f := CCakrakuta遮迦罗矩吒.(*遮迦罗矩吒CakrakutaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1227",
		Title:     "cakrakuta",
		TitleName: "遮迦罗矩吒",
		TitleCode: "c_cakrakuta",
		Baronies:  map[string]feud.Barony{},
	}

	f.遮迦罗矩吒Cakrakuta = BCakrakuta遮迦罗矩吒
	f.遮迦罗矩吒Cakrakuta.SetParent(f)

	f.阇迦陀罗城Jagdalpur = BJagdalpur阇迦陀罗城
	f.阇迦陀罗城Jagdalpur.SetParent(f)

	f.建艺罗Kanker = BKanker建艺罗
	f.建艺罗Kanker.SetParent(f)

	f.富度那揭罗Pudunagram = BPudunagram富度那揭罗
	f.富度那揭罗Pudunagram.SetParent(f)

	f.富波婆菟吉Purbbadulki = BPurbbadulki富波婆菟吉
	f.富波婆菟吉Purbbadulki.SetParent(f)

	f.罗耶姞利呬Raigarh = BRaigarh罗耶姞利呬
	f.罗耶姞利呬Raigarh.SetParent(f)

	f.商羯罗补罗Sankarpur = BSankarpur商羯罗补罗
	f.商羯罗补罗Sankarpur.SetParent(f)

}
