Quick setup — if you’ve done this kind of thing before
or	
https://github.com/tom00ti/go_practice.git
Get started by creating a new file or uploading an existing file. We recommend every repository include a README, LICENSE, and .gitignore.

…or create a new repository on the command line
echo "# go_practice" >> README.md
git init
git add README.md
git commit -m "first commit"
git branch -M main
git remote add origin https://github.com/tom00ti/go_practice.git
git push -u origin main
…or push an existing repository from the command line
git remote add origin https://github.com/tom00ti/go_practice.git
git branch -M main
git push -u origin main
…or import code from another repository
You can initialize this repository with code from a Subversion, Mercurial, or TFS project.


go は　なにがいいのか？
マルチすレディング
pythonやÇを共存させられる。
学びやすい
早い
リソースが少ない
コンパイルが一つのバイナリでできる。


パッケージフォルダを初期化する。

ti  (e) base   master  ~  GitHub  go_practice  booking-app  go mod init booking-app
go: creating new go.mod: module booking-app
go: to add module requirements and sums:
        go mod tidy

goはどこからプログラムを走らせるか宣言する必要がある。
　
変数はどこかで使われてないとエラーになる。厳格性がある

変数はデータタイプが必要

コンパイル時にエラーを見つけられる。ランタイムではない。　

-- 43:05

