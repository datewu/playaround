package main

import (
	"archive/tar"
	"archive/zip"
	"compress/flate"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	// writeTar()
	// readTar()

	// writeZip()
	// readZip()

	// noCompression()

	gzipCompression()
}

var files = []string{
	"proverbs1.txt",
	"proverbs2.txt",
	"proverbs3.txt",
}

func writeTar() {
	tf, err := os.Create("proverbs.tar")
	if err != nil {
		log.Fatalln(err)
	}
	defer tf.Close()

	tw := tar.NewWriter(tf)

	for _, fn := range files {
		f, err := os.Open(fn)
		if err != nil {
			log.Println(err)
			continue
		}
		defer f.Close()
		info, _ := f.Stat()
		h := &tar.Header{
			Name:    info.Name(),
			Mode:    0666,
			Size:    info.Size(),
			ModTime: info.ModTime(),
		}

		if err := tw.WriteHeader(h); err != nil {
			log.Println(err)
			continue
		}
		bs, err := ioutil.ReadAll(f)
		if err != nil {
			log.Println(err)
			continue
		}
		if _, err := tw.Write(bs); err != nil {
			log.Println(err)
		}
	}
	if err := tw.Close(); err != nil {
		log.Fatalln(err)
	}

}

func readTar() {
	f, err := os.Open("proverbs.tar")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	tr := tar.NewReader(f)

	for {
		h, err := tr.Next()
		if err == io.EOF {
			fmt.Println("======= reached the end")
			break
		}
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println(h.Name, "------------")
		if _, err := io.Copy(os.Stdout, tr); err != nil {
			log.Fatalln(err)
		}
	}
}

func writeZip() {
	zf, err := os.Create("proverbs.zip")
	if err != nil {
		log.Fatalln(err)
	}
	defer zf.Close()

	zw := zip.NewWriter(zf)

	for _, fn := range files {
		f, err := os.Open(fn)
		if err != nil {
			log.Println(err)
			continue
		}
		defer f.Close()

		fw, err := zw.Create(f.Name())
		if err != nil {
			log.Println(err)
			continue
		}
		bs, err := ioutil.ReadAll(f)
		if err != nil {
			log.Println(err)
			continue
		}
		if _, err := fw.Write(bs); err != nil {
			log.Println(err)
		}
	}
	if err := zw.Close(); err != nil {
		log.Fatalln(err)
	}
}

func readZip() {
	r, err := zip.OpenReader("proverbs.zip")
	if err != nil {
		log.Fatalln(err)
	}
	defer r.Close()

	for _, f := range r.File {
		fmt.Println(f.Name, "===============")

		rc, err := f.Open()
		if err != nil {
			log.Println(err)
			continue
		}
		if _, err := io.Copy(os.Stdout, rc); err != nil {
			log.Println(err)
		}

		rc.Close()

	}
}

func noCompression() {
	zf, err := os.Create("proverbs-nocompress.zip")
	if err != nil {
		log.Fatalln(err)
	}
	defer zf.Close()

	zw := zip.NewWriter(zf)

	zw.RegisterCompressor(zip.Deflate, func(out io.Writer) (io.WriteCloser, error) {
		return flate.NewWriter(out, flate.NoCompression)
	})

	for _, fn := range files {
		f, err := os.Open(fn)
		if err != nil {
			log.Println(err)
			continue
		}
		defer f.Close()

		fw, err := zw.Create(f.Name())
		if err != nil {
			log.Println(err)
			continue
		}
		bs, err := ioutil.ReadAll(f)
		if err != nil {
			log.Println(err)
			continue
		}
		if _, err := fw.Write(bs); err != nil {
			log.Println(err)
		}
	}
	if err := zw.Close(); err != nil {
		log.Fatalln(err)
	}
}

func gzipCompression() {
	gzfn := "proverbs.txt.gz"

	gzf, err := os.Create(gzfn)
	if err != nil {
		log.Fatalln(err)
	}

	gzw, err := gzip.NewWriterLevel(gzf, gzip.BestCompression)
	if err != nil {
		log.Fatalln(err)
	}

	for _, fn := range files {
		f, err := os.Open(fn)
		if err != nil {
			log.Println(err)
			continue
		}
		defer f.Close()

		bs, err := ioutil.ReadAll(f)
		if err != nil {
			log.Println(err)
			continue
		}

		if _, err := gzw.Write(bs); err != nil {
			log.Println(err)
		}
	}

	if err := gzw.Close(); err != nil {
		log.Fatalln(err)
	}
	if err := gzf.Close(); err != nil {
		log.Fatalln(err)
	}

	// read it back out
	gzf, err = os.Open(gzfn)
	if err != nil {
		log.Fatalln(err)
	}
	gzr, err := gzip.NewReader(gzf)
	if err != nil {
		log.Fatalln(err)
	}
	if _, err := io.Copy(os.Stdout, gzr); err != nil {
		log.Fatalln(err)
	}
	if err := gzw.Close(); err != nil {
		log.Fatalln(err)
	}
	if err := gzf.Close(); err != nil {
		log.Fatalln(err)
	}

}
