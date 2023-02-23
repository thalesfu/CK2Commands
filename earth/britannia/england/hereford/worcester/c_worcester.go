package worcester

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type WorcesterCounty interface {
	feud.County
	BBromsgrove布罗姆斯格罗夫() feud.Barony
	BDroitwich德罗伊特威奇() feud.Barony
	BEvesham伊夫舍姆() feud.Barony
	BKidderminster基德明斯特() feud.Barony
	BLaughern拉恩() feud.Barony
	BMalvern莫尔文() feud.Barony
	BPershore珀肖尔() feud.Barony
	BWorcester伍斯特() feud.Barony
}

type 伍斯特WorcesterCounty struct {
	feud.BaseCounty
	布罗姆斯格罗夫Bromsgrove  feud.Barony
	德罗伊特威奇Droitwich    feud.Barony
	伊夫舍姆Evesham        feud.Barony
	基德明斯特Kidderminster feud.Barony
	拉恩Laughern         feud.Barony
	莫尔文Malvern         feud.Barony
	珀肖尔Pershore        feud.Barony
	伍斯特Worcester       feud.Barony
}

func (c *伍斯特WorcesterCounty) BBromsgrove布罗姆斯格罗夫() feud.Barony {
	return c.布罗姆斯格罗夫Bromsgrove
}

func (c *伍斯特WorcesterCounty) BDroitwich德罗伊特威奇() feud.Barony {
	return c.德罗伊特威奇Droitwich
}

func (c *伍斯特WorcesterCounty) BEvesham伊夫舍姆() feud.Barony {
	return c.伊夫舍姆Evesham
}

func (c *伍斯特WorcesterCounty) BKidderminster基德明斯特() feud.Barony {
	return c.基德明斯特Kidderminster
}

func (c *伍斯特WorcesterCounty) BLaughern拉恩() feud.Barony {
	return c.拉恩Laughern
}

func (c *伍斯特WorcesterCounty) BMalvern莫尔文() feud.Barony {
	return c.莫尔文Malvern
}

func (c *伍斯特WorcesterCounty) BPershore珀肖尔() feud.Barony {
	return c.珀肖尔Pershore
}

func (c *伍斯特WorcesterCounty) BWorcester伍斯特() feud.Barony {
	return c.伍斯特Worcester
}

var CWorcester伍斯特 WorcesterCounty = &伍斯特WorcesterCounty{}

func init() {
	f := CWorcester伍斯特.(*伍斯特WorcesterCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "30",
		Title:     "worcester",
		TitleName: "伍斯特",
		TitleCode: "c_worcester",
		Baronies:  map[string]feud.Barony{},
	}

	f.布罗姆斯格罗夫Bromsgrove = BBromsgrove布罗姆斯格罗夫
	f.布罗姆斯格罗夫Bromsgrove.SetParent(f)

	f.德罗伊特威奇Droitwich = BDroitwich德罗伊特威奇
	f.德罗伊特威奇Droitwich.SetParent(f)

	f.伊夫舍姆Evesham = BEvesham伊夫舍姆
	f.伊夫舍姆Evesham.SetParent(f)

	f.基德明斯特Kidderminster = BKidderminster基德明斯特
	f.基德明斯特Kidderminster.SetParent(f)

	f.拉恩Laughern = BLaughern拉恩
	f.拉恩Laughern.SetParent(f)

	f.莫尔文Malvern = BMalvern莫尔文
	f.莫尔文Malvern.SetParent(f)

	f.珀肖尔Pershore = BPershore珀肖尔
	f.珀肖尔Pershore.SetParent(f)

	f.伍斯特Worcester = BWorcester伍斯特
	f.伍斯特Worcester.SetParent(f)

}
