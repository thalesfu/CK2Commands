package kiri

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KiriCounty interface {
	feud.County
	BDakadiala达卡迪亚拉() feud.Barony
	BDiakha迪亚哈() feud.Barony
	BDianguinare迪昂吉纳雷() feud.Barony
	BKassiola卡西奥拉() feud.Barony
	BKita基塔() feud.Barony
	BNioro尼奥罗() feud.Barony
	BSerne塞尔内() feud.Barony
}

type 基里KiriCounty struct {
	feud.BaseCounty
	达卡迪亚拉Dakadiala   feud.Barony
	迪亚哈Diakha        feud.Barony
	迪昂吉纳雷Dianguinare feud.Barony
	卡西奥拉Kassiola     feud.Barony
	基塔Kita           feud.Barony
	尼奥罗Nioro         feud.Barony
	塞尔内Serne         feud.Barony
}

func (c *基里KiriCounty) BDakadiala达卡迪亚拉() feud.Barony {
	return c.达卡迪亚拉Dakadiala
}

func (c *基里KiriCounty) BDiakha迪亚哈() feud.Barony {
	return c.迪亚哈Diakha
}

func (c *基里KiriCounty) BDianguinare迪昂吉纳雷() feud.Barony {
	return c.迪昂吉纳雷Dianguinare
}

func (c *基里KiriCounty) BKassiola卡西奥拉() feud.Barony {
	return c.卡西奥拉Kassiola
}

func (c *基里KiriCounty) BKita基塔() feud.Barony {
	return c.基塔Kita
}

func (c *基里KiriCounty) BNioro尼奥罗() feud.Barony {
	return c.尼奥罗Nioro
}

func (c *基里KiriCounty) BSerne塞尔内() feud.Barony {
	return c.塞尔内Serne
}

var CKiri基里 KiriCounty = &基里KiriCounty{}

func init() {
	f := CKiri基里.(*基里KiriCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1760",
		Title:     "kiri",
		TitleName: "基里",
		TitleCode: "c_kiri",
		Baronies:  map[string]feud.Barony{},
	}

	f.达卡迪亚拉Dakadiala = BDakadiala达卡迪亚拉
	f.达卡迪亚拉Dakadiala.SetParent(f)

	f.迪亚哈Diakha = BDiakha迪亚哈
	f.迪亚哈Diakha.SetParent(f)

	f.迪昂吉纳雷Dianguinare = BDianguinare迪昂吉纳雷
	f.迪昂吉纳雷Dianguinare.SetParent(f)

	f.卡西奥拉Kassiola = BKassiola卡西奥拉
	f.卡西奥拉Kassiola.SetParent(f)

	f.基塔Kita = BKita基塔
	f.基塔Kita.SetParent(f)

	f.尼奥罗Nioro = BNioro尼奥罗
	f.尼奥罗Nioro.SetParent(f)

	f.塞尔内Serne = BSerne塞尔内
	f.塞尔内Serne.SetParent(f)

}
