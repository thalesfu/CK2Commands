package eilat

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type EilatCounty interface {
	feud.County
	BEilat埃拉特() feud.Barony
	BEliot以禄() feud.Barony
	BKetura基土拉() feud.Barony
	BLotan罗坍() feud.Barony
	BSapir萨皮尔() feud.Barony
	BTzofar琐法() feud.Barony
	BUrim乌里姆() feud.Barony
	BYahel雅赫拉() feud.Barony
}

type 埃拉特EilatCounty struct {
	feud.BaseCounty
	埃拉特Eilat  feud.Barony
	以禄Eliot   feud.Barony
	基土拉Ketura feud.Barony
	罗坍Lotan   feud.Barony
	萨皮尔Sapir  feud.Barony
	琐法Tzofar  feud.Barony
	乌里姆Urim   feud.Barony
	雅赫拉Yahel  feud.Barony
}

func (c *埃拉特EilatCounty) BEilat埃拉特() feud.Barony {
	return c.埃拉特Eilat
}

func (c *埃拉特EilatCounty) BEliot以禄() feud.Barony {
	return c.以禄Eliot
}

func (c *埃拉特EilatCounty) BKetura基土拉() feud.Barony {
	return c.基土拉Ketura
}

func (c *埃拉特EilatCounty) BLotan罗坍() feud.Barony {
	return c.罗坍Lotan
}

func (c *埃拉特EilatCounty) BSapir萨皮尔() feud.Barony {
	return c.萨皮尔Sapir
}

func (c *埃拉特EilatCounty) BTzofar琐法() feud.Barony {
	return c.琐法Tzofar
}

func (c *埃拉特EilatCounty) BUrim乌里姆() feud.Barony {
	return c.乌里姆Urim
}

func (c *埃拉特EilatCounty) BYahel雅赫拉() feud.Barony {
	return c.雅赫拉Yahel
}

var CEilat埃拉特 EilatCounty = &埃拉特EilatCounty{}

func init() {
	f := CEilat埃拉特.(*埃拉特EilatCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "786",
		Title:     "eilat",
		TitleName: "埃拉特",
		TitleCode: "c_eilat",
		Baronies:  map[string]feud.Barony{},
	}

	f.埃拉特Eilat = BEilat埃拉特
	f.埃拉特Eilat.SetParent(f)

	f.以禄Eliot = BEliot以禄
	f.以禄Eliot.SetParent(f)

	f.基土拉Ketura = BKetura基土拉
	f.基土拉Ketura.SetParent(f)

	f.罗坍Lotan = BLotan罗坍
	f.罗坍Lotan.SetParent(f)

	f.萨皮尔Sapir = BSapir萨皮尔
	f.萨皮尔Sapir.SetParent(f)

	f.琐法Tzofar = BTzofar琐法
	f.琐法Tzofar.SetParent(f)

	f.乌里姆Urim = BUrim乌里姆
	f.乌里姆Urim.SetParent(f)

	f.雅赫拉Yahel = BYahel雅赫拉
	f.雅赫拉Yahel.SetParent(f)

}
