package test

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"reflect"
	"testing"

	"github.com/golang/protobuf/proto"
	"gitlab.uuzu.com/war/pbtool/cmd/protoc-gen-clone/example"
	"gitlab.uuzu.com/war/pbtool/cmd/protoc-gen-clone/example/common"
)

func TestGenTestPb(t *testing.T) {
	stu := map[int32]*example.StudentMc{
		10: {Name: "xiaoming"},
		20: {Name: "xiaohong"},
	}

	drawData := map[int32]*example.DrawOnlyEquipBack_DrawData{
		900: {
			SpecialDrawCount: map[string]int32{"son": 10},
			HistoryDrawCount: 88,
			ResId:            3,
		},
	}
	dat := &example.DrawOnlyEquipBack{
		Stu:      stu,
		DrawData: drawData,
		SpecialDrawCount: map[int32]string{
			777: "stru",
			98:  "tea"},
	}
	clone := dat.Clone()
	if !equal(dat, clone) {
		t.Error("not equal")
	}
	clone2 := proto.Clone(dat)
	if !equal(dat, clone2) {
		t.Error("not equal2")
	}
}

func TestGenUserPb(t *testing.T) {
	dat, err := readSrcData()
	if err != nil {
		t.Error(err)
	}

	const testNum = 1234567
	dat.User.Base.HideVipSystems = append(dat.User.Base.HideVipSystems, testNum)

	clone := dat.Clone()
	if clone.User.Base.HideVipSystems[len(clone.User.Base.HideVipSystems)-1] != testNum {
		t.Error("not equal")
	}

	if !equal(dat, clone) {
		t.Error("not equal")
	}
	clone2 := proto.Clone(dat)
	if !equal(dat, clone2) {
		t.Error("not equal2")
	}
}

func readSrcData() (*example.User, error) {
	dat := &example.User{}
	dataJsonStruct := readTestDataFile("./user.json")
	// 解析json
	err := json.Unmarshal([]byte(dataJsonStruct), &dat.User)
	if err != nil {
		return nil, err
	}
	return dat, nil
}

func readTestDataFile(path string) string {
	fi, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	fd, err := ioutil.ReadAll(fi)
	return string(fd)
}

func equal(a proto.Message, b proto.Message) bool {
	return proto.Equal(a, b)
}

func TestValidClone(t *testing.T) {
	originUser, err := readSrcData()
	if err != nil {
		t.Error(err)
	}
	dat := originUser.User

	clone := originUser.Clone().User
	if !equal(dat, clone) {
		t.Error("not equal")
	}
	// string
	clone.Base.ServerId = 2379438
	if reflect.DeepEqual(clone.Base.ServerId, dat.Base.ServerId) {
		t.Error("string error")
	}
	// map
	clone.Base.Counsellor[763] = 83765
	if reflect.DeepEqual(clone.Base.Counsellor, dat.Base.Counsellor) {
		t.Error("map error")
	}
	// array
	clone.Base.HideVipSystems = append(clone.Base.HideVipSystems, 32637846)
	if reflect.DeepEqual(clone.Base.HideVipSystems, dat.Base.HideVipSystems) {
		t.Error("array error")
	}
	// message
	clone.Game.Rebate = &common.Rebate{TakeTime: proto.Int64(36532)}
	if reflect.DeepEqual(clone.Game.Rebate, dat.Game.Rebate) {
		t.Error("message error")
	}
	// *uint32
	clone.Game.DailyBoss.ActivityAdd = proto.Uint32(782327)
	if reflect.DeepEqual(clone.Game.DailyBoss.ActivityAdd, dat.Game.DailyBoss.ActivityAdd) {
		t.Error("uint32 error")
	}
}

func BenchmarkCloneDB_WarClone(b *testing.B) {
	dat, err := readSrcData()
	if err != nil {
		b.Error(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = dat.Clone()
	}
}

func BenchmarkCloneDB_PbClone(b *testing.B) {
	dat, err := readSrcData()
	if err != nil {
		b.Error(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = proto.Clone(dat)
	}
}

func BenchmarkCloneDB_PbMarshal(b *testing.B) {
	dat, err := readSrcData()
	if err != nil {
		b.Error(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = proto.Marshal(dat)
	}
}

// cpu: Intel(R) Core(TM) i5-1038NG7 CPU @ 2.00GHz
// BenchmarkCloneDB_WarClone
// BenchmarkCloneDB_WarClone-8    	    1568	    805789 ns/op	  368876 B/op	   11184 allocs/op
// BenchmarkCloneDB_PbClone
// BenchmarkCloneDB_PbClone-8     	     262	   4602689 ns/op	  571961 B/op	   21785 allocs/op
// BenchmarkCloneDB_PbMarshal
// BenchmarkCloneDB_PbMarshal-8   	     100	  12359954 ns/op	  638916 B/op	   32709 allocs/op
