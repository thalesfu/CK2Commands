package senlis

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SenlisCounty interface {
	feud.County
	BChantilly尚蒂伊() feud.Barony
	BCompiegne贡比涅() feud.Barony
	BCreil克雷伊() feud.Barony
	BCrepy克雷皮() feud.Barony
	BRoissy鲁瓦西() feud.Barony
	BSenlis桑利斯() feud.Barony
}

type 桑利斯SenlisCounty struct {
	feud.BaseCounty
	尚蒂伊Chantilly feud.Barony
	贡比涅Compiegne feud.Barony
	克雷伊Creil     feud.Barony
	克雷皮Crepy     feud.Barony
	鲁瓦西Roissy    feud.Barony
	桑利斯Senlis    feud.Barony
}

func (c *桑利斯SenlisCounty) BChantilly尚蒂伊() feud.Barony {
	return c.尚蒂伊Chantilly
}

func (c *桑利斯SenlisCounty) BCompiegne贡比涅() feud.Barony {
	return c.贡比涅Compiegne
}

func (c *桑利斯SenlisCounty) BCreil克雷伊() feud.Barony {
	return c.克雷伊Creil
}

func (c *桑利斯SenlisCounty) BCrepy克雷皮() feud.Barony {
	return c.克雷皮Crepy
}

func (c *桑利斯SenlisCounty) BRoissy鲁瓦西() feud.Barony {
	return c.鲁瓦西Roissy
}

func (c *桑利斯SenlisCounty) BSenlis桑利斯() feud.Barony {
	return c.桑利斯Senlis
}

var CSenlis桑利斯 SenlisCounty = &桑利斯SenlisCounty{}

func init() {
	f := CSenlis桑利斯.(*桑利斯SenlisCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1963",
		Title:     "senlis",
		TitleName: "桑利斯",
		TitleCode: "c_senlis",
		Baronies:  map[string]feud.Barony{},
	}

	f.尚蒂伊Chantilly = BChantilly尚蒂伊
	f.尚蒂伊Chantilly.SetParent(f)

	f.贡比涅Compiegne = BCompiegne贡比涅
	f.贡比涅Compiegne.SetParent(f)

	f.克雷伊Creil = BCreil克雷伊
	f.克雷伊Creil.SetParent(f)

	f.克雷皮Crepy = BCrepy克雷皮
	f.克雷皮Crepy.SetParent(f)

	f.鲁瓦西Roissy = BRoissy鲁瓦西
	f.鲁瓦西Roissy.SetParent(f)

	f.桑利斯Senlis = BSenlis桑利斯
	f.桑利斯Senlis.SetParent(f)

}
