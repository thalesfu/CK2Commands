package main

import (
	"fmt"
	"github.com/samber/lo"
	"github.com/thalesfu/CK2Commands/family/bi"
	"github.com/thalesfu/CK2Commands/family/chu"
	"github.com/thalesfu/CK2Commands/family/fu"
	"github.com/thalesfu/CK2Commands/family/li"
	"github.com/thalesfu/CK2Commands/family/lin"
	"github.com/thalesfu/CK2Commands/family/lou"
	"github.com/thalesfu/CK2Commands/family/wu"
	"github.com/thalesfu/CK2Commands/family/wuyibu"
	"github.com/thalesfu/CK2Commands/job"
	"github.com/thalesfu/CK2Commands/people"
	"github.com/thalesfu/ck2nebula"
	"github.com/thalesfu/paradoxtools/utils"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

const ck2Folder = "R:\\Thales\\Game\\SteamLibrary\\steamapps\\common\\Crusader Kings II"
const saveFolder = "T:\\OneDrive\\fu.thales@live.com\\OneDrive\\MyDocument\\Paradox Interactive\\Crusader Kings II\\save games"

func main() {

	story, player, err := GetStory(false)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(story.PlayerName)

	coreFamily := map[int]string{
		1000103393: "lou",
		1000103382: "yuan",
		1000103379: "chu",
		1000103374: "wu",
		1000103360: "zhang",
		1000103348: "lin",
		1000103339: "bi",
		1000103336: "yin",
		1051150:    "li",
	}

	people.AutoBuild(ck2nebula.SPACE, player, coreFamily)

	//BuildFriends(player)
}

func BuildFriends(player *ck2nebula.People) {
	fr := player.GetFriends(ck2nebula.SPACE)

	if fr.Ok {
		friends := make([]*ck2nebula.People, 0)
		for _, f := range fr.Data {
			friends = append(friends, f)
		}
		people.BuildHealthAndFertility(ck2nebula.SPACE, friends...)
	}
}

func GetStory(force bool) (*ck2nebula.Story, *ck2nebula.People, error) {

	sr := ck2nebula.GetLatestStory(ck2nebula.SPACE)

	if !sr.Ok {
		return nil, nil, sr.Err
	}

	files, err := os.ReadDir(saveFolder)
	if err != nil {
		return nil, nil, err
	}

	files = lo.Filter(files, func(file os.DirEntry, _ int) bool {
		return !file.IsDir() && strings.HasSuffix(file.Name(), ".ck2")
	})

	sort.Slice(files, func(i, j int) bool {
		xInfo, _ := files[i].Info()
		yInfo, _ := files[j].Info()
		return xInfo.ModTime().Unix() > yInfo.ModTime().Unix()
	})

	filePath := strings.ReplaceAll(filepath.Join(saveFolder, files[0].Name()), "\\", "/")

	if force || !isSameStory(filePath, sr.Data) {
		log.Printf("loading save file \"%s\"", filePath)
		ck2nebula.BuildStory(ck2Folder, filePath)

		sr = ck2nebula.GetLatestStory(ck2nebula.SPACE)

		if !sr.Ok {
			return nil, nil, sr.Err
		}
	}

	pr := sr.Data.GetPlayer(ck2nebula.SPACE)

	if !pr.Ok {
		return nil, nil, pr.Err
	}

	return sr.Data, pr.Data, nil
}

func isSameStory(fp string, story *ck2nebula.Story) bool {

	if fp != story.FilePath {
		return false
	}

	hash, err := utils.GetFileHash(fp)

	if err != nil {
		return false
	}

	if hash != story.FileHash {
		return false
	}

	fileInfo, err := os.Stat(fp)

	if err != nil {
		return false
	}

	return fileInfo.ModTime().Format("2006-01-02 15:04:05") == story.FileUpdateTime.Format("2006-01-02 15:04:05")
}

func buildjobs() {
	people.BuildMyLoad(job.GetMyloads()...)
	people.BuildAmbassador(job.GetAmbassadors()...)
	people.BuildGeneral(job.GetGenerals()...)
	people.BuildManager(job.GetManagers()...)
	people.BuildSpy(job.GetSpys()...)
	people.BuildFather(job.GetFathers()...)
	people.BuildDoctor(job.GetDoctors()...)
	people.BuildRandom(job.GetRandoms()...)
}

func makefriends() {
	people.MakeFriends(getMyBrothersAndSisters(),
		getMySonsAndDaughters(),
		getMyBigFamiliesFriends(),
		bi.GetBiFriends(),
		wu.GetFriends(),
		lou.GetLouFriends(),
		li.GetFriends(),
		chu.GetFriends(),
	)
}

func getMySonsAndDaughters() []int {
	var peoples []int

	peoples = append(peoples, fu.Me)
	peoples = append(peoples, fu.MySon1)
	peoples = append(peoples, fu.MySon2)
	peoples = append(peoples, fu.MySon3)
	peoples = append(peoples, fu.MySon4)
	peoples = append(peoples, fu.MySon5)
	peoples = append(peoples, fu.MySon6)
	peoples = append(peoples, fu.MySon7)
	peoples = append(peoples, fu.MySon8)
	peoples = append(peoples, fu.MySon9)
	peoples = append(peoples, fu.MySon10)
	peoples = append(peoples, fu.MyDaughter1)
	peoples = append(peoples, fu.MyDaughter2)
	peoples = append(peoples, fu.MyDaughter3)
	peoples = append(peoples, fu.MyDaughter4)
	peoples = append(peoples, fu.MyDaughter5)
	peoples = append(peoples, fu.MyDaughter6)
	peoples = append(peoples, fu.MyDaughter7)
	peoples = append(peoples, fu.MyDaughter8)
	peoples = append(peoples, fu.MyDaughter9)
	peoples = append(peoples, fu.MyDaughter10)

	return peoples
}

func getMyBrothersAndSisters() []int {
	var peoples []int

	peoples = append(peoples, fu.Me)
	peoples = append(peoples, fu.MyWife)
	peoples = append(peoples, fu.MyBrother1)
	peoples = append(peoples, fu.MyBrother2)
	peoples = append(peoples, fu.MyBrother3)
	peoples = append(peoples, fu.MyBrother4)
	peoples = append(peoples, fu.MyBrother5)
	peoples = append(peoples, fu.MyBrother6)
	peoples = append(peoples, fu.MyBrother7)
	peoples = append(peoples, fu.MyBrother8)
	peoples = append(peoples, fu.MyBrother9)
	peoples = append(peoples, fu.MyBrother10)
	peoples = append(peoples, fu.MySister1)
	peoples = append(peoples, fu.MySister2)
	peoples = append(peoples, fu.MySister3)
	peoples = append(peoples, fu.MySister4)
	peoples = append(peoples, fu.MySister5)
	peoples = append(peoples, fu.MySister6)
	peoples = append(peoples, fu.MySister7)
	peoples = append(peoples, fu.MySister8)
	peoples = append(peoples, fu.MySister9)
	peoples = append(peoples, fu.MySister10)

	return peoples
}

func getMyBigFamiliesFriends() []int {
	var peoples []int

	peoples = append(peoples, fu.Me)
	peoples = append(peoples, bi.M毕元振2608618)
	peoples = append(peoples, li.M李适217709)
	peoples = append(peoples, li.M李诵217710)
	peoples = append(peoples, lou.M楼彦玮2611890)

	return peoples
}

func getMyBigLoads() []int {
	var peoples []int

	peoples = append(peoples, fu.Me)
	peoples = append(peoples, 2830311)
	peoples = append(peoples, 2840690)
	peoples = append(peoples, lin.M林承庆10960227)

	return peoples
}

func getGermanFriends() []int {
	var peoples []int

	peoples = append(peoples, 2845721)
	peoples = append(peoples, 2840086)

	return peoples
}

func getYiwubuFriends() []int {
	var peoples []int

	peoples = append(peoples, wuyibu.F善理_守礼II)
	peoples = append(peoples, wuyibu.M多梅希_守礼II)
	peoples = append(peoples, 2843524)
	peoples = append(peoples, 2841943)
	peoples = append(peoples, 2838684)
	peoples = append(peoples, 2847861)
	peoples = append(peoples, 2847861)
	peoples = append(peoples, 2864693)
	peoples = append(peoples, 2853577)
	peoples = append(peoples, 2856195)
	peoples = append(peoples, 2850307)

	return peoples
}

func pollinate() {
	//yuan.Hy()
	//pictish.Hy()
	//fu.Hy()
	//lin.Pollinate()
	//wu.Pollinate()
	//bi.Pollinate()
	//lou.Pollinate()
}
