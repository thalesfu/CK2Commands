package pressburg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type PressburgCounty interface {
	feud.County
	BBazin鲍津() feud.Barony
	BDunaszerdahely杜瑙塞尔道海伊() feud.Barony
	BGalanta高兰陶() feud.Barony
	BModor莫多尔() feud.Barony
	BNagyszombat瑙吉松鲍特() feud.Barony
	BPressburg普雷斯堡() feud.Barony
	BSomorja绍莫尔尧() feud.Barony
	BSzentgyorgy圣捷尔吉() feud.Barony
}

type 普雷斯堡PressburgCounty struct {
	feud.BaseCounty
	鲍津Bazin               feud.Barony
	杜瑙塞尔道海伊Dunaszerdahely feud.Barony
	高兰陶Galanta            feud.Barony
	莫多尔Modor              feud.Barony
	瑙吉松鲍特Nagyszombat      feud.Barony
	普雷斯堡Pressburg         feud.Barony
	绍莫尔尧Somorja           feud.Barony
	圣捷尔吉Szentgyorgy       feud.Barony
}

func (c *普雷斯堡PressburgCounty) BBazin鲍津() feud.Barony {
	return c.鲍津Bazin
}

func (c *普雷斯堡PressburgCounty) BDunaszerdahely杜瑙塞尔道海伊() feud.Barony {
	return c.杜瑙塞尔道海伊Dunaszerdahely
}

func (c *普雷斯堡PressburgCounty) BGalanta高兰陶() feud.Barony {
	return c.高兰陶Galanta
}

func (c *普雷斯堡PressburgCounty) BModor莫多尔() feud.Barony {
	return c.莫多尔Modor
}

func (c *普雷斯堡PressburgCounty) BNagyszombat瑙吉松鲍特() feud.Barony {
	return c.瑙吉松鲍特Nagyszombat
}

func (c *普雷斯堡PressburgCounty) BPressburg普雷斯堡() feud.Barony {
	return c.普雷斯堡Pressburg
}

func (c *普雷斯堡PressburgCounty) BSomorja绍莫尔尧() feud.Barony {
	return c.绍莫尔尧Somorja
}

func (c *普雷斯堡PressburgCounty) BSzentgyorgy圣捷尔吉() feud.Barony {
	return c.圣捷尔吉Szentgyorgy
}

var CPressburg普雷斯堡 PressburgCounty = &普雷斯堡PressburgCounty{}

func init() {
	f := CPressburg普雷斯堡.(*普雷斯堡PressburgCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "445",
		Title:     "pressburg",
		TitleName: "普雷斯堡",
		TitleCode: "c_pressburg",
		Baronies:  map[string]feud.Barony{},
	}

	f.鲍津Bazin = BBazin鲍津
	f.鲍津Bazin.SetParent(f)

	f.杜瑙塞尔道海伊Dunaszerdahely = BDunaszerdahely杜瑙塞尔道海伊
	f.杜瑙塞尔道海伊Dunaszerdahely.SetParent(f)

	f.高兰陶Galanta = BGalanta高兰陶
	f.高兰陶Galanta.SetParent(f)

	f.莫多尔Modor = BModor莫多尔
	f.莫多尔Modor.SetParent(f)

	f.瑙吉松鲍特Nagyszombat = BNagyszombat瑙吉松鲍特
	f.瑙吉松鲍特Nagyszombat.SetParent(f)

	f.普雷斯堡Pressburg = BPressburg普雷斯堡
	f.普雷斯堡Pressburg.SetParent(f)

	f.绍莫尔尧Somorja = BSomorja绍莫尔尧
	f.绍莫尔尧Somorja.SetParent(f)

	f.圣捷尔吉Szentgyorgy = BSzentgyorgy圣捷尔吉
	f.圣捷尔吉Szentgyorgy.SetParent(f)

}
