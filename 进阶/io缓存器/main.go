package main

import (
	"fmt"
	"os"
	"time"
)

var (
	text = "recall 74 by radic, recall 0 by milvus_short, recall 0 by milvus_long\n"
)

// 带缓存的FileWriter
type BufferedFileWriter struct {
	cache         []byte   //缓存的内容
	cacheEndIndex int      //cache里有效内容的结束位置
	fout          *os.File //文件句柄
}

func NewWriter(ft *os.File, cacheSize int) *BufferedFileWriter {
	return &BufferedFileWriter{
		cache:         make([]byte, cacheSize), //len=cap=cacheSize
		cacheEndIndex: 0,
		fout:          ft,
	}
}

// 向文件中写入内容。（大概率只是写入了缓存，还没有真正写入磁盘）
func (w *BufferedFileWriter) WriteByte(cont []byte) {
	if len(cont) >= len(w.cache) { //要写的内容比缓存空间还要大，则直接把cont写到文件里去
		w.Flush()
		w.fout.Write(cont)
	} else {
		//先预判cache能否容下cont
		if w.cacheEndIndex+len(cont) > len(w.cache) { //不能容下
			w.Flush()
		}
		// append2(w.cache, w.cacheEndIndex, cont)
		copy(w.cache[w.cacheEndIndex:], cont) //golang内置的copy函数功能上等价于自己写的append2函数，但比append2函数更高效
		w.cacheEndIndex += len(cont)
	}
}

// 把cache里的内容全部写入磁盘文件
func (w *BufferedFileWriter) Flush() {
	w.fout.Write(w.cache[:w.cacheEndIndex]) //把cache里的内容写入文件
	w.cacheEndIndex = 0                     //清空cache

}

// 把src拷贝到dest[index:]里去
func append2(dest []byte, index int, src []byte) {
	for i := 0; i < len(src); i++ {
		dest[index+i] = src[i]
	}
}

// 向文件中写入内容。（大概率只是写入了缓存，还没有真正写入磁盘）
func (writer *BufferedFileWriter) WriteString(content string) {
	writer.WriteByte([]byte(content))
}

// 直接写文件
func WriteDirect(outFile string) {
	fout, err := os.OpenFile(outFile, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	defer fout.Close()

	for i := 0; i < 10000; i++ {
		fout.WriteString(text)
	}
}

// 带缓冲写文件
func WriteWithBuffer(outFile string) {
	fout, err := os.OpenFile(outFile, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	defer fout.Close()

	writer := NewWriter(fout, 4096)
	defer writer.Flush() //最后，务必把缓存里残留的内容写入磁盘
	for i := 0; i < 10000; i++ {
		writer.WriteString(text)
	}
}

func main() {
	begin := time.Now()
	WriteDirect("./a.txt") //耗时24ms
	// WriteWithBuffer("./b.txt") //耗时2ms
	fmt.Printf(" %dms\n", time.Since(begin).Milliseconds())
}