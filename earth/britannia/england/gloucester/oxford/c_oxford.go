package oxford

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type OxfordCounty interface {
	feud.County
	BAbingdon阿宾登() feud.Barony
	BAylesbury艾尔斯伯里() feud.Barony
	BBanbury班伯里() feud.Barony
	BBuckingham白金汉() feud.Barony
	BEynsham恩舍姆() feud.Barony
	BOxford牛津() feud.Barony
	BReading雷丁() feud.Barony
	BWallingford沃灵福德() feud.Barony
}

type 牛津OxfordCounty struct {
	feud.BaseCounty
	阿宾登Abingdon     feud.Barony
	艾尔斯伯里Aylesbury  feud.Barony
	班伯里Banbury      feud.Barony
	白金汉Buckingham   feud.Barony
	恩舍姆Eynsham      feud.Barony
	牛津Oxford        feud.Barony
	雷丁Reading       feud.Barony
	沃灵福德Wallingford feud.Barony
}

func (c *牛津OxfordCounty) BAbingdon阿宾登() feud.Barony {
	return c.阿宾登Abingdon
}

func (c *牛津OxfordCounty) BAylesbury艾尔斯伯里() feud.Barony {
	return c.艾尔斯伯里Aylesbury
}

func (c *牛津OxfordCounty) BBanbury班伯里() feud.Barony {
	return c.班伯里Banbury
}

func (c *牛津OxfordCounty) BBuckingham白金汉() feud.Barony {
	return c.白金汉Buckingham
}

func (c *牛津OxfordCounty) BEynsham恩舍姆() feud.Barony {
	return c.恩舍姆Eynsham
}

func (c *牛津OxfordCounty) BOxford牛津() feud.Barony {
	return c.牛津Oxford
}

func (c *牛津OxfordCounty) BReading雷丁() feud.Barony {
	return c.雷丁Reading
}

func (c *牛津OxfordCounty) BWallingford沃灵福德() feud.Barony {
	return c.沃灵福德Wallingford
}

var COxford牛津 OxfordCounty = &牛津OxfordCounty{}

func init() {
	f := COxford牛津.(*牛津OxfordCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "22",
		Title:     "oxford",
		TitleName: "牛津",
		TitleCode: "c_oxford",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿宾登Abingdon = BAbingdon阿宾登
	f.阿宾登Abingdon.SetParent(f)

	f.艾尔斯伯里Aylesbury = BAylesbury艾尔斯伯里
	f.艾尔斯伯里Aylesbury.SetParent(f)

	f.班伯里Banbury = BBanbury班伯里
	f.班伯里Banbury.SetParent(f)

	f.白金汉Buckingham = BBuckingham白金汉
	f.白金汉Buckingham.SetParent(f)

	f.恩舍姆Eynsham = BEynsham恩舍姆
	f.恩舍姆Eynsham.SetParent(f)

	f.牛津Oxford = BOxford牛津
	f.牛津Oxford.SetParent(f)

	f.雷丁Reading = BReading雷丁
	f.雷丁Reading.SetParent(f)

	f.沃灵福德Wallingford = BWallingford沃灵福德
	f.沃灵福德Wallingford.SetParent(f)

}
