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

	// 移動元ディレクトリのパスを引数から取得
	source_dir := flag.Arg(0)

	// 移動先ディレクトリのパスを引数から取得
	target_dir := flag.Arg(1)

	// 移動元ディレクトリの中の全てのファイル情報を取得
	source_files, err := getFiles(source_dir)

	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	// 拡張子が".JPG"のファイル名を取得する
	source_file_names := getSourceFileNames(source_files)

	// 取得したJPEGファイルのファイル名を一致判定用に".RAF"に変換する
	target_file_names := getTargetFileNames(source_file_names)

	// 指定したファイルを移動元ディレクトリから移動先ディレクトリへ移動
	moveFiles(source_files, target_file_names, source_dir, target_dir)

	fmt.Println("move completed!")
}

func getFiles(source_dir string) ([]os.FileInfo, error) {
	// 移動元ディレクトリの中のファイル全てを取得、格納
	source_files, err := ioutil.ReadDir(source_dir)
	// エラーがあれば"err"を返す
	if err != nil {
		return nil, err
	}

	return source_files, nil
}

func getSourceFileNames(files []os.FileInfo) []string {
	var source_file_names []string
	for _, file := range files {
		// JPEGファイルだけを選別して、ファイル名のみをスライスに格納
		file_extention := filepath.Ext(file.Name())
		if file_extention == ".JPG" || file_extention == ".jpg" || file_extention == ".jpeg" {
			source_file_names = append(source_file_names, file.Name())
		}
	}
	// 格納したファイル名のスライスを返す
	return source_file_names
}

func getTargetFileNames(file_names []string) []string {
	var target_file_names []string
	for _, file_name := range file_names {
		// 拡張子を".JPG"から".RAF"に変換してスライスに格納
		switch filepath.Ext(file_name) {
		case ".JPG":
			target_file_names = append(target_file_names, strings.Replace(file_name, ".JPG", ".RAF", -1))
		case ".jpg":
			target_file_names = append(target_file_names, strings.Replace(file_name, ".jpg", ".RAF", -1))
		case ".jpeg":
			target_file_names = append(target_file_names, strings.Replace(file_name, ".jpeg", ".RAF", -1))
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
