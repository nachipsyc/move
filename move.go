package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// コマンドライン引数を明示
	sourceDirPtr := flag.String("source", "", "移動元ディレクトリのパス")
	targetDirPtr := flag.String("target", "", "移動先ディレクトリのパス")

	// 入力をパース
	flag.Parse()

	// 引数が正しくない場合は実行方法を明示
	if *sourceDirPtr == "" || *targetDirPtr == "" {
		log.Fatal("使用方法: go run move.go -source=<移動元ディレクトリ> -target=<移動先ディレクトリ>")
	}

	sourceDir := *sourceDirPtr
	targetDir := *targetDirPtr

	// 移動元ディレクトリの中の全てのファイル情報を取得
	sourceFiles, err := getFiles(sourceDir)

	if err != nil {
		log.Fatal(err)
	}

	// JPG → RAF のマッピングを作成
	targetFileMap := createTargetFileMap(sourceFiles)

	// RAF ファイルを移動
	moveFiles(sourceFiles, targetFileMap, sourceDir, targetDir)

	log.Println("move completed!")
}

// used by main()
func getFiles(source_dir string) ([]os.DirEntry, error) {
	// 対象ディレクトリの中のファイル全てを取得、格納
	files, err := os.ReadDir(source_dir)

	// エラーがあれば"err"を返す
	if err != nil {
		return nil, err
	}

	return files, nil
}

// used by main()
func createTargetFileMap(files []os.DirEntry) map[string]struct{} {
	targetFileMap := make(map[string]struct{})
	for _, file := range files {
		// JPEGファイルだけを選別して、ファイル名のみをスライスに格納
		ext := strings.ToLower(filepath.Ext(file.Name()))
		if ext == ".jpg" || ext == ".jpeg" {
			raw_file := strings.TrimSuffix(file.Name(), ext) + ".RAF"
			targetFileMap[raw_file] = struct{}{}
		}
	}
	return targetFileMap
}

// used by main()
func moveFiles(files []os.DirEntry, targetFileMap map[string]struct{}, sourceDir, targetDir string) {
	for _, file := range files {
		// ファイル名を用いて移行対象か判定して対象なら移行を実行
		if _, exists := targetFileMap[file.Name()]; exists {
			sourcePath := filepath.Join(sourceDir, file.Name())
			targetPath := filepath.Join(targetDir, file.Name())

			if err := os.Rename(sourcePath, targetPath); err != nil {
				log.Printf("Failed %s: %v", file.Name(), err)
			} else {
				log.Printf("Moved: %s → %s", file.Name(), targetPath)
			}
		}

	}
}
