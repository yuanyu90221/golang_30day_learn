package main

import (
	"container/ring"
	"fmt"
	"time"
)

var printRing = func(v interface{}) {
	fmt.Print(v.(int), "->")
}

func PrintRing(r *ring.Ring) {
	r.Do(printRing)
	fmt.Println()
}
func main() {
	var varRing ring.Ring
	fmt.Println("查看透過var宣告的ring的長度: ", varRing.Len())
	fmt.Println("----------------------")
	//透過New創建10個元素的ring
	r := ring.New(10)
	// 查看透過New()初始化ring的長度
	fmt.Println("查看透過New()初始化ring的長度: ", r.Len())
	// 給ring中每個元素進行走訪並且給值
	for i := 0; i < 10; i++ {
		r.Value = i
		// 取得下一個元素
		r = r.Next()
	}
	fmt.Print("r : ")
	PrintRing(r)

	// 往後移動ring的指向
	r = r.Move(2)
	fmt.Println("ring 向後移動2個位置的元素值:", r.Value)

	// 往前移動ring的指向
	r = r.Move(-8)
	fmt.Println("ring 向前移動8個位置的元素值:", r.Value)

	// 從ring當前指向開始刪除n個元素
	deletedElm := r.Unlink(2)
	fmt.Print("r 所剩下的元素 : ")
	PrintRing(r)
	fmt.Print("從r刪除的元素 : ")
	PrintRing(deletedElm)

	// 準備第2個ring r2
	r2 := ring.New(3)
	for i := 0; i < 3; i++ {
		r2.Value = i + 10
		r2 = r2.Next()
	}
	fmt.Print("r2 : ")
	PrintRing(r2)

	fmt.Println("現在r的指向在 :", r.Value)

	// Link r 跟 r2
	fmt.Print("Link r 跟 r2 : ")
	linkedRing := r.Link(r2)
	PrintRing(r)

	// 以原本r.Next()開始走訪
	fmt.Print("以原本r.Next()開始走訪 : ")
	PrintRing(linkedRing)
	SongListDemo()
}

type song struct {
	name   string
	artist string
	length time.Duration
}

func SongListDemo() {
	var (
		songs = []song{
			{
				name:   "Something Just Like This",
				artist: "The Chainsmokers",
				length: 247,
			},
			{
				name:   "Blame",
				artist: "Calvin Harris",
				length: 214,
			},
			{
				name:   "Wolves",
				artist: "Selena Gomez",
				length: 197,
			},
			{
				name:   "Sing You To Sleep",
				artist: "Matt Cab",
				length: 236,
			},
		}
	)

	songList := ring.New(len(songs))
	repeatedCnt := 0

	for i := 0; i < songList.Len(); i++ {
		songList.Value = songs[i]
		songList = songList.Next()
	}

	for {
		if repeatedCnt == 3 {
			break
		}
		songList.Do(func(v interface{}) {
			time.Sleep((v.(song).length / 100) * time.Second)
			fmt.Printf("現正播放%s, 演唱者為%s\n", v.(song).name, v.(song).artist)
		})
		repeatedCnt++
		fmt.Printf("播放次數 : %d\n", repeatedCnt)
	}
	fmt.Println("播放完畢")
}
