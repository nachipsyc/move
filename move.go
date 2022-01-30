package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	flag.Parse()

	// 移動元ディレクトリを引数から取得
	source_dir := flag.Arg(0)

	// 移動先ディレクトリを引数から取得
	target_dir := flag.Arg(1)

	// 移動元ディレクトリの中の全ファイルを取得
	files, err := getFiles(source_dir)

	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	// 拡張子が".JPG"のファイルの拡張子を移動用に".ARW"に変更して格納
	target_file_names := getTargetFileNames(files)

	// 指定したファイルを移動元ディレクトリから移動先ディレクトリへ移動
	moveFiles(files, target_file_names, source_dir, target_dir)
}

func getFiles(source_dir string) ([]os.FileInfo, error) {
	// 移動元ディレクトリの中のファイル全てを取得、格納
	files, err := ioutil.ReadDir(source_dir)
	// エラーがあれば"err"を返す
	if err != nil {
		return nil, err
	}

	return files, nil
}

func getTargetFileNames(files []os.FileInfo) []string {
	var target_file_names []string
	for _, file := range files {
		// JPEGファイルだけを選別
		if filepath.Ext(file.Name()) == ".JPG" {
			// 拡張子を".ARW"に変えて格納
			target_file_names = append(target_file_names, strings.Replace(file.Name(), ".JPG", ".RAF", -1))
		}
	}
	// 格納したファイル名のスライスを返す
	return target_file_names
}

func moveFiles(files []os.FileInfo, target_file_names []string, source_dir string, target_dir string) {
	for _, file := range files {
		for _, target_file_name := range target_file_names {
			// 移動用の疑似的なファイル名と一致するファイルを選別
			if filepath.Base(file.Name()) == target_file_name {
				// ファイル移動
				err := os.Rename(source_dir+"/"+file.Name(), target_dir+"/"+file.Name())
				// エラーがあればその場で出力
				if err != nil {
					fmt.Println(err)
				}

			}
		}
	}
}
