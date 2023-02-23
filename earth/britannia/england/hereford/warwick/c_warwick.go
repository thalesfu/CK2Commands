package warwick

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type WarwickCounty interface {
	feud.County
	BCoventry考文垂() feud.Barony
	BDudley达德利() feud.Barony
	BKenilworth凯尼尔沃思() feud.Barony
	BLichfield利奇菲尔德() feud.Barony
	BStafford斯塔福德() feud.Barony
	BTamworth塔姆沃思() feud.Barony
	BTutbury塔特伯里() feud.Barony
	BWarwick沃里克() feud.Barony
}

type 沃里克WarwickCounty struct {
	feud.BaseCounty
	考文垂Coventry     feud.Barony
	达德利Dudley       feud.Barony
	凯尼尔沃思Kenilworth feud.Barony
	利奇菲尔德Lichfield  feud.Barony
	斯塔福德Stafford    feud.Barony
	塔姆沃思Tamworth    feud.Barony
	塔特伯里Tutbury     feud.Barony
	沃里克Warwick      feud.Barony
}

func (c *沃里克WarwickCounty) BCoventry考文垂() feud.Barony {
	return c.考文垂Coventry
}

func (c *沃里克WarwickCounty) BDudley达德利() feud.Barony {
	return c.达德利Dudley
}

func (c *沃里克WarwickCounty) BKenilworth凯尼尔沃思() feud.Barony {
	return c.凯尼尔沃思Kenilworth
}

func (c *沃里克WarwickCounty) BLichfield利奇菲尔德() feud.Barony {
	return c.利奇菲尔德Lichfield
}

func (c *沃里克WarwickCounty) BStafford斯塔福德() feud.Barony {
	return c.斯塔福德Stafford
}

func (c *沃里克WarwickCounty) BTamworth塔姆沃思() feud.Barony {
	return c.塔姆沃思Tamworth
}

func (c *沃里克WarwickCounty) BTutbury塔特伯里() feud.Barony {
	return c.塔特伯里Tutbury
}

func (c *沃里克WarwickCounty) BWarwick沃里克() feud.Barony {
	return c.沃里克Warwick
}

var CWarwick沃里克 WarwickCounty = &沃里克WarwickCounty{}

func init() {
	f := CWarwick沃里克.(*沃里克WarwickCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "67",
		Title:     "warwick",
		TitleName: "沃里克",
		TitleCode: "c_warwick",
		Baronies:  map[string]feud.Barony{},
	}

	f.考文垂Coventry = BCoventry考文垂
	f.考文垂Coventry.SetParent(f)

	f.达德利Dudley = BDudley达德利
	f.达德利Dudley.SetParent(f)

	f.凯尼尔沃思Kenilworth = BKenilworth凯尼尔沃思
	f.凯尼尔沃思Kenilworth.SetParent(f)

	f.利奇菲尔德Lichfield = BLichfield利奇菲尔德
	f.利奇菲尔德Lichfield.SetParent(f)

	f.斯塔福德Stafford = BStafford斯塔福德
	f.斯塔福德Stafford.SetParent(f)

	f.塔姆沃思Tamworth = BTamworth塔姆沃思
	f.塔姆沃思Tamworth.SetParent(f)

	f.塔特伯里Tutbury = BTutbury塔特伯里
	f.塔特伯里Tutbury.SetParent(f)

	f.沃里克Warwick = BWarwick沃里克
	f.沃里克Warwick.SetParent(f)

}
