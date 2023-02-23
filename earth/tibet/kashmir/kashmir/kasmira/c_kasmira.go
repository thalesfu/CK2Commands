package kasmira

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KasmiraCounty interface {
	feud.County
	BAmaresvara阿摩丽湿伐罗() feud.Barony
	BBaghsar婆吉萨尔() feud.Barony
	BLoharakota卢诃罗拘吒() feud.Barony
	BPadmapura钵特摩补罗() feud.Barony
	BParihasapura波利诃娑补罗() feud.Barony
	BShankaracharya商羯罗阿遮利耶() feud.Barony
	BSrinagara室利那伽罗() feud.Barony
}

type 迦湿弥罗KasmiraCounty struct {
	feud.BaseCounty
	阿摩丽湿伐罗Amaresvara      feud.Barony
	婆吉萨尔Baghsar           feud.Barony
	卢诃罗拘吒Loharakota       feud.Barony
	钵特摩补罗Padmapura        feud.Barony
	波利诃娑补罗Parihasapura    feud.Barony
	商羯罗阿遮利耶Shankaracharya feud.Barony
	室利那伽罗Srinagara        feud.Barony
}

func (c *迦湿弥罗KasmiraCounty) BAmaresvara阿摩丽湿伐罗() feud.Barony {
	return c.阿摩丽湿伐罗Amaresvara
}

func (c *迦湿弥罗KasmiraCounty) BBaghsar婆吉萨尔() feud.Barony {
	return c.婆吉萨尔Baghsar
}

func (c *迦湿弥罗KasmiraCounty) BLoharakota卢诃罗拘吒() feud.Barony {
	return c.卢诃罗拘吒Loharakota
}

func (c *迦湿弥罗KasmiraCounty) BPadmapura钵特摩补罗() feud.Barony {
	return c.钵特摩补罗Padmapura
}

func (c *迦湿弥罗KasmiraCounty) BParihasapura波利诃娑补罗() feud.Barony {
	return c.波利诃娑补罗Parihasapura
}

func (c *迦湿弥罗KasmiraCounty) BShankaracharya商羯罗阿遮利耶() feud.Barony {
	return c.商羯罗阿遮利耶Shankaracharya
}

func (c *迦湿弥罗KasmiraCounty) BSrinagara室利那伽罗() feud.Barony {
	return c.室利那伽罗Srinagara
}

var CKasmira迦湿弥罗 KasmiraCounty = &迦湿弥罗KasmiraCounty{}

func init() {
	f := CKasmira迦湿弥罗.(*迦湿弥罗KasmiraCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1161",
		Title:     "kasmira",
		TitleName: "迦湿弥罗",
		TitleCode: "c_kasmira",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿摩丽湿伐罗Amaresvara = BAmaresvara阿摩丽湿伐罗
	f.阿摩丽湿伐罗Amaresvara.SetParent(f)

	f.婆吉萨尔Baghsar = BBaghsar婆吉萨尔
	f.婆吉萨尔Baghsar.SetParent(f)

	f.卢诃罗拘吒Loharakota = BLoharakota卢诃罗拘吒
	f.卢诃罗拘吒Loharakota.SetParent(f)

	f.钵特摩补罗Padmapura = BPadmapura钵特摩补罗
	f.钵特摩补罗Padmapura.SetParent(f)

	f.波利诃娑补罗Parihasapura = BParihasapura波利诃娑补罗
	f.波利诃娑补罗Parihasapura.SetParent(f)

	f.商羯罗阿遮利耶Shankaracharya = BShankaracharya商羯罗阿遮利耶
	f.商羯罗阿遮利耶Shankaracharya.SetParent(f)

	f.室利那伽罗Srinagara = BSrinagara室利那伽罗
	f.室利那伽罗Srinagara.SetParent(f)

}
