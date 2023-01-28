package main

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"
)

const (
	head     = `<div class="ebook-home-title">`
	ebook    = `<a class="ebook-home-item" target="_blank" href="/ebook/`
	tail     = `</div>`
	standUrl = `https://developer.aliyun.com/ebook/index/ebook-classify-3-practice__0_0_0_`
	symbol   = `>...<`
	symbol2  = `#ebook-list">`
)

type Books struct {
	raw      []byte
	bookMaps map[int64]struct{}
	page     int
	dbCache  *AliEBook
	finish   bool
}

func main() {
	var ali = Books{page: 0, dbCache: &AliEBook{}}
	//mysql
	books, err := GetAllBooks()
	if err != nil {
		log.Println(err.Error())
		return
	}
	ali.bookMaps = make(map[int64]struct{}, len(books))
	for _, book := range books {
		ali.bookMaps[book.Ebook] = struct{}{}
	}
	log.Println("任务启动，当前数量:" + strconv.Itoa(len(ali.bookMaps)))
	for {
		now := time.Now()
		ali.tasks(now)
		log.Println("执行完成 time:" + time.Now().Sub(now).String() + "  当前数量:" + strconv.Itoa(len(ali.bookMaps)))
		time.Sleep(time.Hour)
	}
}

func (b *Books) tasks(now time.Time) {
	b.finish = false
	b.page = 0
	b.dbCache.Time = now.Format("2006-01-02 15:04:05")

	//1 page
	_url := standUrl + strconv.Itoa(1)
	b.getWebData(_url)
	if b.finish {
		return
	}

	for i := 2; i <= b.page; i++ {
		_url = standUrl + strconv.Itoa(i)
		b.getWebData(_url)
		if b.finish {
			return
		}
		time.Sleep(time.Second)
		println(i)
	}
}

func (b *Books) getWebData(url string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err.Error())
		return
	}

	b.raw, err = io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err.Error())
		return
	}

	b.DataFilter()
	if b.page == 0 {
		b.GetCountPages()
	}
}

func (b *Books) GetCountPages() {
	index := bytes.Index(b.raw, []byte(symbol))
	if index == -1 {
		return
	}
	b.raw = b.raw[index+len(symbol):]
	index = bytes.Index(b.raw, []byte(symbol2))
	if index == -1 {
		return
	}
	b.raw = b.raw[index+len(symbol2):]
	num := make([]byte, 0, 100)
	for _, b := range b.raw {
		if b == '<' {
			break
		}
		num = append(num, b)
	}
	if len(num) > 0 {
		b.page, _ = strconv.Atoi(string(num))
	}
}

func (b *Books) DataFilter() {
	index := bytes.Index(b.raw, []byte(head))
	for ; index != -1; index = bytes.Index(b.raw, []byte(head)) {
		n := index + len(head)
		if index = bytes.Index(b.raw[n:], []byte(tail)); index == -1 {
			continue
		}
		b.dbCache.Name = string(b.raw[n : n+index])
		b.dbCache.Ebook = 0
		_index := bytes.Index(b.raw, []byte(ebook))
		if _index != -1 {
			n2 := _index + len(ebook)
			if _index2 := bytes.IndexByte(b.raw[n2:], '"'); _index2 != -1 {
				bookId := string(b.raw[n2 : n2+_index2])
				_num, _ := strconv.Atoi(bookId)
				b.dbCache.Ebook = int64(_num)
				if _, ok := b.bookMaps[b.dbCache.Ebook]; ok {
					b.finish = true
					return
				}
			}
		}
		err := Mssql.Create(b.dbCache).Error
		if err != nil {
			log.Println(err)
			b.finish = true
			return
		}
		b.bookMaps[b.dbCache.Ebook] = struct{}{}
		b.raw = b.raw[n+index+len(tail):]
	}
	return
}
