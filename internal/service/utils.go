package service

import (
	"book-school/internal/models"
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"os"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"golang.org/x/net/html"
)

func saveImage(filePath string, data, index string) error {
	i := strings.Index(data, ",")
	if i < 0 {
		return errors.New("no comma in base64")
	}
	unbased, err := base64.StdEncoding.DecodeString(data[i+1:])
	if err != nil {
		return fmt.Errorf("error encoding base64: %s", err)
	}
	err = os.WriteFile(filePath+index+".png", []byte(unbased), 0777)
	if err != nil {
		return fmt.Errorf("error ocured by saving image : %s", err)
	}
	return nil
}

func savePreview(path string, img multipart.File) error {
	out, err := os.Create(path)
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, img)
	if err != nil {
		return err
	}
	return nil
}

func hashingId(id string) string {
	h := sha256.New()
	h.Write([]byte(id))
	return fmt.Sprintf("%x", string(h.Sum(nil)))
}

func contentParser(input *models.Book, path, imageUrl string) ([]string, error) {
	var res []string
	var idImage int
	for _, content := range input.Pages {
		content = strings.ReplaceAll(content, `\"`, "")
		doc, err := html.Parse(strings.NewReader(content))
		if err != nil {
			return nil, fmt.Errorf("error occured in parsing string html to html node: %v", err.Error())
		}

		var f func(*html.Node) error
		f = func(n *html.Node) error {
			if n.Type == html.ElementNode && n.Data == "img" {
				var attribute html.Attribute
				attribute.Key = "style"
				attribute.Val = `'width: 50%;'`
				n.Attr = append(n.Attr, attribute)
				for i, attribute := range n.Attr {
					if attribute.Key == "src" {
						idImage++
						strI := strconv.Itoa(idImage)
						if err != nil {
							return err
						}
						if err := saveImage(path, attribute.Val, strI); err != nil {
							return err
						}
						n.Attr[i].Val = "'" + imageUrl + "/static/books/" + input.Hashed_ID + "/" + strI + ".png" + "'"
						break
					}
				}
				log.Println(n.Attr)
			}
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				if err := f(c); err != nil {
					return err
				}
			}
			return nil
		}

		if err := f(doc); err != nil {
			return nil, err
		}
		var page bytes.Buffer
		err = html.Render(&page, doc)
		if err != nil {
			return nil, err
		}
		resPage := strings.ReplaceAll(html.UnescapeString(page.String()), "<html><head></head><body>", "")
		resPage = strings.ReplaceAll(resPage, "</body></html>", "")
		res = append(res, resPage)
	}
	return res, nil
}

var ErrInvalidImage = errors.New("invalid image")

func saveFiles(path, dir string, file *multipart.FileHeader) (string, error) {
	fullPath := path + dir
	if err := os.MkdirAll(fullPath, os.ModePerm); err != nil {
		return "", fmt.Errorf("save files: %w", err)
	}

	if !strings.Contains(file.Header["Content-Type"][0], "image") {
		return "", fmt.Errorf("save files: %w", ErrInvalidImage)
	}

	f, err := file.Open()
	if err != nil {
		return "", fmt.Errorf("save files: %w", err)
	}
	defer f.Close()

	if !validImageType(file.Filename) {
		return "", fmt.Errorf("save files: %w: not supported file type", ErrInvalidImage)
	}

	fileName := uuid.NewString()

	fullPath += "/" + fileName + ".png"

	out, err := os.Create(fullPath)
	if err != nil {
		return "", fmt.Errorf("save files: %w", err)
	}
	defer out.Close()

	_, err = io.Copy(out, f)
	if err != nil {
		return "", fmt.Errorf("save files: %w", err)
	}

	return fullPath, nil
}

func validImageType(file string) bool {
	split := strings.Split(file, ".")
	fileType := split[len(split)-1]

	validImageType := []string{"jpeg", "jpg", "png"}
	for _, t := range validImageType {
		if t == fileType {
			return true
		}
	}

	return false
}
