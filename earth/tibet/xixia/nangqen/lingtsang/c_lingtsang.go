package lingtsang

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type LingtsangCounty interface {
	feud.County
	BArizha阿日扎() feud.Barony
	BDenkok邓柯() feud.Barony
	BDerong德荣马() feud.Barony
	BGemeng格孟() feud.Barony
	BMaga麻呷() feud.Barony
	BSerxu色须() feud.Barony
	BWenbo温波() feud.Barony
}

type 灵藏LingtsangCounty struct {
	feud.BaseCounty
	阿日扎Arizha feud.Barony
	邓柯Denkok  feud.Barony
	德荣马Derong feud.Barony
	格孟Gemeng  feud.Barony
	麻呷Maga    feud.Barony
	色须Serxu   feud.Barony
	温波Wenbo   feud.Barony
}

func (c *灵藏LingtsangCounty) BArizha阿日扎() feud.Barony {
	return c.阿日扎Arizha
}

func (c *灵藏LingtsangCounty) BDenkok邓柯() feud.Barony {
	return c.邓柯Denkok
}

func (c *灵藏LingtsangCounty) BDerong德荣马() feud.Barony {
	return c.德荣马Derong
}

func (c *灵藏LingtsangCounty) BGemeng格孟() feud.Barony {
	return c.格孟Gemeng
}

func (c *灵藏LingtsangCounty) BMaga麻呷() feud.Barony {
	return c.麻呷Maga
}

func (c *灵藏LingtsangCounty) BSerxu色须() feud.Barony {
	return c.色须Serxu
}

func (c *灵藏LingtsangCounty) BWenbo温波() feud.Barony {
	return c.温波Wenbo
}

var CLingtsang灵藏 LingtsangCounty = &灵藏LingtsangCounty{}

func init() {
	f := CLingtsang灵藏.(*灵藏LingtsangCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1504",
		Title:     "lingtsang",
		TitleName: "灵藏",
		TitleCode: "c_lingtsang",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿日扎Arizha = BArizha阿日扎
	f.阿日扎Arizha.SetParent(f)

	f.邓柯Denkok = BDenkok邓柯
	f.邓柯Denkok.SetParent(f)

	f.德荣马Derong = BDerong德荣马
	f.德荣马Derong.SetParent(f)

	f.格孟Gemeng = BGemeng格孟
	f.格孟Gemeng.SetParent(f)

	f.麻呷Maga = BMaga麻呷
	f.麻呷Maga.SetParent(f)

	f.色须Serxu = BSerxu色须
	f.色须Serxu.SetParent(f)

	f.温波Wenbo = BWenbo温波
	f.温波Wenbo.SetParent(f)

}
