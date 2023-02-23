package westmorland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type WestmorlandCounty interface {
	feud.County
	BAppleby阿普尔比() feud.Barony
	BBrough布拉夫() feud.Barony
	BBrougham布鲁厄姆() feud.Barony
	BCartmel卡特梅尔() feud.Barony
	BKendal肯德尔() feud.Barony
	BKirkby柯比() feud.Barony
	BLowther劳瑟() feud.Barony
	BShap沙普() feud.Barony
}

type 威斯特摩兰WestmorlandCounty struct {
	feud.BaseCounty
	阿普尔比Appleby  feud.Barony
	布拉夫Brough    feud.Barony
	布鲁厄姆Brougham feud.Barony
	卡特梅尔Cartmel  feud.Barony
	肯德尔Kendal    feud.Barony
	柯比Kirkby     feud.Barony
	劳瑟Lowther    feud.Barony
	沙普Shap       feud.Barony
}

func (c *威斯特摩兰WestmorlandCounty) BAppleby阿普尔比() feud.Barony {
	return c.阿普尔比Appleby
}

func (c *威斯特摩兰WestmorlandCounty) BBrough布拉夫() feud.Barony {
	return c.布拉夫Brough
}

func (c *威斯特摩兰WestmorlandCounty) BBrougham布鲁厄姆() feud.Barony {
	return c.布鲁厄姆Brougham
}

func (c *威斯特摩兰WestmorlandCounty) BCartmel卡特梅尔() feud.Barony {
	return c.卡特梅尔Cartmel
}

func (c *威斯特摩兰WestmorlandCounty) BKendal肯德尔() feud.Barony {
	return c.肯德尔Kendal
}

func (c *威斯特摩兰WestmorlandCounty) BKirkby柯比() feud.Barony {
	return c.柯比Kirkby
}

func (c *威斯特摩兰WestmorlandCounty) BLowther劳瑟() feud.Barony {
	return c.劳瑟Lowther
}

func (c *威斯特摩兰WestmorlandCounty) BShap沙普() feud.Barony {
	return c.沙普Shap
}

var CWestmorland威斯特摩兰 WestmorlandCounty = &威斯特摩兰WestmorlandCounty{}

func init() {
	f := CWestmorland威斯特摩兰.(*威斯特摩兰WestmorlandCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "55",
		Title:     "westmorland",
		TitleName: "威斯特摩兰",
		TitleCode: "c_westmorland",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿普尔比Appleby = BAppleby阿普尔比
	f.阿普尔比Appleby.SetParent(f)

	f.布拉夫Brough = BBrough布拉夫
	f.布拉夫Brough.SetParent(f)

	f.布鲁厄姆Brougham = BBrougham布鲁厄姆
	f.布鲁厄姆Brougham.SetParent(f)

	f.卡特梅尔Cartmel = BCartmel卡特梅尔
	f.卡特梅尔Cartmel.SetParent(f)

	f.肯德尔Kendal = BKendal肯德尔
	f.肯德尔Kendal.SetParent(f)

	f.柯比Kirkby = BKirkby柯比
	f.柯比Kirkby.SetParent(f)

	f.劳瑟Lowther = BLowther劳瑟
	f.劳瑟Lowther.SetParent(f)

	f.沙普Shap = BShap沙普
	f.沙普Shap.SetParent(f)

}
