package main

import (
	"bufio"
	"io"
)

/*
_, err = fd.Write(p0)
if err != nil {
    return err
}
_, err = fd.Write(p1)
if err != nil {
    return err
}
_, err = fd.Write(p2)
if err != nil {
    return err
}
// and so on
*/


type errWriter struct {
	w io.Writer
	err error
}

func (ew *errWriter) write(buf []byte) {
	if ew.err != nil {
		return 
	}
	_, ew.err = ew.w.Write(buf)
}

// 大部分标准库里存在这种处理方法，例如bufio.Write已经存在
/*
func bufioWrite() {
	b := bufio.NewWriter()
	b.Write()
	b.Write()
	b.Write()

	if b.Flush() != nil {
		return b.Flush()
	}
}
*/
