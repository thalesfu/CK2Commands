package kalat

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KalatCounty interface {
	feud.County
	BBhaminpur婆明补罗() feud.Barony
	BChenga琴加() feud.Barony
	BKalat卡拉特() feud.Barony
	BKatewala伽谛伐罗() feud.Barony
	BNautasal瑙塔萨勒() feud.Barony
	BPalasi帕腊斯() feud.Barony
	BTalq塔拉喀() feud.Barony
}

type 卡拉特KalatCounty struct {
	feud.BaseCounty
	婆明补罗Bhaminpur feud.Barony
	琴加Chenga      feud.Barony
	卡拉特Kalat      feud.Barony
	伽谛伐罗Katewala  feud.Barony
	瑙塔萨勒Nautasal  feud.Barony
	帕腊斯Palasi     feud.Barony
	塔拉喀Talq       feud.Barony
}

func (c *卡拉特KalatCounty) BBhaminpur婆明补罗() feud.Barony {
	return c.婆明补罗Bhaminpur
}

func (c *卡拉特KalatCounty) BChenga琴加() feud.Barony {
	return c.琴加Chenga
}

func (c *卡拉特KalatCounty) BKalat卡拉特() feud.Barony {
	return c.卡拉特Kalat
}

func (c *卡拉特KalatCounty) BKatewala伽谛伐罗() feud.Barony {
	return c.伽谛伐罗Katewala
}

func (c *卡拉特KalatCounty) BNautasal瑙塔萨勒() feud.Barony {
	return c.瑙塔萨勒Nautasal
}

func (c *卡拉特KalatCounty) BPalasi帕腊斯() feud.Barony {
	return c.帕腊斯Palasi
}

func (c *卡拉特KalatCounty) BTalq塔拉喀() feud.Barony {
	return c.塔拉喀Talq
}

var CKalat卡拉特 KalatCounty = &卡拉特KalatCounty{}

func init() {
	f := CKalat卡拉特.(*卡拉特KalatCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1184",
		Title:     "kalat",
		TitleName: "卡拉特",
		TitleCode: "c_kalat",
		Baronies:  map[string]feud.Barony{},
	}

	f.婆明补罗Bhaminpur = BBhaminpur婆明补罗
	f.婆明补罗Bhaminpur.SetParent(f)

	f.琴加Chenga = BChenga琴加
	f.琴加Chenga.SetParent(f)

	f.卡拉特Kalat = BKalat卡拉特
	f.卡拉特Kalat.SetParent(f)

	f.伽谛伐罗Katewala = BKatewala伽谛伐罗
	f.伽谛伐罗Katewala.SetParent(f)

	f.瑙塔萨勒Nautasal = BNautasal瑙塔萨勒
	f.瑙塔萨勒Nautasal.SetParent(f)

	f.帕腊斯Palasi = BPalasi帕腊斯
	f.帕腊斯Palasi.SetParent(f)

	f.塔拉喀Talq = BTalq塔拉喀
	f.塔拉喀Talq.SetParent(f)

}
