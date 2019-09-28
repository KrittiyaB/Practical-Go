package main

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

type Status struct {
	Code   int
	Reason string
}

type Header struct {
	Key, Value string
}

type errWrite struct { // create new type encapsulate io.Write
	io.Writer
	err error
}

func (e *errWrite) Write(buf []byte) (int, error) {
	if e.err != nil {
		return 0, e.err
	}
	var n int
	n, e.err = e.Writer.Write(buf)
	return n, nil
}

func WriteResponse(w io.Writer, st Status, headers []Header, body io.Reader) error {
	ew := &errWrite{
		Writer: w,
		err:    nil,
	}
	fmt.Fprintf(ew, "HTTP/1.1 %d %s\r\n", st.Code, st.Reason)

	for _, h := range headers {
		fmt.Fprintf(ew, "%s: %s\r\n", h.Key, h.Value)
	}

	fmt.Fprintf(ew, "\r\n")
	io.Copy(ew, body)
	return ew.err
}

// you can refactor this code after you find the error pattern
//func WriteResponse(w io.Writer, st Status, headers []Header, body io.Reader) error {
//	_, err := fmt.Fprintf(w, "HTTP/1.1 %d %s\r\n", st.Code, st.Reason)
//	if err != nil {
//		return err
//	}
//
//	for _, h := range headers {
//		_, err := fmt.Fprintf(w, "%s: %s\r\n", h.Key, h.Value)
//		if err != nil {
//			return err
//		}
//	}
//
//	if _, err := fmt.Fprint(w, "\r\n"); err != nil {
//		return err
//	}
//
//	_, err = io.Copy(w, body)
//	return err
//}

func main() {
	var buf bytes.Buffer
	st := Status{Code: 555, Reason: "account not found."}
	headers := []Header{
		{"Content-Type", "application/json"},
	}
	body := strings.NewReader("this is a body")

	err := WriteResponse(&buf, st, headers, body)

	fmt.Printf("buf: \n%s\n", buf)
	fmt.Println("error:", err)
}
