package handler

import (
	"context"
	"fmt"
	"io/ioutil"
	"strings"

	log "github.com/micro/go-micro/v2/logger"

	healer "healer/proto/healer"

	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-micro/v2/config/source/file"
)

type Healer struct{}

func init() {
	if err := config.Load(file.NewSource(
		file.WithPath("./config.json"),
	)); err != nil {
		fmt.Println(err)
		return
	}
	conf := config.Map()
	fmt.Println(conf)
}

// HealList is a single request handler called via client.HealList or the generated client code
func (e *Healer) HealList(ctx context.Context, req *healer.CallRequest, rsp *healer.HealResponse) error {
	log.Info("Received Healer.HealList request")
	var dirPath string = config.Get("dir_path").String("./")
	var data []*healer.HealerGroupData = generateFileList(dirPath)
	rsp.Data = &healer.HealerInnerResponse{HealerData: data}
	return nil
}

func (e *Healer) Categories(ctx context.Context, req *healer.CallRequest, rsp *healer.CateResponse) error {
	log.Info("Received Healer.Categories request")
	var dirPath string = config.Get("dir_path").String("./")
	var files *[]string = getFileList(dirPath)
	data := []*healer.CategoryData{}
	for _, c := range *files {
		data = append(data, &healer.CategoryData{Name: c, Desc: getCategoryDesc(c)})
	}
	rsp.Data = &healer.CategoryInnerResponse{Category: data}
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Healer) Stream(ctx context.Context, req *healer.StreamingRequest, stream healer.Healer_StreamStream) error {
	log.Infof("Received Healer.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&healer.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Healer) PingPong(ctx context.Context, stream healer.Healer_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&healer.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}

func generateFileList(dir string) []*healer.HealerGroupData {
	var serverAddress = config.Get("server_address").String("http://woopsserver/")

	files, err := ioutil.ReadDir(dir)
	groupDatas := []*healer.HealerGroupData{}
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		filename := file.Name()
		groupDataItems := []*healer.HealerSingleData{}
		if file.IsDir() {
			subFiles, subErr := ioutil.ReadDir(dir + "/" + filename)
			if subErr != nil {
				panic(subErr)
			}
			for _, subFile := range subFiles {
				var subName = subFile.Name()
				if !strings.Contains(subName, "DS_Store") {
					mediaType := healer.HealerSingleData_Video
					var mediaUrl string = serverAddress + filename + "/" + subName
					var coverUrl string = ""
					if subFile.IsDir() {
						mediaType = healer.HealerSingleData_Image
						mediaUrl = serverAddress + filename + "/" + subName + "/media.mp4"
						coverUrl = serverAddress + filename + "/" + subName + "/cover.png"
					}
					groupDataItems = append(groupDataItems, &healer.HealerSingleData{
						Name:     subName,
						Url:      mediaUrl,
						CoverUrl: coverUrl,
						Type:     mediaType,
						Desc:     "",
					})
				}
			}
			groupDatas = append(groupDatas, &healer.HealerGroupData{Tag: filename, Items: groupDataItems})
		}
	}
	return groupDatas
}

func getCategoryDesc(category string) string {
	var categoryMaps = config.Get("categories").StringMap(make(map[string]string))
	return categoryMaps[category]
}

func getFileList(dir string) *[]string {
	files, err := ioutil.ReadDir(dir)
	var fileslist []string
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		filename := file.Name()
		if file.IsDir() {
			fileslist = append(fileslist, filename)
		}
	}
	return &fileslist
}
