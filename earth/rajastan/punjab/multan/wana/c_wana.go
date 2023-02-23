package wana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type WanaCounty interface {
	feud.County
	BBori菩里() feud.Barony
	BJua周阿() feud.Barony
	BJunaghur周拿古() feud.Barony
	BKakshar伽叉刹() feud.Barony
	BKalamboli伽蓝菩梨() feud.Barony
	BMun蒙河() feud.Barony
	BNandulia难菟梨() feud.Barony
}

type 弯那WanaCounty struct {
	feud.BaseCounty
	菩里Bori        feud.Barony
	周阿Jua         feud.Barony
	周拿古Junaghur   feud.Barony
	伽叉刹Kakshar    feud.Barony
	伽蓝菩梨Kalamboli feud.Barony
	蒙河Mun         feud.Barony
	难菟梨Nandulia   feud.Barony
}

func (c *弯那WanaCounty) BBori菩里() feud.Barony {
	return c.菩里Bori
}

func (c *弯那WanaCounty) BJua周阿() feud.Barony {
	return c.周阿Jua
}

func (c *弯那WanaCounty) BJunaghur周拿古() feud.Barony {
	return c.周拿古Junaghur
}

func (c *弯那WanaCounty) BKakshar伽叉刹() feud.Barony {
	return c.伽叉刹Kakshar
}

func (c *弯那WanaCounty) BKalamboli伽蓝菩梨() feud.Barony {
	return c.伽蓝菩梨Kalamboli
}

func (c *弯那WanaCounty) BMun蒙河() feud.Barony {
	return c.蒙河Mun
}

func (c *弯那WanaCounty) BNandulia难菟梨() feud.Barony {
	return c.难菟梨Nandulia
}

var CWana弯那 WanaCounty = &弯那WanaCounty{}

func init() {
	f := CWana弯那.(*弯那WanaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1376",
		Title:     "wana",
		TitleName: "弯那",
		TitleCode: "c_wana",
		Baronies:  map[string]feud.Barony{},
	}

	f.菩里Bori = BBori菩里
	f.菩里Bori.SetParent(f)

	f.周阿Jua = BJua周阿
	f.周阿Jua.SetParent(f)

	f.周拿古Junaghur = BJunaghur周拿古
	f.周拿古Junaghur.SetParent(f)

	f.伽叉刹Kakshar = BKakshar伽叉刹
	f.伽叉刹Kakshar.SetParent(f)

	f.伽蓝菩梨Kalamboli = BKalamboli伽蓝菩梨
	f.伽蓝菩梨Kalamboli.SetParent(f)

	f.蒙河Mun = BMun蒙河
	f.蒙河Mun.SetParent(f)

	f.难菟梨Nandulia = BNandulia难菟梨
	f.难菟梨Nandulia.SetParent(f)

}
