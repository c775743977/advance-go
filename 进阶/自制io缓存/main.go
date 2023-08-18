package main

import (
	"fmt"
	"time"
	"os"
)

var (
	log string = "cannot determine module path for source directory\n"
)

type BufferedFileWriter struct {
	fwriter *os.File
	cache []byte
	cacheIndex int
}

func newBufferedFileWriter(file *os.File, max int) *BufferedFileWriter {
	return &BufferedFileWriter{
		fwriter : file,
		cache : make([]byte, max),
		cacheIndex : 0,
	}
}

func(this *BufferedFileWriter) WriteByte(content []byte) {
	// content大于缓存区容量则直接写入
	if len(content) >= cap(this.cache) {
		// 考虑到写入的先后问题，需要先将当前缓存区的内容写入，然后再将本次内容写入
		this.Flush()
		_, err := this.fwriter.Write(content)
		if err != nil {
			fmt.Println("cbufio error:", err)
			return
		}
	} else {
		// 目前缓存区无法存下content，则先将缓存区内容写入，再将content存入
		if len(content) + this.cacheIndex >= cap(this.cache) {
			this.Flush()
		}
		// 不能用append（至少第一次不能用），因为make的时候没有指定len，所以make时会给所有位置上都填充上空字符，所以cache的初始长度就位1024，再append就会扩容
		// this.cache = append(this.cache, content...)
		copy(this.cache[this.cacheIndex:], content)
		this.cacheIndex += len(content)
	}
}

func(this *BufferedFileWriter) Flush() {
	_, err := this.fwriter.Write(this.cache[:this.cacheIndex])
	if err != nil {
		fmt.Println("cbufio error:", err)
		return
	}
	// this.cache = this.cache[0:0]
	this.cacheIndex = 0
}

func(this *BufferedFileWriter) Close() {
	this.Flush()
	this.fwriter.Close()
}

func WriteFile(content []byte) {
	file, err := os.OpenFile("./log.txt", os.O_WRONLY | os.O_CREATE | os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println("open file error:", err)
		return
	}
	defer file.Close()

	_, err = file.Write(content)
	if err != nil {
		fmt.Println("write error:", err)
		return
	}
}

func main() {
	// now := time.Now()

	// file, err := os.OpenFile("./log.txt", os.O_WRONLY | os.O_CREATE | os.O_TRUNC, 0666)
	// if err != nil {
	// 	fmt.Println("open file error:", err)
	// 	return
	// }
	// for i := 0; i < 10000; i++ {
	// 	file.Write([]byte(log))
	// }

	// fmt.Println("直接读写耗时:", time.Since(now))

	now := time.Now()

	file, err := os.OpenFile("./log.txt", os.O_WRONLY | os.O_CREATE | os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println("open file error:", err)
		return
	}
	bwriter := newBufferedFileWriter(file, 4096)
	for i := 0; i < 10000; i++ {
		bwriter.WriteByte([]byte(log))
	}
	defer bwriter.Close()
	fmt.Println("带缓冲的读写耗时:", time.Since(now))
}